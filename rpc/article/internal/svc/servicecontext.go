package svc

import (
	"manlu.org/tao/zrpc"
	"sllt/backend-learning/blog/rpc/article/internal/config"
	"sllt/backend-learning/blog/rpc/reply/replyclient"
)

type ServiceContext struct {
	Config config.Config
	Reply  replyclient.Reply
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Reply:  replyclient.NewReply(zrpc.MustNewClient(c.Reply)),
	}
}
