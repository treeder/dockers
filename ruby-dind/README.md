Ruby and Docker in Docker on Alpine.

## Usage

### Basic usage:

You can use this container like any other, it's got Ruby and Docker installed in it.

```sh
docker run --rm -it --privileged treeder/ruby-dind /bin/sh
```

To start Docker inside the container, run `rc default`.

To start Docker and something else in the container:

```sh
docker run --rm -it --privileged treeder/ruby-dind sh -c 'rc default && /bin/sh'
```


### To start container as a daemon then run commands in it

Start container:

```sh
docker run --name dind -it --privileged -d -p 8080-8090:8080-8090 treeder/ruby-dind
```

Then run commands in it:

```sh
docker exec -it dind ruby -v
```

Or to run an app (assuming you have app.rb in directory you started the dind daemon in):

```sh
docker exec -it dind docker run -e "PORT=8080" -p 8080:8080 treeder/hello-sinatra
```

And how about another one at the same time:

```sh
docker exec -it dind docker run -e "PORT=8081" -p 8081:8081 treeder/hello-sinatra
```

## Building this image

Get Docker in Docker files:

```sh
docker export $(docker create rancher/docker) > files.tar
```

```sh
docker build -t treeder/ruby-dind:latest .
```

Tag versions:

Get versions with:

```
docker run --rm treeder/ruby-dind sh -c 'ruby -v; docker version'
```

Then tag with:

```sh
docker tag treeder/ruby-dind:latest treeder/ruby-dind:rX.Y.Z-dA.B.C
```

Push it:

```sh
docker push treeder/ruby-dind
```
