package logic

import (
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderListLogic {
	return &UserFolderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderListLogic) UserFolderList(req *types.UserFolderListRequest) (resp *types.UserFolderListReply, err error) {
	ur := new(model.UserRepository)
	urList := make([]*types.UserFolder, 0)
	resp = new(types.UserFolderListReply)

	err = l.svcCtx.Db.Where("identity = ?", req.Identity).First(&ur).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	err = l.svcCtx.Db.Model(new(model.UserRepository)).Where("parent_id = ?", ur.ID).Scan(&urList).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	resp.List = urList
	return
}
