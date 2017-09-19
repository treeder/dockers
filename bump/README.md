# Bump

Bumps version files.

## Usage

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump [--filename FILENAME] [CMD]
```

CMD can be one of:

* patch - default
* minor
* major
