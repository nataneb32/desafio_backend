package checkins

type CheckInRepo interface {
	GetCheckIn(*CheckIn) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	GetAllCheckIn(*CheckIn) (error, []CheckIn)
	GetNewestCheckInOf(guestId uint) (error, CheckIn)
	CreateCheckIn(*CheckIn) error
	//	DeleteGuest(Guest) error
	UpdateCheckIn(uint, *CheckIn) error
}
