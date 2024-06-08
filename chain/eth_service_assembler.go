package chain

import (
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
	var from string
	var to string
	var amount *big.Int

	// get sender
	// the signer is used to get the sender of the transaction
	signer := types.NewEIP155Signer(t.ChainId())
	sender, err := types.Sender(signer, t)

	if err != nil {
		logrus.Error("Failed to get sender: ", err)
		return nil
	}
	from = sender.String()
	// end get sender

	// this is a USDT transfer
	if t.To().String() == EthereumUsdtAddress {
		crypto = model.EthereumUSDT
		data := t.Data()
		transfer, err := a.ToUsdtTransferArguments(&data)
		if err != nil {
			logrus.Error("Failed to parse transfer arguments: ", err)
		}

		amount = transfer.Value
		to = transfer.To.String()
	} else { // this is a normal transfer
		crypto = model.Ethereum
		to = t.To().String()
		amount = t.Value()
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
