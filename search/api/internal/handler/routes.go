// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/brickzzhang/zero-helloworld/search/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users/name/:name",
				Handler: GetUserHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
