package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello word!!!")

	// var myName string = "Loc"
	// const mySecondName string = "Roger"
	// myThirdName := "Johnny"
	// fmt.Println(myName)
	// fmt.Println(mySecondName)
	// fmt.Println(myThirdName)

	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello Golang, Fiber!!! ^_^"})
	})

	app.Post("/api/todolist", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todolist/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				// todos[i].Completed = !todos[i].Completed
				todos[i].Completed = true

				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":4000")) // go run .\main.go
}
