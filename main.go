package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gyu-young-park/go_blog/api"
	"github.com/gyu-young-park/go_blog/db"
)

const DSN = "gyu:1234@tcp(127.0.0.1:8002)/blog?parseTime=true"

func main() {
	db := db.StartDB("mysql", DSN)
	server, err := api.NewServer(db)
	if err != nil {
		log.Fatalf("Can't start server [%v]\n", err.Error())
		return
	}
	server.StartServer(":8001")
}
