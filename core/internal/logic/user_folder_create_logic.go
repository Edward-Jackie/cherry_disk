package logic

import (
	"cherry-disk/core/helper"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, identity string) (resp *types.UserFolderCreateReply, err error) {
	var count int64
	err = l.svcCtx.Db.
		Model(new(model.UserRepository)).
		Where("name = ? AND parent_id = ?", req.Name, req.ParentId).
		Count(&count).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	if count >= 1 {
		err = errors.New("该路径已存在")
		return
	}

	folder := &model.UserRepository{
		Identity:     helper.UUID(),
		UserIdentity: identity,
		Name:         req.Name,
		ParentID:     req.ParentId,
	}

	err = l.svcCtx.Db.Create(folder).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}
	res := new(types.UserFolderCreateReply)
	res.Identity = folder.Identity
	return res, nil
}
