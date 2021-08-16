FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o zookeeper_mate .


FROM ttbb/zookeeper:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/zookeeper/mate

COPY --from=build /opt/sh/compile/pkg/zookeeper_mate /opt/sh/zookeeper/mate/zookeeper_mate

WORKDIR /opt/sh/zookeeper

CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/zookeeper/mate/scripts/start.sh"]