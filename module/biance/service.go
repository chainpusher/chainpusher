package biance

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var hosts = []string{
	"https://api.binance.com",
	"https://api-gcp.binance.com",
	"https://api1.binance.com",
	"https://api2.binance.com",
	"https://api3.binance.com",
	"https://api4.binance.com",
}

type Service struct {
}

func (s *Service) GetPrice(symbols ...string) ([]*Price, error) {
	host := hosts[0]

	s2 := make([]string, len(symbols))
	for i, v := range symbols {
		s2[i] = fmt.Sprintf("\"%s\"", v)
	}
	s3 := strings.Join(s2, ",")
	url := fmt.Sprintf("%s/api/v3/ticker/price?symbols=[%s]", host, s3)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	prices := make([]*Price, 0)
	err = json.Unmarshal(body, &prices)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func NewService() *Service {
	return &Service{}
}
