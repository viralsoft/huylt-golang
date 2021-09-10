package main

import (
	"api-fiber/controller/userController"
	"api-fiber/model/user"
	"api-fiber/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)


func setUserRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/users")

	user.Get("/", userController.Index)
	user.Post("/", userController.Store)
	user.Get("/:id", userController.Show)
	user.Delete(":/", userController.Delete)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/fiber-api?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	setUserRoutes(app)

	errPort := app.Listen(":3002")
	if errPort != nil {
		log.Println(errPort)
	}
}
