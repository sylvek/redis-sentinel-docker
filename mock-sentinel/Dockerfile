FROM golang:1.18-alpine AS builder
COPY sentinel.go /
RUN go build /sentinel.go

FROM alpine
EXPOSE 26379
COPY sentinel-get-master-addr-by-name-master-group-name /.
COPY sentinel-sentinels-master-group-name /.
COPY --from=builder /go/sentinel /sentinel
CMD ["/sentinel"]