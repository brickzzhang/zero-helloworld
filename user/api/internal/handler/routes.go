// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/brickzzhang/zero-helloworld/user/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/users/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users",
				Handler: LoginHandler(serverCtx),
			},
		},
	)
}