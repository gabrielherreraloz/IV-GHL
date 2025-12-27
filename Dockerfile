FROM golang:1.25.5-alpine3.23

# Es necesario instalar make y gcc para que "make test" funcione
RUN apk add --no-cache make gcc musl-dev

# Directorio indicado en la documentación del objetivo
WORKDIR /app/test

# Creamos la carpeta de cache y le damos permisos al usuario 1001
RUN mkdir -p /tmp/go-cache && chmod -R 777 /tmp/go-cache
ENV GOCACHE=/tmp/go-cache

COPY . .
RUN go mod download

# Comando ejecutado al lanzar el contenedor
ENTRYPOINT ["make", "test"]