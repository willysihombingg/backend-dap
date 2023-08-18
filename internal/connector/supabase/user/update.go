package user

import (
	"fmt"
	"net/http"
)

func (c *client) UpdateByID(where int, body interface{}) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:     fmt.Sprintf("/user?id=eq.%v", where),
		Method:  http.MethodPatch,
		Payload: body,
	}

	return &requestOption
}
