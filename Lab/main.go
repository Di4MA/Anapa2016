package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// Адреса для функций с контактами
	app.Get("api/v1/contact", InfoAboutContact)
	app.Post("api/v1/contact", GiveNewContact)
	app.Put("api/v1/contact", ChangeInfoContact)
	app.Delete("api/v1/contact", DeleteContactOnId)

	// Адреса для групп
	app.Post("/api/v1/group", GiveNewGroup)
	app.Get("/api/v1/group", InfoAboutGroup)
	app.Put("/api/v1/group", ChangeInfoGroup)
	app.Delete("/api/v1/group", DeleteContactOnId)

	if err := app.Listen(":6080"); err != nil {
		log.Fatal(err)
	}
}
