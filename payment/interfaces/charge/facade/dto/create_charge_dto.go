package dto

type CreateChargeDTO struct {
	SecretKey      string                 `json:"secret_key"`
	WalletId       int64                  `json:"wallet_id"`
	Amount         int64                  `json:"amount"`
	Meta           map[string]interface{} `json:"meta"`
	ValidityPeriod int                    `json:"validity_period"`
}
