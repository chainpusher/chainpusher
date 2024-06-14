package service

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
)

var TronUsdtAddress string = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
var TronTriggerSmartyContract string = "type.googleapis.com/protocol.TriggerSmartContract"

// Converts an Ethereum address to a Tron address.
func ToTronAddress(addr common.Address) address.Address {
	var tronAddress = make([]byte, len(addr)+1)
	tronAddress[0] = 0x41
	copy(tronAddress[1:], addr.Bytes())

	return address.Address(tronAddress)
}

func GetHttpProxyDialOption() (*grpc.DialOption, error) {
	httpProxy := os.Getenv("HTTP_PROXY")

	if len(httpProxy) == 0 {
		return nil, nil
	}

	httpProxyURL, err := url.Parse(httpProxy)
	if err != nil {
		return nil, err
	}

	dialer := &net.Dialer{}
	diaCtx := func(ctx context.Context, network string, addr string) (net.Conn, error) {
		return dialer.DialContext(ctx, network, addr)

	}

	transport := &http.Transport{
		Proxy:       http.ProxyURL(httpProxyURL),
		DialContext: diaCtx,
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	c := func(ctx context.Context, addr string) (net.Conn, error) {
		return httpClient.Transport.(*http.Transport).DialContext(ctx, "tcp", addr)
	}

	d := grpc.WithContextDialer(c)
	return &d, err
}

func NewTronClient() (*client.GrpcClient, error) {
	client := client.NewGrpcClient("")

	options := []grpc.DialOption{grpc.WithInsecure()}

	proxyOption, err := GetHttpProxyDialOption()
	if err != nil {
		return nil, err
	}

	if proxyOption != nil {
		options = append(options, *proxyOption)
	}

	err = client.Start(options...)

	if err != nil {
		return nil, err
	}

	return client, nil
}
