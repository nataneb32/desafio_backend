package checkins

import (
	"time"
)

type CheckIn struct {
	ID               uint      `json:"id"`
	Hospede          uint      `json:"hospede"`
	DataEntrada      time.Time `json:"dataEntrada"`
	DataSaida        time.Time `json:"dataSaida"`
	AdicionalVeiculo bool      `json:"adicionalVeiculo"`
}
