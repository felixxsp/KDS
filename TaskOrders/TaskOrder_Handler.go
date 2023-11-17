package TaskOrder

import (
	"Kitchen/DTO"
	"context"
	"net/http"

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

func (Handler *TaskOrderHandler) Standby(ctx context.Context) {
	Handler.Router.GET("/TaskOrder/all", func(ctx *gin.Context) {
		TaskOrder, err := Handler.UC.getTaskOrders(ctx)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, "")
		} else {
			ctx.IndentedJSON(http.StatusOK, TaskOrder)
		}
	})

	Handler.Router.POST("/TaskOrder", func(ctx *gin.Context) {
		var Incoming DTO.TaskOrder
		ctx.BindJSON(&Incoming)
		err := Handler.UC.insertTaskOrder(ctx, Incoming)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, "")
		} else {
			ctx.IndentedJSON(http.StatusOK, Incoming)
		}
	})
}
