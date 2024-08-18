package factory

import (
	"github.com/chainpusher/chainpusher/trade/application"
	impl2 "github.com/chainpusher/chainpusher/trade/application/impl"
	"github.com/chainpusher/chainpusher/trade/domain/model/price"
	gorm2 "github.com/chainpusher/chainpusher/trade/infrastructure/persistence/gorm"
	"github.com/chainpusher/chainpusher/trade/interfaces/price/facade"
	"github.com/chainpusher/chainpusher/trade/interfaces/price/facade/impl"
	"gorm.io/gorm"
)

type Factory struct {
	db *gorm.DB
}

func (factory *Factory) NewPriceRepository() price.Repository {
	return gorm2.NewPriceRepository(factory.db)
}

func (factory *Factory) CreatePriceService() application.PriceService {
	return impl2.NewPriceService(factory.NewPriceRepository())
}

func (factory *Factory) NewPriceServiceFacade() facade.ServiceFacade {
	return impl.NewPriceServiceFacade(factory.CreatePriceService())
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{db: db}
}
