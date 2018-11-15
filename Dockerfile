FROM golang:1.11.2-alpine

RUN apk add --no-cache git build-base curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/iafoosball/auth-service

ENV GO111MODULE=on

ADD ./go.mod ./go.mod
ADD ./go.sum ./go.sum

RUN go mod download

ADD . .

RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o auth-service
CMD ["./auth-service"]