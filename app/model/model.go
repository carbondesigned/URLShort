package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type URLShort struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	URLShort string `json:"urlshort" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {

	dsn := "host=0.0.0.0 user=admin password=test dbname=admin port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&URLShort{})
	if err != nil {
		fmt.Println(err)
	}
}
