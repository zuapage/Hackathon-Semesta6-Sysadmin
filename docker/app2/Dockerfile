FROM golang:alpine
WORKDIR /app
COPY docker/app2/semesta-app2/* ./
RUN go build -o semesta-go2
EXPOSE 3001
CMD ["./semesta-go2"]
