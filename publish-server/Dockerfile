FROM golang:alpine3.16
WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]