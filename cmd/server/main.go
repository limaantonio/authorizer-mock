package main

import (
	"fmt"

	"github.com/limaantonio/auth/internal/authorization"
	"github.com/limaantonio/auth/internal/clients"
	"github.com/limaantonio/auth/internal/messaging"
	"github.com/limaantonio/auth/internal/parser"
)

func main() {

	parser := parser.ISOParser{}
	req := parser.ParseISO([]byte("ISO_MESSAGE"))

	accountClient := &clients.MockAccountClient{}
	fraudClient := &clients.MockFraudClient{}
	ledgerClient := &clients.MockLedgerClient{}

	// cria publisher primeiro
	publisher, err := messaging.NewClearingPublisher(
		"amqp://admin:admin@localhost:5672/",
	)

	if err != nil {
		fmt.Println("Erro ao criar publisher:", err)
		return
	}

	// injeta publisher no processor
	processor := authorization.AuthorizationProcessor{
		AccountClient: accountClient,
		FraudClient:   fraudClient,
		LedgerClient:  ledgerClient,
		Publisher:     publisher,
	}

	service := authorization.AuthorizationService{
		Processor: &processor,
	}

	resp := service.Authorize(req)

	if resp.Approved {
		fmt.Println("Transação aprovada")
	}

	fmt.Println(resp)
}