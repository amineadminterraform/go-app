#build stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 8080
CMD [ "/app/main" ]
#run stage
FROM alpine:3.19
WORKDIR /app
EXPOSE 8080
COPY . .
COPY --from=builder /app/main .
CMD [ "/app/main" ]
