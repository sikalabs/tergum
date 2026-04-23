ARG MONGO_VERSION=latest
ARG TERGUM_VERSION=v0.44.0
FROM ghcr.io/sikalabs/tergum:$TERGUM_VERSION AS tergum
FROM mongo:$MONGO_VERSION
COPY --from=tergum /usr/local/bin/tergum /usr/local/bin/tergum
