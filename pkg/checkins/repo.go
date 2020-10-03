package checkins

import (
	"nataneb32.live/hospedagem/pkg/checkin"
)

type CheckInRepo interface {
	GetCheckIn(*checkin.CheckIn) error
	//	ListGuests(Guest, uint, uint) (error, []Guest)
	GetAllCheckIn(*checkin.CheckIn) (error, []checkin.CheckIn)
	GetNewestCheckInOf(guestId uint) (error, checkin.CheckIn)
	CreateCheckIn(*checkin.CheckIn) error
	DeleteCheckIn(uint) error
	UpdateCheckIn(uint, *checkin.CheckIn) error
}
