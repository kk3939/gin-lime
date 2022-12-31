package models_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/entity"
	"github.com/kk3939/gin-lime/models"
)

func Test_getTodos(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todos`")).WithArgs().WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Name",
			"Content",
		}))
	if _, err := models.GetTodos(); err != nil {
		t.Errorf("GetTodos is failed. %v", err)
	}
}

func Test_getTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)
	sql := "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
	todo := entity.Todo{}
	var rows []string
	for i := 0; i < reflect.ValueOf(todo).NumField(); i++ {
		rows = append(rows, reflect.TypeOf(todo).Field(i).Name)
	}
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(1).WillReturnRows(sqlmock.NewRows(rows).AddRow(todo.Id, todo.CreatedAt, todo.UpdatedAt, todo.DeletedAt, todo.Name, todo.Content))
	if _, err := models.GetTodo(1); err != nil {
		t.Errorf("GetTodo is failed. %v", err)
	}
}

func Test_createTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)
	sql := "INSERT INTO `todos` (`created_at`,`updated_at`,`deleted_at`,`name`,`content`) VALUES (?,?,?,?,?)"
	mock.ExpectBegin()
	todo := entity.Todo{}
	var rows []string
	for i := 0; i < reflect.ValueOf(todo).NumField(); i++ {
		rows = append(rows, reflect.TypeOf(todo).Field(i).Name)
	}
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs("test", "test").WillReturnRows(sqlmock.NewRows(rows).AddRow(todo.Id, todo.CreatedAt, todo.UpdatedAt, todo.DeletedAt, todo.Name, todo.Content))
	mock.ExpectCommit()
	if err := models.CreateTodo(todo); err != nil {
		t.Errorf("CreateTodo is failed. %v", err)
	}
}
