package usecase

import (
	"TestSteradian/domain"
	"TestSteradian/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrdersUsecase struct {
	ordersRepo repository.OrdersRepository
	carsRepo   repository.CarsRepository
}

func (ou *OrdersUsecase) Create(c *gin.Context) {
	var order domain.Orders
	var response domain.Response

	err := c.ShouldBind(&order)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(order)

	err = order.Validate()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	carsFilter := domain.CarsFilter{}
	limit := 1
	offset := 0
	carsFilter.CarId = &order.CarId
	carsFilter.Limit = &limit
	carsFilter.Offset = &offset

	dataCars, err := ou.carsRepo.Get(carsFilter)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(dataCars) == 0 {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": "cars not found"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataOrder, err := ou.ordersRepo.Insert(order)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = gin.H{"dataOrder": dataOrder, "dataCars": dataCars}
	c.JSON(http.StatusCreated, response)
}

func (ou *OrdersUsecase) Read(c *gin.Context) {
	var orderFilter domain.OrdersFilter
	var response domain.Response
	err := c.ShouldBind(&orderFilter)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataOrder, err := ou.ordersRepo.Get(orderFilter)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(dataOrder) == 0 {
		response.Status = http.StatusNotFound
		response.Data = gin.H{"message": "order not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = gin.H{"dataOrder": dataOrder}
	c.JSON(http.StatusOK, response)
}

func (ou *OrdersUsecase) Update(c *gin.Context) {
	var order domain.Orders
	var response domain.Response
	err := c.ShouldBind(&order)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id := c.Param("orderId")
	orderId, err := strconv.Atoi(id)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataOrder, err := ou.ordersRepo.Update(orderId, order)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = gin.H{"dataOrder": dataOrder}
	c.JSON(http.StatusOK, response)
}

func (ou *OrdersUsecase) Delete(c *gin.Context) {
	var response domain.Response
	var orderFilterDelete domain.OrdersFilterDelete
	err := c.ShouldBind(&orderFilterDelete)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if len(orderFilterDelete.OrderId) < 1 {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"message": "orderId is empty"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ou.ordersRepo.Delete(orderFilterDelete.OrderId)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"message": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = gin.H{"message": "Succesfully deleted"}
	c.JSON(http.StatusOK, response)
}
