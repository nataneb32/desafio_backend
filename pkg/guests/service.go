package guests

import (
	"nataneb32.live/hospedagem/pkg/checkins"
)

type GuestService struct {
	GuestRepo      GuestRepo
	CheckInService *checkins.CheckInService
}

func CreateGuestService(GuestRepo GuestRepo, b *checkins.CheckInService) *GuestService {
	return &GuestService{
		GuestRepo:      GuestRepo,
		CheckInService: b,
	}
}
