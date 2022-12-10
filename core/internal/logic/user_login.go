package logic

import (
	"cherry-disk/core/helper"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogin struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogin {
	return &UserLogin{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserLogin) UserLogin(req *types.LoginReq) (*types.LoginRpl, error) {
	user := new(model.UserBasic)
	err := u.svcCtx.Db.Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).First(&user).Error
	if err != nil || user.ID == 0 {
		u.Logger.Error(fmt.Sprintf("没有 %v用户 err = ", req.Name), err)
		return nil, err
	}

	res := &types.LoginRpl{
		Token:        "CherryNeko",
		RefreshToken: "爱你呀，主人",
	}

	return res, nil
}
