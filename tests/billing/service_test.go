package billing

import (
	"nataneb32.live/hospedagem/pkg/billing"
	"nataneb32.live/hospedagem/pkg/checkins"
	"testing"
	"time"
)

type CheckInRepoMock struct {
	Checkin checkins.CheckIn
}

func (c *CheckInRepoMock) GetCheckIn(checkin *checkins.CheckIn) error {
	*checkin = c.Checkin
	return nil
}

func TestBilling1(t *testing.T) {
	bs := billing.NewBillingService(&CheckInRepoMock{
		Checkin: checkins.CheckIn{
			ID:               1,
			AdicionalVeiculo: false,
			DataEntrada:      time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC),
			DataSaida:        time.Date(2020, 10, 21, 4, 0, 0, 0, time.UTC),
		},
	})
	err, bill := bs.CalculateBillOf(1)
	if bill != 12000 {
		t.Errorf("Expected 12000 got %d", bill)
	}
	if err != nil {
		t.Error(err)
	}
}
func TestBilling2(t *testing.T) {
	bs := billing.NewBillingService(&CheckInRepoMock{
		Checkin: checkins.CheckIn{
			ID:               1,
			AdicionalVeiculo: false,
			DataEntrada:      time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC),
			DataSaida:        time.Date(2020, 10, 22, 4, 0, 0, 0, time.UTC),
		},
	})
	err, bill := bs.CalculateBillOf(1)
	if bill != 24000 {
		t.Errorf("Expected 24000 got %d", bill)
	}
	if err != nil {
		t.Error(err)
	}
}
func TestBilling3(t *testing.T) {
	bs := billing.NewBillingService(&CheckInRepoMock{
		Checkin: checkins.CheckIn{
			ID:               1,
			AdicionalVeiculo: true,
			DataEntrada:      time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC),
			DataSaida:        time.Date(2020, 10, 21, 4, 0, 0, 0, time.UTC),
		},
	})
	err, bill := bs.CalculateBillOf(1)
	if bill != 13500 {
		t.Errorf("Expected 13500 got %d", bill)
	}
	if err != nil {
		t.Error(err)
	}
}
func TestBilling4(t *testing.T) {
	bs := billing.NewBillingService(&CheckInRepoMock{
		Checkin: checkins.CheckIn{
			ID:               1,
			AdicionalVeiculo: false,
			DataEntrada:      time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC),
			DataSaida:        time.Date(2020, 10, 21, 16, 31, 0, 0, time.UTC),
		},
	})
	err, bill := bs.CalculateBillOf(1)
	if bill != 24000 {
		t.Errorf("Expected 24000 got %d", bill)
	}
	if err != nil {
		t.Error(err)
	}
}
