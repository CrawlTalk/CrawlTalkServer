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

	err = db.AutoMigrate(&UserT{}, &FlowT{}, &MessageT{})
	if err != nil {
		panic("failed to migrate database")
	}
	db.Create(&UserT{Uuid: 42, Login: "admin", Password: "admin"})

}
