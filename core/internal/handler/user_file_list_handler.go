package handler

import (
	"cherry-disk/core/internal/logic"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileListHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserFileListLogic(r.Context(), svc)
		res, err := l.UserFileList(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
			return
		} else {
			httpx.OkJson(w, res)
		}
	}
}
