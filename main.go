package main 

import ( 
	"log"

	"tracknsave/database"
	"tracknsave/models"
	"github.com/gofiber/fiber/v2"
)

func main(){
	//connect to the database
	database.Connect()
	
	//migrate models and run them
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	app := fiber.New()
	
	//defien  basic routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the app")
	})

	log.Fatal(app.Listen(":3000"))
}