package external

import (
	"laundry/internal/entity/global"
	"laundry/internal/entity/orders"
	"laundry/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersExternal struct {
	ordersUsecase *usecase.OrdersUsecase
	gin           *gin.Engine
}

func RegiserOrdersExternal(
	ordersUsecase *usecase.OrdersUsecase,
	gin *gin.Engine,
) {
	ext := OrdersExternal{
		ordersUsecase,
		gin,
	}

	group := ext.gin.Group("/orders")
	{
		group.POST("", ext.Create)
	}

}

func (e *OrdersExternal) Create(c *gin.Context) {
	param := orders.CreateOrderParam{}

	if err := c.BindJSON(&param); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": err.Error()})
		return
	}

	data, err := e.ordersUsecase.ProcessOrder(param)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
