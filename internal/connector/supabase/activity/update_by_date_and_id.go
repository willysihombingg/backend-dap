package activity

import (
	"fmt"
	"net/http"
	"time"
)

func (c *client) UpdateByDateAndId(eq time.Time, user_id int, body interface{}) *RequestOptionsSupabase {
	date := fmt.Sprintf("%v-%v-%v", eq.Year(), eq.Month(), eq.Day())
	requestOption := RequestOptionsSupabase{
		URL:     fmt.Sprintf("/activity?and=(date.eq.%v,user_id.eq.%v)", date, user_id),
		Method:  http.MethodPatch,
		Payload: body,
	}

	return &requestOption
}
