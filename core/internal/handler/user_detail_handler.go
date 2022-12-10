package handler

import (
	"cherry-disk/core/internal/logic"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserDetailLogic(r.Context(), svcCtx)
		rep, err := l.UserDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, rep)
		}

	}
}
