package presentations

import "time"

type CreateUser struct {
	// ID       int    `json: "id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Division string `json:"division"`
	Password string `json:"password"`
}

type UpdateUser struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
	Division string `json:"division,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token      string    `json:"token"`
	UserId     int       `json:"user_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
}
