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

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	ur := new(model.UserRepository)
	result := l.svcCtx.Db.Where("identity = ?", req.UserRepositoryIdentity).First(&ur)
	err = result.Error
	count := result.RowsAffected

	if count == 0 {
		err = errors.New("没有该记录")
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	record := &model.ShareBasic{
		Identity:               helper.UUID(),
		UserIdentity:           userIdentity,
		RepositoryIdentity:     req.UserRepositoryIdentity,
		UserRepositoryIdentity: ur.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}

	err = l.svcCtx.Db.Create(result).Error
	if err != nil {
		return nil, err
	}

	resp = &types.ShareBasicCreateReply{
		record.Identity,
	}
	return
}
