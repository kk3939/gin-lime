package db

import (
	"fmt"
	"time"

	"github.com/kk3939/gin-lime/entity"
)

type Seeds struct{}

func (s *Seeds) ToDo() {
	fmt.Println("Todo seeding is starting...")
	time.Sleep(time.Second)

	fmt.Println("Checking whether todo table exist...")
	time.Sleep(time.Second)

	if t := db.Migrator().HasTable(&entity.Todo{}); t {
		fmt.Println("Todo table already exist. Seeds don't create todo table.")
		return
	}

	fmt.Println("Todo table doesn't exist. Seeds is creating todo table...")
	time.Sleep(time.Second)
	if err := db.Migrator().CreateTable(&entity.Todo{}); err != nil {
		panic(err)
	}
	fmt.Println("Created todo table!!")

	fmt.Println("Creating todo data...")
	time.Sleep(time.Second)
	var todos entity.ToDos
	for i := 1; i <= 10; i++ {
		todos = append(todos, entity.Todo{
			Name:    fmt.Sprintf("name_%d", i),
			Content: fmt.Sprintf("content_%d", i),
		})
	}
	if err := db.Create(&todos).Error; err != nil {
		panic(err)
	}
	fmt.Println("Created todo data!!")
}
