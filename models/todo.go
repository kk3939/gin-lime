package models

import (
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/entity"
)

func GetTodos() (entity.ToDos, error) {
	var todos entity.ToDos
	result := db.Db.Find(&todos)
	if err := result.Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodo(id string) (*entity.Todo, error) {
	var todo entity.Todo
	result := db.Db.First(&todo, "ID = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func CreateTodo(td entity.Todo) (*entity.Todo, error) {
	result := db.Db.Create(&td)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &td, nil
}

func UpdateTodo(td entity.Todo) (*entity.Todo, error) {
	result := db.Db.Save(&td)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &td, nil
}

func DeleteTodo(td entity.Todo) error {
	result := db.Db.Delete(&td)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
