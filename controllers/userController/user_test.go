package userController_test

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

func Test_getUsers(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	mock_user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}

	t.Run("Get records from users", func(t *testing.T) {
		sql := "SELECT * FROM `users`"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Email",
				"Password",
			}).AddRow(mock_user.Id, mock_user.Email, mock_user.Password))
	})

	t.Run("Request GET /user", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()

		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var users entity.Users
			if err := json.Unmarshal(body, &users); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if mock_user.Id != users[0].Id {
				t.Error("Id is not expected value.")
			}
			if mock_user.Email != users[0].Email {
				t.Error("Email is not expected value.")
			}
			if mock_user.Password != users[0].Password {
				t.Error("Password is not expected value")
			}
		}
	})
}

func Test_getUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	mock_user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}
	t.Run("Get the record whose id is 1 from DB", func(t *testing.T) {
		sql := "SELECT * FROM `users` WHERE Id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Email",
				"Password",
			}).AddRow(mock_user.Id, mock_user.Email, mock_user.Password))
	})

	t.Run("Request GET /user/1", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/user/1", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()
		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var user entity.User
			if err := json.Unmarshal(body, &user); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if mock_user.Id != user.Id {
				t.Error("Id is not expected value")
			}
			if mock_user.Email != user.Email {
				t.Error("Email is not expected value")
			}
			if mock_user.Password != user.Password {
				t.Error("Password is not expected value")
			}
		}
	})
}

func Test_createUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	mock_user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}

	t.Run("Inset record into DB", func(t *testing.T) {
		sql := "INSERT INTO `users` (`created_at`,`updated_at`,`email`,`password`) VALUES (?,?,?,?)"
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), mock_user.Email, mock_user.Password).WillReturnResult(sqlmock.NewResult(int64(mock_user.Id), 1))
		mock.ExpectCommit()
	})

	t.Run("Request POST /user", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		type RequestBody struct {
			Email    string
			Password string
		}
		rqb := RequestBody{
			Email:    mock_user.Email,
			Password: mock_user.Password,
		}
		if data, err := json.Marshal(rqb); err != nil {
			t.Errorf("Can not marshal json. %v", err)
		} else {
			req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(data))
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
			var user entity.User
			if err := json.Unmarshal(body, &user); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if mock_user.Id != user.Id {
				t.Error("id is not expected value")
			}
			if mock_user.Email != user.Email {
				t.Error("Email is not expected value")
			}
			if mock_user.Password != user.Password {
				t.Error("Password is not expected value")
			}
		}
	})
}

func Test_updateUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	mock_user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}

	var sql string
	t.Run("Get the record whose id is 1 from DB", func(t *testing.T) {
		sql = "SELECT * FROM `users` WHERE Id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Email",
				"Password",
			}).AddRow(mock_user.Id, mock_user.Email, mock_user.Password))
	})
	t.Run("Update the record whose id is 1 from DB", func(t *testing.T) {
		sql = "UPDATE `users` SET `updated_at`=?,`email`=?,`password`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?"
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), mock_user.Email, mock_user.Password, mock_user.Id).WillReturnResult(sqlmock.NewResult(int64(mock_user.Id), 1))
		mock.ExpectCommit()
	})

	t.Run("Request PUT /user/1", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		type Rq struct {
			Email    string
			Password string
		}
		rqb := Rq{
			Email:    mock_user.Email,
			Password: mock_user.Password,
		}
		if data, err := json.Marshal(rqb); err != nil {
			t.Errorf("Can not marshal json. %v", err)
		} else {
			req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(data))
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
			var user entity.User
			if err := json.Unmarshal(body, &user); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if mock_user.Id != user.Id {
				t.Error("Id is not expected value")
			}
			if mock_user.Email != user.Email {
				t.Error("Email is not expected value")
			}
			if mock_user.Password != user.Password {
				t.Error("Password is not expected value")
			}
		}
	})
}

func Test_deleteUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	mock_user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}

	var sql string
	t.Run("Get the record whose id is 1 from DB", func(t *testing.T) {
		sql = "SELECT * FROM `users` WHERE Id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs().WillReturnRows(
			sqlmock.NewRows([]string{
				"Id",
				"Email",
				"Password",
			}).AddRow(mock_user.Id, mock_user.Email, mock_user.Password))
	})

	t.Run("Using `UPDATE`, delete record whose id is 1", func(t *testing.T) {
		sql = "UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL"
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), mock_user.Id).WillReturnResult(sqlmock.NewResult(int64(mock_user.Id), 1))
		mock.ExpectCommit()
	})

	t.Run("Request DELETE /user/1", func(t *testing.T) {
		r := server.GetRouter()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/user/1", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Invalid request. Expected http code is %d but actual code is %d", http.StatusOK, w.Code)
		}

		res := w.Result()
		defer res.Body.Close()
		if body, err := io.ReadAll(res.Body); err != nil {
			t.Errorf("Can not read response body. %v", err)
		} else {
			var user entity.User
			if err := json.Unmarshal(body, &user); err != nil {
				t.Errorf("Can not unmarshal json of response body to struct. %v", err)
			}
			if mock_user.Id != user.Id {
				t.Error("Id is not expected value")
			}
			if user.Email != "" {
				t.Error("Email is not expected value")
			}
			if user.Password != "" {
				t.Error("Password is not expected value")
			}
			if mock_user.DeletedAt.Valid {
				t.Error("DeletedAt is not expected value")
			}
		}
	})
}
