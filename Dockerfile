FROM golang:1.16.5-alpine3.14 as builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
ENV APP_PORT=8080
COPY --from=builder /app/main /app/
EXPOSE 8080
CMD ["/app/main"]
