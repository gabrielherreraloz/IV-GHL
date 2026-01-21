FROM bitnami/golang

RUN install_packages make

RUN mkdir -p /tmp/go-cache && chmod 777 /tmp/go-cache
ENV GOCACHE=/tmp/go-cache

USER 1001

WORKDIR /app/test

ENTRYPOINT ["make", "test"]