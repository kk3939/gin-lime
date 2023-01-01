package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/entity"
	"github.com/kk3939/gin-lime/server"
)

func Test_getTodos(t *testing.T) {
	mockDB, err := db.ConnectTestDB("getTodos")
	if err != nil {
		t.Error(err)
	}
	db.SetDB(mockDB)

	todos := []entity.Todo{
		{Name: "1", Content: "1"},
		{Name: "2", Content: "2"},
	}
	db := db.GetDB()
	if err := db.Create(&todos).Error; err != nil {
		panic(err)
	}

	r := server.GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/todo", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Error("invalid request")
	}
}
