package kitchen

import (
	domain "KDS/domain/entity"
	interfaces "KDS/domain/interface"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskOrderHandler struct {
	Router *gin.Engine
	uc     interfaces.TaskOrderUC
}

func InitTaskOrderHandler(router *gin.Engine, uc interfaces.TaskOrderUC) *TaskOrderHandler {
	return &TaskOrderHandler{
		Router: router,
		uc:     uc,
	}
}

func (Handler *TaskOrderHandler) Standby(c context.Context) {
	Handler.Router.POST("/newOrderGroup/:id", Handler.NewOrderGroup)
	Handler.Router.POST("/kitchenTask/:uuid", Handler.NewKitchenTask)
	Handler.Router.GET("/kitchenTask/:uuid", Handler.GetKitchenTask)
}

func (handler *TaskOrderHandler) NewOrderGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("Table ID parsing failed: %s", err.Error()))
		return
	}

	uuid, err := handler.uc.NewOrderGroup(c, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, uuid)
}

func (handler *TaskOrderHandler) NewKitchenTask(c *gin.Context) {
	var Incoming domain.KitchenTask
	uuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("UUID parsing failed: %s", err.Error()))
		return
	}

	err = c.BindJSON(&Incoming)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("JSON binding failed: %s", err.Error()))
		return
	}

	err = handler.uc.NewKitchenTask(c, Incoming, uuid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "Succesfully added kitchen task")
}

func (handler *TaskOrderHandler) GetKitchenTask(c *gin.Context) {
	uuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("UUID parsing failed: %s", err.Error()))
		return
	}

	output, err := handler.uc.GetKitchenTasks(c, uuid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, output)
}
