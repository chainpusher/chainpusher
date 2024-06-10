package chain

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/chainpusher/chainpusher/infrastructure"
	"github.com/chainpusher/chainpusher/model"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	tcAbi "github.com/fbsobreira/gotron-sdk/pkg/abi"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var TronUsdtAddress string = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
var TronTriggerSmartyContract string = "type.googleapis.com/protocol.TriggerSmartContract"

type TronBlockChainService struct {
	Client                *client.GrpcClient
	Listner               TransactionListener
	SmartContractAbi      *core.SmartContract_ABI
	UsdtTransferArguments *abi.Arguments
	Channel               chan interface{}
}

func (service *TronBlockChainService) GetNowBlock() (*api.BlockExtention, []*model.Transaction, error) {
	block, err := service.Client.GetNowBlock()
	if err != nil {
		return nil, nil, err
	}

	go func(block *api.BlockExtention) {
		if service.Channel == nil {
			return
		}

		service.Channel <- block
	}(block)

	return block, ToTransactions(service.UsdtTransferArguments, block), nil
}

func (service *TronBlockChainService) GetBlock(number int64) ([]*model.Transaction, error) {

	logrus.Debug("Fetching block ", number)
	for i := 0; i < 10; i++ {
		block, err := service.Client.GetBlockByNum(number)

		if err != nil {
			logrus.Warnf("Error getting block %d: %v", number, err)
		} else if block.BlockHeader != nil {

			go func(block *api.BlockExtention) {
				if service.Channel == nil {
					return
				}

				service.Channel <- block
			}(block)

			return ToTransactions(service.UsdtTransferArguments, block), nil
		} else {
			logrus.Errorf("Block not found: %v. block description: %v", number, block)
			time.Sleep(1 * time.Second)
			return nil, errors.New("block not found")
		}

		time.Sleep(1 * time.Second)
	}

	return nil, errors.New("error getting block on 10 attempts")
}

func GetUsdtSmartContract(client *client.GrpcClient) (*core.SmartContract_ABI, error) {
	var contractAbi *core.SmartContract_ABI
	cacheKey := fmt.Sprintf("proto_%s", TronUsdtAddress)
	contractAbiBytes, err := infrastructure.GetKey(cacheKey)

	if err != nil {
		log.Printf("Error getting contract ABI from cache: %v", err)
	} else {
		log.Printf("Contract ABI fetched from cache")
	}

	// TODO: Implement the following code
	// if err == nil {
	// 	err := proto.Unmarshal(contractAbiBytes, contractAbi)
	// 	if err != nil {
	// 		log.Fatalf("Error unmarshaling contract ABI: %v", err)
	// 	}
	// }

	contractAbi, err = client.GetContractABI(TronUsdtAddress)
	if err != nil {
		log.Fatalf("Error getting contract ABI: %v", err)
		return nil, err
	}

	contractAbiBytes, err = proto.Marshal(contractAbi)
	if err != nil {
		log.Fatalf("Error marshaling contract ABI: %v", err)
	}
	infrastructure.SetKey(cacheKey, contractAbiBytes)
	return contractAbi, err
}

func NewTronBlockChainService(
	listener TransactionListener,
	smartContractAbi *core.SmartContract_ABI,
	client *client.GrpcClient,
	channel chan interface{}) *TronBlockChainService {

	args, err := tcAbi.GetInputsParser(smartContractAbi, "transfer")

	if err != nil {
		log.Fatalf("Error getting inputs parser: %v", err)
		panic(err)
	}

	return &TronBlockChainService{
		Client:                client,
		Listner:               listener,
		SmartContractAbi:      smartContractAbi,
		UsdtTransferArguments: &args,
		Channel:               channel,
	}
}

// ToTronAddress converts an Ethereum address to a Tron address.
func ToTronAddress(addr common.Address) address.Address {
	var tronAddress = make([]byte, len(addr)+1)
	tronAddress[0] = 0x41
	copy(tronAddress[1:], addr.Bytes())

	return address.Address(tronAddress)
}
