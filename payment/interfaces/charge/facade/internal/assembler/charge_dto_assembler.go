package assembler

import (
	"github.com/chainpusher/chainpusher/payment/domain/model/account"
	"github.com/chainpusher/chainpusher/payment/domain/model/charge"
	"github.com/chainpusher/chainpusher/payment/interfaces/charge/facade/dto"
)

type ChargeDTOAssembler struct {
}

func (assembler *ChargeDTOAssembler) ToCharge(a *account.Account, dto *dto.CreateChargeDTO) (*charge.Charge, error) {
	var c charge.Charge

	c = charge.Charge{
		AccountId: a.ID,
		Amount:    dto.Amount,
	}

	return &c, nil
}

func (assembler *ChargeDTOAssembler) ToDTO(c *charge.Charge) *dto.ChargeDTO {
	return &dto.ChargeDTO{
		Id: c.ID,
	}
}
