package main

import (
	"flag"
	"fmt"

	"sllt/backend-learning/blog/api/internal/config"
	"sllt/backend-learning/blog/api/internal/handler"
	"sllt/backend-learning/blog/api/internal/svc"

	"manlu.org/tao/core/conf"
	"manlu.org/tao/rest"
)

var configFile = flag.String("f", "etc/blog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
