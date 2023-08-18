package activity

import "net/http"

func (c *client) Insert(body interface{}) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:     "/activity",
		Method:  http.MethodPost,
		Payload: body,
	}

	return &requestOption
}
