package user

import (
	"fmt"
	"net/http"
)

func (c *client) Delete(value int) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:    fmt.Sprintf("/user?id=eq.%v", value),
		Method: http.MethodDelete,
	}
	return &requestOption
}
