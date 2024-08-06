package impl

import (
	"github.com/chainpusher/chainpusher/payment/application"
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"time"
)

type DefaultChargeService struct {
	repository charge.Repository
}

func (svc *DefaultChargeService) Charge(a *account.Account, c *charge.Charge) (*charge.Charge, error) {

	a.PickWallet(c)

	if err := svc.repository.Save(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (svc *DefaultChargeService) Charged(c *charge.Charge) error {
	c.Status = charge.PAID
	c.PaidAt = time.Now()

	if err := svc.repository.Save(c); err != nil {
		return err
	}
	return nil
}

func NewChargeService(repository charge.Repository) application.ChargeService {
	return &DefaultChargeService{
		repository: repository,
	}
}
