package handler

import (
	"net/http"

	"Alredomn/cmd/internal/logic"
	"Alredomn/cmd/internal/svc"
	"Alredomn/cmd/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserOrdersRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserOrdersLogic(r.Context(), svcCtx)
		resp, err := l.UserOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
