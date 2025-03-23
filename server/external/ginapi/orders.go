package external

import (
	"encoding/json"
	"fmt"
	"laundry/internal/entity/global"
	"laundry/internal/entity/orders"
	"laundry/internal/usecase"
	"laundry/redisclient"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersExternal struct {
	ordersUsecase *usecase.OrdersUsecase
	redisClient   *redisclient.RedisClient
	gin           *gin.Engine
}

func RegiserOrdersExternal(
	ordersUsecase *usecase.OrdersUsecase,
	redisClient *redisclient.RedisClient,
	gin *gin.Engine,
) {
	ext := OrdersExternal{
		ordersUsecase,
		redisClient,
		gin,
	}

	group := ext.gin.Group("/orders")
	{
		group.POST("/calculate", ext.Calculate)
		group.POST("/create/:id", ext.CreateOrder)
		group.GET("/today", ext.GetTodayOrders)
	}
}

// Calculate Метод принимает JSON с параметрами расчета заказа, выполняет расчет через ordersUsecase.CalculateOrder,
// сериализует результат и сохраняет его в Redis. Затем отправляет рассчитанные данные клиенту.
func (e *OrdersExternal) Calculate(c *gin.Context) {
	param := orders.CalculateOrderParam{}

	if err := c.BindJSON(&param); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": err.Error()})
		return
	}

	data, err := e.ordersUsecase.CalculateOrder(param)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInternalError], gin.H{"message": global.ErrInternalError.Error()})
		return
	}

	err = e.redisClient.Set(data.TemporaryID, bytes)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// CreateOrder Метод получает ID предварительно рассчитанного заказа, извлекает данные из Redis,
// принимает пользовательские данные из JSON, объединяет их с расчетными данными и создает заказ через ordersUsecase.CreateOrder.
// Возвращает ID созданного заказа.
func (e *OrdersExternal) CreateOrder(c *gin.Context) {
	id := c.Param("id")

	data, err := e.redisClient.Get(id)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	userParam := orders.CreateOrderParam{}

	if err := c.BindJSON(&userParam); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": err.Error()})
		return
	}

	var result orders.CalculateOrderResponse

	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInternalError], gin.H{"message": global.ErrInternalError.Error()})
		return
	}

	createId, err := e.ordersUsecase.CreateOrder(orders.CreateOrderParamWithPreCalculatedData{
		UserParam: orders.CreateOrderDbParam{
			UserName:     userParam.UserName,
			PhoneNumber:  userParam.PhoneNumber,
			CreationDate: time.Now(),
			Total:        result.Total,
			Final:        result.Final,
		},
		PreCalculatedData: result,
	})
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": createId})
}

// GetTodayOrders Метод запрашивает заказы, оформленные сегодня, через ordersUsecase.FindTodayOrders и отправляет их клиенту.
func (e *OrdersExternal) GetTodayOrders(c *gin.Context) {
	data, err := e.ordersUsecase.FindTodayOrders()
	fmt.Println("data: ", data)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
