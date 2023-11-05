package controller

import (
	"log"

	"github.com/Geun-Oh/fiber-react/server/database"
	"github.com/Geun-Oh/fiber-react/server/model"
	"github.com/gofiber/fiber/v2"
)

// Blog list
func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog List",
	}

	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)

}

// Add blog
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog Create",
	}

	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request")
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
	}

	res := database.DBConn.Create(record)

	if res.Error != nil {
		log.Println("Error in saving data")
	}

	context["msg"] = "Record is saved successfully"
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// Update blog contents
func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog Update",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["statusText"] = ""
		context["msg"] = "Record not found"
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
	}

	context["msg"] = "Record updated successfully"
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

// Delete blog
func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog Delete",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record Not Found.")
		context["statusText"] = ""
		context["msg"] = "Record Not Found."

		return c.Status(400).JSON(context)
	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		context["message"] = "Something went wrong."

		return c.Status(400).JSON(context)
	}

	context["msg"] = "Record deleted successfully"
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}
