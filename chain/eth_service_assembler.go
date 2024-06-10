package chain

import (
	"encoding/hex"
	"math/big"
	"strings"
	"time"

	"github.com/chainpusher/chainpusher/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

type EthereumServiceAssembler struct {
	ParsedABI     abi.ABI
	TansferMethod *abi.Method
}

func NewEthereumServiceAssembler() (*EthereumServiceAssembler, error) {

	abi, err := abi.JSON(strings.NewReader(EthereumUsdtAbi))
	if err != nil {
		logrus.Error("Failed to parse Ethereum USDT ABI: ", err)
	}

	method, err := abi.MethodById(EthereumUsdtMethodTransfer)
	if err != nil {
		return nil, err
	}

	return &EthereumServiceAssembler{
		ParsedABI:     abi,
		TansferMethod: method,
	}, nil
}

func (a *EthereumServiceAssembler) ToUsdtTransferArguments(data *[]byte) (*EthereumContractUsdtTransfer, error) {

	var to common.Address
	var amount *big.Int

	b := (*data)
	args, err := a.TansferMethod.Inputs.Unpack(b[4:])
	if err != nil {
		return nil, err
	}

	to = args[0].(common.Address)
	amount = args[1].(*big.Int)

	return &EthereumContractUsdtTransfer{
		To:    &to,
		Value: amount,
	}, nil
}

func (a *EthereumServiceAssembler) ToTransaction(t *types.Transaction) *model.Transaction {

	var crypto model.CryptoCurrency
	var from string = PraseEthereumTransactionFromAddress(t)
	var to string
	var amount *big.Int

	// this is a USDT transfer
	txTo := t.To()
	if txTo == nil {
		logrus.Errorf("Transaction to is nil: %v", t.Hash().String())
		return nil
	}
	if t.To().String() == EthereumUsdtAddress {
		crypto = model.EthereumUSDT
		data := t.Data()
		transfer, err := a.ToUsdtTransferArguments(&data)
		if err != nil {
			logrus.Error("Failed to parse transfer arguments: ", err)
			logrus.Debugf("Data: %s", hex.EncodeToString(data))
			return nil
		}

		amount = transfer.Value
		to = transfer.To.String()
	} else if t.Value().Cmp(big.NewInt(0)) == 1 { // this is a normal transfer
		crypto = model.Ethereum
		to = t.To().String()
		amount = t.Value()
	} else { // this is a contract call
		logrus.Tracef("Contract call that is not USDT contract: %v", t.Hash().String())
		return nil
	}

	return &model.Transaction{
		Platform:       model.PlatformEthereum,
		CryptoCurrency: crypto,
		Payee:          from,
		Payer:          to,
		Amount:         *amount,
		CreatedAt:      time.Now(),
	}
}

func (a *EthereumServiceAssembler) BlockToTransactions(block *types.Block) []*model.Transaction {
	transactions := make([]*model.Transaction, 0)

	for _, tx := range block.Transactions() {
		t := a.ToTransaction(tx)
		if t == nil {
			continue
		}
		transactions = append(transactions, t)
	}

	return transactions

}

func PraseEthereumTransactionFromAddress(t *types.Transaction) string {
	var signer types.Signer
	if t.Type() == types.AccessListTxType {
		signer = types.NewEIP2930Signer(t.ChainId())
	} else if t.Type() == types.DynamicFeeTxType {
		signer = types.NewLondonSigner(t.ChainId())
	} else if t.Type() == types.BlobTxType {
		logrus.Tracef("Blob transaction: %v", t.Hash().String())
		return EthereumEmptyAddress
	} else {
		signer = types.NewEIP155Signer(t.ChainId())
	}

	sender, err := types.Sender(signer, t)

	if err != nil {
		logrus.Error("Failed to get sender: ", err, t.Type())
		return EthereumEmptyAddress
	}
	return sender.String()
}
