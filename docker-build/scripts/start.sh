#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
mkdir /opt/sh/zookeeper/data
mkdir /opt/sh/zookeeper/logs
$ZOOKEEPER_HOME/mate/config_gen
bash -x $DIR/start-daemon.sh
tail -f /dev/null