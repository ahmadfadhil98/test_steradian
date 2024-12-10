package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Cars struct {
	CarId     *int    `json:"car_id" form:"car_id" gorm:"primary_key;AUTO_INCREMENT"`
	CarName   string  `json:"car_name" form:"car_name" gorm:"type:varchar(50);not null"`
	DayRate   float32 `json:"day_rate" form:"day_rate" gorm:"not null"`
	MonthRate float32 `json:"month_rate" form:"month_rate" gorm:"not null"`
	Image     string  `json:"image" form:"image" gorm:"type:varchar(256);not null"`
}

type CarsFilter struct {
	CarId  *int `json:"car_id" form:"car_id"`
	Limit  *int `json:"limit" form:"limit"`
	Offset *int `json:"offset" form:"offset"`
}

type CarsFilterDelete struct {
	CarId []int `json:"car_id" form:"car_id"`
}

func (car Cars) Validate() error {
	return validation.ValidateStruct(&car,
		validation.Field(&car.CarName, validation.Required),
		validation.Field(&car.DayRate, validation.Required),
		validation.Field(&car.MonthRate, validation.Required),
		validation.Field(&car.Image, validation.Required, is.URL),
	)
}
