package main


//reorganizing repo
import (
	"api"
	"cmd"
	"handlers"
  "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("main")

  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{"msg": "Valid"})
  })

  log.Fatal(app.listen(":8080"))
}
