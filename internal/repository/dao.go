package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type DAO interface {
	NewPostQuery() PostQuery
	NewTagQuery() TagQuery
}

type dao struct {
	pgdb *gorm.DB
}

func NewDAO(db *gorm.DB) DAO {
	return &dao{
		pgdb: db,
	}
}

func SetupDB() *gorm.DB {
	var dbHost = os.Getenv("DB_HOST")
	var dbName = os.Getenv("POSTGRES_DB")
	var dbUsername = os.Getenv("POSTGRES_USER")
	var dbPassword = os.Getenv("POSTGRES_PASSWORD")
	var dbPort = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)
	fmt.Println(dsn)

	var db *gorm.DB
	var err error
	maxAttempts := 6
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("[PostgreSQL] Connected to DB instance")
			sqlDB, err := db.DB()
			if err != nil {
				panic("Failed to get DB instance: " + err.Error())
			}
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			return db
		}
		log.Printf("Attempt %d: failed to connect to database: %s", attempts, err.Error())
		time.Sleep(5 * time.Second)
	}
	panic("failed to connect to database after several attempts: " + err.Error())
}

func (d *dao) NewPostQuery() PostQuery {
	return NewPostQuery(d.pgdb)
}
func (d *dao) NewTagQuery() TagQuery {
	return NewTagQuery(d.pgdb)
}
