ARG POSTGRES_VERSION=13
ARG TERGUM_VERSION=v0.18.0
FROM sikalabs/tergum:$TERGUM_VERSION as tergum
FROM postgres:$POSTGRES_VERSION
COPY --from=tergum /usr/local/bin/tergum /usr/local/bin/tergum
