package todoController_test

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

func Test_getTodos(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	test := entity.Todo{
		Id:      1,
		Name:    "test",
		Content: "test",
	}

	t.Run("Get records from DB", func(t *testing.T) {
		sql := "SELECT * FROM `todos`"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Name",
				"Content",
			}).AddRow(test.Id, test.Name, test.Content))
	})

	t.Run("Request GET /todo", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todo", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()

		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var td entity.ToDos
			if err := json.Unmarshal(body, &td); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if test.Id != td[0].Id {
				t.Error("Id is not expected value.")
			}
			if test.Name != td[0].Name {
				t.Error("Name is not expected value.")
			}
			if test.Content != td[0].Content {
				t.Error("Content is not expected value")
			}
		}
	})
}

func Test_getTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	test := entity.Todo{
		Id:      1,
		Name:    "test",
		Content: "test",
	}
	t.Run("Get the record whose id is 1 from DB", func(t *testing.T) {
		sql := "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Name",
				"Content",
			}).AddRow(test.Id, test.Name, test.Content))
	})

	t.Run("Request GET /todo/1", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/todo/1", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()
		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var td entity.Todo
			if err := json.Unmarshal(body, &td); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if test.Id != td.Id {
				t.Error("Id is not expected value")
			}
			if test.Name != td.Name {
				t.Error("Name is not expected value")
			}
			if test.Content != td.Content {
				t.Error("Content is not expected value")
			}
		}
	})
}

func Test_createTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	test := entity.Todo{
		Id:      1,
		Name:    "test",
		Content: "test",
	}

	t.Run("Inset record into DB", func(t *testing.T) {
		sql := "INSERT INTO `todos` (`created_at`,`updated_at`,`name`,`content`) VALUES (?,?,?,?)"
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), test.Name, test.Content).WillReturnResult(sqlmock.NewResult(int64(test.Id), 1))
		mock.ExpectCommit()
	})

	t.Run("Request POST /todo", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		type RequestBody struct {
			Name    string
			Content string
		}
		rqb := RequestBody{
			Name:    test.Name,
			Content: test.Content,
		}
		if data, err := json.Marshal(rqb); err != nil {
			t.Errorf("Can not marshal json. %v", err)
		} else {
			req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(data))
			r.ServeHTTP(w, req)
		}

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()
		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var td entity.Todo
			if err := json.Unmarshal(body, &td); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if test.Id != td.Id {
				t.Error("id is not expected value")
			}
			if test.Name != td.Name {
				t.Error("Name is not expected value")
			}
			if test.Content != td.Content {
				t.Error("Content is not expected value")
			}
		}
	})
}

func Test_updateTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	todo := entity.Todo{
		Id:      1,
		Name:    "test",
		Content: "test",
	}

	var sql string
	t.Run("Get the record whose id is 1 from DB", func(t *testing.T) {
		sql = "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Name",
				"Content",
			}).AddRow(todo.Id, todo.Name, todo.Content))
	})
	t.Run("Update the record whose id is 1 from DB", func(t *testing.T) {
		sql = "UPDATE `todos` SET `updated_at`=?,`name`=?,`content`=? WHERE `todos`.`deleted_at` IS NULL AND `id` = ?"
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), todo.Name, todo.Content, todo.Id).WillReturnResult(sqlmock.NewResult(int64(todo.Id), 1))
		mock.ExpectCommit()
	})

	t.Run("Request PUT /todo/1", func(t *testing.T) {
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
			t.Errorf("Can not marshal json. %v", err)
		} else {
			req, _ := http.NewRequest("PUT", "/todo/1", bytes.NewBuffer(data))
			r.ServeHTTP(w, req)
		}

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()
		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var td entity.Todo
			if err := json.Unmarshal(body, &td); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if todo.Id != td.Id {
				t.Error("Id is not expected value")
			}
			if todo.Name != td.Name {
				t.Error("Name is not expected value")
			}
			if todo.Content != td.Content {
				t.Error("Content is not expected value")
			}
		}
	})
}

func Test_deleteTodo(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	todo := entity.Todo{
		Id:      1,
		Name:    "test",
		Content: "test",
	}

	var sql string
	t.Run("Get the record whose id is 1 from DB", func(t *testing.T) {
		sql = "SELECT * FROM `todos` WHERE Id = ? AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Name",
				"Content",
			}).AddRow(todo.Id, todo.Name, todo.Content))
	})

	t.Run("Using `UPDATE`, delete record whose id is 1", func(t *testing.T) {
		sql = "UPDATE `todos` SET `deleted_at`=? WHERE `todos`.`id` = ? AND `todos`.`deleted_at` IS NULL"
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), todo.Id).WillReturnResult(sqlmock.NewResult(int64(todo.Id), 1))
		mock.ExpectCommit()
	})

	t.Run("Request DELETE /todo/1", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/todo/1", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()
		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var td entity.Todo
			if err := json.Unmarshal(body, &td); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if todo.Id != td.Id {
				t.Error("Id is not expected value")
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
	})
}
