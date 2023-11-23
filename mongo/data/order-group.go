package myMongo

import (
	"time"

	"github.com/google/uuid"
)

type OrderGroup struct {
	UUID     uuid.UUID     `json:"order_group_uuid" bson:"order_group_uuid"`
	Space    *SpaceInfo    `json:"space" bson:"space"`
	Queue    *QueueInfo    `json:"queue" bson:"queue"`
	Delivery *DeliveryInfo `json:"delivery" bson:"delivery"`
	Tasks    []Task        `json:"tasks" bson:"tasks"`
}

type SpaceInfo struct {
	Number    int    `json:"space_number" bson:"space_number"`
	GroupCode string `json:"group_code" bson:"group_code"`
	ZoneName  string `json:"zone_name" bson:"zone_name"`
}

type DeliveryInfo struct {
	Number      int       `json:"number" bson:"number"`
	Partner     string    `json:"partner" bson:"partner"`
	Driver      string    `json:"driver" bson:"driver"`
	ScheduledAt time.Time `json:"scheduled_at" bson:"scheduled_at"`

	Customer Customer `json:"customer" bson:"customer"`
}

type QueueInfo struct {
	UUID     uuid.UUID `json:"space_uuid" bson:"space_uuid"`
	Number   int       `json:"number" bson:"number"`
	Customer Customer  `json:"customer" bson:"customer"`

	ScheduledAt time.Time `json:"scheduled_at" bson:"scheduled_at"`
}
