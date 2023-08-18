package activity

import (
	"fmt"
	"net/http"
	"time"
)

type ActivityByMonth struct {
	ID        uint64    `json:"id"`
	Date      time.Time `json:"dates"`
	Verified  bool      `json:"verified"`
	Highlight Highlight `json:"highlight"`
}

type Highlight struct {
	Color    string `json:"color"`
	Fillmode string `json:"fillMode"`
}

func (c *client) GetAllActivityByMonth(user_id string) *RequestOptionsSupabase {
	requestOption := RequestOptionsSupabase{
		URL:    fmt.Sprintf("/activity?user_id=eq.%v", user_id),
		Method: http.MethodGet,
	}
	return &requestOption
}
