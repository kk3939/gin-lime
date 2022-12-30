package main

import (
	"github.com/kk3939/gin-lime/db"
	"github.com/kk3939/gin-lime/server"
)

func main() {
	rt := 30
	db.Connect(rt)
	db.Seeds()
	r := server.GetRouter()
	r.Run(":3000")
}
