package kitchen

import (
	"Kitchen/domain"
	"context"
	"time"

	"github.com/google/uuid"
)

type TaskOrderUC_Mongo struct {
	Repo *TaskOrderRepo_mongo
}

type TaskOrderUC_别的数据库 struct {
	Repo *TaskOrderRepo_别的数据库
}

func (uc *TaskOrderUC_别的数据库) NewOrderGroup(ctx context.Context, tableID int) int {
	if !uc.Repo.CheckOrderGroup(ctx, tableID) {
		uc.Repo.NewOrderGroup(ctx, tableID)
		return 200
	} else {
		return 500
	}
}

type TaskOrderUC interface {
	NewOrderGroup(ctx context.Context, tableID int) int
	NewKitchenTask(ctx context.Context, item domain.KitchenTask, id int)
	GetKitchenTask(ctx context.Context, tableID int) ([]domain.KitchenTask, error)
}

func New_TaskOrderUsecase_mongo(repo *TaskOrderRepo_mongo) *TaskOrderUC_Mongo {
	return &TaskOrderUC_Mongo{
		Repo: repo,
	}
}

func New_TaskOrderUsecase_别的数据库(repo *TaskOrderRepo_别的数据库) *TaskOrderUC_别的数据库 {
	//邪教代码
	return &TaskOrderUC_别的数据库{
		Repo: repo,
	}
	//邪教代码
}

func (uc *TaskOrderUC_Mongo) NewOrderGroup(ctx context.Context, tableID int) int {
	if !uc.Repo.CheckOrderGroup(ctx, tableID) {
		uc.Repo.NewOrderGroup(ctx, tableID)
		return 200
	} else {
		return 500
	}
}

func (uc *TaskOrderUC_Mongo) NewKitchenTask(ctx context.Context, item domain.KitchenTask, id int) {
	item.UUID, _ = uuid.NewUUID()
	for i := range item.Orders {
		item.Orders[i].UUID, _ = uuid.NewUUID()
		item.Orders[i].CreatedAt = time.Now().Unix()
	}
	uc.Repo.NewKitchenTask(ctx, item, id)
}

func (uc *TaskOrderUC_Mongo) GetKitchenTask(ctx context.Context, tableID int) ([]domain.KitchenTask, error) {
	return uc.Repo.GetKitchenTask(ctx, tableID)
}
