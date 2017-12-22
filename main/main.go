package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main()  {

	fmt.Println("hello")

	//数据库链接
	db, err := gorm.Open("postgres", "host=localhost  dbname=xcmy sslmode=disable ")
	defer db.Close()
	if err == nil {
		fmt.Println("db connected successful")
	} else {
		fmt.Println(err)
	}

	//路由配置
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})
	r.Run(":3000")

}
