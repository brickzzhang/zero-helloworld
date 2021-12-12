package logic

import (
	"context"
	"strings"

	"github.com/brickzzhang/zero-helloworld/common/errorx"
	"github.com/brickzzhang/zero-helloworld/user/api/internal/svc"
	"github.com/brickzzhang/zero-helloworld/user/api/internal/types"
	"github.com/brickzzhang/zero-helloworld/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterRequest) (*types.RegisterResponse, error) {
	if len(strings.TrimSpace(req.Name)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	if _, err := l.svcCtx.Model.Insert(&model.User{Name: req.Name, Password: req.Password}); err != nil {
		return nil, errorx.NewDefaultError("注册失败")
	}

	user, err := l.svcCtx.Model.FindOneByName(req.Name)
	if err != nil {
		return nil, errorx.NewDefaultError("注册失败")
	}

	return &types.RegisterResponse{ID: int(user.Id)}, nil
}
