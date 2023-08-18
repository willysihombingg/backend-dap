package user

import (
	"fmt"
	"net/http"
)

func (c *client) GetByEmail(email string) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:    fmt.Sprintf("/user?email=eq.%s&select=*", email),
		Method: http.MethodGet,
	}
	return &requestOption
}
