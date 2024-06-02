package monitor

import (
	"github.com/chainpusher/chainpusher/chain"
	"github.com/chainpusher/chainpusher/model"
)

type Monitor interface {
	chain.TransactionListener

	StartPlatform(platform model.Platform)

	Start()

	Stop()
}
