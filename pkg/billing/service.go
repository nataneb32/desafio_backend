package billing

type BillingService struct {
	CheckInRepo CheckInRepo
}

func NewBillingService(cr CheckInRepo) *BillingService {
	return &BillingService{
		CheckInRepo: cr,
	}
}
