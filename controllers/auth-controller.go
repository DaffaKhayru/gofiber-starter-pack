package controllers

import "github.com/gofiber/fiber/v2"

func Signup(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"msg": "hello",
	})
}