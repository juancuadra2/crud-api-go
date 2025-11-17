# Etapa de construcción
CMD ["./main"]
# Comando para ejecutar la aplicación

EXPOSE 8080
# Exponer el puerto de la aplicación

COPY --from=builder /app/main .
# Copiar el binario compilado desde la etapa de construcción

WORKDIR /root/

RUN apk --no-cache add ca-certificates
# Instalar certificados CA para HTTPS

FROM alpine:latest
# Etapa de producción

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api
# Compilar la aplicación

COPY . .
# Copiar el código fuente

RUN go mod download
# Descargar las dependencias

COPY go.mod go.sum ./
# Copiar los archivos de dependencias

WORKDIR /app
# Establecer el directorio de trabajo

RUN apk add --no-cache git ca-certificates
# Instalar dependencias necesarias

FROM golang:1.25.3-alpine AS builder

