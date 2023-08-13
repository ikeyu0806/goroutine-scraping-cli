FROM golang:1.21-bullseye

COPY ./ /app/frontend

WORKDIR /app

RUN cd /app/frontend/webassembly && GOARCH=wasm GOOS=js go build -o main.wasm

CMD ["go","run","/app/frontend/server/main.go"]
