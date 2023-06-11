package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=muhammadazrifatihahsusanto dbname=bwastartupdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	input := user.LoginInput{
		Email:    "email@domain.com",
		Password: "p@ssw0rd",
	}

	user, err := userService.Login(input)

	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(user.Email)
		fmt.Println(user.Name)
	}

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()

}
