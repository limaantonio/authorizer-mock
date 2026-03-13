package clients

type Account struct {
	ID     string
	Status string
	Limit  float64
}

type AccountClient interface {
	GetAccount(cardNumber string) (Account, error)
}