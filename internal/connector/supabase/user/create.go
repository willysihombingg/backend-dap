package user

import "net/http"

func (c *client) Insert(body interface{}) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:     "/user",
		Method:  http.MethodPost,
		Payload: body,
	}

	return &requestOption
}
