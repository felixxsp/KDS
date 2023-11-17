package TaskOrder

import (
	"Kitchen/DTO"
	"context"
)

type TaskOrderUsecase struct {
	Repo *TaskOrderRepo
}

func InitTaskOrderUsecase(repo *TaskOrderRepo) *TaskOrderUsecase {
	return &TaskOrderUsecase{
		Repo: repo,
	}
}

func (uc *TaskOrderUsecase) getTaskOrders(ctx context.Context) ([]DTO.TaskOrder, error) {
	return uc.Repo.getTaskOrders(ctx)
}

func (uc *TaskOrderUsecase) insertTaskOrder(ctx context.Context, task DTO.TaskOrder) error {
	return uc.Repo.insertTaskOrder(ctx, task)
}
