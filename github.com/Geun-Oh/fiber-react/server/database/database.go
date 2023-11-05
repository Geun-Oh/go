package database

import (
	"log"
	"os"

	"github.com/Geun-Oh/fiber-react/server/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	// db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("connection successed")
	// defer db.Close()

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Error),
	// })
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed")
	}

	log.Println("Connection Successful")

	db.AutoMigrate(new(model.Blog))

	DBConn = db
}
