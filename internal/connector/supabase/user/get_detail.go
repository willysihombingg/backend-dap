package user

import (
	"fmt"
	"net/http"
)

func (c *client) GetDetailByID(value int) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:    fmt.Sprintf("/user?id=eq.%v", value),
		Method: http.MethodGet,
	}
	return &requestOption
}
