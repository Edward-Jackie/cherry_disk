package logic

import (
	"cherry-disk/core/define"
	"cherry-disk/core/helper"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"

	"github.com/tencentyun/cos-go-sdk-v5"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteReply, err error) {
	ob := make([]cos.Object, 0)
	for _, v := range req.CosObjects {
		ob = append(ob, cos.Object{
			ETag:       v.Etag,
			PartNumber: v.PartNumber,
		})
	}

	err = helper.CosPartUploadComplete(req.Key, req.UploadId, ob)
	rp := &model.RepositoryPool{
		Identity: helper.UUID(),
		Hash:     req.Md5,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     define.CosBucket + "/" + req.Key,
	}

	err = l.svcCtx.Db.Create(rp).Error
	if err != nil {
		return nil, err
	}
	return
}
