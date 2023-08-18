package user

import "context"

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
	GetAll() *RequestOptionsSupabase
	GetDetailByID(value int) *RequestOptionsSupabase
	GetByEmail(email string) *RequestOptionsSupabase
	UpdateByID(where int, body interface{}) *RequestOptionsSupabase
	Delete(value int) *RequestOptionsSupabase
}
