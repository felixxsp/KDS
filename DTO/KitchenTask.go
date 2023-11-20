package dto

import (
	"time"

	"github.com/google/uuid"
)

type OrderModifier struct {
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}

type TaskOrder struct {
	UUID            uuid.UUID       `json:"uuid" bson:"uuid"`
	ItemName        string          `json:"item_name" bson:"item_name"`
	ItemVariantName string          `json:"item_variant_name" bson:"item_variant_name"`
	Quantity        int             `json:"quantity" bson:"quantity"`
	Notes           string          `json:"notes" bson:"notes"`
	Modifier        []OrderModifier `json:"modifiers" bson:"modifiers"`

	CreatedAt   int64  `json:"created_at" bson:"created_at"`
	CompletedAt *int64 `json:"completed_at" bson:"completed_at"`
}

type KitchenTask struct {
	UUID     uuid.UUID     `json:"uuid" bson:"uuid"`
	Name     string        `json:"name" bson:"name"`
	Space    *SpaceInfo    `json:"space" bson:"space"`
	Delivery *DeliveryInfo `json:"delivery" bson:"delivery"`
	Queue    *QueueInfo    `json:"queue" bson:"queue"`
	Orders   []TaskOrder   `json:"orders" bson:"orders"`
	DoingAt  int64         `json:"doing_at" bson:"doing_at"`
}

type SpaceInfo struct {
	UUID      uuid.UUID `json:"space_uuid" bson:"space_uuid"`
	Number    int       `json:"space_number" bson:"space_number"`
	ZoneName  string    `json:"zone_name" bson:"zone_name"`
	GroupCode string    `json:"group_code" bson:"group_code"`
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

type Customer struct {
	Name    string `json:"name" bson:"name"`
	Phone   Phone  `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`
}

type Phone struct {
	CountryCode string `json:"country_code" bson:"country_code"`
	Number      string `json:"number" bson:"number"`
}
