package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goshort struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect"`
	Goshort  string `json:"goshort" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	dsn := "host=172.17.0.2 user=admin password=test dbname=admin port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
