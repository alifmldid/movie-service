package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBCon() *gorm.DB{
	dbCon := "host=localhost user=developer password=devspassword dbname=movie port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dbCon), &gorm.Config{})

	if(err != nil){
		panic(err)
	}

	return db
}