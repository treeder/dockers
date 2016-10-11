set -ex

USERNAME=treeder
IMAGE=bump

docker run --rm -v "$PWD":/go/src/github.com/treeder/bump -w /go/src/github.com/treeder/bump iron/go:dev sh -c 'go build -o bump'
docker build -t $USERNAME/$IMAGE:latest .
