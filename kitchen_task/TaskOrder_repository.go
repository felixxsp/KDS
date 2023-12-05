package kitchen

import (
	myMongo "KDS/mongo/data"
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

func New_TaskOrderRepo_Mongo(database *mongo.Database) *TaskOrderRepo_mongo {
	return &TaskOrderRepo_mongo{
		database:   database,
		collection: database.Collection("taskOrder"),
	}
}

func (repo *TaskOrderRepo_mongo) NewOrderGroup(ctx context.Context, obj myMongo.OrderGroup) error {
	result, err := repo.collection.InsertOne(ctx, obj)
	fmt.Println(result)
	if err != nil {
		return fmt.Errorf("error at: repository/NewOrderGroup/InsertOne; error: %s", err.Error())
	}
	return nil
}

func (repo *TaskOrderRepo_mongo) OrderGroup_exists(ctx context.Context, uuid uuid.UUID) (bool, error) {
	check, err := repo.collection.CountDocuments(ctx, bson.M{"order_group_uuid": uuid})
	if err != nil {
		return false, fmt.Errorf("error at: repository/OrderGroup_exists/count_documents; error: %s", err.Error())
	} else if check >= 1 {
		return true, nil
	}
	return false, nil
}

func (repo *TaskOrderRepo_mongo) GetOrderGroup(ctx context.Context, uuid uuid.UUID) (myMongo.OrderGroup, error) {
	var result myMongo.OrderGroup
	err := repo.collection.FindOne(ctx, bson.M{"order_group_uuid": uuid}).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("error at: repository/GetOrderGroup/decoding; error: %s", err.Error())
	}
	return result, nil
}

func (repo *TaskOrderRepo_mongo) NewKitchenTask(ctx context.Context, obj myMongo.KitchenTask, group_uuid uuid.UUID) error {
	OrderGroup_parent, err := repo.GetOrderGroup(ctx, group_uuid)
	if err != nil {
		return err
	}
	OrderGroup_parent.Tasks = append(OrderGroup_parent.Tasks, obj)
	result, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"order_group_uuid": group_uuid},
		bson.M{"$set": OrderGroup_parent},
	)
	fmt.Println(result)
	if err != nil {
		return fmt.Errorf("error at: repository/NewKitchenTask/UpdateOne; error: %s", err.Error())

	}
	return nil
}

func (repo *TaskOrderRepo_mongo) GetKitchenTasks(ctx context.Context, group_uuid uuid.UUID) ([]myMongo.KitchenTask, error) {
	OrderGroup_parent, err := repo.GetOrderGroup(ctx, group_uuid)
	if err != nil {
		return OrderGroup_parent.Tasks, err
	}
	return OrderGroup_parent.Tasks, nil
}
