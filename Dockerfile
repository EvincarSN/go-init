FROM golang:latest
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=windows go build -ldflags="-w -s" -o api cmd/api/main.go
CMD ["./api"]