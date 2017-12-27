package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	)

func Connect()  {

	//数据库链接
	db, err := gorm.Open("postgres", "host=localhost  dbname=xcmy sslmode=disable ")
	defer db.Close()
	if err == nil {
		fmt.Println("db connected successful")
	} else {
		fmt.Println(err)
	}
}
