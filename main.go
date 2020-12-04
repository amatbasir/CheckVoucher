package main

import (
	"CheckVoucherStatus/DBConfig"
	"CheckVoucherStatus/Models"
	"CheckVoucherStatus/Routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func main() {
	var err error
	Config.DB, err = gorm.Open(sqlite.Open(Config.DbUrl(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil{
		panic("Cannot connect to DB")
	}

	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Voucher{})

	r := Routes.SetupRouter()

	r.Run()
}