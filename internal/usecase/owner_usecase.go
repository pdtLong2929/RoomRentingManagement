package usecase

import (
	"Renting/internal/domain"
	"context"
)

type OwnerUseCase interface {
	CreateRoom(ctx context.Context, room *domain.Room, pricings []domain.RoomPricing) error
	ListRooms(ctx context.Context, ownerID string) ([]*domain.Room, error)
	GetRoomDetail(ctx context.Context, ownerID string, roomID string) (*domain.RoomDetail, error)
	UpdateRoomCost(ctx context.Context, ownerID string, pricing *domain.RoomPricing) error
	DeleteRoom(ctx context.Context, ownerID string, roomID string) error
	GetPaymentHistory(ctx context.Context, ownerID string) ([]*domain.Payment, error)
}
