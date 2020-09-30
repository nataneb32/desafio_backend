package checkins

import (
	"time"
)

type CheckIn struct {
	ID               uint
	Hospede          uint
	DataEntrada      time.Time
	DataSaida        time.Time
	AdicionalVeiculo bool
}
