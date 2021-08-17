package zk

import (
	"github.com/paashzj/gutil"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"zookeeper_mate_go/pkg/config"
	"zookeeper_mate_go/pkg/path"
)

func Start() {
	err := generateZkConfig()
	if err != nil {
		logrus.Error("can't start zk ", err)
	}
	startZk()
}

func startZk() {
	command := exec.Command("/bin/bash", path.ZooKeeperStartScript)
	err := command.Start()
	if err != nil {
		logrus.Error("start zk server failed ", err)
	}
	err = command.Wait()
	if err != nil {
		logrus.Error("command wait error ", err)
	}
}

func generateZkConfig() error {
	configProp := gutil.ConfigProperties{}
	configProp.Init()
	configProp.Set("tickTime", "2000")
	configProp.Set("initLimit", "10")
	configProp.Set("syncLimit", "5")
	configProp.Set("dataDir", "/opt/sh/zookeeper/data")
	configProp.Set("clientPort", "2181")
	configProp.Set("maxClientCnxns", "60")
	configProp.Set("4lw.commands.whitelist", "stat,ruok,conf,isro,mntr,srvr")
	configProp.Set("metricsProvider.className", "org.apache.zookeeper.metrics.prometheus.PrometheusMetricsProvider")
	configProp.Set("metricsProvider.exportJvmInfo", "true")
	if config.ClusterEnable {
		configProp.Set("standaloneEnabled", "false")
		configProp.Set("reconfigEnabled", "false")
		// cluster mode config server
		hostname := os.Getenv("HOSTNAME")
		index := strings.LastIndex(hostname, "-")
		prefix := hostname[0:index]
		for i := 0; i < 3; i++ {
			configProp.Set("server."+strconv.Itoa(i+1), prefix+"-"+strconv.Itoa(i)+".zookeeper:2888:3888")
		}
		zkIndex := hostname[index+1:]
		index, _ = strconv.Atoi(zkIndex)
		err := ioutil.WriteFile(path.ZooKeeperMyId, []byte(strconv.Itoa(index+1)), os.FileMode(0666))
		if err != nil {
			return err
		}
	}
	return configProp.Write(path.ZooKeeperConfig)
}
