FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=on

COPY . /app

RUN go build -o main .

EXPOSE 444

CMD ["/app/main"]
