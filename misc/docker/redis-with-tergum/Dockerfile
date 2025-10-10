ARG REDIS_VERSION=8
ARG TERGUM_VERSION=v0.42.0
FROM sikalabs/tergum:$TERGUM_VERSION AS tergum
FROM redis:$REDIS_VERSION
COPY --from=tergum /usr/local/bin/tergum /usr/local/bin/tergum
RUN apt-get update && apt-get install -y ca-certificates redis
