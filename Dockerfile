FROM debian:10-slim
COPY tergum /
ENTRYPOINT [ "/tergum" ]
