package main

import (
	"crowdproject/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func handler(c *gin.Context)  {
	dsn := "host=localhost user=iqbalpradipta password=mbangg12 dbname=cfproject port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []users.Users
	db.Find(&users)

	c.JSON(http.StatusOK, users)

}

func main()  {
	// dsn := "host=localhost user=iqbalpradipta password=mbangg12 dbname=cfproject port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connect to database successfuly")

	// var users []users.Users

	// db.Find(&users)
	
	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("================")
	// }

	router := gin.Default()
	router.GET("/users", handler)
	router.Run()
}