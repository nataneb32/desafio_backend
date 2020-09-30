package checkins

type CheckInService struct {
	CheckInRepo CheckInRepo
}

func CreateCheckInService(cr CheckInRepo) *CheckInService {
	return &CheckInService{
		CheckInRepo: cr,
	}
}
