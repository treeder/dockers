Docker in Docker on Alpine Linux.

## Usage

Start container:

```sh
docker run --name dind -it --privileged -d treeder/dind
```

Then run commands in it:

```sh
docker exec -it dind docker ps
```

## Thanks

To RancherOS for making this: https://github.com/rancher/docker-from-scratch

## Building this image

Get Docker in Docker files:

```sh
docker export $(docker create rancher/docker:1.8.1) > files.tar
```

```sh
docker build -t treeder/dind:latest .
```

Tag with Docker version too:

```sh
docker tag treeder/dind:latest treeder/dind:X.Y.Z
```

Push it:

```sh
docker push treeder/dind
```
