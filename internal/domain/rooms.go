package domain

import "context"

type Room struct {
	RoomID      string  `json:"room_id"`
	RoomOwnerID string  `json:"room_owner_id"`
	RoomNumber  int     `json:"room_number"`
	Status      bool    `json:"status"`
	BaseRent    float64 `json:"base_rent"`
}

type CostType struct {
	CostTypeID string  `json:"cost_type_id"`
	OwnerID    string  `json:"owner_id"`
	RoomID     string  `json:"room_id"`
	CostName   string  `json:"cost_name"`
	ChargeType string  `json:"charge_type"`
	Cost       float32 `json: "cost"`
}

type RoomPricing struct {
	PricingID  string  `json:"pricing_id"`
	RoomID     string  `json:"room_id"`
	CostTypeID string  `json:"cost_type_id"`
	CostValue  float64 `json:"cost_value"`
}

type RoomDetail struct {
	Room
	Pricings []RoomPricing `json:"pricings"`
}

type RoomRepository interface {
	Create(ctx context.Context, room *Room) error
	GetByID(ctx context.Context, id string) (*Room, error)
	UpdateByID(ctx context.Context, room *Room) error
	DeleteByID(ctx context.Context, id string) error
}
