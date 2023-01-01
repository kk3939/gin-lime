package models

import (
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/entity"
)

func GetTodos() (entity.ToDos, error) {
	var todos entity.ToDos
	db := db.GetDB()
	result := db.Find(&todos)
	if err := result.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodo(id string) (*entity.Todo, error) {
	var todo entity.Todo
	db := db.GetDB()
	result := db.First(&todo, "Id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func CreateTodo(td *entity.Todo) error {
	db := db.GetDB()
	result := db.Select("Name", "Content").Create(&td)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(td *entity.Todo) error {
	db := db.GetDB()
	result := db.Model(&td).Updates(entity.Todo{Name: td.Name, Content: td.Content})
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(td *entity.Todo) error {
	db := db.GetDB()
	result := db.Delete(&td)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
