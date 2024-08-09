package charge

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/price"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Charge struct {
	gorm.Model
	ID             int64
	Amount         int64
	AccountId      int64
	Price          price.Price
	PriceId        int64
	Meta           datatypes.JSON `gorm:"type:json"`
	ValidityPeriod int
	Status         Status
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ExpiredAt      time.Time
	PaidAt         time.Time
}

func (charge *Charge) AssignWallets(wallets []Wallet) {

}
