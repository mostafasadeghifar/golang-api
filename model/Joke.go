package model
import (
  	"gorm.io/gorm"
	"myjoke/config"
)
var (
	db *gorm.DB
)
type Joke struct {
	Title string `json:"title"`
	Content string `json:"content"`
	gorm.Model
}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Joke{})

}

func (b *Joke) CreateJoke()  *Joke{
	db.Create(&b)
	return b
}