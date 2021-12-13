package zk

import (
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"zookeeper_mate_go/pkg/config"
	"zookeeper_mate_go/pkg/path"
	"zookeeper_mate_go/pkg/util"
)

func Start() {
	startZk()
	go func() {
		for {
			time.Sleep(10 * time.Second)
			processExists, err := gutil.ProcessExists("org.apache.zookeeper.server.quorum.QuorumPeerMain")
			if err != nil {
				util.Logger().Error("check process exists error ", zap.Error(err))
				continue
			}
			if !processExists {
				startZk()
			}
		}
	}()
}

func startZk() {
	stdout, stderr, err := gutil.CallScript(path.ZooKeeperStartScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	if err != nil {
		util.Logger().Error("command wait error ", zap.Error(err))
	}
}

func Config() error {
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
	if config.TlsEnable {
		configProp.SetInt("secureClientPort", 2182)
	}
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
