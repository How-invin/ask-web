package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type Table struct {
	gorm.Model
	Name           string
	Telephone      string
	Gender         string
	Birthday       string
	Education      string
	Nationality    string
	Identity       string
	Register_place string
	Birth_place    string
	Reside_place   string
	Marital_status string
}

func main() {
	db, _ := gorm.Open(sqlite.Open("ask.db"), &gorm.Config{})

	db.AutoMigrate(&Table{})

	//db.Create(&Table{
	//Name:               "虚拟人物",
	//Telephone:          "1333444555",
	//Gender:             "male",
	//Birthday:           "2021-01-01",
	//Nationality:        "汉族",
	//Identity:           "33399200001010000",
	//Education:           "大学",
	//Registration_place: "山西大同",
	//Birth_place:        "山西大同",
	//Reside_place:       "山西大同",
	//Marital_status:     "未婚",
	//})

	r := gin.Default()
	r.StaticFS("/ask", http.Dir("./static"))
	r.POST("/ask", func(c *gin.Context) {
		db.Create(&Table{
			Name:           c.PostForm("name"),
			Telephone:      c.PostForm("telephone"),
			Gender:         c.PostForm("gender"),
			Birthday:       c.PostForm("birthday"),
			Nationality:    c.PostForm("nationality"),
			Education:      c.PostForm("education"),
			Identity:       c.PostForm("identity"),
			Register_place: c.PostForm("register-place"),
			Birth_place:    c.PostForm("birth-place"),
			Reside_place:   c.PostForm("reside-place"),
			Marital_status: c.PostForm("marital-status"),
		})

		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run()
}
