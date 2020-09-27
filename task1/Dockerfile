FROM golang:latest AS builder
WORKDIR /go/src/github.com/speakerkfm/highload2020/task1
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/weather_service ./cmd/weather_service/main.go \
    && mv ./build/weather_service /exe

FROM scratch
COPY --from=builder /exe /
ENV WEATHER_PORT=8000
ENV OPEN_WEATHER_MAP_HOST=community-open-weather-map.p.rapidapi.com
ENV OPEN_WEATHER_MAP_API_KEY=67acb54007mshaceca7633d2cfe2p1e6770jsnae5a84b6050f
ENTRYPOINT ["/exe"]