package postgres

import (
	"context"
	"database/sql"
	"errors"

	// This relative import path bypasses any go.mod capitalization mismatches!
	"Renting/internal/domain"
)

// PostgresOwnerRepository implements domain.OwnerRepository
type PostgresOwnerRepository struct {
	db *sql.DB
}

// NewPostgresOwnerRepository instantiates your repository worker
func NewPostgresOwnerRepository(db *sql.DB) *PostgresOwnerRepository {
	return &PostgresOwnerRepository{db: db}
}

// Create inserts a brand new landlord into your database table
func (r *PostgresOwnerRepository) Create(ctx context.Context, owner *domain.Owner) error {
	query := `
		INSERT INTO Owner (ID, FullName, Email, PhoneNumber)
		VALUES ($1, $2, $3, $4);
	`
	_, err := r.db.ExecContext(ctx, query, owner.ID, owner.FullName, owner.Email, owner.PhoneNumber)
	return err
}

// GetByID looks up an owner profile by their unique ID
func (r *PostgresOwnerRepository) GetByID(ctx context.Context, id string) (*domain.Owner, error) {
	query := `
		SELECT ID, FullName, Email, PhoneNumber 
		FROM Owner 
		WHERE ID = $1;
	`

	var owner domain.Owner

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&owner.ID,
		&owner.FullName,
		&owner.Email,
		&owner.PhoneNumber,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("owner profile not found")
		}
		return nil, err
	}

	return &owner, nil
}
