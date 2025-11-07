FROM ghcr.io/sikalabs/slu:v0.94.0 AS tergum-downloader
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
ARG TERGUM_VERSION=v0.42.0
RUN slu install-bin tergum --version ${TERGUM_VERSION}

FROM debian:13-slim
COPY --from=tergum-downloader /usr/local/bin/tergum /usr/local/bin/tergum
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
