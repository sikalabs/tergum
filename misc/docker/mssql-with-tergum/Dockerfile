ARG TERGUM_VERSION=v0.44.0
FROM sikalabs/tergum:$TERGUM_VERSION AS tergum
FROM ubuntu:22.04
RUN apt-get update \
    && apt-get install -y curl unzip libicu70 \
    && rm -rf /var/lib/apt/lists/*
RUN curl -sSL https://aka.ms/sqlpackage-linux -o /tmp/sqlpackage.zip \
    && unzip /tmp/sqlpackage.zip -d /opt/sqlpackage \
    && chmod +x /opt/sqlpackage/sqlpackage \
    && ln -sf /opt/sqlpackage/sqlpackage /usr/local/bin/sqlpackage \
    && rm /tmp/sqlpackage.zip
COPY --from=tergum /usr/local/bin/tergum /usr/local/bin/tergum
