package user

import (
	"net/http"
	"time"
)

type ResponseGetUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Division  string    `json:"division"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *client) GetAll() *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:    "/user?select=*",
		Method: http.MethodGet,
	}
	return &requestOption
}
