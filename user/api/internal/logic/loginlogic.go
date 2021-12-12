package logic

import (
	"context"
	"strings"
	"time"

	"github.com/brickzzhang/zero-helloworld/common/errorx"
	"github.com/brickzzhang/zero-helloworld/user/api/internal/svc"
	"github.com/brickzzhang/zero-helloworld/user/api/internal/types"
	"github.com/brickzzhang/zero-helloworld/user/model"
	"github.com/golang-jwt/jwt"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (*types.LoginResponse, error) {
	if len(strings.TrimSpace(req.Name)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	userInfo, err := l.svcCtx.Model.FindOneByName(req.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errorx.NewDefaultError("用户密码不正确")
	}

	// generate token
	now := time.Now().Unix()
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{Token: jwtToken}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
