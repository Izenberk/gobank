package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Izenberk/gobank/internal/domain"
)

type PostgresAccountRepository struct {
	db 	*sql.DB
}

var _ AccountRepository = (*PostgresAccountRepository) (nil)

func NewPostgresAccountRepository(db *sql.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) Create(ctx context.Context, acc *domain.Account) (int64, error) {
	query := `
		INSERT INTO accounts (fullname, balance, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int64

	err := r.db.QueryRowContext(ctx, query,
	acc.Fullname, acc.Balance, acc.CreatedAt).Scan(&id)
	return id, err
}

func (r *PostgresAccountRepository) GetByID(ctx context.Context, id int64) (*domain.Account, error) {
	query := `
		SELECT id, fullname, balance, created_at
		FROM accounts WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	acc := &domain.Account{}
	err := row.Scan(&acc.ID, &acc.Fullname, &acc.Balance, &acc.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return acc, nil
}

func (r *PostgresAccountRepository) UpdateByID(ctx context.Context, acc *domain.Account) error {
	query := `
		UPDATE accounts SET fullname=$1, balance=$2, created_at=$3 WHERE id=$4
	`

	_, err := r.db.ExecContext(ctx, query, acc.Fullname, acc.Balance, acc.CreatedAt, acc.ID)
	return err
}