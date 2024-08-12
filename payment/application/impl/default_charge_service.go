package impl

import (
	"github.com/chainpusher/chainpusher/payment/application"
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/domain/model/transaction"
	"github.com/chainpusher/chainpusher/payment/domain/service"
	"github.com/chainpusher/chainpusher/payment/domain/shared"
)

type DefaultChargeService struct {
	chargeService *service.ChargeService
	repository    charge.Repository
}

func (svc *DefaultChargeService) Charge(a *account.Account, c *charge.Charge) (*charge.Charge, error) {
	var err error
	if c, err = svc.chargeService.Charge(a, c); err != nil {
		return nil, err
	}
	if err := svc.repository.Save(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (svc *DefaultChargeService) Charged(transactions shared.Slice[*transaction.Transaction]) error {
	var exceptions = make(shared.Slice[error], 0)

	charges, err := svc.repository.FindChargingByTransactions(transactions)
	exceptions.Add(err)

	charges = charges.MatchingTransactionsHaveBeenPaid(transactions)
	err = svc.repository.Complete(charges)
	exceptions.Add(err)

	return nil
}

func NewChargeService(chargeService *service.ChargeService, repository charge.Repository) application.ChargeService {
	return &DefaultChargeService{
		chargeService: chargeService,
		repository:    repository,
	}
}
