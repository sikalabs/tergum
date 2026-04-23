ARG POSTGRES_VERSION=17
ARG TERGUM_VERSION=v0.44.0
FROM ghcr.io/tergum-extra/postgres-with-tergum:${POSTGRES_VERSION}-tergum-${TERGUM_VERSION}
# Install mysqldump
RUN apt-get update && apt-get install -y mariadb-client
# Install mongodump (direct .deb; MongoDB apt repo unusable on Debian 13/trixie due to SHA1 key rejection by sqv)
ARG MONGO_TOOLS_VERSION=100.16.0
RUN \
  apt-get install -y curl && \
  ARCH=$(dpkg --print-architecture) && \
  case "$ARCH" in \
    amd64) MONGO_ARCH=x86_64 ;; \
    arm64) MONGO_ARCH=arm64 ;; \
    *) echo "Unsupported arch: $ARCH" && exit 1 ;; \
  esac && \
  curl -fsSL "https://fastdl.mongodb.org/tools/db/mongodb-database-tools-ubuntu2204-${MONGO_ARCH}-${MONGO_TOOLS_VERSION}.deb" -o /tmp/mdt.deb && \
  apt-get install -y /tmp/mdt.deb && \
  rm /tmp/mdt.deb
