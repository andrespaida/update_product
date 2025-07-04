FROM golang:1.24
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o update_product
EXPOSE 4003
CMD ["./update_product"]