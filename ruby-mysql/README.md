This one has Ruby libmysql in it. 

## Using

```sh
docker run -it --rm treeder/ruby-mysql ruby -v
```

## Building

```sh
docker build -t treeder/ruby-mysql:latest .
```

Tag with versions, check with `docker run -it --rm treeder/ruby-imagemagick convert -version` and
`docker run -it --rm treeder/ruby-imagemagick ruby -v`

```sh
docker tag treeder/ruby-imagemagick:latest treeder/ruby-imagemagick:X.Y.Z-A.B.C
```

X.Y.Z being Ruby version and A.B.C being imagemagic version.


Push:

```sh
docker push treeder/ruby-imagemagick
```
