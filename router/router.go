package router

import (
	"github.com/gin-gonic/gin"
	"time"
)


type Public struct {
	ID        	uint 			`gorm:"primary_key",valid:"-"`
	CreatedAt 	time.Time		`valid:"-"`
	UpdatedAt 	time.Time		`valid:"-"`
	DeletedAt 	*time.Time 		`sql:"index",valid:"-"`
}


func Router()  {

	//路由配置
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})
	r.Run(":3000")

}