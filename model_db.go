package main

type MessageT struct {
	Id       int
	Text     string
	FromUser int
	FromFlow int
}

type UserT struct {
	Uuid         int `gorm:"primaryKey"`
	Login        string
	Password     string
	HashPassword string
	Username     string
	IsBot        bool   `gorm:"column:isbot"`
	AuthId       string `gorm:"column:authid"`
	Avatar       string
	Bio          string
	Salt         string
	Key          string
	Message      MessageT `gorm:"foreignKey:FromUser;references:Uuid"`
}

type FlowT struct {
	FlowId      int `gorm:"primaryKey"`
	TimeCreated int
	FlowType    string
	Title       string
	Info        string
	Message     MessageT `gorm:"foreignKey:FromFlow;references:FlowId"`
}

func (UserT) TableName() string {
	return "UserConfig"
}

func (FlowT) TableName() string {
	return "Flows"
}

func (MessageT) TableName() string {
	return "Messages"
}
