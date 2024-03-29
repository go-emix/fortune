FROM golang:latest as build

WORKDIR /go/src/fortune

COPY . .

ENV GOPROXY https://proxy.golang.com.cn,direct

RUN go get

RUN go build

FROM ubuntu:20.04

WORKDIR /opt/fortune

COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /opt/fortune/zoneinfo.zip

COPY --from=build /go/src/fortune/fortune  /opt/fortune/fortune
COPY --from=build /go/src/fortune/conf  /opt/fortune/conf
COPY --from=build /go/src/fortune/frontend  /opt/fortune/frontend

#RUN sed -i 's#http://archive.ubuntu.com#http://mirrors.aliyun.com#g' /etc/apt/sources.list

RUN apt install -y --no-install-recommends ca-certificates curl

ENV GIN_MODE release

ENV ZONEINFO /opt/fortune/zoneinfo.zip

ENTRYPOINT ["./fortune","run"]