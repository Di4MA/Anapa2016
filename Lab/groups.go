package main

import (
	"github.com/gofiber/fiber/v2"
)

type Group struct {
	ID          int
	Title       string
	Description string
	Contacts    []int
}

func GroupNoArgs() *Group {
	return &Group{}
}

func GiveNewGroup(ctx *fiber.Ctx) error {
	var group Group
	if err := ctx.BodyParser(&group); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}
	return ctx.JSON(group)
}

func InfoAboutGroup(ctx *fiber.Ctx) error {
	idForNoArgsStruct := ctx.QueryInt("id")
	GroupNoArgs := GroupNoArgs()
	GroupNoArgs.ID = idForNoArgsStruct
	return ctx.JSON(GroupNoArgs)
}

func ChangeInfoGroup(ctx *fiber.Ctx) error {
	idForNoArgsStruct := ctx.QueryInt("id")
	GroupNoArgs := GroupNoArgs()
	GroupNoArgs.ID = idForNoArgsStruct
	return ctx.JSON(GroupNoArgs)
}
func DeleteGroupOnId(ctx *fiber.Ctx) error {
	idForNoArgsStruct := ctx.QueryInt("id")
	GroupNoArgs := GroupNoArgs()
	GroupNoArgs.ID = idForNoArgsStruct
	return ctx.JSON(GroupNoArgs)
}
