
This is a Docker image to help you develop in Go (golang). The great thing is you don't need
to have anything install except Docker, you don't even need Go installed. See [this post about developing with Docker](https://medium.com/iron-io-blog/why-and-how-to-use-docker-for-development-a156c1de3b24).

This image can perform the following functions:

* vendor - vendors your dependencies into a /vendor directory.  
* build - builds your program using the vendored dependencies.
* cross - cross compile your program into a variety of platforms. Based on [this](https://medium.com/iron-io-blog/how-to-cross-compile-go-programs-using-docker-beaa102a316d#95d9).
* static - statically compile your program. This is great for making the [tiniest Docker image possible](http://www.iron.io/blog/2015/07/an-easier-way-to-create-tiny-golang-docker-images.html).
* remote - this one will produce a binary from a github repo. Equivalent to cloning, vendoring and building.

## Usage

### Vendor dependencies:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go vendor
```

### Build:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go build -o app
```

### Run it:

```sh
docker run --rm -v $PWD:/app -w /app golang ./app
```

### Cross compile:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go cross
```

### Build static binary:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go static
```

### Build a remote git repo:

This one will get a remote repo and produce a binary:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go remote https://github.com/treeder/go-docker.git
```

You'll end up with a binary called `app` which you can run with the same command as above.

### Check Go version:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go version
```

## TODO:

Cross compile works if we use the big golang:1.4.2-cross. But do you need to cross compile if you're just using it inside Docker?
