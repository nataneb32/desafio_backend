package guests

type Guest struct {
	ID        uint   `json:"id"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Documento string `json:"documento"`
}
