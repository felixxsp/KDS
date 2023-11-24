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

type TaskOrderRepo_mongo struct {
	database   *mongo.Database
	collection *mongo.Collection
}

type TaskOrderRepo_别的数据库 struct {
	database   *mongo.Database
	collection *mongo.Collection
}

type TaskOrderRepo interface {
	NewOrderGroup(ctx context.Context, tableID int)
	CheckOrderGroup(ctx context.Context, tableID int) bool
	NewKitchenTask(ctx context.Context, item domain.KitchenTask, tableID int)
	GetKitchenTask(ctx context.Context, tableID int) ([]domain.KitchenTask, error)
}

func New_TaskOrderRepo_Mongo(database *mongo.Database) *TaskOrderRepo_mongo {
	return &TaskOrderRepo_mongo{
		database:   database,
		collection: database.Collection("taskOrder"),
	}
}

func New_TaskOrderRepo_别的数据库(database *mongo.Database) *TaskOrderRepo_别的数据库 {
	return &TaskOrderRepo_别的数据库{
		database:   database,
		collection: database.Collection("taskOrder"),
	}
}

func (repo *TaskOrderRepo_别的数据库) NewOrderGroup(ctx context.Context, tableID int) {
	//邪教代码
	newUUID, _ := uuid.NewUUID()
	newOrder := &myMongo.OrderGroup{
		UUID: newUUID,
		Space: &myMongo.SpaceInfo{
			Number: tableID,
		},
	}
	repo.collection.InsertOne(ctx, newOrder)
	fmt.Println(newOrder)
	//邪教代码
}

func (repo *TaskOrderRepo_mongo) NewOrderGroup(ctx context.Context, tableID int) {
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

func (repo *TaskOrderRepo_别的数据库) CheckOrderGroup(ctx context.Context, tableID int) bool {
	//邪教代码
	check, _ := repo.collection.CountDocuments(ctx, bson.M{"space.space_number": tableID})
	return check >= 1
	//邪教代码
}

func (repo *TaskOrderRepo_mongo) CheckOrderGroup(ctx context.Context, tableID int) bool {
	check, _ := repo.collection.CountDocuments(ctx, bson.M{"space.space_number": tableID})
	return check >= 1
}

func (repo *TaskOrderRepo_mongo) GetOrderGroup(ctx context.Context, tableID int) myMongo.OrderGroup {
	var result myMongo.OrderGroup
	repo.collection.FindOne(ctx, bson.M{"space.space_number": tableID}).Decode(&result)
	return result
}

func (repo *TaskOrderRepo_mongo) NewKitchenTask(ctx context.Context, item domain.KitchenTask, tableID int) {
	finalTask := &myMongo.Task{
		UUID:      item.UUID,
		CreatedAt: item.Orders[0].CreatedAt,
	}
	for _, iterate := range item.Orders {
		orderTransfer := myMongo.Order{
			UUID:            iterate.UUID,
			ItemName:        iterate.ItemName,
			ItemVariantName: iterate.ItemVariantName,
			Quantity:        iterate.Quantity,
		}
		finalTask.Orders = append(finalTask.Orders, orderTransfer)
	}
	finalForm := repo.GetOrderGroup(ctx, tableID)
	finalForm.Tasks = append(finalForm.Tasks, *finalTask)
	repo.collection.UpdateOne(
		ctx,
		bson.M{"space.space_number": tableID},
		bson.M{"$set": finalForm},
	)
}

func (repo *TaskOrderRepo_mongo) GetKitchenTask(ctx context.Context, tableID int) ([]domain.KitchenTask, error) {
	var Results []domain.KitchenTask
	group := repo.GetOrderGroup(ctx, tableID)

	for _, groupTask := range group.Tasks {
		var result domain.KitchenTask
		result.UUID = groupTask.UUID

		for _, order := range groupTask.Orders {
			result.Space.Number = tableID
			temp := domain.Order{
				UUID:            order.UUID,
				ItemName:        order.ItemName,
				ItemVariantName: order.ItemVariantName,
				Quantity:        order.Quantity,
				CreatedAt:       groupTask.CreatedAt,
			}
			result.Orders = append(result.Orders, temp)
		}
		Results = append(Results, result)
	}
	return Results, nil
}
