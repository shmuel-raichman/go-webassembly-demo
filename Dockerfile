FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN GOOS=js GOARCH=wasm go build -o http/wa.wasm

WORKDIR /app/http

CMD go run http.go