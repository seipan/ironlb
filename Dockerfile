FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /docker-load-balancer ./cmd/load-balancer

FROM alpine:3.18

COPY --from=builder /docker-load-balancer /docker-load-balancer

CMD ["/docker-load-balancer"]
