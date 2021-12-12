package handler

import (
	"net/http"

	"github.com/brickzzhang/zero-helloworld/search/api/internal/logic"
	"github.com/brickzzhang/zero-helloworld/search/api/internal/svc"
	"github.com/brickzzhang/zero-helloworld/search/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), ctx)
		resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
