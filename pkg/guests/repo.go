package guests

type GuestRepo interface {
	GetGuest(*Guest) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	//	GetAllGuest(Guest) (error, Guest)
	UpdateGuest(uint, *Guest) error
	CreateGuest(*Guest) error
	DeleteGuest(uint) error
	SearchGuest(query GuestQuery) struct {
		Guests     []Guest
		TotalPages uint
	}
	SearchInHotelGuest(query GuestQuery) struct {
		Guests     []Guest
		TotalPages uint
	}
}
