package main

import (
	"my-project/controller"
	"my-project/db"
	"my-project/repository"
	"my-project/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(dbConnection)
	sentEmailRepository := repository.NewSentEmailRepository(dbConnection)

	userUseCase := usecase.NewUserUseCase(userRepository)
	sentEmailUseCase := usecase.NewSentEmailUseCase(sentEmailRepository)

	userController := controller.UserController{
		UserUseCase: userUseCase,
	}

	sentEmailController := controller.SentEmailController{
		SentEmailUseCase: sentEmailUseCase,
	}

	server.GET("/users", userController.GetUsers)
	server.POST("/users", userController.CreateUser)
	server.PUT("/users/:id", userController.UpdateUser)
	server.DELETE("/users/:id", userController.DeleteUser)

	server.GET("/sentemails", sentEmailController.GetSentEmails)
	server.POST("/sentemails", sentEmailController.CreateSentEmail)
	server.PUT("/sentemails/:id", sentEmailController.UpdateSentEmail)
	server.DELETE("/sentemails/:id", sentEmailController.DeleteSentEmail)

	server.Run(":8800")
}
