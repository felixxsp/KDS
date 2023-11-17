package main

import (
	"Kitchen/TaskOrder"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mainCtx := context.Background()

	router := gin.Default()

	client, _ := mongo.Connect(mainCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	database := client.Database("Kitchen")

	TaskOrderRepo := TaskOrder.InitTaskOrderRepo(database)
	TaskorderUsecase := TaskOrder.InitTaskOrderUsecase(TaskOrderRepo)
	TaskOrderHandler := TaskOrder.InitTaskOrderHandler(router, TaskorderUsecase)
	TaskOrderHandler.Standby(mainCtx)

	router.Run(":6969")
}
