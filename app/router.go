package app

import (
	"log"

	"github.com/haerul-umam/capstone-project-mikti/controller"
	"github.com/haerul-umam/capstone-project-mikti/helper"
	customMiddleware "github.com/haerul-umam/capstone-project-mikti/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(
	authController controller.AuthController,
	orderController controller.OrderController,
) *echo.Echo {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Validator = helper.NewValidator()
	e.HTTPErrorHandler = helper.BindAndValidate

	// Auth Controller
	e.POST("/v1/login", authController.Login)
	e.POST("/v1/register", authController.Register)

	adminRoutes := e.Group("/api/admin")
	adminRoutes.Use(customMiddleware.JWTProtection())
	adminRoutes.Use(customMiddleware.JWTAuthRole("ADMIN"))
	adminRoutes.GET("/v1/order", orderController.GetOrdersPage)

	buyerRoutes := e.Group("/api")
	buyerRoutes.Use(customMiddleware.JWTProtection())
	buyerRoutes.Use(customMiddleware.JWTAuthRole("BUYER"))
	buyerRoutes.POST("/v1/order", orderController.CreateOrder)

	return e
}
