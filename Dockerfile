# Usar una imagen base oficial de Go 1.15
FROM golang:1.15 as builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum para manejar dependencias
# Asumiendo que go.mod y go.sum están en el directorio raíz del proyecto
COPY go.mod go.sum ./

# Descargar todas las dependencias
RUN go mod download

# Copiar el resto del código fuente de la aplicación
# Ajustar esta parte para que apunte al subdirectorio donde está el código
COPY cdm/ ./cdm/

# Cambiar al directorio donde está el main.go
WORKDIR /app/cdm/api/app

# Compilar la aplicación. Asegúrate de adaptar el nombre del archivo main si es necesario.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Etapa final para obtener una imagen limpia y ligera
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar el ejecutable compilado desde la primera etapa
COPY --from=builder /app/cdm/api/app/myapp .

# Exponer el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./myapp"]
