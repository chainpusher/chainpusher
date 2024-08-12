package impl_test

import (
	"github.com/chainpusher/chainpusher/payment/application/impl"
	account2 "github.com/chainpusher/chainpusher/payment/domain/model/account"
	charge2 "github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/price"
	"github.com/chainpusher/chainpusher/payment/domain/model/test"
	"github.com/chainpusher/chainpusher/payment/domain/service"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
	"github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDefaultChargeService_Charge(t *testing.T) {
	db := test.SetupTestDB()
	accountRepository := gorm.NewAccountRepository(db)
	counterRepository := gorm.NewCounterRepository(db)
	priceRepository := gorm.NewPriceRepository(db)
	chargeRepository := gorm.NewChargeRepository(db)
	factory := price.NewFactory(counterRepository, priceRepository)
	domainService := service.NewChargeService(factory)
	applicationService := impl.NewChargeService(domainService, chargeRepository)

	var a *account2.Account
	var c *charge2.Charge
	var err error

	db.Preload("Secrets").First(&a)
	a, err = accountRepository.FindBySecretKey(a.Secrets[0].Key)

	c = &charge2.Charge{
		Amount: 100,
	}

	c, err = applicationService.Charge(a, c)

	assert.Nil(t, err)
	assert.NotNil(t, c)

	assert.Equal(t, int64(100), c.Amount)
	assert.Equal(t, int64(100), c.Price.Amount)
	assert.Equal(t, int((time.Hour * 24).Seconds()), c.ValidityPeriod)
	assert.Equal(t, charge2.Unpaid, c.Status)

	eth := shared.Of(c.Wallets...).Filter(func(w charge2.Wallet) bool {
		return w.Blockchain == "ETHEREUM"
	})[0]

	tron := shared.Of(c.Wallets...).Filter(func(w charge2.Wallet) bool {
		return w.Blockchain == "TRON"
	})[0]

	assert.Equal(t, "1", eth.Address)
	assert.Equal(t, "2", tron.Address)
}
