package lateReason

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	mds "github.com/kiji/late_reason/models"
	dbInfo "github.com/kiji/late_reason/db"
	"os/exec"
	"fmt"
	"reflect"
)

var db *gorm.DB
func Handler(c *gin.Context) {
	var (
		reason = []mds.LateReason{}
		jsonMap map[string]interface{} = make(map[string]interface{})
		count = 0
	)
	db = dbInfo.DBconnect()
	defer db.Close()
	db.Find(&reason).Count(&count)
	text := reason[rand.Intn(count)].Text
	err := exec.Command("curl", "-X", "POST", "--data-urlencode", "payload={\"text\": \"" + text + "\"}", "web hook api").Run()
	if err != nil {
		fmt.Println(err)
	}
	jsonMap["reason"] = text
	c.JSON(200, jsonMap)
}
