# Usa la imagen oficial de Go como base
FROM golang:1.22.2-bullseye

# Crear el directorio de trabajo (donde el profesor montará el volumen)
WORKDIR /app/test

# Copiar archivos de dependencias primero para aprovechar la caché
COPY go.mod ./
# Si tienes go.sum, descomenta la siguiente línea:
# COPY go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# IMPORTANTE: Para que Go no intente escribir en la carpeta montada (solo lectura)
# redirigimos el cache de los tests a una carpeta temporal con permisos
ENV GOCACHE=/tmp/go-cache

# El comando que ejecutará los tests
CMD ["go", "test", "./..."]