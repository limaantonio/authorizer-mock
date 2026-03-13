package parser

import "github.com/limaantonio/auth/internal/domain"

type ISOParser struct{}

func (p *ISOParser) ParseISO(message []byte) domain.AuthorizationRequest {
	// simulando parse ISO8583
	return domain.AuthorizationRequest{
		CardNumber: "12334566",
		Amount: 100.0,
		MCC: "5411",
		MerchantID: "Merchant123",
	}
}