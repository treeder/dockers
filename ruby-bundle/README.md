This image is for building your Ruby dependencies. Just run it and it will update your gems and vendor your dependencies
for you.

You should use treeder/ruby to run them (way, way smaller image).

## Using

Just run this image in the directory where you Gemfile is.

```sh
docker run --rm -v $PWD:/app -w /app treeder/bundle
```

That will do a `bundle update` and `bundle install --standalone`.

You can also run commands explicitly too if you'd like, for example: 

```sh
docker run --rm -v $PWD:/app -w /app treeder/bundle bundle update
docker run --rm -v $PWD:/app -w /app treeder/bundle bundle install --standalone --clean
docker run --rm -v $PWD:/app -w /app treeder/bundle chmod -R a+rw .bundle
docker run --rm -v $PWD:/app -w /app treeder/bundle chmod -R a+rw bundle
```

## Building

```sh
docker build -t treeder/bundle:latest .
```

Push:

```sh
docker push treeder/bundle
```
