package guests

type GuestRepo interface {
	GetGuest(*Guest) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	//	GetAllGuest(Guest) (error, Guest)
	CreateGuest(*Guest) error
	//	DeleteGuest(Guest) error
}
