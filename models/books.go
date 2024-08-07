package models

import "gorm.io/gorm"

type Books struct {
	Id        uint    `gorm:"primary_key; auto_increment" json:"id"`
	Author    *string `gorm:"type:varchar(255)"`
	Title     *string `gorm:"type:varchar(255)"`
	Publisher *string `gorm:"type:varchar(255)"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	if err != nil {
		return err
	}
	return nil
}
