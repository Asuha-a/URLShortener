package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// DB instance
	DB  *gorm.DB
	err error
)

// Init DB
func Init() {
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=user_db user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Tokyo",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	autoMigration()
}

// Close DB
func Close() {
	dbSQL, err := DB.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func autoMigration() {
	DB.AutoMigrate(&User{})
}
