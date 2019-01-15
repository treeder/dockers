set -ex

USERNAME=treeder
IMAGE=go-dev

docker build -t $USERNAME/$IMAGE:latest .
GOVERS=`docker run --rm -it $USERNAME/$IMAGE:latest go version`
echo $GOVERS
GOVERS="$(cut -d' ' -f3 <<<$GOVERS )"
echo $GOVERS

docker tag $USERNAME/$IMAGE:latest $USERNAME/$IMAGE:$GOVERS

docker push $USERNAME/$IMAGE:$GOVERS
docker push $USERNAME/$IMAGE:latest
