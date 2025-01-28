# Build stage
FROM golang:1.18.2-alpine3.16 
WORKDIR /app
COPY . .
RUN go build -o main main.go
CMD ["/app/main"]
