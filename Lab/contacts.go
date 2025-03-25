package main

import "github.com/gofiber/fiber/v2"

// Спросить как сделать названия для полей JSON'a
type Phone struct {
	TypeID      int
	CountryCode int
	Operator    int
	Number      int
}

type Contact struct {
	ID         int
	Username   string
	GivenName  string
	FamilyName string
	Phones     []Phone
	Emails     []string
	Birthdate  string
}

func ContactNoArgs() *Contact {
	return &Contact{}
}

// Из тела запроса достаём информацию и сопоставляем её с полями структуры Contact.
// Ответом шлём ту же структуру с заполненными(а может и пустыми) полями в формате JSON
func GiveNewContact(ctx *fiber.Ctx) error {
	var contact Contact
	if err := ctx.BodyParser(&contact); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(contact)
}

func InfoAboutContact(ctx *fiber.Ctx) error {
	ContactNoArgs := ContactNoArgs()
	id := ctx.QueryInt("id")
	ContactNoArgs.ID = id
	return ctx.JSON(ContactNoArgs)
}

func ChangeInfoContact(ctx *fiber.Ctx) error {
	ContactNoArgs := ContactNoArgs()
	id := ctx.QueryInt("id")
	ContactNoArgs.ID = id
	return ctx.JSON(ContactNoArgs)
}

func DeleteContactOnId(ctx *fiber.Ctx) error {
	ContactNoArgs := ContactNoArgs()
	id := ctx.QueryInt("id")
	ContactNoArgs.ID = id
	return ctx.JSON(ContactNoArgs)
}
