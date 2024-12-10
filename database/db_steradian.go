package database

import (
	"TestSteradian/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var Database DBSterdian

type DBSterdian struct {
	DB *gorm.DB
}

func (dbs *DBSterdian) DBInit() error {
	err := godotenv.Load("../TestSteradian/database/.env")
	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"), os.Getenv("DB_PORT"))
	fmt.Println(dataSourceName)
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(domain.Cars{}, domain.Orders{})
	dbs.DB = db
	return err
}
