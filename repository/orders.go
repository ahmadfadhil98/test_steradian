package repository

import (
	"TestSteradian/database"
	"TestSteradian/domain"
	"net/http"
)

type OrdersRepository struct {
}

func (o *OrdersRepository) Insert(order domain.Orders) (domain.Response, error) {
	result := domain.Response{}

	err := database.Database.DB.Debug().Table("orders").Create(&order).Error
	if err != nil {
		return result, err
	}

	result.Status = http.StatusCreated
	result.Data = order

	return result, nil
}

func (o *OrdersRepository) Get(orderFilter domain.OrdersFilter) ([]domain.Orders, error) {
	var order []domain.Orders
	get := database.Database.DB.Debug().Table("orders")

	if orderFilter.OrderId != nil {
		get = get.Where("order_id=?", *orderFilter.OrderId)
	}

	if orderFilter.CarId != nil {
		for _, id := range *orderFilter.CarId {
			get = get.Or("car_id=?", id)
		}
	}

	if orderFilter.Limit != nil {
		if *orderFilter.Limit > 0 {
			get = get.Limit(*orderFilter.Limit)
		}

		if orderFilter.Offset != nil {
			get = get.Offset(*orderFilter.Offset)
		}
	}

	err := get.Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (o *OrdersRepository) Update(orderId int, order domain.Orders) (*domain.Orders, error) {
	err := database.Database.DB.Debug().Table("orders").Where("order_id=?", orderId).Updates(&order).Scan(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil

}

func (o *OrdersRepository) Delete(orderId []int) error {
	get := database.Database.DB.Debug().Table("orders")

	for _, id := range orderId {
		get = get.Where("order_id=?", id)
	}

	err := get.Delete(&domain.Orders{}).Error
	if err != nil {
		return err
	}

	return nil
}
