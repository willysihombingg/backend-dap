package activity

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseGetActivity struct {
	ID        uint64    `json:"id"`
	UserId    uint      `json:"user_id"`
	Activity  string    `json:"activity"`
	Time      uint      `json:"time"`
	Verified  bool      `json:"verified"`
	Date      time.Time `json:"date"`
	CreatedAt string    `json:"created_at"`
}

func (c *client) GetAllByDate(eq time.Time, user_id int) *RequestOptionsSupabase {
	date := fmt.Sprintf("%v-%v-%v", eq.Year(), eq.Month(), eq.Day())
	fmt.Println("Cek Date :", date)
	fmt.Println("User :", user_id)
	requestOption := RequestOptionsSupabase{
		URL:    fmt.Sprintf("/activity?and=(date.eq.%v,user_id.eq.%v)", date, user_id),
		Method: http.MethodGet,
	}
	return &requestOption
}
