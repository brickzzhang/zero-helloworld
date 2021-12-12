package svc

import (
	"github.com/brickzzhang/zero-helloworld/user/model"
	"github.com/brickzzhang/zero-helloworld/user/rpc/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
