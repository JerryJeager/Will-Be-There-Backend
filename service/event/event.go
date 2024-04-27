package event

import (
	"time"

	"github.com/JerryJeager/will-be-there-backend/service"
	"github.com/google/uuid"
)

type Event struct {
	service.BaseModel
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	UserID      uuid.UUID  `json:"user_id" binding:"required"`
	Country     string     `json:"country"`
	State       string     `json:"state"`
	Date        *time.Time `json:"date" binding:"required"`
	Venue       string     `json:"venue" binding:"required"`
	ImageUrl	string	   `json:"image_url"`
}

type Events []Event
