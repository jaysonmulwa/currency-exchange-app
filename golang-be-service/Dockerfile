FROM golang:1.18.3-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN go build -o app cmd/server/main.go

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build /app/

WORKDIR /app
CMD ["./app"]