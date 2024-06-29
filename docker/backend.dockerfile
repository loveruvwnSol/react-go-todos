FROM golang:1.19.1-alpine

COPY ./backend/go.mod /go/src/app/go.mod
COPY ./backend/go.sum /go/src/app/go.sum

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@v1.29.0

WORKDIR /go/src/app

CMD ["air"]