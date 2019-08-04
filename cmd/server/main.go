package main

import (
	"fmt"
	"github.com/caicloud/nirvana/rest"
	"os"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/log"
	"github.com/spf13/pflag"

	"github.com/caicloud/nirvana-practice/pkg/apis"
	"github.com/caicloud/nirvana-practice/pkg/info"
)

var (
	httpPort uint16
	cacheEndPoint string
	cachePort uint16
	version  bool
)

func init() {
	pflag.Uint16VarP(&httpPort, "port", "p", 8080, "the HTTP port used by the server")
	pflag.StringVarP(&cacheEndPoint, "cache-endpoint", "", "127.0.0.1", "cache client endpoint")
	pflag.Uint16VarP(&cachePort, "cache-port", "", 8081, "cache port")
	pflag.BoolVarP(&version, "version", "v", false, "show version info")
	pflag.Parse()
}

func main() {
	if version {
		fmt.Printf("practice-server, %s\n", info.Info())
		os.Exit(0)
	}

	// initialize Server config
	config := nirvana.NewDefaultConfig().Configure(nirvana.Port(httpPort))

	cacheConf := &rest.Config{
		Scheme:   "http",
		Host:     fmt.Sprintf("%s:%d", cacheEndPoint, cachePort),
		Executor: nil,
	}

	// install APIs
	apis.Install(config, cacheConf)

	// create the server and server
	server := nirvana.NewServer(config)

	if err := server.Serve(); err != nil {
		log.Errorf("server failed with error: %s", err.Error())
	}

}
