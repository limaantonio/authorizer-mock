package clients

type FraudResult struct {
	Score    int
	Approved bool
}

type FraudClient interface {
	CheckFraud(cardNumber string, amount float64) (FraudResult, error)
}