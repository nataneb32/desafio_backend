package checkins

type CheckInRepo interface {
	// GetCheckIn(*CheckIn) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	//	GetAllGuest(Guest) (error, Guest)
	CreateCheckIn(*CheckIn) error
	//	DeleteGuest(Guest) error
}
