package interfaces

import (
	domain "KDS/domain/entity"
	"context"

	"github.com/google/uuid"
)

type TaskOrderUC interface {
	NewOrderGroup(context.Context, int) (uuid.UUID, error)
	NewKitchenTask(context.Context, domain.KitchenTask, uuid.UUID) error
	GetKitchenTasks(context.Context, uuid.UUID) ([]domain.KitchenTask, error)
}
