package main

import (
	interfaces "KDS/domain/interface"
	kitchen "KDS/kitchen_task"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mainCtx := context.Background()

	router := gin.Default()

	client, _ := mongo.Connect(mainCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	database := client.Database("kitchen")

	var TaskOrderRepo interfaces.TaskOrderRepo = kitchen.New_TaskOrderRepo_Mongo(database)
	var TaskorderUsecase interfaces.TaskOrderUC = kitchen.New_TaskOrderUsecase_mongo(TaskOrderRepo)
	TaskOrderHandler := kitchen.InitTaskOrderHandler(router, TaskorderUsecase)
	TaskOrderHandler.Standby(mainCtx)

	router.Run(":8080")
}
