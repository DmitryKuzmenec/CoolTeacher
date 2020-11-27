package model

import (
	"encoding/json"
	"strings"
	"time"
)

type UserData struct {
	ID          int       `json:"id"`
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	Gender      string    `json:"gender"`
	Email       string    `json:"email,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Description string    `json:"description,omitempty"`
	BirthDate   BirthDate `json:"birthDate"`
}

type BirthDate time.Time

func (d *BirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("01/02/2006", s)
	if err != nil {
		return err
	}
	*d = BirthDate(t)
	return nil
}

func (b *BirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}
