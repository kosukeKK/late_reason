package lateReason

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	mds "github.com/kiji/late_reason/models"
	dbInfo "github.com/kiji/late_reason/db"
)

var db *gorm.DB
func Handler(c *gin.Context) {
	db = dbInfo.DBconnect()
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
