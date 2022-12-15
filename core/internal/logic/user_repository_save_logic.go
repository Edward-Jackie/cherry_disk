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

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, identity string) (*types.UserRegisterReply, error) {
	rp := new(model.RepositoryPool)
	err := l.svcCtx.Db.Where("identity = ?", req.RepositoryIdentity).First(&rp).Error
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	u := new(model.UserBasic)
	err = l.svcCtx.Db.Where("identity = ?", identity).First(&u).Error
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	if u.NowVolume+rp.Size > u.TotalVolume {
		err = errors.New("已超出当前容量")
		return nil, err
	}

	err = l.svcCtx.Db.Update("now_volume = ?", u.NowVolume+rp.Size).Where("identity = ?", identity).Error
	if err != nil {
		return nil, err
	}

	ur := &model.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       identity,
		ParentID:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}

	//新增关联记录
	err = l.svcCtx.Db.Create(ur).Error
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	return &types.UserRegisterReply{Message: "保存成功"}, nil
}
