package infrastructure

import (
	"TestSteradian/usecase"
	"github.com/gin-gonic/gin"
)

var Route Router

type Router struct {
	CarsUsecase  usecase.CarsUsecase
	OrderUsecase usecase.OrdersUsecase
}

func (r *Router) RouterInit() (*gin.Engine, error) {

	router := gin.Default()

	cars := router.Group("/cars")
	{
		cars.POST("/", r.CarsUsecase.Create)
		cars.GET("/", r.CarsUsecase.Read)
		cars.PUT("/:carId", r.CarsUsecase.Update)
		cars.DELETE("/", r.CarsUsecase.Delete)
	}

	orders := router.Group("/orders")
	{
		orders.POST("/", r.OrderUsecase.Create)
		orders.GET("/", r.OrderUsecase.Read)
		orders.PUT("/:orderId", r.OrderUsecase.Update)
		orders.DELETE("/", r.OrderUsecase.Delete)
	}

	err := router.Run(":8001")
	if err != nil {
		return nil, err
	}
	return router, nil
}
