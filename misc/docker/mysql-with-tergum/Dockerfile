ARG MYSQL_VERSION=latest
ARG TERGUM_VERSION=v0.46.0
FROM ghcr.io/sikalabs/tergum:$TERGUM_VERSION AS tergum
FROM mysql:$MYSQL_VERSION
COPY --from=tergum /usr/local/bin/tergum /usr/local/bin/tergum
