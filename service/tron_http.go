package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type GetContractCommand struct {
	Value   string `json:"value"`
	Visible bool   `json:"visible"`
}

type GetContractAbiResponse struct {
	Entrys []interface{} `json:"entrys"`
}

type GetContractResponse struct {
	Bytecode string                 `json:"bytecode"`
	Abi      GetContractAbiResponse `json:"abi"`
}

type TronHttpClient struct {
}

func NewTronHttpClient() *TronHttpClient {
	return &TronHttpClient{}
}

func (c *TronHttpClient) GetContract(id string) (*GetContractResponse, error) {
	cmd := &GetContractCommand{Value: id, Visible: true}
	request, err := json.Marshal(cmd)
	if err != nil {
		return nil, err
	}
	response, err := http.Post(
		"https://api.trongrid.io/wallet/getcontract",
		"application/json",
		bytes.NewReader(request),
	)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var contract GetContractResponse

	err = json.Unmarshal(body, &contract)
	if err != nil {
		return nil, err
	}

	return &contract, nil
}
