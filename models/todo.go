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
