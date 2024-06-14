package service

import (
	"encoding/hex"
	"errors"
	"math/big"
	"time"

	"github.com/chainpusher/blockchain/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	tcAbi "github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/sirupsen/logrus"
)

type TronBlockChainAssembler struct {

	// The arguments it's of transfer function of contract
	arguments *abi.Arguments
}

func (a *TronBlockChainAssembler) ToTransactions(args *abi.Arguments, block *api.BlockExtention) []*model.Transaction {
	logrus.Debugf("Block transactions: %d", len(block.Transactions))
	var transactions []*model.Transaction
	for _, transaction := range block.Transactions {
		contracts := transaction.GetTransaction().GetRawData().GetContract()
		if len(contracts) == 0 {
			continue
		}

		for _, contract := range contracts {
			transfer, err := a.ToTransaction(args, transaction, contract)
			if err == nil {
				transactions = append(transactions, transfer)
			}
		}

	}

	return transactions
}

func (a *TronBlockChainAssembler) GetTransferArguments(
	service *TronBlockChainService,
) (*abi.Arguments, error) {

	if a.arguments != nil {
		return a.arguments, nil
	}

	abi, err := service.GetSmartContractABI(TronUsdtAddress)
	if err != nil {
		return nil, err
	}
	args, err := tcAbi.GetInputsParser(abi, "transfer")
	if err != nil {
		return nil, err
	}

	a.arguments = &args

	return &args, nil
}

func (a *TronBlockChainAssembler) ToBlock(
	block *api.BlockExtention,
	service *TronBlockChainService,
) (*model.Block, error) {

	if block.BlockHeader == nil {
		return nil, errors.New("block header is nil")
	}
	args, err := a.GetTransferArguments(service)
	if err != nil {
		return nil, err
	}
	transactions := a.ToTransactions(args, block)

	aBlock := &model.Block{
		Height:       big.NewInt(block.BlockHeader.RawData.Number),
		Id:           hex.EncodeToString(block.Blockid),
		Transactions: transactions,
	}
	return aBlock, nil
}

func (a *TronBlockChainAssembler) TransactionFromTransfer(t *api.TransactionExtention, tc *core.Transaction_Contract) *model.Transaction {
	var transferContract core.TransferContract

	tc.Parameter.UnmarshalTo(&transferContract)

	var owner address.Address = transferContract.OwnerAddress
	var to address.Address = transferContract.ToAddress
	amount := big.NewInt(transferContract.Amount)

	transfer := model.Transaction{
		Platform:       model.PlatformTron,
		CryptoCurrency: model.TRX,
		Amount:         *amount,
		Payer:          owner.String(),
		Payee:          to.String(),
		CreatedAt:      time.Unix(t.Transaction.RawData.Timestamp, 0),
	}

	return &transfer
}

func (a *TronBlockChainAssembler) TransactionFromContract(
	arg *abi.Arguments,
	t *api.TransactionExtention,
	tc *core.Transaction_Contract) (*model.Transaction, error) {

	var contract core.TriggerSmartContract
	err := tc.Parameter.UnmarshalTo(&contract)
	if err != nil {
		return nil, err
	}

	if address.Address(contract.ContractAddress).String() != TronUsdtAddress {
		return nil, errors.New("contract address is not USDT")
	}

	unpacked, err := arg.Unpack(contract.Data[4:])

	if err != nil {
		return nil, err
	}

	ethAddress := unpacked[0].(common.Address)
	amount := unpacked[1].(*big.Int)

	tronAddress := ToTronAddress(ethAddress)

	return &model.Transaction{
			Platform:       model.PlatformTron,
			CryptoCurrency: model.TronUSDT,
			Amount:         *amount,
			Payer:          address.Address(contract.OwnerAddress).String(),
			Payee:          tronAddress.String(),
			CreatedAt:      time.Unix(t.Transaction.RawData.Timestamp, 0),
		},
		nil
}

func (a *TronBlockChainAssembler) ToTransaction(args *abi.Arguments, t *api.TransactionExtention, tc *core.Transaction_Contract) (*model.Transaction, error) {

	if tc.GetType() == core.Transaction_Contract_TransferContract {
		return a.TransactionFromTransfer(t, tc), nil
	}

	if tc.GetType() == core.Transaction_Contract_TriggerSmartContract {
		return a.TransactionFromContract(args, t, tc)
	}

	return nil, errors.New("unknown contract type")
}

func NewTronBlockChainAssembler() *TronBlockChainAssembler {
	return &TronBlockChainAssembler{}
}
