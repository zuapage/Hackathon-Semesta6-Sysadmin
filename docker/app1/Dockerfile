FROM golang:alpine
WORKDIR /app
COPY docker/app1/semesta-app1/* ./
RUN go build -o semesta-go
EXPOSE 3000
CMD ["./semesta-go"]
