package chain

import (
	"errors"
	"math/big"

	"github.com/chainpusher/chainpusher/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

func FromTronTransferContract(tc *core.Transaction_Contract) *model.Transaction {
	var transferContract core.TransferContract

	tc.Parameter.UnmarshalTo(&transferContract)

	var owner address.Address = transferContract.OwnerAddress
	var to address.Address = transferContract.ToAddress
	amount := big.NewInt(transferContract.Amount)

	transfer := model.Transaction{
		Platform:       model.PlatformTron,
		CryptoCurrency: model.Tron,
		Amount:         *amount,
		Payer:          owner.String(),
		Payee:          to.String(),
	}

	return &transfer
}

func FromTronTriggerSmartContract(arg *abi.Arguments, tc *core.Transaction_Contract) (*model.Transaction, error) {

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
		},
		nil
}

func FromTronContractToTransaction(args *abi.Arguments, tc *core.Transaction_Contract) (*model.Transaction, error) {

	if tc.GetType() == core.Transaction_Contract_TransferContract {
		return FromTronTransferContract(tc), nil
	}

	if tc.GetType() == core.Transaction_Contract_TriggerSmartContract {
		return FromTronTriggerSmartContract(args, tc)
	}

	return nil, errors.New("unknown contract type")
}

func ToTransactions(args *abi.Arguments, block *api.BlockExtention) []*model.Transaction {

	var transactions []*model.Transaction
	for _, transaction := range block.Transactions {
		contracts := transaction.GetTransaction().GetRawData().GetContract()
		if len(contracts) == 0 {
			continue
		}

		for _, contract := range contracts {
			transfer, err := FromTronContractToTransaction(args, contract)
			if err == nil {
				transactions = append(transactions, transfer)
			}
		}

	}

	return transactions
}
