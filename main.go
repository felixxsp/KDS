package main

import (
	"Kitchen/kitchen"
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

	TaskOrderRepo := kitchen.InitTaskOrderRepo(database)
	TaskorderUsecase := kitchen.InitTaskOrderUsecase(TaskOrderRepo)
	TaskOrderHandler := kitchen.InitTaskOrderHandler(router, TaskorderUsecase)
	TaskOrderHandler.Standby(mainCtx)

	//fmt.Println(TaskOrderHandler.UC.GetKitchenTask(mainCtx, 20))
	router.Run(":8080")

}
