package guests

type GuestRepo interface {
	GetGuest(*Guest) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	//	GetAllGuest(Guest) (error, Guest)
	CreateGuest(*Guest) error
	//	DeleteGuest(Guest) error
	SearchGuest(query struct {
		Documento string
		Nome      string
		Limit     uint
		Page      uint
	}) struct {
		Guests     []Guest
		TotalPages uint
	}
}
