package logic

import (
	"cherry-disk/core/define"
	"cherry-disk/core/internal/model"
	"cherry-disk/core/internal/svc"
	"cherry-disk/core/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, identity string) (res *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	//分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}

	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size

	ur := new(model.UserRepository)
	err = l.svcCtx.Db.Where("identity = ?", req.Identity).First(&ur).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	//err = l.svcCtx.Db.Table("user_repository").
	//	Where("parent_id = ? AND identity = ?", ur.ID, identity).
	//	Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,"+
	//		"user_repository.name, repository_pool.path, repository_pool.size").
	//	Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
	//	Where("user_repository.deleted_at IS NULL ").
	//	Where("LIMIT ? , ?", size, offset).
	//	Find(&ur).Error

	err = l.svcCtx.Db.Exec(" SELECT "+
		"user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,user_repository.name, repository_pool.path, repository_pool.size "+
		"FROM `user_repository` "+
		"left join repository_pool "+
		"on user_repository.repository_identity = repository_pool.identity "+
		"WHERE (parent_id = ? "+
		"AND identity = ?)"+
		" AND user_repository.deleted_at IS NULL  "+
		"AND `user_repository`.`deleted_at` IS NULL "+
		"AND `user_repository`.`id` = ? "+
		"LIMIT ? , ? ", ur.ID, ur.Identity, ur.ID, size, offset).Find(&uf).Error

	if err != nil {
		l.Logger.Error(err)
		return
	}

	var count int64
	err = l.svcCtx.Db.Where("parent_id = ? AND user_identity = ?", ur.ID, identity).Model(new(model.UserRepository)).Count(&count).Error
	if err != nil {
		l.Logger.Error(err)
		return
	}

	res.List = uf
	res.Count = count

	return
}
