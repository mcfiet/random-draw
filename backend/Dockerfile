FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM scratch

COPY --from=builder /docker-gs-ping /docker-gs-ping

EXPOSE 3000

CMD ["/docker-gs-ping"]
