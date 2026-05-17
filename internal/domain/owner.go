package domain

import "context"

// Owner defines the data structure for a property landlord
type Owner struct {
	ID          string `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

// OwnerRepository defines the behavioral rules for handling owner data
type OwnerRepository interface {
	Create(ctx context.Context, owner *Owner) error
	GetByID(ctx context.Context, id string) (*Owner, error)
}
