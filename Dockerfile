FROM golang:alpine as builder

WORKDIR /build

COPY go.mod .
COPY cmd/main.go .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/app

FROM alpine

WORKDIR /app

COPY --from=builder /build/bin/app /bin
ENTRYPOINT ["./bin"]
