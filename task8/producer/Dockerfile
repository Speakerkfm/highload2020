FROM golang:1.14.2-alpine as builder
WORKDIR /go/src/github.com/speakerkfm/highload2020/task8/producer
COPY . .
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories \
    && apk --no-cache add -U ca-certificates make librdkafka-dev>=1.4.0-r0 pkgconf gcc libc-dev librdkafka>=1.4.0-r0
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o ./build/producer ./cmd/main.go \
    && mv ./build/producer /exe

FROM scratch
COPY --from=builder /exe /
EXPOSE 8080
ENTRYPOINT ["/exe"]