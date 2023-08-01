package routers

import (
	"kv/tross/database"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func createUser(c *fiber.Ctx) error {
	db := database.DB

	var user database.User

	err := c.BodyParser(&user)

	if err != nil {
		c.Status(http.StatusBadRequest).SendString("Failed to parse request into user.")
	}

	_, err = db.CreateUser(c.Context(), database.CreateUserParams{
		ID:   uuid.NewString(),
		Name: user.Name,
	})

	if err != nil {
		panic(err)
	}

	return c.SendString("Created the user.")
}

func getAllUsers(c *fiber.Ctx) error {
	db := database.DB
	users, err := db.GetAllUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func RegisterUserRoutes(app *fiber.App) {
	usersRouter := app.Group("/users")

	usersRouter.Post("/create", createUser)
	usersRouter.Get("/", getAllUsers)
}
