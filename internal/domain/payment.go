package domain

import "time"

type Payment struct {
	PaymentID     string    `json:"payment_id"`
	RoomID        string    `json:"room_id"`
	RoomNumber    int       `json:"room_number"`    // Included so the owner can easily see which room paid on their dashboard dashboard
	AmountPaid    float64   `json:"amount_paid"`    // Total dollar/VND amount collected
	BillingMonth  time.Time `json:"billing_month"`  // e.g., May 2026 2026
	PaidAt        time.Time `json:"paid_at"`        // The exact date/time the money arrived
	PaymentMethod string    `json:"payment_method"` // e.g., "Bank Transfer", "Cash", "Momo"
	Status        string    `json:"status"`         // e.g., "Completed", "Pending", "Failed"
}
