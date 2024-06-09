package monitor_test

import (
	"sync"
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/chain"
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/monitor"
	"github.com/chainpusher/chainpusher/sys"
	"github.com/sirupsen/logrus"
)

func TestEthereumWatcher(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	infuraKey, err := sys.GetEnv("INFURA_KEY")
	if err != nil {
		t.Log("Failed to get Infura key: ", err)
		return
	}

	cfg := config.Config{
		Wallets:   []string{},
		InfuraKey: chain.GetInfuraApiUrlV2(infuraKey),
	}
	service, err := chain.NewEthereumBlockChainService(chain.GetInfuraApiUrlV2(infuraKey))
	if err != nil {
		t.Fatal("Failed to create Ethereum block chain service: ", err)
	}
	application := application.NewTransactionService(&cfg)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	p := monitor.NewPlatformWatcherEthereum(15*time.Second, &waitGroup, service, application)
	go p.Start()
	time.Sleep(10 * time.Second)
	p.Stop()
	waitGroup.Wait()
}
