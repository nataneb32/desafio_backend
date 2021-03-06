package checkin

import (
	"encoding/json"
	"fmt"
	"time"
)

type CheckIn struct {
	ID               uint       `json:"id" gorm:"primaryKey"`
	Hospede          uint       `json:"hospede"`
	DataEntrada      *time.Time `json:"dataEntrada"`
	DataSaida        *time.Time `json:"dataSaida"`
	AdicionalVeiculo bool       `json:"adicionalVeiculo"`
}

func (ci *CheckIn) UnmarshalJSON(j []byte) error {
	var raw struct {
		ID               uint    `json:"id"`
		Hospede          uint    `json:"hospede"`
		DataEntrada      *string `json:"dataEntrada"`
		DataSaida        *string `json:"dataSaida"`
		AdicionalVeiculo bool    `json:"adicionalVeiculo"`
	}

	err := json.Unmarshal(j, &raw)
	if err != nil {
		fmt.Println("asd")
		fmt.Println(err)
		return err
	}

	ci.ID = raw.ID

	ci.Hospede = raw.Hospede

	if raw.DataEntrada != nil {
		t, err := time.Parse(time.RFC3339, *raw.DataEntrada)
		if err != nil {
			return err
		}

		ci.DataEntrada = &t
	}

	if raw.DataSaida != nil {
		t, err := time.Parse(time.RFC3339, *raw.DataSaida)

		if err != nil {
			return err
		}
		ci.DataSaida = &t
	}
	ci.AdicionalVeiculo = raw.AdicionalVeiculo

	return nil
}
