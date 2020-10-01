package guests

import (
	"nataneb32.live/hospedagem/pkg/checkins"
)

type Guest struct {
	ID        uint               `json:"id"`
	Nome      string             `json:"nome"`
	Telefone  string             `json:"telefone"`
	Documento string             `json:"documento"`
	CheckIns  []checkins.CheckIn `json:"checkins"`
}
