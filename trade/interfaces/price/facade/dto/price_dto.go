package dto

type PriceDTO struct {
	BlockChain string  `json:"blockchain"`
	Crypto     string  `json:"crypto"`
	Price      float64 `json:"price"`
}
