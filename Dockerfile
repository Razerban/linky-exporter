FROM golang:1.19-alpine AS builder

COPY . /go/src/linky-exporter

WORKDIR /go/src/linky-exporter

RUN apk add --no-cache make curl tar && \
    go mod vendor && \
    make build

FROM alpine:latest
LABEL maintainer="Ahmed Abdelkafi <abdelkafiahmed@yahoo.fr>"

COPY --from=builder /go/src/linky-exporter/linky-exporter /bin/linky-exporter

ENTRYPOINT ["/bin/linky-exporter"]
EXPOSE     9901
