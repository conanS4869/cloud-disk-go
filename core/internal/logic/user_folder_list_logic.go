package logic

import (
	"cloud-disk-go/core/models"
	"context"

	"cloud-disk-go/core/internal/svc"
	"cloud-disk-go/core/internal/types"

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

func (l *UserFolderListLogic) UserFolderList(req *types.UserFolderListRequest) (resp *types.UserFolderListResponse, err error) {
	var (
		ufl = make([]*types.UserFolder, 0)
		ur  = new(models.UserRepository)
	)
	resp = new(types.UserFolderListResponse)
	// 获取当前文件夹ID
	_, err = l.svcCtx.Engine.Table("user_repository").Select("id").
		Where("identity = ?", req.Identity).Get(ur)
	if err != nil {
		return
	}
	// 获取文件夹列表
	err = l.svcCtx.Engine.Table("user_repository").Select("identity,name").
		Where("parent_id = ?", ur.Id).Find(&ufl)
	if err != nil {
		return
	}

	resp.List = ufl
	return
}
