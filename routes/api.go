package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sizata-siege/finance-management/account"
	"github.com/sizata-siege/finance-management/auth"
)

func SetupAPI(app *fiber.App) {
	app.Get("/", indexHome)
	app.Static("/", "./public", fiber.Static{MaxAge: 0})

	api := app.Group("/api", callNext)

	v1 := api.Group("/v1", callNext)

	/* =-=-=-=-=-=-= Accounts =-=-=-=-=-=-= */
	accounts := v1.Group("/accounts", auth.Middleware)
	accounts.Get("/", account.IndexAccounts)
	accounts.Post("/", account.StoreAccount)
	accounts.Get("/:id", account.ShowAccount)
	accounts.Patch("/:id", account.UpdateAccount)
	accounts.Delete("/:id", account.DeleteAccount)

	/* =-=-=-=-=-=-= Session & User =-=-=-=-=-=-= */
	// v1.Get("/auth")    // get loged in user
	v1.Post("/session", auth.Login)                     // login
	v1.Delete("/session", auth.Middleware, auth.Logout) // logout / smiliar to /logout
	v1.Post("/users", auth.CreateNewUser)

	/* =-=-=-=-=-=-= Transactions =-=-=-=-=-=-= */
	// transactions := v1.Group("/transaction", auth.Middleware)
	// transactions.Get("/", )

	/* =-=-=-=-=-=-= Test =-=-=-=-=-=-= */
	app.Get("/test", testHandler)
}

const welcomeMessage = "Welcome to SIZATA's Finance Management System"

func indexHome(c *fiber.Ctx) error {
	/* if request doesn't accepts html */
	if c.Accepts("text/html") == "" {
		return c.JSON(fiber.Map{"message": welcomeMessage})
	}
	return c.Next()
}

func callNext(c *fiber.Ctx) error { return c.Next() }

func testHandler (c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name: "Foo",
		Value: "Bar",
		Expires: time.Now().Add(time.Minute * 10),
	})
	return c.SendString("OK")
}
