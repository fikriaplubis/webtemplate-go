package config

import (
	"github.com/gin-contrib/cors"
	// "webtemplate/internal/middlewares/validator"
)

func (s *server) SetupRouter() error {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	}))

	handlers := s.SetupHandlers()
	// fileValidation, dataValidation := s.ImplementValidators()

	s.Router.POST("/auth/login", handlers.auth.Login)

	s.Router.POST("/menu", handlers.menu.Create)
	s.Router.GET("/menu", handlers.menu.GetAll)
	// s.Router.GET("/menu/:menu_id", handlers.menu.GetByMenuId)

	s.Router.POST("/role", handlers.role.Create)

	s.Router.POST("/user", handlers.user.Create)

	// s.Router.POST("/direct_transfer/upload", validator.File(fileValidation.direct_transfer), validator.Data(dataValidation.direct_transfer), handlers.direct_transfer.Create)

	return nil
}
