package clients

type MockAccountClient struct{}
type MockFraudClient struct{}
type MockLedgerClient struct{}

func (m *MockAccountClient) GetAccount(cardNumber string) (Account, error) {

	return Account{
		ID:     "acc123",
		Status: "ACTIVE",
		Limit:  1000,
	}, nil
}

func (m *MockLedgerClient) CheckBalance(accountID string, amount float64) (bool, error) {

	return true, nil
}

func (m *MockFraudClient) CheckFraud(cardNumber string, amount float64) (FraudResult, error) {

	return FraudResult{
		Score:    10,
		Approved: true,
	}, nil
}