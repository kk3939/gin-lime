package db

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kk3939/gin-lime/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func Connect(count int) {
	err = godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	up := fmt.Sprintf("%s:%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"))
	dsn := fmt.Sprintf("%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local", up, os.Getenv("MYSQL_DATABASE"))
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	if t := Db.Migrator().HasTable(&entity.Todo{}); t {
		fmt.Println("already Todo table exist. Seeds does not create table.")
		return
	}
	if err := Db.Migrator().CreateTable(&entity.Todo{}); err != nil {
		panic(err)
	}
	var todos entity.ToDos
	for i := 1; i <= 10; i++ {
		todos = append(todos, entity.Todo{
			Name:    fmt.Sprintf("name_%d", i),
			Content: fmt.Sprintf("content_%d", i),
		})
	}
	Db.Create(&todos)
	fmt.Println("create ten todo data!")
}
