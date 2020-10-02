package guests

import (
	"nataneb32.live/hospedagem/pkg/billing"
)

type GuestService struct {
	GuestRepo      GuestRepo
	BillingService *billing.BillingService
}

func CreateGuestService(GuestRepo GuestRepo, b *billing.BillingService) *GuestService {
	return &GuestService{
		GuestRepo:      GuestRepo,
		BillingService: b,
	}
}
