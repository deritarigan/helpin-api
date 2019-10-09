package main

import (
	"fmt"
	"helpin/controller"
	"helpin/driver"

	"context"
	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const version = "v1"

var ctx = context.Background()
var psqlInfo string
var db *sql.DB
var err error

func setupRouter(db *sql.DB) *gin.Engine {

	router := gin.Default()
	helpinController := controller.Controllers{}

	router.POST("/"+version+"/auth/signup", helpinController.SignUp(db))
	router.POST("/"+version+"/auth/login", helpinController.Login(db))

	return router
}

func main() {
	db = driver.ConnectDBLocal()

	r := setupRouter(db)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
	fmt.Println("Successfully connected!")
	// staticPort := "38585"
}
