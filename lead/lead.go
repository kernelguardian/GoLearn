package lead

import (
	"fiberproject/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   string // Changed to string
}

func GetLeads(c *fiber.Ctx) error {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	return c.Status(fiber.StatusOK).JSON(leads)
}

func NewLead(c *fiber.Ctx) error {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(lead)
	return c.JSON(lead)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	if result := db.First(&lead, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lead not found"})
	}
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	if result := db.First(&lead, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("No lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}
