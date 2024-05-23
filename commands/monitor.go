package commands

import (
	"log"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"google.golang.org/grpc"
)

type MonitorCommand struct {
	Client      *client.GrpcClient
	BlockNumber int64
}

func (m *MonitorCommand) Execute() error {

	go m.MonitorLatest()
	time.Sleep(3 * time.Second)
	m.MonitorBlock()

	return nil
}

// Monitor the block by number.
//
// Fetch the block every 3 seconds
func (m *MonitorCommand) MonitorBlock() {
	for {
		log.Println("monitoring block number: ", m.BlockNumber)
		go func() {
			if block := m.FetchBlock(m.BlockNumber); block != nil {
				log.Println("new block: ", m.BlockNumber)
				m.BlockNumber = block.BlockHeader.RawData.Number + 1
			}
		}()
		time.Sleep(3 * time.Second)
	}
}

// Fetch the block by number.
func (m *MonitorCommand) FetchBlock(number int64) *api.BlockExtention {
	block, err := m.Client.GetBlockByNum(number)

	if err != nil {
		log.Fatalf("failed to fetch block: %v", err)
	}

	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("Recovered in f", r)
		}
	}()

	return block
}

// Fetch the latest block from the network.
func (m *MonitorCommand) FetchLatestBlock() {
	log.Println("fetching latest block...")
	block, err := m.Client.GetNowBlock()
	if err != nil {
		log.Fatalf("failed to fetch latest block: %v", err)
	}

	if m.BlockNumber == 0 {
		m.BlockNumber = block.BlockHeader.RawData.Number
		log.Println("set current block number to: ", block.BlockHeader.RawData.Number)
	}

	log.Println("now block number is: ", m.BlockNumber)

	// TODO: if currently fetch is too late to get the latest block, need to fix this issue later

}

// Monitor the latest block.
//
// Fetch the latest block every 3 seconds
func (m *MonitorCommand) MonitorLatest() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("Recovered in f", r)
		}
	}()

	for {
		go m.FetchLatestBlock()
		time.Sleep(3 * time.Second)
	}
}

func NewMonitorCommand() *MonitorCommand {
	client := client.NewGrpcClient("")
	err := client.Start(grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	return &MonitorCommand{
		Client:      client,
		BlockNumber: 0,
	}

}
