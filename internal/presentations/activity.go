package presentations

import "time"

type Activity struct {
	UserId   int       `json:"user_id"`
	Activity string    `json:"activity"`
	Time     uint      `json:"time"`
	Date     time.Time `json:"date"`
}

type CreateActivity struct {
	UserId   int    `json:"user_id"`
	Activity string `json:"activity"`
	Time     uint   `json:"time"`
	Date     string `json:"date"`
}

type UpdateActivity struct {
	ID       int    `json:"id,omitempty"`
	UserId   int    `json:"user_id"`
	Activity string `json:"activity,omitempty"`
	Time     uint   `json:"time,omitempty"`
	Date     string `json:"date,omitempty"`
}

type VerifyActivity struct {
	UserId int    `json:"user_id"`
	Date   string `json:"date"`
}

type UpdateVerifyBody struct {
	Verified bool `json:"verified"`
}
