package postgres

import (
	"context"
	"database/sql"
	"errors"

	"Renting/internal/domain"
)

type PostgresOwnerRepository struct {
	db *sql.DB
}

func NewPostgresOwnerRepository(db *sql.DB) *PostgresOwnerRepository {
	return &PostgresOwnerRepository{db: db}
}

func (r *PostgresOwnerRepository) Create(ctx context.Context, owner *domain.Owner) error {
	query := `
		INSERT INTO Owner (ID, FullName, Email, PhoneNumber)
		VALUES ($1, $2, $3, $4);
	`
	_, err := r.db.ExecContext(ctx, query, owner.ID, owner.FullName, owner.Email, owner.PhoneNumber)
	return err
}

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
			return nil, domain.ErrOwnerNotFound
		}
		return nil, err
	}

	return &owner, nil
}

func (r *PostgresOwnerRepository) UpdateByID(ctx context.Context, owner *domain.Owner) (*domain.Owner, error) {
	query := `
		UPDATE Owner
		SET FullName = $2, PhoneNumber = $3, Email = $4
		WHERE  = $1
		RETURNING ID, Full_Name, Phone_Number, Email;
	`
	var ResponeOwner domain.Owner

	err := r.db.QueryRowContext(ctx, query, owner.ID, owner.FullName, owner.PhoneNumber, owner.Email).
		Scan(&ResponeOwner.ID, &ResponeOwner.FullName, &ResponeOwner.PhoneNumber, &ResponeOwner.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrOwnerNotFound
		}
	}
	return &ResponeOwner, nil
}

func (r *PostgresOwnerRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM Owner
		WHERE ID = $1
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
