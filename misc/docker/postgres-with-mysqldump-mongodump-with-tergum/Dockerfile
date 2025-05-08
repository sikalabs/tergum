ARG POSTGRES_VERSION=16
ARG TERGUM_VERSION=v0.40.0
FROM ghcr.io/tergum-extra/postgres-with-tergum:${POSTGRES_VERSION}-tergum-${TERGUM_VERSION}
# Install mysqldump
RUN apt-get update && apt-get install -y mariadb-client
# Install mongodump
RUN \
  apt-get install -y gnupg curl && \
  curl -fsSL https://www.mongodb.org/static/pgp/server-7.0.asc | gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor && \
  echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-7.0.list && \
  apt-get update && \
  apt-get install -y mongodb-org-tools
