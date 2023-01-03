package main

import (
	"fmt"
	"os"

	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/server"
)

func main() {
	rt := 30
	db.Connect(rt)
	if os.Getenv("APP_ENV") == "local" {
		fmt.Println("This is local env.")
		var seeds db.Seeds
		seeds.ToDo()
	} else if os.Getenv("APP_ENV") == "dev" {
		fmt.Println("This is staging env.")
	} else if os.Getenv("APP_ENV") == "prod" {
		fmt.Println("This is production env.")
	} else {
		panic("APP_ENV environmental variable is not considered value.")
	}
	r := server.GetRouter()
	r.Run(":3000")
}
