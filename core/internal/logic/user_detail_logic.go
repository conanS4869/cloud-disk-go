package logic

import (
	"cloud-disk-go/core/models"
	"context"
	"errors"

	"cloud-disk-go/core/internal/svc"
	"cloud-disk-go/core/internal/types"

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

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	resp = &types.UserDetailResponse{}
	ub := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.Identity).Get(ub)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}
