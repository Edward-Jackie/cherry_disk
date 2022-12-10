package logic

import (
	"cherry-disk/core/define"
	"cherry-disk/core/helper"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svc *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svc,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (*types.MailCodeSendReply, error) {
	var cnt int64
	err := l.svcCtx.Db.Model(&model.UserBasic{}).Where("email = ?", req.Email).Count(&cnt).Error
	if err != nil || cnt > 0 {
		l.Logger.Error("邮箱注册报错 err=", err)
		err = errors.New("该邮箱已被注册")
		return nil, err
	}

	codeTTL, err := l.svcCtx.Rc.TTL(req.Email).Result()
	if err != nil {
		return nil, errors.New("the verify code has not expired, RcDB err")
	}

	if codeTTL.Seconds() > 0 || codeTTL.Seconds() == -1 {
		return nil, errors.New("the verify code has not expired")
	}

	code := helper.RandCode()
	l.svcCtx.Rc.Set(req.Email, code, time.Duration(define.CodeExpire)*time.Second)

	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return &types.MailCodeSendReply{Message: "验证码发送失败"}, err
	}

	return &types.MailCodeSendReply{Message: "验证码发送成功"}, nil

}
