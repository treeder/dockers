
This is a Docker image to help you develop in Go (golang). The great thing is you don't need
to have anything install except Docker, you don't even need Go installed. See [this post about developing with Docker](https://medium.com/iron-io-blog/why-and-how-to-use-docker-for-development-a156c1de3b24).

This image can perform the following functions:

* vendor - vendors your dependencies into a /vendor directory.  
* build - builds your program using the vendored dependencies.
* cross - cross compile your program into a variety of platforms. Based on [this](https://medium.com/iron-io-blog/how-to-cross-compile-go-programs-using-docker-beaa102a316d#95d9).
* static - statically compile your program. This is great for making the [tiniest Docker image possible](http://www.iron.io/blog/2015/07/an-easier-way-to-create-tiny-golang-docker-images.html).

## Usage

Vendor dependencies:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go vendor
```

Build:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go build -o myapp
```

Cross compile:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go cross
```

Build static binary:

```sh
docker run --rm -v $PWD:/app -w /app treeder/go static
```
