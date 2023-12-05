package interfaces

import (
	myMongo "KDS/mongo/data"
	"context"

	"github.com/google/uuid"
)

type TaskOrderRepo interface {
	NewOrderGroup(context.Context, myMongo.OrderGroup) error
	OrderGroup_exists(context.Context, uuid.UUID) (bool, error)
	NewKitchenTask(context.Context, myMongo.KitchenTask, uuid.UUID) error
	GetKitchenTasks(context.Context, uuid.UUID) ([]myMongo.KitchenTask, error)
}
