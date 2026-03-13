package authorization

import (
	"github.com/limaantonio/auth/internal/domain"
)

type AuthorizationService struct {
	Processor *AuthorizationProcessor
}

func (s *AuthorizationService) Authorize(req domain.AuthorizationRequest) domain.AuthorizationResponse {

	return s.Processor.Process(req)
}