package logic

import (
	"context"

	"github.com/brickzzhang/zero-helloworld/common/errorx"
	"github.com/brickzzhang/zero-helloworld/user/rpc/internal/svc"
	"github.com/brickzzhang/zero-helloworld/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPasswordLogic {
	return &GetUserPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPasswordLogic) GetUserPassword(in *user.GetUserPasswordRequest) (*user.GetUserPasswordResponse, error) {
	one, err := l.svcCtx.Model.FindOneByName(in.Name)
	if err != nil {
		return nil, errorx.NewDefaultError("获取用户信息失败")
	}

	return &user.GetUserPasswordResponse{Password: one.Password}, nil
}
