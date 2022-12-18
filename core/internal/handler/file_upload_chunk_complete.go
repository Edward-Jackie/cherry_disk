package handler

import (
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkCompleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkCompleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// 判断是否已达用户容量上限
		userIdentity := r.Header.Get("UserIdentity")
		ub := new(model.UserBasic)
		err := svcCtx.Db.Where("identity = ?", userIdentity).First(&ub).Error
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if req.Size+ub.NowVolume > ub.TotalVolume {
			httpx.Error(w, errors.New("已超出当前容量"))
			return
		}
		//l := logic.NewFileUploadChunkCompleteLogic(r.Context(), svcCtx)
		//resp, err := l.FileUploadChunkComplete(&req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}
