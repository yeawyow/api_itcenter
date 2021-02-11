package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	"github.com/gin-gonic/gin"
)

func setupUserAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{

		authenAPI.POST("/edituser", edituser)
		authenAPI.POST("/getuser", interceptor.JwtVerify, getuser)
	}
}

func getuser(c *gin.Context) {
	var employee []model.Employee
	db.GetDB().Find(&employee)
	c.JSON(200, gin.H{"data": employee})
}

func edituser(c *gin.Context) {
	c.JSON(200, gin.H{"stus": "ok"})
}
