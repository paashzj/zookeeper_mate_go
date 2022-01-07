#!/bin/bash

if [ $ZOOKEEPER_TLS_ENABLE == "true" ]; then
    CERT_DIR=$ZOOKEEPER_HOME/mate/cert
    export SERVER_JVMFLAGS="
    -Dzookeeper.serverCnxnFactory=org.apache.zookeeper.server.NettyServerCnxnFactory
    -Dzookeeper.ssl.keyStore.location=$CERT_DIR/zk_server_key.jks
    -Dzookeeper.ssl.keyStore.password=zk_server_pwd
    -Dzookeeper.ssl.trustStore.location=$CERT_DIR/zk_server_trust.jks
    -Dzookeeper.ssl.trustStore.password=zk_server_pwd
    -Dzookeeper.quorum.keyStore.location=$CERT_DIR/zk_server_key.jks
    -Dzookeeper.quorum.keyStore.password=zk_server_pwd
    -Dzookeeper.quorum.trustStore.location=$CERT_DIR/zk_server_trust.jks
    -Dzookeeper.quorum.trustStore.password=zk_server_pwd"
fi
/bin/bash $ZOOKEEPER_HOME/bin/zkServer.sh start
