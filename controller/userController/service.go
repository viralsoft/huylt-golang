package userController

import (
	"api-fiber/database"
	"api-fiber/model/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type User = user.User

type UserValidate struct {
	Name  string `validate:"required,min=3,max=32"`
	Email string `validate:"required,email,min=6,max=32"`
	Status int `validate:"required,number"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func Index(c *fiber.Ctx) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}
func ValidateStruct(user UserValidate) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func Show(c *fiber.Ctx) error  {
	id := c.Params("id")
	db := database.DBConn
	var user User
	db.Find(&user, id)
	return c.JSON(user)
}

func Store(c *fiber.Ctx) error {
	 input := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Status int `json:"status"`
	}{}

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	inputValidate := new(UserValidate)

	inputValidate.Name = input.Name
	inputValidate.Email = input.Email
	inputValidate.Status = input.Status

	errors := ValidateStruct(*inputValidate)
	if errors != nil {
		return c.JSON(errors)
	}
	db := database.DBConn
	var user User
	user.Name = input.Name
	user.Email = input.Email
	user.Status = input.Status

	db.Create(&user)
	return c.JSON(user)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var user User
	db.First(&user, id)
	if user.Email == "" {
		return c.Send([]byte("No Book Found with ID"))
	}
	db.Delete(&user)
	return c.Send([]byte("User Successfully deleted"))
}