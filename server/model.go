package main

import (
	"time"

	log "github.com/asim/go-micro/v3/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnector struct {
	db *gorm.DB
}

type Danmaku struct {
	ID        uint   `gorm:"primaryKey"`
	ChannelID uint64 `gorm:"index;not null;"`
	Author    string
	Time      float64
	Text      string
	Color     uint32
	Type      uint8 `default:"0"`
	CreatedAt time.Time
}

func InitDB() *DBConnector {
	dsn := "root:qwerty@tcp(10.4.20.15:3306)/olddanmaku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Danmaku{})

	return &DBConnector{db: db}
}

func (db *DBConnector) AddDanmaku(dmk Danmaku) *Danmaku {

	log.Infof("%+v", dmk)
	result := db.db.Create(&dmk)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return &dmk
}

func (db *DBConnector) GetDanmakuListByChannel(channelID uint64) []Danmaku {
	var danmakus []Danmaku
	db.db.Where(&Danmaku{ChannelID: channelID}).Find(&danmakus)

	return danmakus
}

func fromDanmakuPost(channelID uint64, dmk danmaku) Danmaku {
	return Danmaku{
		ChannelID: channelID,
		Author:    dmk.Author,
		Time:      dmk.Time,
		Text:      dmk.Text,
		Color:     dmk.Color,
		Type:      dmk.Type,
	}
}
