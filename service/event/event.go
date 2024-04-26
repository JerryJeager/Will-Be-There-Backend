package event

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/JerryJeager/will-be-there-backend/service"
	"github.com/google/uuid"
)

type Event struct {
	service.BaseModel
	Description string     `json:"description" binding:"required"`
	EventTypes  EventTypes `json:"event_types"`
	UserID      uuid.UUID  `json:"user_id" binding:"required"`
	Country     string     `json:"country" binding:"required"`
	State       string     `json:"state" binding:"required"`
}

type Events []Event

type EventType struct {
	Name        string     `json:"name" binding:"required"`
	Venue       string     `json:"venue"`
	Description string     `json:"description"`
	Date        *time.Time `json:"date" binding:"required"`
}

type EventTypes map[string]EventType

func (c EventTypes) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *EventTypes) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		// Unmarshal JSON data into the Values
		return json.Unmarshal(v, c)
	default:
		return fmt.Errorf("unsupported type for Values: %T", v)
	}
}

func (o *Event) MarshalJSON() ([]byte, error) {

	Event := map[string]interface{}{
		"id":          o.ID,
		"description": o.Description,
		"user_id":     o.UserID,
		"country":     o.Country,
		"state":       o.State,
		"event_types": o.EventTypes,
	}

	return json.Marshal(Event)
}
