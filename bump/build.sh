set -ex

USERNAME=treeder
IMAGE=bump

docker build -t $USERNAME/$IMAGE:latest .
