package tests_controllers

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDB(models ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect memory database: %v", err)
	}
	if err = db.AutoMigrate(models...); err != nil {
		return nil, fmt.Errorf("failed to migrate model: %v", err)
	}
	return db, nil
}

func dropTable(db *gorm.DB, tables ...interface{}) error {
	if err := db.Migrator().DropTable(tables...); err != nil {
		return fmt.Errorf("failed to drop table: %v", err)
	}
	return nil
}
