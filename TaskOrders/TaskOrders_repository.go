package kitchen

import (
	"Kitchen/domain"
	myMongo "Kitchen/mongo/data"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskOrderRepoImpl struct {
	collection *mongo.Collection
}

func InitTaskOrderRepo(database *mongo.Database) *TaskOrderRepoImpl {
	return &TaskOrderRepoImpl{
		collection: database.Collection("taskOrder"),
	}
}

func (repo *TaskOrderRepoImpl) NewOrderGroup(ctx context.Context, tableID int) {
	newUUID, _ := uuid.NewUUID()
	newOrder := &myMongo.OrderGroup{
		UUID: newUUID,
		Space: &myMongo.SpaceInfo{
			Number: tableID,
		},
	}
	repo.collection.InsertOne(ctx, newOrder)
	fmt.Println(newOrder)
}

func (repo *TaskOrderRepoImpl) CheckOrderGroup(ctx context.Context, tableID int) bool {
	check, _ := repo.collection.CountDocuments(ctx, bson.M{"space.space_number": tableID})
	return check >= 1
}

func (repo *TaskOrderRepoImpl) GetOrderGroup(ctx context.Context, tableID int) myMongo.OrderGroup {
	var result myMongo.OrderGroup
	repo.collection.FindOne(ctx, bson.M{"space.space_number": tableID}).Decode(&result)
	return result
}

func (repo *TaskOrderRepoImpl) NewKitchenTask(ctx context.Context, item domain.KitchenTask, tableID int) {
	finalTask := &myMongo.Task{
		UUID:      item.UUID,
		CreatedAt: item.Orders[0].CreatedAt,
	}
	for _, iterate := range item.Orders {
		orderTransfer := &myMongo.Order{
			UUID:            iterate.UUID,
			ItemName:        iterate.ItemName,
			ItemVariantName: iterate.ItemVariantName,
			Quantity:        iterate.Quantity,
		}
		finalTask.Orders = append(finalTask.Orders, *orderTransfer)
	}
	finalForm := repo.GetOrderGroup(ctx, tableID)
	finalForm.Tasks = append(finalForm.Tasks, *finalTask)
	repo.collection.UpdateOne(
		ctx,
		bson.M{"space.space_number": tableID},
		bson.M{"$set": finalForm},
	)
}

func (repo *TaskOrderRepoImpl) GetKitchenTask(ctx context.Context, tableID int) ([]domain.KitchenTask, error) {
	var Results []domain.KitchenTask
	group := repo.GetOrderGroup(ctx, tableID)

	for _, groupTask := range group.Tasks {
		var result domain.KitchenTask
		result.UUID = groupTask.UUID

		for _, order := range groupTask.Orders {
			result.Space.Number = tableID
			temp := &domain.Order{
				UUID:            order.UUID,
				ItemName:        order.ItemName,
				ItemVariantName: order.ItemVariantName,
				Quantity:        order.Quantity,
				CreatedAt:       groupTask.CreatedAt,
			}
			result.Orders = append(result.Orders, *temp)
		}
		Results = append(Results, result)
	}
	return Results, nil
}
