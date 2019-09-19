FROM golang:1.12-alpine3.10  AS Builder

ENV GOPROXY="https://goproxy.io"
ENV CGO_ENABLED=0

RUN echo "https://mirrors.aliyun.com/alpine/v3.10/main/" > /etc/apk/repositories
RUN echo "https://mirrors.aliyun.com/alpine/v3.10/community/" >> /etc/apk/repositories
RUN apk add tzdata
COPY . /app
WORKDIR /app
RUN go build .

FROM alpine:3.10

WORKDIR /app
COPY --from=Builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=Builder /app/gin_server /app/bin/gin_server
COPY --from=Builder /app/conf /app/conf
EXPOSE 5022
ENTRYPOINT ["/app/bin/gin_server"]