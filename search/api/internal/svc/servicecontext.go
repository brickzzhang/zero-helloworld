package svc

import (
	"github.com/brickzzhang/zero-helloworld/search/api/internal/config"
	"github.com/brickzzhang/zero-helloworld/user/rpc/userclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.User)),
	}
}
