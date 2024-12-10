package repository

import (
	"TestSteradian/database"
	"TestSteradian/domain"
)

type CarsRepository struct {
}

func (c *CarsRepository) Insert(cars domain.Cars) (*domain.Cars, error) {

	err := database.Database.DB.Debug().Table("cars").Create(&cars).Error
	if err != nil {
		return nil, err
	}

	return &cars, nil
}

func (c *CarsRepository) Get(carsFilter domain.CarsFilter) ([]domain.Cars, error) {
	var result []domain.Cars
	get := database.Database.DB.Debug().Table("cars")

	if carsFilter.CarId != nil {
		get = get.Where("car_id = ?", *carsFilter.CarId)
	}

	if carsFilter.Limit != nil {
		if *carsFilter.Limit > 0 {
			get = get.Limit(*carsFilter.Limit)
		}
		if carsFilter.Offset != nil {
			get = get.Offset(*carsFilter.Offset)
		}
	}

	err := get.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *CarsRepository) Update(id int, cars domain.Cars) (*domain.Cars, error) {

	err := database.Database.DB.Debug().Table("cars").Where("car_id = ?", id).Update(&cars).Scan(&cars).Error
	if err != nil {
		return nil, err
	}

	return &cars, nil
}

func (c *CarsRepository) Delete(carId []int) error {
	get := database.Database.DB.Debug().Table("cars")

	for _, id := range carId {
		get = get.Or("car_id = ?", id)
	}

	err := get.Delete(&domain.Cars{}).Error
	if err != nil {
		return err
	}

	return nil
}
