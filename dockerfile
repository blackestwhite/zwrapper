FROM golang:1.19.0-alpine3.16 as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /app/app .

FROM alpine:3.13.1
WORKDIR /app
COPY .env .
COPY templates/ templates/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["/app/app"]