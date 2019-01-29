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
	UUID   uuid.UUID `gorm:"primary_key"`
	Time   time.Time
	JSONM  string `gorm:"type:text"`
	JSONP  string `gorm:"type:text"`
	RawXML string `gorm:"type:text"`
}

var db *gorm.DB

// Add info to MySQL
func Add(u uuid.UUID, t time.Time, jm string, jp string, x string) {
	db.Create(&Data{
		UUID:   u,
		Time:   t,
		JSONM:  jm,
		JSONP:  jp,
		RawXML: x,
	})
}

func init() {
	var err error
	db, err = gorm.Open("mysql", os.Getenv("KISHOW_DB"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Data{})
}
