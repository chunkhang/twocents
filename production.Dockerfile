FROM golang:1.11.2-alpine AS builder

RUN apk update && apk add --no-cache git

COPY . /twocents
WORKDIR /twocents

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /twocents/main

FROM scratch

COPY --from=builder /twocents/main /twocents/main

ENTRYPOINT ["/twocents/main"]
