package kitchen

import (
	"Kitchen/domain"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskOrderHandler struct {
	Router   *gin.Engine
	UC_mongo *TaskOrderUC_Mongo
	UC_别的数据库 *TaskOrderUC_别的数据库
}

func InitTaskOrderHandler(router *gin.Engine, uc1 *TaskOrderUC_Mongo, uc2 *TaskOrderUC_别的数据库) *TaskOrderHandler {
	return &TaskOrderHandler{
		Router:   router,
		UC_mongo: uc1,
		UC_别的数据库: uc2,
	}
}

func (Handler *TaskOrderHandler) Standby(c context.Context) {
	Handler.Router.POST("/orderGroup/:id", Handler.NewOrderGroup)
	Handler.Router.POST("/taskOrder/:id", Handler.NewKitchenTask)
	Handler.Router.GET("/taskOrder/:id", Handler.GetKitchenTask)

}

func (handler *TaskOrderHandler) NewOrderGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	errInt := handler.UC_mongo.NewOrderGroup(c, id)
	if errInt == 200 {
		c.IndentedJSON(http.StatusOK, "Created new order group on Mongo")
	} else {
		c.IndentedJSON(http.StatusBadRequest, "Duplicate table ID")
	}

	//邪教代码
	errInt = handler.UC_别的数据库.NewOrderGroup(c, id)
	if errInt == 200 {
		c.IndentedJSON(http.StatusOK, "Created new order group on 别的数据库")
	} else {
		c.IndentedJSON(http.StatusBadRequest, "Duplicate table ID")
	}
	//邪教代码
}

func (handler *TaskOrderHandler) NewKitchenTask(c *gin.Context) {
	var Incoming domain.KitchenTask
	id, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&Incoming)
	handler.UC_mongo.NewKitchenTask(c, Incoming, id)
}

func (handler *TaskOrderHandler) GetKitchenTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	output, _ := handler.UC_mongo.GetKitchenTask(c, id)
	c.IndentedJSON(http.StatusOK, output)
}
