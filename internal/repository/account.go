package repository

import (
	"context"

	"github.com/Izenberk/gobank/internal/domain"
)



type AccountRepository interface {
	Create(ctx context.Context, acc *domain.Account) error
	GetByID(ctx context.Context, id int64) (*domain.Account, error)
	UpdateByID(ctx context.Context, acc *domain.Account) error
}