package model

import "time"

type (
	User struct {
		Username string
		Password string
	}

	Session struct {
		Username string
		Expiry   time.Time
	}

	Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

//override session
func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
