package router

import (
	"crypto/sha256"
	"crypto/subtle"
	"os"

	main_app "blogv2/app"
	"blogv2/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup() {
	main_app.Core.Server.Setup()
	setupRouter(main_app.Core.Server.App)
	main_app.Core.Server.Start()
}
func setupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())
	//v1 for UI calling
	v1 := api.Group("/v1")
	v1.Get("/ping", handler.GetPong)
	//account api
	account := api.Group("/account")
	account.Use(keyauth.New(keyauth.Config{
		KeyLookup: "session:x-api-key",
		Validator: validateAccountApiKey,
	}))
	account.Post("/register.json", handler.Register)
	account.Post("/login.json", handler.Login)

	//manager for CMS admin calling
	system := v1.Group("/system")
	system.Get("/login.json", handler.Login)
	system.Post("/post.json", handler.CreatePost)
}

func validateAccountApiKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAdminApiKey := sha256.Sum256([]byte(os.Getenv("X_ADMIN_API_KEY")))
	hashedKey := sha256.Sum256([]byte(key))
	if subtle.ConstantTimeCompare(hashedAdminApiKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}
