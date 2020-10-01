package billing

import (
	"fmt"
	"nataneb32.live/hospedagem/pkg/checkins"
	"time"
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
	bill := uint(0)
	var checkin checkins.CheckIn
	checkin.ID = checkinID
	err := bs.CheckInRepo.GetCheckIn(&checkin)
	if err != nil {
		return err, 0
	}

	entrada := checkin.DataEntrada
	saida := checkin.DataSaida

	// extra dairy if saida as after 16:30hrs
	// saida after 16:30hrs is the same as saida + 1 day
	if saida.After(time.Date(saida.Year(), saida.Month(), saida.Day(), 16, 30, 0, 0, time.UTC)) {
		saida = time.Date(saida.Year(), saida.Month(), saida.Day(), 0, 0, 0, 0, time.UTC).Add(24 * time.Hour)
	}

	nWeekends, nWeekdays := countWeekendsAndWeekdayBetween(entrada, saida)

	bill += nWeekends*15000 + nWeekdays*12000

	// charging extra for parking
	if checkin.AdicionalVeiculo {
		bill += nWeekends*2000 + nWeekdays*1500
	}

	return nil, bill

}

func countWeekendsAndWeekdayBetween(in time.Time, out time.Time) (uint, uint) {
	nWeekends := uint(0)
	nWeekdays := uint(0)
	din := time.Date(in.Year(), in.Month(), in.Day(), 0, 0, 0, 0, time.UTC)
	dout := time.Date(out.Year(), out.Month(), out.Day(), 0, 0, 0, 0, time.UTC)
	dn := int(dout.Sub(din).Hours() / 24)

	for i := int(0); i < dn; i++ {
		d := din.Add(time.Duration(i*24) * time.Hour)
		if d.Weekday() == time.Sunday {
			nWeekends++
			continue
		}
		if d.Weekday() == time.Saturday {
			nWeekends++
			continue
		}
		nWeekdays++
	}

	fmt.Println(dn)
	return nWeekends, nWeekdays
}
