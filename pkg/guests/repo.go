package guests

type GuestRepo interface {
	GetGuest(*Guest) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	//	GetAllGuest(Guest) (error, Guest)
	UpdateGuest(uint, *Guest) error
	CreateGuest(*Guest) error
	DeleteGuest(uint) error
	SearchGuest(query struct {
		Documento string
		Nome      string
		Limit     uint
		Page      uint
	}) struct {
		Guests     []Guest
		TotalPages uint
	}
	SearchInHotelGuest(query struct {
		Documento string
		Nome      string
		Limit     uint
		Page      uint
	}) struct {
		Guests     []Guest
		TotalPages uint
	}
}
