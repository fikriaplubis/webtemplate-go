package config

import (
	roleServices "webtemplate/internal/services/role"
	roleHandlers "webtemplate/internal/transport/rest/role"

	userServices "webtemplate/internal/services/user"
	userHandlers "webtemplate/internal/transport/rest/user"

	authServices "webtemplate/internal/services/auth"
	authHandlers "webtemplate/internal/transport/rest/auth"

	menuServices "webtemplate/internal/services/menu"
	menuHandlers "webtemplate/internal/transport/rest/menu"

	directTransferServices "webtemplate/internal/services/direct_transfer"
	directTransferHandlers "webtemplate/internal/transport/rest/direct_transfer"
)

type registeredHandlers struct {
	role            *roleHandlers.Handler
	user            *userHandlers.Handler
	auth            *authHandlers.Handler
	menu            *menuHandlers.Handler
	direct_transfer *directTransferHandlers.Handler
}

func (s *server) SetupHandlers() *registeredHandlers {
	return &registeredHandlers{
		role:            roleHandlers.NewHandler(roleServices.NewService(s.DB)),
		user:            userHandlers.NewHandler(userServices.NewService(s.DB)),
		auth:            authHandlers.NewHandler(authServices.NewService(s.DB)),
		menu:            menuHandlers.NewHandler(menuServices.NewService(s.DB)),
		direct_transfer: directTransferHandlers.NewHandler(directTransferServices.NewService(s.DB)),
	}
}
