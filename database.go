package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func databaseInit() {
	var err error
	db, err = gorm.Open(sqlite.Open("CrawlTalkServer.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserT{}, &FlowT{}, &MessageT{})
	db.Create(&UserT{Uuid: 42, Login: "admin", Password: "admin"})

}
