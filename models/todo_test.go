package models_test

import (
	"regexp"
	"strconv"
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
		t.Error(err)
	}
}

func Test_getTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	todo := entity.Todo{
		Id:      1,
		Name:    "test_name",
		Content: "test_content",
	}
	sql := "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(strconv.FormatUint(uint64(todo.Id), 10)).WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Name",
			"Content",
		}).AddRow(1, "test_name", "test_content"))
	if _, err := models.GetTodo("1"); err != nil {
		t.Error(err)
	}
}

func Test_createTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	todo := entity.Todo{
		Id:      1,
		Name:    "test_name",
		Content: "test_content",
	}
	sql := "INSERT INTO `todos` (`created_at`,`updated_at`,`name`,`content`) VALUES (?,?,?,?)"
	// https://github.com/DATA-DOG/go-sqlmock/issues/117
	// https://github.com/DATA-DOG/go-sqlmock/issues/224
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), todo.Name, todo.Content).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	if err := models.CreateTodo(&todo); err != nil {
		t.Error(err)
	}
}

func Test_updateToDo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	todo := entity.Todo{
		Id:      1,
		Name:    "test_name",
		Content: "test_content",
	}
	sql := "UPDATE `todos` SET `updated_at`=?,`name`=?,`content`=? WHERE `todos`.`deleted_at` IS NULL AND `id` = ?"
	// https://github.com/DATA-DOG/go-sqlmock/issues/117
	// https://github.com/DATA-DOG/go-sqlmock/issues/224
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), todo.Name, todo.Content, todo.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	if err := models.UpdateTodo(&todo); err != nil {
		t.Error(err)
	}
}

func Test_deleteToDo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	todo := entity.Todo{
		Id:      1,
		Name:    "test_name",
		Content: "test_content",
	}
	sql := "UPDATE `todos` SET `deleted_at`=? WHERE `todos`.`id` = ? AND `todos`.`deleted_at` IS NULL"
	// https://github.com/DATA-DOG/go-sqlmock/issues/117
	// https://github.com/DATA-DOG/go-sqlmock/issues/224
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), todo.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	if err := models.DeleteTodo(&todo); err != nil {
		t.Error(err)
	}
}
