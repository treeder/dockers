set -ex

docker build -t treeder/glide:latest .
docker push treeder/glide:latest
