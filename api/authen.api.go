package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(passwordjwt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordjwt), 14)
	return string(bytes), err
}
func setupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{
      
		authenAPI.POST("/login", login)
		authenAPI.GET("/login", test)
		authenAPI.POST("/register", register)
	}
}
func login(c *gin.Context) {
	var employee model.Employee
	if c.ShouldBind(&employee) == nil {
		var queryUser model.Employee
		if err := db.GetDB().First(&queryUser, "username=?", employee.Username).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else if CheckPasswordHash(employee.Passwordjwt, queryUser.Passwordjwt) == false {
			c.JSON(200, gin.H{"result": "nok", "error": "invalid password"})
		} else {

			token := interceptor.JwtSign(queryUser)

			c.JSON(200, gin.H{"result": "ok", "token": token})
		}

	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}

}

func register(c *gin.Context) {
	var employee model.Employee

	if c.ShouldBind(&employee) == nil {
		employee.Passwordjwt, _ = hashPassword(employee.Passwordjwt)
		if err := db.GetDB().Create(&employee).Error; err != nil {
			c.JSON(200, gin.H{"status": "nok", "error": err})
		} else {
			c.JSON(200, gin.H{"result": "ok", "data": employee})
		}

	} else {
		c.JSON(200, gin.H{"data": "unable to bind data"})
	}

}
func test(c *gin.Context){
	c.JSON(200,gin.H{"status":"ok"})
}
