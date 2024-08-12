package service

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/price"
)

type ChargeService struct {
	factory *price.Factory
}

func (svc *ChargeService) Charge(a *account.Account, c *charge.Charge) (*charge.Charge, error) {
	var p *price.Price
	var err error
	if p, err = svc.factory.AskForPrice(c.Amount); err != nil {
		return nil, err
	}

	a.PickWallets()

	c.AssignDefaultValidityPeriod()
	c.Price = *p

	return c, nil
}

func NewChargeService(factory *price.Factory) *ChargeService {
	return &ChargeService{
		factory: factory,
	}
}
