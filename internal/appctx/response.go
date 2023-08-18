// Package appctx
package appctx

import (
	"sync"

	"gitlab.com/willysihombing/task-c3/pkg/msg"
)

var rsp *Response
var oneRsp sync.Once

// Response presentation contract object
type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Entity  string      `json:"entity,omitempty"`
	State   string      `json:"state,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Lang    string      `json:"-"`
	Meta    interface{} `json:"meta,omitempty"`
}

// MetaData represent meta data response for multi data
type MetaData struct {
	Page       uint64 `json:"page"`
	Limit      uint64 `json:"limit"`
	TotalPage  uint64 `json:"total_page"`
	TotalCount uint64 `json:"total_count"`
}

// GetCode method to transform response name var to http status
func (r *Response) GetCode() int {
	return msg.GetCode(r.Code)
}

// WithStatus method to transform response
func (r *Response) WithStatus(s string) *Response {
	r.Status = s
	return r
}

// WithEntity method to transform response
func (r *Response) WithEntity(s string) *Response {
	r.Entity = s
	return r
}

// WithState method to transform response
func (r *Response) WithState(s string) *Response {
	r.State = s
	return r
}

// GetMessage method to transform response name var to message detail
func (r *Response) GetMessage() string {
	return msg.Get(r.Code, r.Lang)
}

// GenerateMessage setter message
func (r *Response) GenerateMessage() {
	r.Message = msg.Get(r.Code, r.Lang)
}

// WithCode setter response var name
func (r *Response) WithCode(c int) *Response {
	r.Code = c
	return r
}

// WithData setter data response
func (r *Response) WithData(v interface{}) *Response {
	r.Data = v
	return r
}

// WithError setter error messages
func (r *Response) WithError(v interface{}) *Response {
	r.Errors = v
	return r
}

// WithMeta setter meta data response
func (r *Response) WithMeta(v interface{}) *Response {
	r.Meta = v
	return r
}

// WithMessage setter custom message response
func (r *Response) WithMessage(v interface{}) *Response {
	if v != nil {
		r.Message = v
	}
	return r
}

// NewResponse initialize response
func NewResponse() *Response {
	oneRsp.Do(func() {
		rsp = &Response{}
	})

	// clone response
	x := *rsp

	return &x
}
