package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"zookeeper_mate_go/pkg/zk"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	logrus.Info("this is info message")
	logrus.Error("this is error message")
	zk.Start()
	for {
		select {
		case <-interrupt:
			return
		}
	}

}
