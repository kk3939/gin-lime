package controllers_test

import (
	"bytes"
	"encoding/json"
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

func Test_createTodo(t *testing.T) {
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

	sql := "INSERT INTO `todos` (`created_at`,`updated_at`,`name`,`content`) VALUES (?,?,?,?)"
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), todo.Name, todo.Content).WillReturnResult(sqlmock.NewResult(int64(todo.Id), 1))
	mock.ExpectCommit()

	r := server.GetRouter()
	w := httptest.NewRecorder()
	type Rq struct {
		Name    string
		Content string
	}
	rqb := Rq{
		Name:    todo.Name,
		Content: todo.Content,
	}
	if data, err := json.Marshal(rqb); err != nil {
		t.Errorf("can not marshal json. %v", err)
	} else {
		req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(data))
		r.ServeHTTP(w, req)
	}
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

func Test_updateTodo(t *testing.T) {
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

	var sql string
	sql = "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Name",
			"Content",
		}).AddRow(todo.Id, todo.Name, todo.Content))

	sql = "UPDATE `todos` SET `updated_at`=?,`name`=?,`content`=? WHERE `todos`.`deleted_at` IS NULL AND `id` = ?"
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), todo.Name, todo.Content, todo.Id).WillReturnResult(sqlmock.NewResult(int64(todo.Id), 1))
	mock.ExpectCommit()

	r := server.GetRouter()
	w := httptest.NewRecorder()
	type Rq struct {
		Name    string
		Content string
	}
	rqb := Rq{
		Name:    todo.Name,
		Content: todo.Content,
	}
	if data, err := json.Marshal(rqb); err != nil {
		t.Errorf("can not marshal json. %v", err)
	} else {
		req, _ := http.NewRequest("PUT", "/todo/1", bytes.NewBuffer(data))
		r.ServeHTTP(w, req)
	}
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

func Test_deleteTodo(t *testing.T) {
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

	var sql string
	sql = "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Name",
			"Content",
		}).AddRow(todo.Id, todo.Name, todo.Content))

	sql = "UPDATE `todos` SET `deleted_at`=? WHERE `todos`.`id` = ? AND `todos`.`deleted_at` IS NULL"
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), todo.Id).WillReturnResult(sqlmock.NewResult(int64(todo.Id), 1))
	mock.ExpectCommit()

	r := server.GetRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/todo/1", nil)
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
			t.Error("id is not expected value")
		}
		if td.Name != "" {
			t.Error("Name is not expected value")
		}
		if td.Content != "" {
			t.Error("Content is not expected value")
		}
		if todo.DeletedAt.Valid {
			t.Error("DeletedAt is not expected value")
		}
	}
}
