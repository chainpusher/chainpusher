package chain

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthereumTransactions []*EthereumTransaction

type EthereumTransaction struct {
	BlockHash   string `json:"block_hash"`
	BlockNumber uint64 `json:"block_number"`

	Type  uint8  `json:"type"`
	Hash  string `json:"hash"`
	Nonce uint64 `json:"nonce"`

	ChainID    *big.Int         `json:"chain_id"`
	AccessList types.AccessList `json:"access_list"`
	Data       string           `json:"data"`
	Gas        uint64           `json:"gas"`
	GasPrice   *big.Int         `json:"gas_price"`
	GasTipCap  *big.Int         `json:"gas_tip_cap"`
	GasFeeCap  *big.Int         `json:"gas_fee_cap"`
	From       string           `json:"from"`
	To         string           `json:"to"`
	Value      *big.Int         `json:"value"`
	Time       time.Time        `json:"time"`
}

// FromTransactions converts a list of Ethereum transactions to a list of EthereumTransaction
func (t *EthereumTransactions) FromTransactions(block *types.Block, txs types.Transactions) *EthereumTransactions {
	for i, tx := range txs {
		var s EthereumTransaction
		s.FromTransaction(block, tx)
		(*t)[i] = &s
	}
	return t
}

func (t *EthereumTransaction) FromTransaction(block *types.Block, tx *types.Transaction) {
	t.BlockHash = block.Hash().Hex()
	t.BlockNumber = block.Header().Number.Uint64()

	t.Type = tx.Type()
	t.Hash = tx.Hash().Hex()
	t.Nonce = tx.Nonce()

	t.ChainID = tx.ChainId()
	t.AccessList = tx.AccessList()
	t.Data = common.Bytes2Hex(tx.Data())
	t.Gas = tx.Gas()
	t.GasPrice = tx.GasPrice()
	t.GasTipCap = tx.GasTipCap()
	t.GasFeeCap = tx.GasFeeCap()

	if tx.To() != nil {
		t.To = tx.To().Hex()
	} else {
		t.To = EthereumEmptyAddress
	}
	t.From = PraseEthereumTransactionFromAddress(tx)

	t.Value = tx.Value()
	t.Time = tx.Time()

}
