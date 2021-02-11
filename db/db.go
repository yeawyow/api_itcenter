package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
func GetDB() *gorm.DB{
	return db
}

func SetupDB(){

	dsn := "yeaw:481725209@tcp(202.80.228.46:3306)/airoffice?charset=utf8mb4&parseTime=True&loc=Local"
	database,err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //nochang table name
		},
	})

	//database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil{
          panic("fail to connet databese")
	}
  
	/*database.AutoMigrate(&model.Employee{})*/
	
	db=database
}