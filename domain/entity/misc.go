package domain

import (
	"github.com/google/uuid"
)

type SpaceInfo struct {
	UUID      uuid.UUID `json:"space_uuid" bson:"space_uuid"`
	Number    int       `json:"space_number" bson:"space_number"`
	ZoneName  string    `json:"zone_name" bson:"zone_name"`
	GroupCode string    `json:"group_code" bson:"group_code"`
}

type DeliveryInfo struct {
	Number      int    `json:"number" bson:"number"`
	Partner     string `json:"partner" bson:"partner"`
	Driver      string `json:"driver" bson:"driver"`
	ScheduledAt int64  `json:"scheduled_at" bson:"scheduled_at"`

	Customer Customer `json:"customer" bson:"customer"`
}

type QueueInfo struct {
	UUID     uuid.UUID `json:"space_uuid" bson:"space_uuid"`
	Number   int       `json:"number" bson:"number"`
	Customer Customer  `json:"customer" bson:"customer"`

	ScheduledAt int64 `json:"scheduled_at" bson:"scheduled_at"`
}

type Customer struct {
	Name    string `json:"name" bson:"name"`
	Phone   Phone  `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
}

type Phone struct {
	CountryCode string `json:"country_code" bson:"country_code"`
	Number      string `json:"number" bson:"number"`
}
