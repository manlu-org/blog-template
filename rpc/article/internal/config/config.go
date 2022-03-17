package config

import (
	"manlu.org/tao/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Reply zrpc.RpcClientConf
}
