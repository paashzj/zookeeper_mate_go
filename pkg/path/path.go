package path

import (
	"os"
	"path/filepath"
)

// zookeeper
var (
	ZooKeeperHome      = os.Getenv("ZOOKEEPER_HOME")
	ZooKeeperConfigDir = filepath.FromSlash(ZooKeeperHome + "/conf")
	ZooKeeperConfig    = filepath.FromSlash(ZooKeeperConfigDir + "/zoo.cfg")
	ZooKeeperDataDir = filepath.FromSlash(ZooKeeperHome + "/data")
	ZooKeeperMyId = filepath.FromSlash(ZooKeeperDataDir + "/myid")
)

// mate
var (
	ZooKeeperMatePath    = filepath.FromSlash(ZooKeeperHome + "/mate")
	ZooKeeperScripts     = filepath.FromSlash(ZooKeeperMatePath + "/scripts")
	ZooKeeperStartScript = filepath.FromSlash(ZooKeeperScripts + "/start-zookeeper.sh")
)
