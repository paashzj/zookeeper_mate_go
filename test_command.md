### start
```bash
docker run --rm ttbb/zookeeper:mate
```
### start with port expose
```bash
docker run -p 2181:2181 --rm -d ttbb/zookeeper:mate
```
### start tls,port expose
```bash
docker run -p 2181:2181 -p 2182:2182 -e ZOOKEEPER_TLS_ENABLE="true" --rm -d ttbb/zookeeper:mate
```