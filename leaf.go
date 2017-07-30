package leaf

import (
	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/console"
	"github.com/golang/glog"
	"github.com/name5566/leaf/module"
	"os"
	"os/signal"
)

func Run(mods ...module.Module) {
	// logger
	//glog.SetFlags(conf.LogFlag)
	glog.SetLogDir(conf.LogPath)
	glog.SetLogLevel(conf.LogLevel)

	glog.Infof("Leaf %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	glog.Infof("Leaf closing down (signal: %v)", sig)
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
