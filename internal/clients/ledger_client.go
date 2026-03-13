package clients

type LedgerClient interface {
	CheckBalance(accountID string, amount float64) (bool, error)
}