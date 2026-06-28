package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/Izenberk/gobank/internal/domain"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestPostgresAccountRepository_Create(t *testing.T) {
	db := newTestDB(t)
	repo := NewPostgresAccountRepository(db)

	account := &domain.Account{
		Fullname: "John Doe",
		Balance: 	1000,
	}

	ctx := context.Background()

	ID, err := repo.Create(ctx, account)
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}

	acc, err := repo.GetByID(ctx, ID)
	if err != nil {
      t.Fatalf("GetByID() error: %v", err)
  }
	if acc.Fullname != account.Fullname {
		t.Errorf("Fullname = %v, want %v", acc.Fullname, account.Fullname)
	}
	if acc.Balance != account.Balance {
		t.Errorf("Balance = %v, want %v", acc.Balance, account.Balance)
	}
}

func TestPostgresAccountRepository_GetByID_notFound(t *testing.T) {
	db := newTestDB(t)
	repo := NewPostgresAccountRepository(db)

	ctx := context.Background()

	_, err := repo.GetByID(ctx, -1)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func newTestDB(t *testing.T) *sql.DB {
	connStr := "postgres://postgres:password@localhost:5435/gobank?sslmode=disable"
	ctx := context.Background()

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.PingContext(ctx); err != nil {
		t.Fatal(err)
	}

	return db
}

