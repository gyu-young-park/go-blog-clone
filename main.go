package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gyu-young-park/go_blog/api"
	"github.com/gyu-young-park/go_blog/db"
)

const DSN = "gyu:1234@tcp(127.0.0.1:8002)/blog"

func main() {
	db := db.StartDB("mysql", DSN)
	server := api.NewServer(db)
	server.StartServer(":8001")
}
