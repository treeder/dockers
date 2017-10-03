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

Or to pull it out of your last git commit, you can add `[bump major]` or `[bump minor]` to your git commit message, then use:

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump --filename $version_file "$(git log -1 --pretty=%B)"
```
