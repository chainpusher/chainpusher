package charge

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/wallet"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Charge struct {
	gorm.Model
	ID             int64
	Amount         int64
	AccountId      int64
	Pool           wallet.Pool
	PoolId         int64
	Meta           datatypes.JSON `gorm:"type:json"`
	ValidityPeriod int
	Status         Status
	CreatedAt      time.Time
	UpdatedAt      time.Time
	PaidAt         time.Time
}
