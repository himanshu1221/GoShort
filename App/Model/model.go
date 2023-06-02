package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type GoShort struct {
	ID       uint64 `json:"id" gorm:"primaryKey"` //Go Object Relational Mapping,
	Redirect string `json:"redirect" gorm:"not null"`
	GoShort  string `json:"goshort" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	//IP Address of postgres sql running using docker
	//DSN represent data source name
	//It is used to establish a conection between postgres database
	//SSL (Secure Sockets Layer) is a cryptographic protocol used to secure communications over a network, such as the internet. In the context of the sslmode parameter in the PostgreSQL connection string, it determines how the Go application will handle SSL/TLS encryption when connecting to the PostgreSQL database.
	dsn := "host=172.17.0.2 user=admin password=test dbname=admin port=5432 sslmode=disable"
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
