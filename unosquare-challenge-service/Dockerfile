FROM golang:1.17.11 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o unosquare

FROM alpine:3.14.0 as release

WORKDIR /app

COPY --from=builder /app/unosquare /app/

COPY ./api/swagger ./api/swagger

CMD ["./unosquare"]