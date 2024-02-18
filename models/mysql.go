package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	dsn := "root:wm123@tcp(127.0.0.1::3306)/WM?charset=utf8mb4&parseTime=true&loc=Local"
	dB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	dB.Debug()
	fmt.Println("Mysql Database successfully")
}
