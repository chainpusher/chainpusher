package service

import (
	"math/big"

	"github.com/chainpusher/blockchain/model"
)

type DataSendingBlockListener struct {
	channel chan interface{}
}

func (l *DataSendingBlockListener) BeforeQuerying(height *big.Int) {

}

func (l *DataSendingBlockListener) AfterRawQuerying(block interface{}, err error) {
	if l.channel == nil {
		return
	}

	l.channel <- block
}

func (l *DataSendingBlockListener) AfterQuerying(blok *model.Block, err error) {

}

func NewDataSendingBlockListener(channel chan interface{}) *DataSendingBlockListener {
	return &DataSendingBlockListener{channel: channel}
}
