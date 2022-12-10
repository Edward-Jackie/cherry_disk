package logic

import (
	"cherry-disk/core/helper"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegister struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegister {
	return &UserRegister{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserRegister) UserReg(req *types.UserRegisterRequest) (*types.UserRegisterReply, error) {

	code, err := u.svcCtx.Rc.Get(req.Email).Result()
	if req.Code != code {
		err = errors.New(fmt.Sprintf("验证码不一致"))
		u.Logger.Error(err)
		return &types.UserRegisterReply{Message: err.Error()}, err
	}

	rows := u.svcCtx.Db.Where("name = ?", req.Name).Find(&model.UserBasic{}).RowsAffected
	if rows >= 1 {
		err = errors.New(fmt.Sprintf("%v 用户名已存在", req.Name))
		u.Logger.Error(err)
		return &types.UserRegisterReply{Message: err.Error()}, err
	}

	user := &model.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}

	err = u.svcCtx.Db.Create(&user).Error
	if err != nil {
		u.Logger.Error("用户注册失败 err =", err)
		return &types.UserRegisterReply{Message: "用户注册失败"}, err
	}
	return &types.UserRegisterReply{Message: "用户注册成功"}, nil
}
