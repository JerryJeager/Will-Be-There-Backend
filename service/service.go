package service

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type User struct {
	BaseModel
	Email     string `json:"email" binding:"required" gorm:"unique"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Password  string `json:"password" binding:"required"`
	IsToWed   bool   `json:"is_to_wed"`
}

type Event struct {
	BaseModel
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	UserID      uuid.UUID  `json:"user_id" binding:"required"`
	Country     string     `json:"country"`
	State       string     `json:"state"`
	Date        *time.Time `json:"date" binding:"required"`
	Venue       string     `json:"venue" binding:"required"`
	ImageUrl	string	   `json:"image_url"`
}

type Invitee struct {
	BaseModel
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" binding:"required" gorm:"unique"`
	Status    Status    `json:"status" binding:"required"`
	PlusOnes  *PlusOnes `json:"plus_ones"`
	EventID   uuid.UUID `json:"event_id" binding:"required"`
}

type Status string

type PlusOne struct {
	Name string
}

type PlusOnes []PlusOne

func (c PlusOnes) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *PlusOnes) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		// Unmarshal JSON data into the Values
		return json.Unmarshal(v, c)
	default:
		return fmt.Errorf("unsupported type for Values: %T", v)
	}
}

func (o *Invitee) MarshalJSON() ([]byte, error) {

	invitee := map[string]interface{}{
		"id":         o.ID,
		"first_name": o.FirstName,
		"last_name":  o.LastName,
		"email":      o.Email,
		"status":     o.Status,
		"event_id":   o.EventID,
		"plus_ones":  o.PlusOnes,
	}

	return json.Marshal(invitee)
}
