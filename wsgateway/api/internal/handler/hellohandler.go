package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-ws/wsgateway/api/internal/logic"
	"go-zero-ws/wsgateway/api/internal/svc"
	"go-zero-ws/wsgateway/api/internal/types"
)

func helloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
