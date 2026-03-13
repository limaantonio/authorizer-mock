package repository

import "github.com/limaantonio/auth/internal/domain"

type AuthorizationRepository interface {
	Save(auth domain.Authorization) error
}