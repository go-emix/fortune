FROM golang:latest as build

WORKDIR /go/src/fortune

COPY . .

ENV GOPROXY https://goproxy.io,direct

RUN go get

RUN go build

FROM ubuntu:20.04

WORKDIR /opt/fortune

COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /opt/fortune/zoneinfo.zip

COPY --from=build /go/src/fortune/fortune  /opt/fortune/fortune
COPY --from=build /go/src/fortune/conf  /opt/fortune/conf
COPY --from=build /go/src/fortune/frontend  /opt/fortune/frontend

RUN apt update

RUN apt install -y --no-install-recommends ca-certificates curl

ENV GIN_MODE release

ENV ZONEINFO /opt/fortune/zoneinfo.zip

ENTRYPOINT ["./fortune","run"]