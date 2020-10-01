package billing

import (
	"nataneb32.live/hospedagem/pkg/checkins"
)

type CheckInRepo interface {
	//  GetCheckIn(*CheckIn) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	//	GetAllGuest(Guest) (error, Guest)
	CreateCheckIn(*checkins.CheckIn) error
	//	DeleteGuest(Guest) error

}
