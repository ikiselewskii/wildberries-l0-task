FROM golang:alpine-latest
WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080 8081

CMD ["./main"   ]