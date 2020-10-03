package billing

import (
	"nataneb32.live/hospedagem/pkg/billing"
	"testing"
	"time"
)

func TestBilling(t *testing.T) {
	bs := billing.CreateBillingService(1500, 12000, 2000, 15000)
	dataEntrada := time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC)
	dataSaida := time.Date(2020, 10, 21, 4, 0, 0, 0, time.UTC)

	bill := bs.CalculateBillOf(dataEntrada, dataSaida, false)
	if bill != 12000 {
		t.Errorf("Expected 12000 got %d", bill)
	}
}

func TestBillingTwoDays(t *testing.T) {
	bs := billing.CreateBillingService(1500, 12000, 2000, 15000)
	dataEntrada := time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC)
	dataSaida := time.Date(2020, 10, 22, 4, 0, 0, 0, time.UTC)
	bill := bs.CalculateBillOf(dataEntrada, dataSaida, false)
	if bill != 24000 {
		t.Errorf("Expected 24000 got %d", bill)
	}

}

func TestBillingVeiculo(t *testing.T) {
	bs := billing.CreateBillingService(1500, 12000, 2000, 15000)
	dataEntrada := time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC)
	dataSaida := time.Date(2020, 10, 21, 4, 0, 0, 0, time.UTC)
	bill := bs.CalculateBillOf(dataEntrada, dataSaida, true)
	if bill != 13500 {
		t.Errorf("Expected 13500 got %d", bill)
	}
}
func TestBillingExtraDairy(t *testing.T) {
	bs := billing.CreateBillingService(1500, 12000, 2000, 15000)
	dataEntrada := time.Date(2020, 10, 20, 3, 0, 0, 0, time.UTC)
	dataSaida := time.Date(2020, 10, 21, 16, 31, 0, 0, time.UTC)
	bill := bs.CalculateBillOf(dataEntrada, dataSaida, false)
	if bill != 24000 {
		t.Errorf("Expected 24000 got %d", bill)
	}
}
