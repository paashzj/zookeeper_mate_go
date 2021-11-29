#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
$ZOOKEEPER_HOME/mate/config_gen
bash -x $DIR/start-daemon.sh
tail -f /dev/null