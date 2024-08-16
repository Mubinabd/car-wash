FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gateway .

FROM alpine:latest
WORKDIR /app
COPY .env .

COPY --from=builder /app/gateway .
COPY --from=builder /app//load/model.conf ./load/
COPY --from=builder /app/load/policy.csv ./load/

EXPOSE 8070

CMD ["./gateway"]