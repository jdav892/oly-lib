package main

//reorganizing repo
import (
  "fmt"
  "log"
  "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Server Running")
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{"msg": "Valid"})
  })

  log.Fatal(app.Listen(":8080"))
}
