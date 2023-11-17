package TaskOrder

import (
	"Kitchen/DTO"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskOrderRepo struct {
	Collection *mongo.Collection
}

func InitTaskOrderRepo(database *mongo.Database) *TaskOrderRepo {
	return &TaskOrderRepo{
		Collection: database.Collection("TaskOrder"),
	}
}

func (repo *TaskOrderRepo) getTaskOrders(ctx context.Context) ([]DTO.TaskOrder, error) {
	var Results []DTO.TaskOrder

	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var Result DTO.TaskOrder
		err := cursor.Decode(&Result)
		if err != nil {
			return nil, err
		}
		Results = append(Results, Result)
	}

	return Results, nil
}

func (repo *TaskOrderRepo) insertTaskOrder(ctx context.Context, task DTO.TaskOrder) error {
	_, err := repo.Collection.InsertOne(ctx, task)
	return err
}
