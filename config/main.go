package config

import (
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
	"fmt"
)

func init(){
	dsn := "root:@tcp(127.0.0.1:3306)/go-joke?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)
	fmt.Println(err)
	fmt.Println("from init file")
}

func Connect() {
	fmt.Printf("connect to database")
}
