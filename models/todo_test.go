package models_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kk3939/gin-lime/db"
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
