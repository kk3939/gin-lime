package db

import (
	"fmt"
	"os"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func SetDB(argDB *gorm.DB) {
	db = argDB
}

func Connect(count int) {
	err = godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	userPass := fmt.Sprintf("%s:%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"))
	dsn := fmt.Sprintf("%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local", userPass, os.Getenv("MYSQL_DATABASE"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// https://gorm.io/ja_JP/docs/logger.html
		Logger: logger.Default.LogMode(logger.Info),
	})
	// If can not connect to mysql, retry.
	if err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			Connect(count)
			return
		}
		panic(err)
	}
}

func Mock_DB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})
	return gormDb, mock, err
}

func ConnectTestDB(name string) (*gorm.DB, error) {
	userPass := fmt.Sprintf("%s:%s", "test_user", "test_password")
	dsn := fmt.Sprintf("%s@tcp(db-test)/%s?charset=utf8mb4&parseTime=True&loc=Local", userPass, "gin-lime-test")
	txdb.Register(name, "mysql", dsn)
	dialector := mysql.New(mysql.Config{
		DriverName: name,
		DSN:        dsn,
	})

	return gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
