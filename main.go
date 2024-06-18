package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello word!!!")

	// var myName string = "Loc"
	// const mySecondName string = "Roger"
	// myThirdName := "Johnny"
	// fmt.Println(myName)
	// fmt.Println(mySecondName)
	// fmt.Println(myThirdName)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello Golang, Fiber!!! ^_^"})
	})

	log.Fatal(app.Listen(":4000")) // go run .\main.go
}
