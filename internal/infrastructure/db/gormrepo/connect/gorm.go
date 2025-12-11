package dbconnect

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	var currentDB, currentSchema, currentUser string

	err = db.Raw(`
    SELECT current_database(), current_schema(), current_user
`).Row().Scan(&currentDB, &currentSchema, &currentUser)

	log.Println("======================================")
	log.Println(" Connected to DB      :", currentDB)
	log.Println(" Connected schema     :", currentSchema)
	log.Println(" Connected as user    :", currentUser)
	log.Println("======================================")

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	log.Println("✅ Connected to PostgreSQL")
	return db, nil
}
