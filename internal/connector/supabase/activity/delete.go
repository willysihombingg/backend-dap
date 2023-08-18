package activity

import (
	"fmt"
	"net/http"
)

func (c *client) Delete(value int) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:    fmt.Sprintf("/activity?id=eq.%v", value),
		Method: http.MethodDelete,
	}
	return &requestOption
}
