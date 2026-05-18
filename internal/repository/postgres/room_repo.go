package postgres

import (
	"Renting/internal/domain"
	"context"
	"database/sql"
	"errors"
)

type PostgresRoomRepository struct {
	db *sql.DB
}

func NewPostgresRoomRepository(db *sql.DB) *PostgresRoomRepository {
	return &PostgresRoomRepository{db: db}
}

func (r *PostgresRoomRepository) Create(ctx context.Context, room *domain.Room) error {
	query := `
		INSERT INTO Rooms (ID, RoomOwnerID, RoomNumber, Status, BaseRent)
		VALUES ($1, $2, $3, $4, $5);
	`
	_, err := r.db.ExecContext(ctx, query,
		room.RoomID,
		room.RoomOwnerID,
		room.RoomNumber,
		room.Status,
		room.BaseRent,
	)

	return err
}

func (r *PostgresRoomRepository) GetByID(ctx context.Context, roomID string) (*domain.Room, error) {
	query := `
		SELECT ID, RoomOwnerID, RoomNumber, Status, BaseRent
		FROM Rooms
		WHERE ID = $1;
	`

	var room domain.Room

	err := r.db.QueryRowContext(ctx, query, roomID).Scan(
		&room.RoomID,
		&room.RoomOwnerID,
		&room.RoomNumber,
		&room.Status,
		&room.BaseRent,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}

	return &room, nil
}

func (r *PostgresRoomRepository) Update(ctx context.Context, room *domain.Room) error {
	query := `
		UPDATE Rooms
		SET RoomNumber = $2, Status = $3, BaseRent = $4
		WHERE ID = $1;
	`

	_, err := r.db.ExecContext(ctx, query,
		room.RoomID,
		room.RoomNumber,
		room.Status,
		room.BaseRent,
	)

	return err
}
