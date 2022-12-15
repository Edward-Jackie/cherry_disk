package handler

import (
	"cherry-disk/core/internal/logic"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRepositorySaveHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRepositorySaveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRepositorySaveLogic(r.Context(), svc)
		resp, err := l.UserRepositorySave(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
			return
		} else {
			httpx.OkJson(w, resp)
			return
		}
	}
}
