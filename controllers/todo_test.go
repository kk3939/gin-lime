package controllers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/entity"
	"github.com/kk3939/gin-lime/server"
)

// this is intgration test
func Test_getTodos(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Error("mock is failed")
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	todo := entity.Todo{
		Id:      1,
		Name:    "test",
		Content: "test",
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todos`")).WithArgs().WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Name",
			"Content",
		}).AddRow(todo.Id, todo.Name, todo.Content))

	r := server.GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todo", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("invalid request")
	}
	res := w.Result()
	defer res.Body.Close()
	if body, err := io.ReadAll(res.Body); err != nil {
		t.Errorf("can not read res body. %v", err)
	} else {
		var td entity.ToDos
		if err := json.Unmarshal(body, &td); err != nil {
			t.Errorf("can not map res json to struct. %v", err)
		}
		if todo.Id != td[0].Id {
			fmt.Println(td[0].Id)
			t.Error("id is not expected value")
		}
		if todo.Name != td[0].Name {
			t.Error("Name is not expected value")
		}
		if todo.Content != td[0].Content {
			t.Error("Content is not expected value")
		}
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
		Name:    "test",
		Content: "test",
	}
	sql := "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Name",
			"Content",
		}).AddRow(todo.Id, todo.Name, todo.Content))

	r := server.GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todo/1", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("invalid request")
	}
	res := w.Result()
	defer res.Body.Close()
	if body, err := io.ReadAll(res.Body); err != nil {
		t.Errorf("can not read res body. %v", err)
	} else {
		var td entity.Todo
		if err := json.Unmarshal(body, &td); err != nil {
			t.Errorf("can not map res json to struct. %v", err)
		}
		if todo.Id != td.Id {
			fmt.Println(td.Id)
			t.Error("id is not expected value")
		}
		if todo.Name != td.Name {
			t.Error("Name is not expected value")
		}
		if todo.Content != td.Content {
			t.Error("Content is not expected value")
		}
	}
}
