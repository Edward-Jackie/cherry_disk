package logic

import (
	"cherry-disk/core/define"
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

	//生成TOKEN
	token, err := helper.TokenBuilder(user.ID, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("%v用户 生成Token失败 err = ", req.Name), err)
		return nil, err
	}

	//用于刷新TOKEN的TOKEN
	refresht, err := helper.TokenBuilder(user.ID, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		u.Logger.Error(fmt.Sprintf("%v用户 生成RefreshToken失败 err = ", req.Name), err)
		return nil, err
	}

	res := &types.LoginRpl{
		Token:        token,
		RefreshToken: refresht,
	}

	return res, nil
}
