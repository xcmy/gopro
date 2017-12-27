package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/asaskevich/govalidator"
	"fmt"
	"time"
	"reflect"
)



type Public struct {
	ID        	uint 			`valid:"-" gorm:"primary_key"`
	CreatedAt 	time.Time		`valid:"-"`
	UpdatedAt 	time.Time		`valid:"-"`
	DeletedAt 	*time.Time 		`valid:"-" sql:"index"`

}


type Job struct {
	gorm.Model
	Name 		string  `form:"name" json:"name" binding:"required"`
}


type User struct {
	Public		struct{}    `valid:"-"`
	Account  	string		`valid:"length(4|18)~account:,required~account can not be empty",gorm:"size:255"`
	Password 	string		`valid:"length(6|18),required~password can not be empty"`
	Nickname  	string		`valid:"-"`
	Gender 		bool		`valid:"-"`
	Phone 		string		`valid:"dialstring,optional"`
	Email 		string		`valid:"email,optional"`
	HeadImg 	string		`valid:"-"`
	Job 		Job			`valid:"-"`
	Status 		string			`valid:"-"`

}




var regStruct map[string]interface{}

func Init(db *gorm.DB)  {


	regStruct = make(map[string]interface{})
	regStruct["job"] = Job{}

	//db.Set("gorm:table_options", "ENGINE=PostgreSQL").AutoMigrate(&Job{},&User{})
	db.AutoMigrate(&Job{},&User{})

	userInit(db)
}

func userInit(db *gorm.DB)  {

	//super := Job{Name:"管理员"}
	//db.Create(&super)

	admin := User{Account:"adn",Password:"123456",Email:"67575@qq.com",Gender:true,Status:"启用"}

	result,err := govalidator.ValidateStruct(&admin)

	_,_ = govalidator.ValidateStruct(&admin)
	if err != nil {
		fmt.Println("err:",err.Error())
	}else {
		fmt.Println(result)
		//db.Create(&admin)
	}
	db.Create(&admin)

}



func Get(db *gorm.DB) (result []User) {

	fmt.Println("==>",db.HasTable("acl_user"))
	//acl_user := db.Table("acl_user")

	fmt.Println(db.Table("acl_user"))

	//db.AutoMigrate(db.Table("acl_user"))
	var users []User

	db.Find(&users)
	fmt.Println(users[0].Account)
	return  users

}

func Create(db *gorm.DB,c *gin.Context)(result Job) {

	var job Job
	if err := c.Bind(&job); err == nil {
		db.Create(&job)
	}else {
		fmt.Println(err)
	}
	return job
}
func GetById(db *gorm.DB,id int)(result Job) {

	var job Job
	db.First(&job, id)
	return job
}

func UpdateById(db *gorm.DB,c *gin.Context,id int) (result Job) {
	var job Job
	db.First(&job, id)
	if err := c.Bind(&job); err == nil {
		db.Save(&job)
	}else {
		fmt.Println(err)
	}
	return job
}

func DeleteById(db *gorm.DB,id int)(result string) {

	var job Job
	db.Delete(&job, id)
	return "success"
}


func GetList(db *gorm.DB,resource string)(result []Job) {

	str := resource
	if regStruct[str] != nil {
		t := reflect.ValueOf(regStruct[str]).Type()
		v := reflect.New(t).Elem()


		fmt.Println("de=>",Job{})
		fmt.Println("de=>",reflect.ValueOf(regStruct[str]).Type().String())
		fmt.Println("t=>",t)
		fmt.Println("v=>",v)

		
		var jobs []Job
		db.Find(&jobs)
		return jobs
	}else {
		return nil
	}

}


