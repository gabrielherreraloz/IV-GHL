FROM bitnami/golang

RUN install_packages make

USER 1001

WORKDIR /app/test

ENTRYPOINT ["make", "test"]