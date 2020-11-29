FROM golang:latest AS builder
WORKDIR /go/src/github.com/speakerkfm/highload2020/task8/consumer
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/consumer ./cmd/main.go \
    && mv ./build/consumer /exe

FROM scratch
COPY --from=builder /exe /
ENTRYPOINT ["/exe"]