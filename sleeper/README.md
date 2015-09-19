Sleeps for ENV['SLEEP'] amount of time (or 60 seconds).

```sh
docker run -it --rm -e "SLEEP=10" treeder/sleeper
```

## Building this image

```sh
docker build -t treeder/sleeper:latest .
```

```sh
docker push treeder/sleeper
```
