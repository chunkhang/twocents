FROM golang:1.11.2-alpine

RUN apk update && apk add --no-cache git

COPY . /go/src/github.com/chunkhang/twocents
WORKDIR /go/src/github.com/chunkhang/twocents

RUN go get -d -v

RUN go get github.com/pilu/fresh

CMD ["fresh"]
