FROM golang:1.21-bullseye

WORKDIR /app

COPY frontend /app/frontend

RUN cd /app/frontend/webassembly && GOARCH=wasm GOOS=js go build -o main.wasm

CMD ["go","run","/app/frontend/cmd/main.go"]
