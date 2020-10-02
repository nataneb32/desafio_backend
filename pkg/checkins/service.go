package checkins

import (
	"errors"
	"nataneb32.live/hospedagem/pkg/billing"
	"time"
)

type CheckInService struct {
	CheckInRepo    CheckInRepo
	BillingService *billing.BillingService
}

// Create a checkin where data_entrada is null and data_saida is now.
func (cs *CheckInService) DoCheckIn(guestId uint, adicionalVeiculo bool) error {
	now := time.Now()
	checkin := CheckIn{
		Hospede:          guestId,
		AdicionalVeiculo: adicionalVeiculo,
		DataEntrada:      &now,
	}
	return cs.CheckInRepo.CreateCheckIn(&checkin)
}

// Update a checkin to data_saida = now.
func (cs *CheckInService) DoCheckOut(checkinID uint, adicionalVeiculo bool) error {
	now := time.Now()
	checkin := CheckIn{
		DataSaida: &now,
	}

	return cs.CheckInRepo.UpdateCheckIn(checkinID, &checkin)
}

func (cs *CheckInService) CalculateBill(checkinID uint) (error, uint) {
	var checkin CheckIn
	checkin.ID = checkinID
	err := cs.CheckInRepo.GetCheckIn(&checkin)
	if err != nil {
		return err, 0
	}

	if checkin.DataEntrada == nil {
		return errors.New("Cannot calculate the bill of a checin when the data_entrada is null."), 0
	}

	if checkin.DataSaida == nil {
		return errors.New("Cannot calculate the bill of a checkin when the data_saida is null. Please do the checkout."), 0
	}

	bill := cs.BillingService.CalculateBillOf(*checkin.DataEntrada, *checkin.DataSaida, checkin.AdicionalVeiculo)
	return nil, bill
}

func (cs *CheckInService) SumBillOf(checkins []CheckIn) uint {
	bill := uint(0)
	for _, checkin := range checkins {
		if checkin.DataEntrada == nil {
			continue
		}
		if checkin.DataSaida == nil {
			continue
		}

		bill += cs.BillingService.CalculateBillOf(*checkin.DataEntrada, *checkin.DataSaida, checkin.AdicionalVeiculo)
	}
	return bill
}

func (cs *CheckInService) NewestBillOf(checkins []CheckIn) uint {
	var newest CheckIn
	for _, checkin := range checkins {
		if checkin.DataEntrada == nil {
			continue
		}
		if checkin.DataSaida == nil {
			continue
		}
		if newest.DataEntrada == nil {
			newest = checkin
			continue
		}
		if newest.DataEntrada.Before(*checkin.DataEntrada) {
			newest = checkin
		}
	}
	if newest.DataEntrada == nil || newest.DataSaida == nil {
		return 0
	}
	return cs.BillingService.CalculateBillOf(*newest.DataEntrada, *newest.DataSaida, newest.AdicionalVeiculo)
}

func (cs *CheckInService) CalculateTotalSpendBy(guestId uint) (error, uint) {
	err, checkins := cs.CheckInRepo.GetAllCheckIn(&CheckIn{Hospede: guestId})
	if err != nil {
		return err, 0
	}
	bill := uint(0)
	for _, checkin := range checkins {
		if checkin.DataEntrada == nil {
			continue
		}
		if checkin.DataSaida == nil {
			continue
		}

		bill += cs.BillingService.CalculateBillOf(*checkin.DataEntrada, *checkin.DataSaida, checkin.AdicionalVeiculo)
	}
	return err, bill
}

func (cs *CheckInService) CalculateLastBillBy(guestId uint) (error, uint) {
	err, checkin := cs.CheckInRepo.GetNewestCheckInOf(guestId)
	if checkin.DataEntrada == nil {
		return nil, 0
	}
	if checkin.DataSaida == nil {
		return nil, 0
	}
	bill := cs.BillingService.CalculateBillOf(*checkin.DataEntrada, *checkin.DataSaida, checkin.AdicionalVeiculo)
	return err, bill
}

func CreateCheckInService(cr CheckInRepo, b *billing.BillingService) *CheckInService {
	return &CheckInService{
		CheckInRepo:    cr,
		BillingService: b,
	}
}
