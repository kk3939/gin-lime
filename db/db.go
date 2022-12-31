package db

import (
	"fmt"
	"os"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/kk3939/gin-lime/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

func Seeds() {
	if t := db.Migrator().HasTable(&entity.Todo{}); t {
		fmt.Println("already Todo table exist. Seeds does not create table.")
		return
	}
	if err := db.Migrator().CreateTable(&entity.Todo{}); err != nil {
		panic(err)
	}
	var todos entity.ToDos
	for i := 1; i <= 10; i++ {
		todos = append(todos, entity.Todo{
			Name:    fmt.Sprintf("name_%d", i),
			Content: fmt.Sprintf("content_%d", i),
		})
	}
	db.Create(&todos)
	fmt.Println("create ten todo data!")
}

func Mock_DB() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})
	return gormDb, mock, err
}
