package handler

import (
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func MailCodeSendRegisterHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailCodeSendRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

	}
}
