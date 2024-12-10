package domain

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Orders struct {
	OrderId         *int   `gorm:"primary_key;AUTO_INCREMENT"`
	CarId           int    `json:"car_id" form:"car_id" gorm:"not null"`
	OrderDate       string `json:"order_date" form:"order_date" gorm:"type:date;not null"`
	PickupDate      string `json:"pickup_date" form:"pickup_date" gorm:"type:date;not null"`
	DropoffDate     string `json:"dropoff_date" form:"dropoff_date" gorm:"type:date;not null"`
	PickupLocation  string `json:"pickup_location" form:"pickup_location" gorm:"type:varchar(50);not null"`
	DropoffLocation string `json:"dropoff_location" form:"dropoff_location" gorm:"type:varchar(50);not null"`
	//Cars            Cars   `gorm:"foreignkey:CarId;references:CarId;"`
}

type OrdersFilter struct {
	OrderId *int   `json:"order_id" form:"order_id"`
	CarId   *[]int `json:"car_id" form:"car_id"`
	Limit   *int   `json:"limit" form:"limit"`
	Offset  *int   `json:"offset" form:"offset"`
}

type OrdersFilterDelete struct {
	OrderId []int `json:"order_id" form:"order_id"`
}

func (order Orders) Validate() error {

	err := validation.ValidateStruct(&order,
		validation.Field(&order.CarId, validation.Required),
		validation.Field(&order.OrderDate, validation.Required),
		validation.Field(&order.PickupDate, validation.Required),
		validation.Field(&order.DropoffDate, validation.Required),
		validation.Field(&order.DropoffLocation, validation.Required),
		validation.Field(&order.DropoffLocation, validation.Required),
	)

	err = validation.Date("2006-01-02").Validate(order.OrderDate)
	err = validation.Date("2006-01-02").Validate(order.PickupDate)
	err = validation.Date("2006-01-02").Validate(order.DropoffDate)

	if !(order.OrderDate <= order.PickupDate && order.PickupLocation <= order.DropoffDate) {
		err = errors.New("You miss order date or pickup date or dropoff date")
	}

	return err
}
