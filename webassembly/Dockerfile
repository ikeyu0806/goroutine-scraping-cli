FROM golang:1.21-bullseye

WORKDIR /app

RUN go mod init go-webassembly

RUN GOARCH=wasm GOOS=js go build -o main.wasm
