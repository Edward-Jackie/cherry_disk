package logic

import (
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (rep *types.UserDetailReply, err error) {
	u := new(model.UserBasic)
	err = l.svcCtx.Db.Where("identity = ?", req.Identity).First(&u).Error
	if err != nil {
		l.Logger.Error(fmt.Sprintf("没有 %v identity 的记录", req.Identity))
		return nil, err
	}
	rep = &types.UserDetailReply{}
	rep.Email = u.Email
	rep.Name = u.Name
	return
}
