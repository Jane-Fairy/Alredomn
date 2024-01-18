package handler

import (
	"net/http"

	"Alredomn/cmd/internal/logic"
	"Alredomn/cmd/internal/svc"
	"Alredomn/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProductLogic(r.Context(), svcCtx)
		resp, err := l.Product(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
