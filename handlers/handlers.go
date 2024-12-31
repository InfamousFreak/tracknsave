package handlers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"tracknsave/config"
	"tracknsave/models"
	"tracknsave/repository"
)

//login route
func Login(c *fiber.Ctx) error {
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

//find the user by credentials 
user, err := repository.Find(loginRequest.Email, loginRequest.Password)
if err != nil {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": err.Error(),
	})
}
day := time.Hour * 24
//create the jwt claims, which includes the user ID and expiry time
claims := jtoken.MapClaims{
	"ID": user.ID,
	"email": user.Email,
	"exp": time.Now().Add(day * 1).Unix(),
}
//create token
token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
//generate encoded token and send it as response
t, err := token.SignedString([]byte(config.Secret))
if err != nil {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}
//return the token
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}
//protected route
func Protected(c *fiber.Ctx) error {
	//get the userfrom the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	return c.SendString("Welcome" + email)
}