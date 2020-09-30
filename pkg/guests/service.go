package guests

type GuestService struct {
	GuestRepo GuestRepo
}

func CreateGuestService(GuestRepo GuestRepo) *GuestService {
	return &GuestService{
		GuestRepo: GuestRepo,
	}
}
