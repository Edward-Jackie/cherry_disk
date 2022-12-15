package handler

import (
	"cherry-disk/core/helper"
	"cherry-disk/core/internal/logic"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(types.FileUploadRequest)
		// 没用上传参
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}

		userIdentity := r.Header.Get("UserIdentity")
		u := new(model.UserBasic)
		err = svc.Db.Where("identity  = ?", userIdentity).First(u).Error
		if err != nil {
			httpx.Error(w, err)
			return
		}

		if fileHeader.Size+u.NowVolume > u.TotalVolume {
			httpx.Error(w, errors.New("已超过当前容量!"))
			return
		}

		//判断文件是否已存在
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(model.RepositoryPool)
		result := svc.Db.Where("hash = ?", hash).First(&rp)
		var count int64
		err = result.Error
		result.Count(&count)

		//if rp.Identity != "" {
		//	httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
		//	return
		//}
		//var count int64
		//err = svc.Db.Where("hash = ?", hash).Find(&rp).Count(&count).Error

		if count > 0 {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}

		if err != nil && count > 0 {
			httpx.Error(w, err)
			return
		}
		// 上传文件
		var filepath string
		filepath, err = helper.UploadTX(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = int32(fileHeader.Size)
		req.Hash = hash
		req.Path = filepath

		l := logic.NewFileUploadLogic(r.Context(), svc)
		resp, err := l.FileUpload(req)

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
