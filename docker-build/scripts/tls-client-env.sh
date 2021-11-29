#!/bin/bash

CERT_DIR=$ZOOKEEPER_HOME/mate/cert
export CLIENT_JVMFLAGS="
-Dzookeeper.clientCnxnSocket=org.apache.zookeeper.ClientCnxnSocketNetty
-Dzookeeper.client.secure=true
-Dzookeeper.ssl.hostnameVerification=false
-Dzookeeper.ssl.keyStore.location=$CERT_DIR/zk_client_key.jks
-Dzookeeper.ssl.keyStore.password=zk_client_pwd
-Dzookeeper.ssl.trustStore.location=$CERT_DIR/zk_client_trust.jks
-Dzookeeper.ssl.trustStore.password=zk_client_pwd"