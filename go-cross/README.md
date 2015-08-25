
This is simply a helper for this Docker based Go tool: https://github.com/treeder/dockers/tree/master/go

Just for Cross Compiling. See the link above for full set of commands.

### Cross compile:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go-cross cross
```

### Check Go version:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go-cross version
```

## Building this image

```sh
docker build -t treeder/go-cross:latest .
```

Tag it with Go version too (can check with `docker run --rm treeder/go-cross version`):

```sh
docker tag treeder/go-cross:latest treeder/go-cross:1.4.2
```

Push:

```sh
docker push treeder/go-cross
```
