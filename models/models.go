package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// Init - initialize the models system
func Init(appDB *gorm.DB) {
	db = appDB
}
