FROM golang:1.14

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /go/src/github.com/dell/nigha-monitor
WORKDIR /go/src/github.com/dell/nigha-monitor

RUN ls -l
RUN go build .

FROM alpine
MAINTAINER  The Prometheus Authors <prometheus-developers@googlegroups.com>

COPY --from=0 /go/src/github.com/dell/nigha-monitor  /bin/nigha-monitor
ENTRYPOINT ["/bin/nigha-monitor"]