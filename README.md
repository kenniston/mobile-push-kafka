# Mobile Push with Kafka

This document describes the steps needed to build and run the Mobile Push Sender using Kafka.

*Updated: 24 Jan 2021*

## Developer Environment

### Tools

* Go (1.15.7 - https://golang.org/dl/)
* GoLand (Optional - https://www.jetbrains.com/go/download)
* Visual Studio Code (Optional - https://code.visualstudio.com/Download)
  - Visual Studio Code Extensions for Go
    - Go for Visual Studio Code (Microsoft - https://github.com/Microsoft/vscode-go.git)

### GoLang Sender

#### How to build the *GoLang* Sender project

```
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a \
    -o mobilepushsender \
    -ldflags \
    "-s -w \
     -extldflags '-static' \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S%p') \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.GitCommit=$(git rev-parse HEAD) \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.Version='0.1'" main.go && \
     upx --ultra-brute -v /home/app/build/mobilepushsender && \
     upx -t /home/app/build/mobilepushsender
```

#### How to optimize Push Sender executable size

* Install UPX (Ultimate Packer for eXecutables - Compress/expand executable files)
    - MacOS X and Homebrew:
        ```
        brew install upx
        ```
    - Linux:
        ```
        apt-get install -y upx
        ```    
    -  Windows:
        ```
        choco install upx
        ```
    
* Compress the server executable with UPX:

    ```
    upx --ultra-brute -v ./mobilepushsender && upx -t ./mobilepushsender
    ```    
    or
    ```
    upx -9 -v ./mobilepushsender && upx -t ./mobilepushsender
    ```    

#### Build from Docker Container (GoLang Container)
```
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
    -o mobilepushsender \
    -ldflags \
    "-s -w \
     -extldflags '-static' \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S%p') \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.GitCommit=$(git rev-parse HEAD) \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.Version='0.1'" main.go && \
     upx --ultra-brute -v /home/app/build/mobilepushsender && \
     upx -t /home/app/build/mobilepushsender

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /home/app/build/mobilepushsender /home/app/mobilepushsender

USER app

ENTRYPOINT ["/home/app/mobilepushsender", "run"]
```

#### Sender Command Line

```
    ./mobilepushsender run \
        -p [PORT - Optional. Default 6001]
        --log-level [Default: info - (debug, error, trace, info, warning, panic, fatal)]
```

### Java Sender

#### How to build the *Java* Sender project

*Not implemented yet*

### Kafka Container

*Not implemented yet*

### Prometheus Container

*Not implemented yet*

<p align="center">
  <img src="images/prometheus-architecture.svg"  alt="Prometheus Architecture"/>
  https://hub.docker.com/r/prom/prometheus
</p>

### Grafana Container

*Not implemented yet*

## Docker Compose File

```
version: '3.5'

networks:
  mobilepush:

services:
  zookeeper:
    image: zookeeper:3.6.2
    container_name: zookeper-service
    restart: always
    ports:
      - 2181:2181
    environment:
      ZOOKEEPER_TICK_TIME: 2000
    networks:
      - mobilepush

  kafka:
    build:
      context: ./kafka
      dockerfile: Dockerfile
    image: kenniston/kafka:latest
    container_name: kafka-service
    ports:
      - 9092:9092
      - 8082:8082
    environment:
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_HOST_NAME: 192.168.1.111
      KAFKA_LISTENERS: PLAINTEXT://:9092
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:9092
      KAFKA_CREATE_TOPICS: "MobileSendPush:4:1,MobilePushResult:4:1"
      EXTRA_ARGS: -javaagent:/usr/app/jmx_prometheus_javaagent.jar=8082:/usr/app/prom-jmx-agent-config.yml
    networks:
      - mobilepush
    depends_on:
      - zookeeper

  kafdrop:
    image: obsidiandynamics/kafdrop:3.27.0
    container_name: kafdrop-service
    ports:
      - 19000:9000
    environment:
      KAFKA_BROKERCONNECT: kafka:9092    
    networks: 
      - mobilepush
    depends_on:
      - kafka
    
  prometheus:
    build:
      context: ./prometheus
      dockerfile: Dockerfile
    image: kenniston/prometheus:latest
    container_name: prometheus-service
    ports:
      - "9090:9090"
    networks:
      - mobilepush
    depends_on:
      - kafka

  grafana:
    image: grafana/grafana:7.3.7
    container_name: grafana-service
    ports:
      - "3000:3000"
    networks:
      - mobilepush
    depends_on:
      - prometheus

  golang-sender:
    build:
      context: ./golang
      dockerfile: Dockerfile
    image: golang-sender:latest
    container_name: golang-sender-service
    stdin_open: true
    tty: true
    ports:
      - "6001:6001"
    environment:
      - SRV_LOG_LEVEL=debug
    restart: on-failure:3
    networks:
      - mobilepush
```


