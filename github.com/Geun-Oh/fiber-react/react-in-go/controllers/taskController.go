package controllers

import (
	"log"

	"github.com/Geun-Oh/fiber-react/react-in-go/initializers"
	"github.com/Geun-Oh/fiber-react/react-in-go/model"
	"github.com/gofiber/fiber/v2"
)

func FetchTasks(c *fiber.Ctx) error {
	var tasks []model.Task
	initializers.DBConn.Order("created_at desc").Find(&tasks)

	return c.JSON(fiber.Map{
		"tasks": tasks,
	})
}

func CreateTask(c *fiber.Ctx) error {
	var record model.Task

	if err := c.BodyParser(&record); err != nil {
		log.Println("Parse Error at create task")
		return c.Status(400).JSON(model.Task{
			Title: "Parse error",
			Body:  "",
		})
	}

	result := initializers.DBConn.Create(record)

	if result.Error != nil {
		log.Println("Record Err")
		return c.Status(400).JSON(model.Task{
			Title: "Record error",
			Body:  "",
		})
	}

	return c.Status(201).JSON(record)
}

func FetchTask(c *fiber.Ctx) error {
	var task model.Task

	id := c.Params("id")

	record := initializers.DBConn.First(&task, id)

	if record.Error != nil {
		log.Println("Record not found.")
		return c.Status(400).JSON(model.Task{
			Title: "Record not found",
			Body:  "",
		})
	}

	return c.Status(200).JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	var task model.Task

	id := c.Params("id")

	record := initializers.DBConn.Delete(&task, id)

	if record.Error != nil {
		log.Println("Record not found.")
		return c.Status(400).JSON(model.Task{
			Title: "Record not found",
			Body:  "",
		})
	}

	return c.Status(200).JSON(task)
}
