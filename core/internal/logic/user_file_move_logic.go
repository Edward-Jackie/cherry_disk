package logic

import (
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	parent := new(model.UserRepository)
	err = l.svcCtx.Db.Where("identity = ? AND user_identity = ?", req.ParentIdnetity, userIdentity).First(&parent).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	err = l.svcCtx.Db.Model(new(model.UserRepository)).Where("identity = ?", req.Idnetity).UpdateColumn("parent_id", parent.ID).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp = &types.UserFileMoveReply{Message: "修改成功"}

	return
}
