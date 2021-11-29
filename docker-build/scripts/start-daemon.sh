#!/bin/bash

export REMOTE_MODE=false
mkdir /opt/sh/zookeeper/data
nohup $ZOOKEEPER_HOME/mate/zookeeper_mate >$ZOOKEEPER_HOME/zookeeper_mate.log 2>$ZOOKEEPER_HOME/zookeeper_mate_error.log