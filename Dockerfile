FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY ./src .
RUN go mod download
RUN go build 

FROM alpine:latest AS runner
WORKDIR /app
RUN apk add curl
COPY --from=builder /app/web-service-gin .
EXPOSE 8010
ENTRYPOINT ["./web-service-gin"]
