package domain

import (
	"context"
)

type Users struct {
	ID          string `json: "id"`
	RoomOwnerId string `json: "room_owner_id"`
	RoomID      string `json: "room_id"`
	FullName    string `json: "full_name"`
	PhoneNumber string `json: "phone_number"`
	Email       string `json: "email"`
}

type UsersRepository interface {
	Create(ctx *context.Context, user *Users) error
	GetByID(ctx *context.Context, id string) (*Users, error)
	UpdateByID(ctx *context.Context, user *Users) (*Users, error)
	Delete(ctx *context.Context, id string) error
}
