package main

import (
	"fiberproject/database"
	"fiberproject/lead"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Connected opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":8080")
	defer database.DBConn.Close()
}
