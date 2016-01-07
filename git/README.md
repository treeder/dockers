Git Docker image. Just like git, but as an image. ie: you don't need to install git. 

## Usage

```sh
docker run --rm treeder/git commit -am "yo, this is my commit"
```

## Building this image

```sh
docker build -t treeder/git:latest .
```

Tag with curl version too (check with `docker run --rm treeder/git --version`)

```sh
docker tag treeder/git:latest treeder/git:X.Y.Z
```

Push it:

```sh
docker push treeder/git
```
