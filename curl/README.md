Curl Docker image. Just like curl, but as an image.

## Usage

```sh
docker run --rm treeder/curl -X POST -d "fizz=buzz" http://requestb.in/uaf4hdua
```

## Building this image

```sh
docker build -t treeder/curl:latest .
```

Tag with curl version too (check with `docker run --rm treeder/curl --version`)

```sh
docker tag treeder/curl:latest treeder/curl:X.Y.Z
```

Push it:

```sh
docker push treeder/go-dind
```
