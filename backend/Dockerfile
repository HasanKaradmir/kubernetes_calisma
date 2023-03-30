FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY . .

RUN go build -o webserver .

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/app/ /app/

EXPOSE 444

CMD [ "./webserver" ]