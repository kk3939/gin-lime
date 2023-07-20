package userModel_test

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/entity"
	"github.com/kk3939/gin-lime/models/userModel"
)

func Test_getUsers(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WithArgs().WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Email",
			"Password",
		}))
	if _, err := userModel.GetUsers(); err != nil {
		t.Error(err)
	}
}

func Test_getUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}
	sql := "SELECT * FROM `users` WHERE Id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(strconv.FormatUint(uint64(user.Id), 10)).WillReturnRows(
		sqlmock.NewRows([]string{
			"Id",
			"Email",
			"Password",
		}).AddRow(1, "test_email", "test_password"))
	if _, err := userModel.GetUser("1"); err != nil {
		t.Error(err)
	}
}

func Test_createUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}
	sql := "INSERT INTO `users` (`created_at`,`updated_at`,`email`,`password`) VALUES (?,?,?,?)"
	// https://github.com/DATA-DOG/go-sqlmock/issues/117
	// https://github.com/DATA-DOG/go-sqlmock/issues/224
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), user.Email, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	if err := userModel.CreateUser(&user); err != nil {
		t.Error(err)
	}
}

func Test_updateUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}
	sql := "UPDATE `users` SET `updated_at`=?,`email`=?,`password`=? WHERE `users`.`deleted_at` IS NULL AND `id` = ?"
	// https://github.com/DATA-DOG/go-sqlmock/issues/117
	// https://github.com/DATA-DOG/go-sqlmock/issues/224
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), user.Email, user.Password, user.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	if err := userModel.UpdateUser(&user); err != nil {
		t.Error(err)
	}
}

func Test_deleteUser(t *testing.T) {
	if _, _, err := db.Mock_DB(); err != nil {
		t.Errorf("Mocking database is failed. %v", err)
	}
	mockDB, mock, _ := db.Mock_DB()
	db.SetDB(mockDB)

	user := entity.User{
		Id:       1,
		Email:    "test_email",
		Password: "test_password",
	}
	sql := "UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL"
	// https://github.com/DATA-DOG/go-sqlmock/issues/117
	// https://github.com/DATA-DOG/go-sqlmock/issues/224
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sql)).WithArgs(sqlmock.AnyArg(), user.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	if err := userModel.DeleteUser(&user); err != nil {
		t.Error(err)
	}
}
