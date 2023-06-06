FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/application
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
RUN apk add --no-cache make
RUN make build


FROM alpine

RUN echo -e https://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/application/bin /app/bin
COPY --from=builder /build/application/conf /app/conf

RUN chmod +x bin

EXPOSE 8082

CMD ["./bin/app"]
