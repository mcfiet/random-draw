FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM alpine:3.20

COPY entrypoint.sh /

COPY --from=builder /docker-gs-ping /docker-gs-ping

EXPOSE 3000

ENTRYPOINT [ "/entrypoint.sh" ]

CMD ["/docker-gs-ping"]
