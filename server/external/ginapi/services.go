package external

import (
	"laundry/internal/entity/global"
	"laundry/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServicesExternal struct {
	serviceUsecase *usecase.ServicesUsecase
	gin            *gin.Engine
}

func RegiserServicesExternal(
	serviceUsecase *usecase.ServicesUsecase,
	gin *gin.Engine,
) {
	ext := ServicesExternal{
		serviceUsecase,
		gin,
	}

	group := ext.gin.Group("/services")
	{
		group.GET("", ext.GetServices)
		group.GET("/:id/items", ext.GetServiceItemsById)
		group.GET("/:id/sub", ext.GetServiceSubServiceById)
		group.GET("/sub/:id/items", ext.GetSubServiceItemsById)
		// group.POST("/logout", ext.Logout)
	}

}

// GetServices получение доступных услуг
func (e *ServicesExternal) GetServices(c *gin.Context) {
	data, err := e.serviceUsecase.FindAllServices()
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetServiceItemsById поиск списка доступной одежды в услуге по id
func (e *ServicesExternal) GetServiceItemsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	data, err := e.serviceUsecase.FindServiceItemsByID(id, false)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetServiceItemsById поиск списка доступной одежды в под услуге услуге по id
func (e *ServicesExternal) GetSubServiceItemsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	data, err := e.serviceUsecase.FindServiceItemsByID(id, true)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetServiceItemsById поиск доп. услуг у основной услуги по id
func (e *ServicesExternal) GetServiceSubServiceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	data, err := e.serviceUsecase.FindServiceSubServiceById(id)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
