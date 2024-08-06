package facade

import "github.com/chainpusher/chainpusher/payment/interfaces/charge/facade/dto"

type ChargeServiceFacade interface {
	Charge(cmd *dto.CreateChargeDTO) (*dto.ChargeDTO, error)

	Charged(id int64)
}
