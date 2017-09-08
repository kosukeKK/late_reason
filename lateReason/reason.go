package lateReason

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	mds "github.com/kiji/late_reason/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"math/rand"
)

var db *gorm.DB
func Handler(c *gin.Context) {
	db, err := gorm.Open("postgres","user=go_user dbname=ratereason sslmode=disable")
	if err != nil {
		log.Print("no connet")
	}
	defer db.Close()
	var (
		reason = []mds.LateReason{}
		jsonMap map[string]interface{} = make(map[string]interface{})
		count = 0
	)

	db.Find(&reason).Count(&count)
	jsonMap["reason"] = reason[rand.Intn(count)].Text
	c.JSON(200, jsonMap)
}
