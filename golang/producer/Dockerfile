FROM golang:1.15.7-alpine as builder

RUN apk update && \
    apk add --no-cache build-base && \
    apk add --no-cache upx git ca-certificates tzdata && \
    update-ca-certificates && \
    addgroup --system app && adduser -S -G app app

WORKDIR /home/app/build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a \
    -o mobilepushproducer \
    -ldflags \
    "-s -w \
     -extldflags '-static' \
     -X github.com/kenniston/mobile-push-kafka/golang/producer/cmd.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S%p') \
     -X github.com/kenniston/mobile-push-kafka/golang/producer/cmd.GitCommit=$(git rev-parse HEAD) \
     -X github.com/kenniston/mobile-push-kafka/golang/producer/cmd.Version='0.1'" main.go && \
     upx --ultra-brute -v /home/app/build/mobilepushproducer && \
     upx -t /home/app/build/mobilepushproducer

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /home/app/build/mobilepushproducer /home/app/mobilepushproducer

USER app

ENTRYPOINT ["/home/app/mobilepushproducer", "run"]