FROM golang:alpine

RUN apk add --no-cache git
RUN go get -u github.com/golang/dep/cmd/dep

RUN dep version
ENTRYPOINT ["dep"]
