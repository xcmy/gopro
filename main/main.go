package main

import (
	"fmt"
	"gopro/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/asaskevich/govalidator"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"strconv"
)

func init()  {
	govalidator.SetFieldsRequiredByDefault(true)
}

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

	model.Init(db)

	//路由配置
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(200,model.Get(db))
		//c.String(200,"hello world")
	})

	r.POST("/job/create", func(c *gin.Context) {
		c.JSON(200,model.Create(db,c))
	})


	r.POST("/job/update/:id", func(c *gin.Context) {
		id,err := strconv.Atoi(c.Param("id"))
		if err == nil {
			c.JSON(200,model.UpdateById(db,c,id))
		}
	})



	r.GET("/job/:id", func(c *gin.Context) {
		id,err := strconv.Atoi(c.Param("id"))
		if err == nil {
			c.JSON(200,model.GetById(db,id))
		}
	})


	r.GET("/job", func(c *gin.Context) {

		//resource := c.Param("resource")

		c.JSON(200,model.GetList(db,"job"))
	})


	r.DELETE("/job/:id", func(c *gin.Context) {
		id,err := strconv.Atoi(c.Param("id"))
		if err == nil {
			c.JSON(200,model.DeleteById(db,id))
		}
	})




	r.Run(":3000")

}
