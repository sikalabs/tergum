ARG POSTGRES_VERSION=17
ARG TERGUM_VERSION=v0.41.0
FROM ghcr.io/tergum-extra/postgres-with-tergum:${POSTGRES_VERSION}-tergum-${TERGUM_VERSION}
# Install mysqldump
RUN apt-get update && apt-get install -y mariadb-client
