package app

import (
	"github.com/dzikrisyafi/kursusvirtual_middleware/middleware"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/enrolls"
	"github.com/dzikrisyafi/kursusvirtual_users-api/src/controllers/users"
)

func mapUrls() {
	router.POST("/users/login", users.Login)

	// users group end point
	usersGroup := router.Group("/users")
	usersGroup.Use(middleware.Auth())
	{
		usersGroup.POST("/", users.Create)
		usersGroup.GET("/", users.GetAll)
		usersGroup.GET("/:user_id", users.Get)
		usersGroup.PUT("/:user_id", users.Update)
		usersGroup.PATCH("/:user_id", users.Update)
		usersGroup.DELETE("/:user_id", users.Delete)
	}

	// internal end point
	internalGroup := router.Group("/internal")
	internalGroup.Use(middleware.Auth())
	{
		router.GET("/internal/users/search", users.Search)
		router.GET("/internal/enrolls/:course_id", enrolls.Get)
		router.POST("/internal/enrolls", enrolls.Create)
	}
}