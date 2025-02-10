package data

import "time"

type JwtUser struct {
	Username string    `json:"username"`
	Expires  time.Time `json:"expires"`
}
