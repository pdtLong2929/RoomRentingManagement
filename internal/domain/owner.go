package domain

type owner struct {
	ID          string `json: "id"`
	FullName    string `json: "full_name"`
	Email       string `json: "email"`
	PhoneNumber string `json: "phone_number"`
}
