package main

import (
	"os"
	"os/signal"
	"zookeeper_mate_go/pkg/config"
	"zookeeper_mate_go/pkg/util"
	"zookeeper_mate_go/pkg/zk"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	util.Logger().Info("this is info message")
	util.Logger().Error("this is error message")
	if !config.RemoteMode {
		zk.Start()
	} else {
		util.Logger().Info("remote mode")
	}
	for {
		select {
		case <-interrupt:
			return
		}
	}

}
