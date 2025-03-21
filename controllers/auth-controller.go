package controllers

import (
	"github.com/DaffaKhayru/gofiber-starter-pack/config"
	"github.com/DaffaKhayru/gofiber-starter-pack/models"
	"github.com/DaffaKhayru/gofiber-starter-pack/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Signup(ctx *fiber.Ctx) error {
	// create variable to store user body req
	var User models.User

	// parsing body request to user variables
	if err := ctx.BodyParser(&User).Error; err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Invalid body request",
		});
	}

	// check existing user
	findUser := config.DB.Where("email = ?", User.Email).First(&User)

	if findUser.Error != nil {
		if findUser.Error == gorm.ErrRecordNotFound {
			// hashing password
			hashPassword,err := util.HashPassword(User.Password)

			// if error occured while hashing password
			if err != nil {
				return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
					"error": "Error hashing password",
				})
			}

			// create new user
			newUser := models.User{
				Username: User.Username,
				Email: User.Email,
				Password: hashPassword,
			}

			// create user to database
			config.DB.Create(&newUser)
		}else {
			return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": "Error checking user",
			})
		}
	}else {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "User already exist",
		})
	}
	
	// return value
	return ctx.Status(200).JSON(fiber.Map{
		"msg": "Signup success",
	})
}

func Login(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"msg": "hello",
	})
}

func Signout(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"msg": "hello",
	})
}