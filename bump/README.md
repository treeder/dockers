Bumps version files. 


## Usage

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump CMD
```

CMD can be any:

* patch
* minor
* major

## Building 

```sh
docker build -t treeder/bump .
docker push treeder/bump
```
