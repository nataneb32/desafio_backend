package guests

type GuestQuery struct {
	Documento string `form:"documento"`
	Nome      string `form:"nome"`
	Telefone  string `form:"telefone"`
	Limit     uint   `form:"limit"`
	Page      uint   `form:"page"`
}
