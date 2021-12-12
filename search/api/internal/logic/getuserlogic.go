package logic

import (
	"context"

	"github.com/brickzzhang/zero-helloworld/common/errorx"
	"github.com/brickzzhang/zero-helloworld/search/api/internal/svc"
	"github.com/brickzzhang/zero-helloworld/search/api/internal/types"
	"github.com/brickzzhang/zero-helloworld/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.Request) (resp *types.Response, err error) {
	res, err := l.svcCtx.User.GetUserPassword(l.ctx, &user.GetUserPasswordRequest{Name: req.Name})
	if err != nil {
		return nil, errorx.NewDefaultError("get user password failed")
	}

	return &types.Response{Password: res.Password}, nil
}
