#!/bin/bash

export REMOTE_MODE=false
nohup $ZOOKEEPER_HOME/mate/zookeeper_mate >$ZOOKEEPER_HOME/logs/zookeeper_mate.log 2>$ZOOKEEPER_HOME/logs/zookeeper_mate_error.log