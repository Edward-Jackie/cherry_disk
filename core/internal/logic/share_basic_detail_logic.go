package logic

import (
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (*types.ShareBasicDetailReply, error) {
	//
	err := l.svcCtx.Db.Model(&model.ShareBasic{}).
		UpdateColumn("click_num", gorm.Expr("click_num + ?", 1)).
		Where("identity = ?", req.Identity).
		Error

	if err != nil {
		return nil, err
	}

	resp := new(types.ShareBasicDetailReply)

	err = l.svcCtx.Db.Model(&model.ShareBasic{}).
		Select("share_basic.repository_identity, user_repository.name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Joins("left join repository_poll on share_basic.repository_identity = repository_pool.identity").
		Joins("left join user_repository on user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", req.Identity).Scan(&resp).Error

	if err != nil {
		return nil, err
	}

	return resp, nil
}
