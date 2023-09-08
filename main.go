package main

import (
	"crowdproject/users"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=iqbalpradipta password=mbangg12 dbname=cfproject port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)

	userInput := users.RegisterInput{}
	userInput.Name = "Pak Ikebal"
	userInput.Email = "cobacoba@gmail.com"
	userInput.Occupation = "kang bensin"
	userInput.Password = "123"

	userService.RegisterUser(userInput)

}
