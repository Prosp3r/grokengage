package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=dbadmin password=Akomeno123, dbname=engage_and_earn port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Auto-migrate models
	db.AutoMigrate(&User{}, &Task{}, &CheckIn{}, &Reward{})
	return db
}
