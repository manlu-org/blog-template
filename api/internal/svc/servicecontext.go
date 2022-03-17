package svc

import (
	"manlu.org/tao/zrpc"
	"sllt/backend-learning/blog/api/internal/config"
	"sllt/backend-learning/blog/rpc/article/articleclient"
)

type ServiceContext struct {
	Config  config.Config
	Article articleclient.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Article: articleclient.NewArticle(zrpc.MustNewClient(c.Article)),
	}
}
