package myMongo

import (
	"github.com/google/uuid"
)

type KitchenTask struct {
	UUID        uuid.UUID `json:"uuid" bson:"uuid"`
	Name        string    `json:"name" bson:"name"`
	Priority    int       `json:"priority" bson:"priority"`
	Orders      []Order   `json:"orders" bson:"orders"`
	CreatedAt   int64     `json:"created_at" bson:"created_at"`
	DoingAt     *int64    `json:"doing_at" bson:"doing_at"`
	CompletedAt *int64    `json:"completed_at" bson:"completed_at"`
}

type Order struct {
	UUID            uuid.UUID       `json:"uuid" bson:"uuid"`
	ItemName        string          `json:"name" bson:"name"`
	ItemVariantName string          `json:"variant_name" bson:"variant_name"`
	Quantity        int             `json:"quantity" bson:"quantity"`
	Notes           string          `json:"notes" bson:"notes"`
	Modifiers       []OrderModifier `json:"modifiers" bson:"modifiers"`

	CompletedAt *int64 `json:"completed_at" bson:"completed_at"`
}
