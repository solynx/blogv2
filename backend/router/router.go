package router

import (
	"crypto/sha256"
	"crypto/subtle"
	"os"

	main_app "blogv2/app"
	"blogv2/handler"
	"blogv2/repositories"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func Setup() {
	main_app.Core.Server.Setup()
	setupRouter(main_app.Core.Server.App)
	main_app.Core.Server.Start()
}
func setupRouter(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:3001",
		AllowHeaders: "Origin, Content-Type, Accept, x-api-key, Authorization",
	}))
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
	// system.Get("/login.json", handler.Login)
	// system.Post("/post.json", handler.CreatePost)
	system.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    main_app.Core.Server.RouterPrivateKey.Public(),
		},
		SuccessHandler: func(c *fiber.Ctx) error {

			return checkAuthenticationUserLogin(c)
		},
		ContextKey: "user_token",
	}))
	system.Post("/post.json", handler.CreatePost)
	//category rest api
	system.Post("/category.json", handler.CreateCategory)
	system.Get("/category.json", handler.GetCategory)
	system.Patch("/category.json", handler.UpdateCategory)
	system.Delete("/category.json", handler.DeleteCategory)
}

func validateAccountApiKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAdminApiKey := sha256.Sum256([]byte(os.Getenv("X_API_KEY")))
	hashedKey := sha256.Sum256([]byte(key))
	if subtle.ConstantTimeCompare(hashedAdminApiKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func checkAuthenticationUserLogin(c *fiber.Ctx) error {
	userToken := c.Locals("user_token").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(string)
	lastLoginId := claims["last_login_id"].(string)
	role := claims["role"].(string)
	if role != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	user, row := repositories.GetUserByUidAndLastLoginId(uuid.MustParse(userId), uuid.MustParse(lastLoginId), role)
	if row < 1 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	c.Locals("user", user)
	return c.Next()
}
