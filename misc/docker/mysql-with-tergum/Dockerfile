ARG MYSQL_VERSION=latest
ARG TERGUM_VERSION=v0.39.0
FROM sikalabs/tergum:$TERGUM_VERSION as tergum
FROM mysql:$MYSQL_VERSION
COPY --from=tergum /usr/local/bin/tergum /usr/local/bin/tergum
