package domain

import (
	dto "Kitchen/DTO"

	"github.com/google/uuid"
)

type OrderGroup struct {
	UUID     uuid.UUID         `json:"order_group_uuid" bson:"order_group_uuid"`
	Space    Space             `json:"space" bson:"space"`
	Queue    *dto.QueueInfo    `json:"queue" bson:"queue"`
	Delivery *dto.DeliveryInfo `json:"delivery" bson:"delivery"`
	Tasks    []Tasks           `json:"tasks" bson:"tasks"`
}

type Space struct {
	Number    int    `json:"space_number" bson:"space_number"`
	GroupCode string `json:"group_code" bson:"group_code"`
	ZoneName  string `json:"zone_name" bson:"zone_name"`
}

type Tasks struct {
	UUID        uuid.UUID `json:"uuid" bson:"uuid"`
	Name        string    `json:"name" bson:"name"`
	Priority    int       `json:"priority" bson:"priority"`
	Orders      Orders    `json:"orders" bson:"orders"`
	CreatedAt   int64     `json:"created_at" bson:"created_at"`
	DoingAt     int64     `json:"doing_at" bson:"doing_at"`
	CompletedAt int64     `json:"completed_at" bson:"completed_at"`
}

type Orders struct {
	UUID            uuid.UUID           `json:"uuid" bson:"uuid"`
	ItemName        string              `json:"name" bson:"name"`
	ItemVariantName string              `json:"variant_name" bson:"variant_name"`
	Quantity        int                 `json:"quantity" bson:"quantity"`
	Notes           string              `json:"notes" bson:"notes"`
	Modifier        []dto.OrderModifier `json:"modifiers" bson:"modifiers"`

	CompletedAt *int64 `json:"completed_at" bson:"completed_at"`
}
