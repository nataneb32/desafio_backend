package guests

import (
	"nataneb32.live/hospedagem/pkg/checkin"
)

type Guest struct {
	ID        uint              `json:"id" gorm:"primaryKey"`
	Nome      string            `json:"nome"`
	Telefone  string            `json:"telefone"`
	Documento string            `json:"documento"`
	CheckIns  []checkin.CheckIn `json:"checkins" gorm:"foreignKey:Hospede"` //TODO: delete on cascade
}
