package billing

import (
	"fmt"
	"nataneb32.live/hospedagem/pkg/checkin"
	"time"
)

type BillingService struct {
	ParkingFee        uint
	DairyFee          uint
	WeekendDairyFee   uint
	WeekendParkingFee uint
}

// Create a billing service. fee is in cents.
func NewBillingService(pf, df, wpf, wdf uint) *BillingService {
	return &BillingService{
		DairyFee:          df,
		WeekendDairyFee:   wdf,
		ParkingFee:        pf,
		WeekendParkingFee: wpf,
	}
}

//Returns the price in cents of the checkin.
func (bs *BillingService) CalculateBillOf(entrada, saida time.Time, adicionalVeiculo bool) uint {
	bill := uint(0)

	// extra dairy if saida as after 16:30hrs
	// saida after 16:30hrs is the same as saida + 1 day
	if saida.After(time.Date(saida.Year(), saida.Month(), saida.Day(), 16, 30, 0, 0, time.UTC)) {
		saida = time.Date(saida.Year(), saida.Month(), saida.Day(), 0, 0, 0, 0, time.UTC).Add(24 * time.Hour)
	}

	nWeekends, nWeekdays := countWeekendsAndWeekdayBetween(entrada, saida)

	bill += nWeekends*bs.WeekendDairyFee + nWeekdays*bs.DairyFee

	// charging extra for parking
	if adicionalVeiculo {
		bill += nWeekends*bs.WeekendParkingFee + nWeekdays*bs.ParkingFee
	}

	return bill
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
func (bs *BillingService) SumBillOf(checkins []checkin.CheckIn) uint {
	bill := uint(0)
	for _, checkin := range checkins {
		if checkin.DataEntrada == nil {
			continue
		}
		if checkin.DataSaida == nil {
			continue
		}

		bill += bs.CalculateBillOf(*checkin.DataEntrada, *checkin.DataSaida, checkin.AdicionalVeiculo)
	}
	return bill
}
