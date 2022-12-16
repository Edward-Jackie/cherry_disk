package logic

import (
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	ur := new(model.UserRepository)
	result := l.svcCtx.Db.
		Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).
		First(ur)

	count := result.RowsAffected
	err = result.Error
	if count != 0 && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查询文件出错 或 该名称已存在")
		l.Logger.Error(err)
		return
	}

	err = l.svcCtx.
		Db.Model(&model.UserRepository{}).
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		UpdateColumn("name", req.Name).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp.Message = "修改成功！"

	return
}
