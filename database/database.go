package database

import (
    "log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/TasteLabGroup?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed")
	}

	log.Println("Connnection successfully")

    // Configure the connection pool
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("failed to get database instance: %v", err)
    }

    // Set the maximum number of idle connections
    sqlDB.SetMaxIdleConns(10)
    // Set the maximum number of open connections
    sqlDB.SetMaxOpenConns(100)
    // Set the maximum lifetime of a connection
    sqlDB.SetConnMaxLifetime(0) // 0 means connections are not closed

	DBConn = db
}

func Migrate(db *gorm.DB, models ...interface{}) error {
    return db.AutoMigrate(models...)
}
