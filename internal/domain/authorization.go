package domain

type AuthorizationRequest struct {
	CardNumber string
	Amount     float64
	MCC        string
	MerchantID string
}

type AuthorizationResponse struct {
	Approved     bool
	ResponseCode string
}

type Authorization struct {
	ID     string
	Amount float64
	Status string
}

