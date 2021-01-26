# **Mobile Push with Kafka**

This document describes the steps needed to build and run the Mobile Push Sender using Kafka.

*Updated: 25 Jan 2021*

<br/>

## Architecture

<br/>

<p align="center">
  <img src="images/architecture.png"  alt="Architecture"/>
</p>

<br/>

## How to run the project

<br/>

All docker images can be created by docker-compose. To start all containers, use the ***docker-compose up*** command. This command will build all the necessary images and create the containers. 

Below are the hosts and ports for the all containers (docker on localhost):

* ZooKeeper - http://localhost:2181
* Kakfa - http://localhost:9092
* Kafdrop - http://localhost:19000
* Prometheus - http://localhost:9090
* Grafana - http://localhost:3000

<br/>

## Developer Environment

### Tools

* Docker (https://www.docker.com/products/docker-desktop)
* Go (1.15.7 - https://golang.org/dl)
* GoLand (Optional - https://www.jetbrains.com/go/download)
* Intellij IDEA (Optional - https://www.jetbrains.com/pt-br/idea/download)
* Visual Studio Code (Optional - https://code.visualstudio.com/Download)
  - Visual Studio Code Extensions for Go
    - Go for Visual Studio Code (Microsoft - https://github.com/Microsoft/vscode-go.git)

<br/>

# GoLang Projects

***project-name = Procuder or Consumer***

## How to build the *GoLang* project

```
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a \
    -o project-name \
    -ldflags \
    "-s -w \
     -extldflags '-static' \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S%p') \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.GitCommit=$(git rev-parse HEAD) \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.Version='0.1'" main.go && \
     upx --ultra-brute -v /home/app/build/project-name && \
     upx -t /home/app/build/project-name
```

<br/>

## How to optimize the executable size

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
    upx --ultra-brute -v ./project-name && upx -t ./project-name
    ```    
    or
    ```
    upx -9 -v ./project-name && upx -t ./project-name
    ```    

<br/>

## Build from Docker Container (GoLang Container)

The **docker-composer.yml** file has the build section for all GoLang project. This section uses the Dockerfile below to generate the project's executable and optimize it using the upx tool.

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
    -o project-name \
    -ldflags \
    "-s -w \
     -extldflags '-static' \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.BuildTime=$(date -u '+%Y-%m-%d_%H:%M:%S%p') \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.GitCommit=$(git rev-parse HEAD) \
     -X github.com/kenniston/mobile-push-kafka/golang/cmd.Version='0.1'" main.go && \
     upx --ultra-brute -v /home/app/build/project-name && \
     upx -t /home/app/build/project-name

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /home/app/build/project-name /home/app/project-name

USER app

ENTRYPOINT ["/home/app/project-name", "run"]
```

<br/>

## Project Command Line

*The default port can change per project*

```
./project-name run \
    -p [PORT - Optional. Default 6001]
    --log-level [Default: info - (debug, error, trace, info, warning, panic, fatal)]
```

<br/>

# Java Projects

## How to build the *Java* project

*Not implemented yet*

<br/>

# Kafka

*Not implemented yet*

<br/>

# Prometheus

## Prometheus Container

This project uses the official Prometheus container to build a new image. The new image has the config file used to link Prometheus with Kafka JXM port to receive the Kafka's metrics.

Kafka uses the Prometheus JMX Agent (port 8082 in this project) to expose Prometheus metrics through his host.

## Prometheus Architecture
<p align="center">
  <img src="images/prometheus-architecture.svg"  alt="Prometheus Architecture"/>
  <p align="center">https://hub.docker.com/r/prom/prometheus</p>
</p>

<br/>

# Grafana

This project uses the official Grafana container. After run the project using the ***docker-compose up*** the Prometheus's Datasource must be configured.
Use the following URL (docker on localhost) to configure the Prometheus Datasource:

<br/>

http://localhost:3000/datasources/new?utm_source=grafana_gettingstarted

<br/>

Select Prometheus in the Time series database section:

<p align="center">
  <img src="images/grafana-prometheus-datasource-step1.png"  alt="Grafana Setup Step 1"/>
  <p align="center">Grafana Datasources</p>
</p>

<br/>
<br/>

Fill in the fields as the image below: 

<p align="center">
  <img src="images/grafana-prometheus-datasource-step2.png"  alt="Grafana Setup Step 2"/>
  <p align="center">Prometheus Datasource Info</p>
</p>

<br/>
<br/>

Last but not least, set up the Kafka Dashboard in Grafana.
Import the dashboard file from the kakfa folder (grafana-dashboard-kafka-metrics_rev4.json) using a URL below: 

<br/>

http://localhost:3000/dashboard/import

<br/>


<p align="center">
  <img src="images/grafana-prometheus-datasource-step3.png"  alt="Grafana Setup Step 3"/>
  <p align="center">Import the Kafka's Dashboard using the Upload JSON file button</p>
</p>

<br/>
<br/>

# Graylog

*Not implemented yet*

<br/>
