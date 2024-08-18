package facade

type PriceServiceFacadeFactory interface {
	NewPriceServiceFacade() ServiceFacade
}
