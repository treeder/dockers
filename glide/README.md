For Mastermind glide. 

## Usage

```sh
docker run --rm -it -v $PWD:/go/src/github.com/USERNAME/REPO -w /go/src/github.com/USERNAME/REPO treeder/glide CMD
```

CMD can be any glide command. 

## Building 

```sh
docker build -t treeder/glide .
docker push treeder/glide
```
