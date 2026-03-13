package authorization

import (
	"github.com/limaantonio/auth/internal/clients"
	"github.com/limaantonio/auth/internal/domain"
	"github.com/limaantonio/auth/internal/messaging"
)

type AuthorizationProcessor struct {
	AccountClient clients.AccountClient
	FraudClient   clients.FraudClient
	LedgerClient  clients.LedgerClient
	Publisher     *messaging.ClearingPublisher
}

func (p *AuthorizationProcessor) Process(req domain.AuthorizationRequest) domain.AuthorizationResponse {

	account, _ := p.AccountClient.GetAccount(req.CardNumber)

	if account.Status != "ACTIVE" {
		return domain.AuthorizationResponse{
			Approved:     false,
			ResponseCode: "05",
		}
	}

	fraud, _ := p.FraudClient.CheckFraud(req.CardNumber, req.Amount)

	if !fraud.Approved {
		return domain.AuthorizationResponse{
			Approved:     false,
			ResponseCode: "59",
		}
	}

	ok, _ := p.LedgerClient.CheckBalance(account.ID, req.Amount)

	if !ok {
		return domain.AuthorizationResponse{
			Approved:     false,
			ResponseCode: "51",
		}
	}


	if ok {
		event := domain.ClearingEvent{
			AuthorizationID: "some-id",
			Amount:          req.Amount,
			CardNumber:      req.CardNumber,
			Status:          "APPROVED",
		}

		p.Publisher.Publish(event)
	}
		

	return domain.AuthorizationResponse{
		Approved:     true,
		ResponseCode: "00",
	}	
}