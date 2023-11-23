package kitchen

import (
	"Kitchen/domain"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskOrderHandler struct {
	Router *gin.Engine
	UC     *TaskOrderUsecase
}

func InitTaskOrderHandler(router *gin.Engine, uc *TaskOrderUsecase) *TaskOrderHandler {
	return &TaskOrderHandler{
		Router: router,
		UC:     uc,
	}
}

func (Handler *TaskOrderHandler) Standby(c context.Context) {
	Handler.Router.POST("/orderGroup/:id", Handler.NewOrderGroup)
	Handler.Router.POST("/taskOrder/:id", Handler.NewKitchenTask)
	Handler.Router.GET("/taskOrder/:id", Handler.GetKitchenTask)

}

func (handler *TaskOrderHandler) NewOrderGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	errInt := handler.UC.NewOrderGroup(c, id)
	if errInt == 200 {
		c.IndentedJSON(http.StatusOK, "Created new order group")
	} else {
		c.IndentedJSON(http.StatusBadRequest, "Duplicate table ID")
	}
}

func (handler *TaskOrderHandler) NewKitchenTask(c *gin.Context) {
	var Incoming domain.KitchenTask
	id, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&Incoming)
	handler.UC.NewKitchenTask(c, Incoming, id)
}

func (handler *TaskOrderHandler) GetKitchenTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	output, _ := handler.UC.GetKitchenTask(c, id)
	c.IndentedJSON(http.StatusOK, output)
}
