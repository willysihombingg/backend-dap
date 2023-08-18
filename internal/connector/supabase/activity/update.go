package activity

import (
	"fmt"
	"net/http"
)

func (c *client) UpdateByID(where int, body interface{}) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:     fmt.Sprintf("/activity?id=eq.%v", where),
		Method:  http.MethodPatch,
		Payload: body,
	}

	return &requestOption
}
