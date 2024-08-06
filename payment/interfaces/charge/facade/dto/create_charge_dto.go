package dto

type CreateChargeDTO struct {
	SecretKey      string                 `json:"secret_key"`
	Amount         int64                  `json:"amount"`
	Meta           map[string]interface{} `json:"meta"`
	ValidityPeriod int                    `json:"validity_period"`
}
