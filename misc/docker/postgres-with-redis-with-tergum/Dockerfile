ARG POSTGRES_VERSION=18
ARG TERGUM_VERSION=v0.46.0
FROM ghcr.io/sikalabs/tergum:$TERGUM_VERSION AS tergum
FROM postgres:$POSTGRES_VERSION
COPY --from=tergum /usr/local/bin/tergum /usr/local/bin/tergum
RUN apt-get update && apt-get install -y ca-certificates redis
