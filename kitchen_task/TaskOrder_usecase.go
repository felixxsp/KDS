package kitchen

import (
	domain "KDS/domain/entity"
	interfaces "KDS/domain/interface"
	myMongo "KDS/mongo/data"
	"fmt"

	"context"
	"time"

	"github.com/google/uuid"
)

type TaskOrderUC_Mongo struct {
	Repo interfaces.TaskOrderRepo
}

func New_TaskOrderUsecase_mongo(repo interfaces.TaskOrderRepo) *TaskOrderUC_Mongo {
	return &TaskOrderUC_Mongo{
		Repo: repo,
	}
}

func (uc *TaskOrderUC_Mongo) NewOrderGroup(ctx context.Context, tableID int) (uuid.UUID, error) {
	newUUID, _ := uuid.NewUUID() // what could go wrong when we're just simply generating a UUID (ツ)
	newOrder := myMongo.OrderGroup{
		UUID: newUUID,
		Space: &myMongo.SpaceInfo{
			Number: tableID,
		},
	}
	err := uc.Repo.NewOrderGroup(ctx, newOrder)
	if err != nil {
		return newUUID, err
	}
	return newUUID, nil
}

func (uc *TaskOrderUC_Mongo) NewKitchenTask(ctx context.Context, kitchenTask_req domain.KitchenTask, group_uuid uuid.UUID) error {
	newuuid, _ := uuid.NewUUID()
	kitchenTask_input := myMongo.KitchenTask{
		UUID:      newuuid,
		CreatedAt: time.Now().Unix(),
	}
	newuuid, _ = uuid.NewUUID()
	for _, iterate := range kitchenTask_req.Orders {
		orderTransfer := myMongo.Order{
			UUID:            newuuid,
			ItemName:        iterate.ItemName,
			ItemVariantName: iterate.ItemVariantName,
			Quantity:        iterate.Quantity,
		}
		kitchenTask_input.Orders = append(kitchenTask_input.Orders, orderTransfer)
	}
	return uc.Repo.NewKitchenTask(ctx, kitchenTask_input, group_uuid)
}

func (uc *TaskOrderUC_Mongo) GetKitchenTasks(ctx context.Context, group_uuid uuid.UUID) ([]domain.KitchenTask, error) {
	KitchenTask_db, err := uc.Repo.GetKitchenTasks(ctx, group_uuid)
	var Results []domain.KitchenTask
	if err != nil {
		return Results, err
	}

	for _, task_db := range KitchenTask_db {
		var result domain.KitchenTask
		result.UUID = task_db.UUID
		result.CreatedAt = task_db.CreatedAt
		//result.DoingAt = task_db.DoingAt
		for _, orders_db := range task_db.Orders {
			var order domain.Order
			order.UUID = orders_db.UUID
			order.ItemName = orders_db.ItemName
			order.ItemVariantName = orders_db.ItemVariantName
			order.Notes = orders_db.Notes
			order.Quantity = orders_db.Quantity
			// order.Modifier = orders_db.Modifiers
			order.CompletedAt = orders_db.CompletedAt
			result.Orders = append(result.Orders, order)
		}
		Results = append(Results, result)
	}
	fmt.Println(Results)
	return Results, nil
}
