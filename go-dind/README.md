Go and Docker in Docker on Alpine.


### Basic usage:

You can use this container like any other, it's got Ruby and Docker installed in it.

```sh
docker run --rm -it --privileged treeder/go-dind /bin/sh
```

To start Docker inside the container, run `rc default`.

To start Docker and something else in the container:

```sh
docker run --rm -it --privileged treeder/go-dind sh -c 'rc default && /bin/sh'
```

### To start container as a daemon then run commands in it

Start container:

```sh
docker run --name dind -it --privileged -d -p 8080-8090:8080-8090 treeder/go-dind
```

Then run commands in it:

```sh
docker exec -it dind go version
```

Or to run an app (assuming you have app.rb in directory you started the dind daemon in):

```sh
docker exec -it dind docker run -e "PORT=8080" -p 8080:8080 treeder/hello.go
```

And how about another one at the same time:

```sh
docker exec -it dind docker run -e "PORT=8081" -p 8081:8081 treeder/hello.go
```

## Building this image

Get Docker in Docker files:

```sh
docker build -t treeder/go-dind:latest .
```

Tag with Docker version too:

```sh
docker tag treeder/go-dind:latest treeder/go-dind:g1.4.2-d1.8.1
```

Push it:

```sh
docker push treeder/go-dind
```
