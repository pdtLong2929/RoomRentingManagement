package domain

type Users struct {
	ID            string `json: "id"`
	Room_Owner_Id string `json: "room_owner_id"`
	Room_ID       string `json: "room_id"`
	FullName      string `json: "full_name"`
	PhoneNumber   string `json: "phone_number"`
	Email         string `json: "email"`
}
