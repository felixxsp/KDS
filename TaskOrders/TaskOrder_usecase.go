package kitchen

import (
	"Kitchen/domain"
	"context"
	"time"

	"github.com/google/uuid"
)

type TaskOrderUsecase struct {
	Repo *TaskOrderRepoImpl
}

func InitTaskOrderUsecase(repo *TaskOrderRepoImpl) *TaskOrderUsecase {
	return &TaskOrderUsecase{
		Repo: repo,
	}
}

func (uc *TaskOrderUsecase) NewOrderGroup(ctx context.Context, tableID int) int {
	if !uc.Repo.CheckOrderGroup(ctx, tableID) {
		uc.Repo.NewOrderGroup(ctx, tableID)
		return 200
	} else {
		return 500
	}
}

func (uc *TaskOrderUsecase) NewKitchenTask(ctx context.Context, item domain.KitchenTask, id int) {
	item.UUID, _ = uuid.NewUUID()
	for i := range item.Orders {
		item.Orders[i].UUID, _ = uuid.NewUUID()
		item.Orders[i].CreatedAt = time.Now().Unix()
	}
	uc.Repo.NewKitchenTask(ctx, item, id)
}

func (uc *TaskOrderUsecase) GetKitchenTask(ctx context.Context, tableID int) ([]domain.KitchenTask, error) {
	return uc.Repo.GetKitchenTask(ctx, tableID)
}
