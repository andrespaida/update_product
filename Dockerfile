# Imagen base
FROM golang:1.24

# Directorio de trabaj
WORKDIR /app

# Copiar los archivos del proyecto
COPY . .

# Descargar dependencias
RUN go mod download

# Compilar el binario
RUN go build -o update_product

# Exponer el puerto
EXPOSE 4003

# Comando para ejecutar el microservicio
CMD ["./update_product"]