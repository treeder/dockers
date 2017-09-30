set -ex

USERNAME=treeder
IMAGE=strings

docker run --rm -v "$PWD":/go/src/github.com/treeder/strings -w /go/src/github.com/treeder/strings golang:alpine sh -c 'go build -o strings'
docker build -t $USERNAME/$IMAGE:latest .
