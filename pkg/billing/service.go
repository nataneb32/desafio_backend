package billing

import (
	"fmt"
	"nataneb32.live/hospedagem/pkg/checkins"
)

type BillingService struct {
	CheckInRepo CheckInRepo
}

func NewBillingService(cr CheckInRepo) *BillingService {
	return &BillingService{
		CheckInRepo: cr,
	}
}

//Returns the price of the checkin.
func (bs *BillingService) CalculateBillOf(checkinID uint) (error, uint) {
	var checkin checkins.CheckIn
	err := bs.CheckInRepo.GetCheckIn(&checkin)
	if err != nil {
		return err, 0
	}

	entrada := checkin.DataEntrada
	saida := checkin.DataSaida

	fmt.Println(entrada)
	fmt.Println(saida)

	return nil, 0
}
