package db

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	// GORM MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Data model
type Data struct {
	UUID uuid.UUID `gorm:"primary_key"`
	Time time.Time
	JSON string `gorm:"type:longtext"`
	XML  string `gorm:"type:longtext"`
}

var db *gorm.DB

// Add info to MySQL
func Add(u uuid.UUID, t time.Time, j string, x string) {
	db.Create(&Data{
		UUID: u,
		Time: t,
		JSON: j,
		XML:  x,
	})
}

// Get info from MySQL
func Get(u uuid.UUID) Data {
	d := new(Data)
	db.Where("uuid = ?", u).First(d)
	return *d
}

func init() {
	var err error
	db, err = gorm.Open("mysql", os.Getenv("KISHOW_DB"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Data{})
}
