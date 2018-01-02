package dbinfo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func DBconnect() *gorm.DB {
	db, err := gorm.Open("mysql","root:@/late_reason?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Print("no connet", err)
	}
	return db
}
