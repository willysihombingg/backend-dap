package activity

import (
	"context"
	"time"
)

type Headers struct {
	ApiKey      string `json:"apikey"`
	BearerToken string `json:"Authorization"`
}

type RequestOptionsSupabase struct {
	Payload       interface{}
	URL           string
	Header        Headers
	Method        string
	TimeoutSecond int
	Context       context.Context
}

type Connectorer interface {
	Send(ctx context.Context, requestOption *RequestOptionsSupabase, response interface{}) error
	Insert(body interface{}) *RequestOptionsSupabase
	GetAllByDate(eq time.Time, user_id int) *RequestOptionsSupabase
	GetAllActivityByMonth(user_id string) *RequestOptionsSupabase
	GetDetailByID(value int) *RequestOptionsSupabase
	UpdateByID(where int, body interface{}) *RequestOptionsSupabase
	UpdateByDateAndId(eq time.Time, user_id int, body interface{}) *RequestOptionsSupabase
	Delete(value int) *RequestOptionsSupabase
}
