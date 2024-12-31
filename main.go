package main 

import ( 
	"log"

	"tracknsave/database"
	"tracknsave/models"
	"github.com/gofiber/fiber/v2"
	"tracknsave/handlers"
	"tracknsave/config"
	"tracknsave/middleware"
)

func main(){

	app := fiber.New()
	//connect to the database
	database.Connect()
	
	//migrate models and run them
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	jwt := middleware.NewAuthMiddleware(config.Secret)
	
	//define  basic routes
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Welcome")
		return err
	})

	app.Post("/login", handlers.Login)
	app.Get("/protected", jwt, handlers.Protected)

	log.Fatal(app.Listen(":3000"))
}