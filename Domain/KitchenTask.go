package domain

import (
	"github.com/google/uuid"
)

type KitchenTask struct {
	UUID     uuid.UUID     `json:"uuid" bson:"uuid"`
	Space    *SpaceInfo    `json:"space" bson:"space"`
	Delivery *DeliveryInfo `json:"delivery" bson:"delivery"`
	Queue    *QueueInfo    `json:"queue" bson:"queue"`
	Orders   []Order       `json:"orders" bson:"orders"`
	DoingAt  int64         `json:"doing_at" bson:"doing_at"`
}

type Order struct {
	UUID            uuid.UUID       `json:"uuid" bson:"uuid"`
	ItemName        string          `json:"item_name" bson:"item_name"`
	ItemVariantName string          `json:"item_variant_name" bson:"item_variant_name"`
	Quantity        int             `json:"quantity" bson:"quantity"`
	Notes           string          `json:"notes" bson:"notes"`
	Modifier        []OrderModifier `json:"modifiers" bson:"modifiers"`

	CreatedAt   int64  `json:"created_at" bson:"created_at"`
	CompletedAt *int64 `json:"completed_at" bson:"completed_at"`
}

type OrderModifier struct {
	Name     string `json:"name" bson:"name"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
