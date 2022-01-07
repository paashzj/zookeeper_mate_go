#!/bin/bash

if [ $ZOOKEEPER_TLS_ENABLE == "true" ]; then
    CERT_DIR=$ZOOKEEPER_HOME/mate/cert
    export SERVER_JVMFLAGS="
    -Dzookeeper.serverCnxnFactory=org.apache.zookeeper.server.NettyServerCnxnFactory
    -Dzookeeper.ssl.keyStore.location=$CERT_DIR/zk_server_key.jks
    -Dzookeeper.ssl.keyStore.password=zk_server_pwd
    -Dzookeeper.ssl.trustStore.location=$CERT_DIR/zk_server_trust.jks
    -Dzookeeper.ssl.trustStore.password=zk_server_pwd
    -Dzookeeper.ssl.hostnameVerification=false
    -Dzookeeper.ssl.quorum.keyStore.location=$CERT_DIR/zk_server_key.jks
    -Dzookeeper.ssl.quorum.keyStore.password=zk_server_pwd
    -Dzookeeper.ssl.quorum.trustStore.location=$CERT_DIR/zk_server_trust.jks
    -Dzookeeper.ssl.quorum.trustStore.password=zk_server_pwd
    -Dzookeeper.ssl.quorum.hostnameVerification=false
    "
fi
/bin/bash $ZOOKEEPER_HOME/bin/zkServer.sh start
