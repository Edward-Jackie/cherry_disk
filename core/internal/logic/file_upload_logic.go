package logic

import (
	"cherry-disk/core/helper"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (*types.FileUploadReply, error) {
	rp := &model.RepositoryPool{
		Identity: helper.UUID(),
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}

	err := l.svcCtx.Db.Create(rp).Error
	if err != nil {
		return nil, err
	}

	resp := &types.FileUploadReply{
		Identity: rp.Identity,
		Ext:      rp.Ext,
		Name:     rp.Name,
	}
	return resp, nil
}
