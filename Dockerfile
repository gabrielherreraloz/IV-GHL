FROM bitnami/golang

RUN install_packages make

USER 1001

RUN mkdir -p /tmp/go-cache && chmod 777 /tmp/go-cache 
ENV GOCACHE=/tmp/go-cache

WORKDIR /app/test

ENTRYPOINT ["make", "test"] 