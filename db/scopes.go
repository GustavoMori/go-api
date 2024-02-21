package db

import "gorm.io/gorm"

func NotBeRonaldinho(db *gorm.DB) *gorm.DB {
	return db.Where("name != 'Ronaldinho'")
}
