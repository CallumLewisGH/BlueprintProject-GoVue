package database

import (
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	GormDB *gorm.DB
}

var (
	dbInstance  *Database
	once        sync.Once
	testConnStr string
)

func SetTestModeConnectionString(connStr string) {
	testConnStr = connStr

}
func GetDatabase() *Database {
	once.Do(func() {
		dbInstance = &Database{}

		environment := os.Getenv("ENVIRONMENT")

		switch environment {
		case "dev", "development":
			dbInstance.InitialiseDevDB()

		case "test", "testing":
			dbInstance.InitialiseTestDB(testConnStr)

		default:
			dbInstance.InitialiseProdDB()
		}
	})
	return dbInstance
}

func (db *Database) InitialiseTestDB(connStr string) {
	log.Printf("Connecting to Database with GORM...")

	var gormDB *gorm.DB

	for i := range 5 {
		if i > 0 {
			time.Sleep(1 * time.Second)
			log.Printf("Retrying connection to database, attempt %d", i+1)
		}
		var err error

		gormDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			PrepareStmt: true,
		})
		if err == nil {
			db.GormDB = gormDB
			log.Printf("GORM Database Connection Succeeded")
			return
		}
		log.Printf("Failed to open database: %v", err)
	}

	if gormDB == nil {
		log.Fatalf("Failed to connect to database after 5 attempts")
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	// Capped conservatively: if you deploy this across multiple
	// autoscaled instances, each with its own pool, and your Postgres has
	// a typical max_connections=100, an uncapped pool per instance risks
	// exhausting that under load and failing every instance at once.
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	db.GormDB = gormDB
	log.Printf("GORM Database Connection Succeeded")
}

func (db *Database) InitialiseDevDB() {
	log.Printf("Connecting to Database with GORM...")

	var gormDB *gorm.DB

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")

	for i := range 5 {
		if i > 0 {
			time.Sleep(1 * time.Second)
			log.Printf("Retrying connection to database, attempt %d", i+1)
		}
		var err error

		gormDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			PrepareStmt: true,
		})
		if err == nil {
			db.GormDB = gormDB
			log.Printf("GORM Database Connection Succeeded")
			return
		}
		log.Printf("Failed to open database: %v", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	// Capped conservatively: if you deploy this across multiple
	// autoscaled instances, each with its own pool, and your Postgres has
	// a typical max_connections=100, an uncapped pool per instance risks
	// exhausting that under load and failing every instance at once.
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	db.GormDB = gormDB
	log.Printf("GORM Database Connection Succeeded")
}

func (db *Database) InitialiseProdDB() {
	log.Printf("Connecting to Database with GORM...")
	var gormDB *gorm.DB

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")

	for i := range 5 {
		if i > 0 {
			time.Sleep(1 * time.Second)
			log.Printf("Retrying connection to database, attempt %d", i+1)
		}
		var err error

		gormDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			PrepareStmt: true,
		})
		if err == nil {
			db.GormDB = gormDB
			log.Printf("GORM Database Connection Succeeded")
			return
		}
		log.Printf("Failed to open database: %v", err)
	}

	if gormDB == nil {
		log.Fatalf("Failed to connect to database after 5 attempts")
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	// Capped conservatively: if you deploy this across multiple
	// autoscaled instances, each with its own pool, and your Postgres has
	// a typical max_connections=100, an uncapped pool per instance risks
	// exhausting that under load and failing every instance at once.
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	db.GormDB = gormDB
	log.Printf("GORM Database Connection Succeeded")
}

func (db *Database) CheckDatabaseHealth() error {
	sqlDB, err := db.GormDB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (db *Database) CloseDatabase() error {
	sqlDB, err := db.GormDB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *Database) GetGormDatabase() *gorm.DB {
	return db.GormDB
}

func (db *Database) RunMigrations() error {
	for _, model := range GetModelRegistry().models {
		if err := GetDatabase().GetGormDatabase().AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) ClearAllTables() error {
	for _, model := range GetModelRegistry().models {
		if err := db.GormDB.Unscoped().Where("1 = 1").Delete(model).Error; err != nil {
			return err
		}
	}
	return nil
}
