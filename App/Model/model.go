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

	// DSN standds for data source name which consisit information like host,user,pass,dbname port sslmode

	dsn := "host=172.17.0.2 user=admin password=test dbname=admin port=5432 sslmode=disable"

	var err error

	// gorm.Open is a function provided by GORM for creating a database connection.
	// &gorm.Config{} is used to configure the behavior of the GORM instance. In this case, it uses the default configuration.
	// postgres.Open(dsn) specifies that the PostgreSQL driver should be used, and it takes the DSN string dsn as an argument to establish the connection.

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
