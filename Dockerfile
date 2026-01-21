FROM bitnami/golang

RUN install_packages make

WORKDIR /app/test

ENTRYPOINT ["make", "test"]