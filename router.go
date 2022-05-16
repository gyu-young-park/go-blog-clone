package main

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func (app *App) ping() {
// 	app.Router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// }

// func (app *App) getUserData() {
// 	app.Router.GET("/id/:id", func(c *gin.Context) {
// 		var name string
// 		var email string
// 		var password string
// 		id, err := strconv.Atoi(c.Param("id"))
// 		if err != nil {
// 			fmt.Errorf("Failed convert id string to number")
// 		}
// 		err = app.DB.QueryRow("SELECT user_name, email, password FROM User WHERE id = ?", id).Scan(&name, &email, &password)
// 		if err != nil {
// 			fmt.Errorf("Failed read user[%v]\n", err)
// 		}
// 		c.JSON(200, gin.H{"name": name, "email": email, "password": password})
// 	})
// }

// func (app *App) getAllUserData() {
// 	app.Router.GET("/api/users", func(c *gin.Context) {
// 		var name [5]string
// 		var email [5]string
// 		var password [5]string

// 		rows, err := app.DB.Query("SELECT user_name, email, password FROM User")
// 		if err != nil {
// 			fmt.Errorf("Failed read user[%v]\n", err)
// 		}
// 		defer rows.Close()

// 		index := 0
// 		for rows.Next() && index < 5 {
// 			err := rows.Scan(&name[index], &email[index], &password[index])
// 			if err != nil {
// 				fmt.Errorf("Failed read user[%v]\n", err)
// 			}
// 			index++
// 		}
// 		c.JSON(200, gin.H{"name": name, "email": email, "password": password})
// 	})
// }

// func (app *App) insertUserData() {
// 	app.Router.POST("/user", func(c *gin.Context) {
// 		stmt, err := app.DB.Prepare("INSERT into User (user_name, email, password) VALUES( ?, ?, ?)")
// 		if err != nil {
// 			fmt.Errorf("Failed prepared statement user:[%v]\n", err)
// 		}
// 		defer stmt.Close()

// 		res1, err := stmt.Exec("ww", "wwww@ww", "www")
// 		if err != nil {
// 			fmt.Errorf("Failed insert user:[%v]\n", err)
// 		}

// 		res2, err := stmt.Exec("qq", "qqqq@qq", "qqq")
// 		if err != nil {
// 			fmt.Errorf("Failed insert user:[%v]\n", err)
// 		}

// 		n1, err := res1.RowsAffected()
// 		n2, err := res2.RowsAffected()
// 		if n1 == 1 && n2 == 1 {
// 			fmt.Println("success")
// 		}
// 		c.JSON(200, gin.H{"data": "success"})
// 	})
// }

// func (app *App) RoutingMethodSetup() {
// 	app.ping()
// 	app.getUserData()
// 	app.getAllUserData()
// 	app.insertUserData()
// }
