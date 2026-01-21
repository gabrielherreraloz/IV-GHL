FROM golang:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    make \
    && rm -rf /var/lib/apt/lists/

RUN useradd -r test_usr

USER test_usr

WORKDIR /app/test

ENTRYPOINT ["make", "test"]