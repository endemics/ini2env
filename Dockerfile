FROM golang:1-buster as builder
COPY main.go /go/src/ini2env/
WORKDIR /go/src/ini2env
RUN go get && go build

FROM alpine:3
RUN apk add --no-cache libc6-compat
COPY --from=builder /go/bin/ini2env /bin/ini2env
ENTRYPOINT [ "/bin/ini2env" ]
