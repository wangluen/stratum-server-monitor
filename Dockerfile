FROM golang:alpine AS build-env
RUN apk --no-cache add git
RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/github.com/dubuqingfeng/stratum-server-monitor
RUN cd /go/src/github.com/dubuqingfeng/stratum-server-monitor && \
   dep ensure -v && \
   go build -v -o /src/bin/stratum-server-monitor main.go
   
FROM alpine
RUN apk --no-cache add openssl ca-certificates tzdata
WORKDIR /app
COPY --from=build-env /src/bin /app/
COPY --from=build-env /go/src/github.com/dubuqingfeng/stratum-server-monitor/configs /app/configs
COPY --from=build-env /go/src/github.com/dubuqingfeng/stratum-server-monitor/logs /app/logs
ENTRYPOINT ./stratum-server-monitor
