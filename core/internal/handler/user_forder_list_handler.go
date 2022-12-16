package handler

import (
	"cherry-disk/core/internal/logic"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFolderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFolderListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserFolderListLogic(r.Context(), svcCtx)
		resp, err := l.UserFolderList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
