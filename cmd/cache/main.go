package main

import (
	"fmt"
	"os"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/log"
	"github.com/spf13/pflag"

	"github.com/caicloud/nirvana-practice/pkg/apis"
	"github.com/caicloud/nirvana-practice/pkg/info"
)

var (
	httpPort uint16
	version bool
)

func init() {
	pflag.Uint16VarP(&httpPort, "port", "p", 8081,"the http port used by the server")
	pflag.BoolVarP(&version,"version", "v", false, "show version info")
	pflag.Parse()
}

func main() {
	if version {
		fmt.Printf("practice-cache-server %s\n", info.Info())
		os.Exit(0)
	}

	// initialize Cache server config
	config := nirvana.NewDefaultConfig().Configure(nirvana.Port(httpPort))

	// install CacheAPIs
	apis.InstallCache(config)

	// create the server and serve
	server := nirvana.NewServer(config)
	if err := server.Serve(); err != nil {
		log.Errorf("server failed with error: %s", err.Error())
	}


}
