#!/bin/bash

export REMOTE_MODE=false
mkdir $ZOOKEEPER_HOME/data
mkdir $ZOOKEEPER_HOME/logs
$ZOOKEEPER_HOME/mate/config_gen
nohup $ZOOKEEPER_HOME/mate/zookeeper_mate >>$ZOOKEEPER_HOME/logs/zookeeper_mate.stdout.log 2>>$ZOOKEEPER_HOME/logs/zookeeper_mate.stderr.log

