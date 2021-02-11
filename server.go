package main

import (
	"main/api"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main(){

	router := gin.Default()
	//setup CORS midleware Option
	config := cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 1 * time.Minute,
		Credentials: true,
		ValidateHeaders: false,
	}
	router.Use(cors.Middleware(config))
	api.Setup(router)
	router.Run(":8081")
}