package config

import "github.com/paashzj/gutil"

// zk config
var (
	ClusterEnable bool
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
}
