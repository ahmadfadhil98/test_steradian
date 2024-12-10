package usecase

import (
	"TestSteradian/domain"
	"TestSteradian/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CarsUsecase struct {
	carsRepo   repository.CarsRepository
	ordersRepo repository.OrdersRepository
}

func (cu *CarsUsecase) Create(c *gin.Context) {
	var car domain.Cars
	var response domain.Response
	err := c.ShouldBind(&car)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = car.Validate()
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataCars, err := cu.carsRepo.Insert(car)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusCreated
	response.Data = gin.H{"data": dataCars}
	c.JSON(http.StatusCreated, response)

}

func (cu *CarsUsecase) Read(c *gin.Context) {
	var carFilter domain.CarsFilter
	var response domain.Response
	err := c.ShouldBind(&carFilter)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataCars, err := cu.carsRepo.Get(carFilter)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(dataCars) == 0 {
		response.Status = http.StatusNotFound
		response.Data = gin.H{"error": "car not found"}
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = gin.H{"data": dataCars}
	c.JSON(http.StatusOK, response)

}

func (cu *CarsUsecase) Update(c *gin.Context) {
	var car domain.Cars
	var response domain.Response
	err := c.ShouldBind(&car)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	id := c.Param("carId")
	carId, err := strconv.Atoi(id)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataCars, err := cu.carsRepo.Update(carId, car)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = gin.H{"data": dataCars}
	c.JSON(http.StatusOK, response)

}

func (cu *CarsUsecase) Delete(c *gin.Context) {
	var response domain.Response
	var carsFilterDelete domain.CarsFilterDelete
	err := c.ShouldBind(&carsFilterDelete)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fmt.Println(carsFilterDelete)

	if len(carsFilterDelete.CarId) < 1 {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": "orderId is empty"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	filterOrder := domain.OrdersFilter{}
	limit := len(carsFilterDelete.CarId)
	offset := 0

	filterOrder.CarId = &carsFilterDelete.CarId
	filterOrder.Limit = &limit
	filterOrder.Offset = &offset

	dataOrder, err := cu.ordersRepo.Get(filterOrder)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(dataOrder) > 0 {
		response.Status = http.StatusBadRequest
		response.Data = gin.H{"error": "the car id is still in order"}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = cu.carsRepo.Delete(carsFilterDelete.CarId)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Data = gin.H{"error": err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = gin.H{"message": "Succesfully deleted"}
	c.JSON(http.StatusOK, response)

}
