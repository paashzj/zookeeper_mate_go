package main

import "zookeeper_mate_go/pkg/zk"

func main() {
	err := zk.Config()
	if err != nil {
		panic(err)
	}
}
