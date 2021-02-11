package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	db.SetupDB()
	setupAuthenAPI(router)
	setupUserAPI(router)
}
