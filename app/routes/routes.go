package routes

import (
	controller "backend-golang/app/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRouter(e *echo.Echo, ctrl *controller.Main) {

	v1 := e.Group("/v1")

	v1.GET("/swagger/*", echoSwagger.WrapHandler)

	
	users := v1.Group("/users")
	{
		users.POST("", ctrl.User.Create)
		users.GET("", ctrl.User.Get)
		users.GET("/:id", ctrl.User.GetDetail)
		users.PATCH("/:id", ctrl.User.Update)
		users.DELETE("/:id", ctrl.User.Delete)
	}


	pakets := v1.Group("/paket-data")
	{
		pakets.POST("", ctrl.PaketData.Create)
		pakets.GET("", ctrl.PaketData.Get)
		pakets.GET("/:id", ctrl.PaketData.GetDetail)
		pakets.PATCH("/:id", ctrl.PaketData.Update)
		pakets.DELETE("/:id", ctrl.PaketData.Delete)
	}


	transactions := v1.Group("/transactions")
	{
		transactions.POST("", ctrl.Transaction.Create)
		transactions.GET("", ctrl.Transaction.Get)
		transactions.GET("/:id", ctrl.Transaction.GetDetail)
	}
}