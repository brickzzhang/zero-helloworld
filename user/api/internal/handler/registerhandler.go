package handler

import (
	"net/http"

	"github.com/brickzzhang/zero-helloworld/user/api/internal/logic"
	"github.com/brickzzhang/zero-helloworld/user/api/internal/svc"
	"github.com/brickzzhang/zero-helloworld/user/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
