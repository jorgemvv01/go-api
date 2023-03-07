package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		var err error
		dsn := `host=` + os.Getenv("PGHOST") + ` user=` + os.Getenv("PGUSER") + ` password=` + os.Getenv("PGPASSWORD") + `dbname=` + os.Getenv("PGDATABASE") + ` port=` + os.Getenv("PGPORT") + ` sslmode=disable`
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db
}
