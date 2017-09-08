package main

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	mds "github.com/kiji/late_reason/models"
)


func main(){
	db, err := gorm.Open("postgres","user=go_user dbname=ratereason sslmode=disable")
	if err != nil {
		log.Print("コネクションエラー")
	}
	// go run db/migrate.go select * from products
	db.AutoMigrate(&mds.LateReason{})
	defer db.Close()
	//lateReason := &mds.LateReason{Title: "勤怠連絡", Text:"すいません、寝坊しました。10:30までに出社致します。"}
	//lateReason2 := &mds.LateReason{Title: "勤怠連絡", Text:"おはようございます。山手線が止ってしまい遅刻するかもしれません。"}
	//lateReason3 := &mds.LateReason{Title: "勤怠連絡", Text:"トイレ混んでて5分ほど遅れます"}
	//lateReason4 := &mds.LateReason{Title: "勤怠連絡", Text:"おはようございます。トイレ行ってたら長引いてしまい、5分ほど遅刻しそうです。"}
	//lateReason5 := &mds.LateReason{Title: "勤怠連絡", Text:"おはようございます。すいません、考え事をしていたら寝坊しました。10:00までに出社致します。"}
	//db.Create(lateReason)
	//db.Create(lateReason2)
	//db.Create(lateReason3)
	//db.Create(lateReason4)
	//db.Create(lateReason5)
}