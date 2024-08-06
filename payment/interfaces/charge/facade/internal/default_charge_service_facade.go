package internal

import (
	"github.com/chainpusher/chainpusher/payment/application"
	"github.com/chainpusher/chainpusher/payment/application/impl"
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	gorm2 "github.com/chainpusher/chainpusher/payment/infrastructure/gorm"
	"github.com/chainpusher/chainpusher/payment/interfaces/charge/facade"
	"github.com/chainpusher/chainpusher/payment/interfaces/charge/facade/dto"
	"github.com/chainpusher/chainpusher/payment/interfaces/charge/facade/internal/assembler"
	"gorm.io/gorm"
)

type DefaultChargeServiceFacade struct {
	accountRepository account.Repository
	chargeRepository  charge.Repository
	assembler         *assembler.ChargeDTOAssembler
	service           application.ChargeService
}

func (f *DefaultChargeServiceFacade) Charge(cmd *dto.CreateChargeDTO) (*dto.ChargeDTO, error) {
	a, err := f.accountRepository.FindBySecretKey(cmd.SecretKey)
	if err != nil {
		return nil, err
	}
	c, err := f.assembler.ToCharge(a, cmd)
	if err != nil {
		return nil, err
	}

	c, err = f.service.Charge(a, c)
	if err != nil {
		return nil, err
	}

	return f.assembler.ToDTO(c), nil

}

func (f *DefaultChargeServiceFacade) Charged(id int64) {
	c, err := f.chargeRepository.Find(id)
	if err != nil {

	}
	if err := f.service.Charged(c); err != nil {

	}
}

//func NewChargeServiceFacade(
//	accountRepository account.Repository,
//	assembler *assembler.ChargeDTOAssembler,
//	service application.ChargeService) facade.ChargeServiceFacade {
//
//	return &DefaultChargeServiceFacade{
//		accountRepository: accountRepository,
//		assembler:  assembler,
//		service:    service,
//	}
//}

func NewChargeServiceFacade(db *gorm.DB) facade.ChargeServiceFacade {
	r := gorm2.NewAccountRepository(db)
	chargeRepository := gorm2.NewChargeRepository(db)
	return &DefaultChargeServiceFacade{
		accountRepository: r,
		chargeRepository:  chargeRepository,
		assembler:         &assembler.ChargeDTOAssembler{},
		service:           impl.NewChargeService(gorm2.NewChargeRepository(db)),
	}
}
