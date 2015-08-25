# To test this script before building an image:
#  docker run --rm -v $PWD:/app -w /app golang:1.4.2-cross sh go.sh vendor

# Original command to vendor
# docker run --rm -it -v "$PWD":/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang go get
# Original command to build
# docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang go build -o hello
# Original command to cross compile:

cmd="$1"
echo "Args: $*"
if [ "$#" -gt 0 ]
then
    echo "There's Beans"
else
    echo "No command provided."
    exit 1
fi

wd=$PWD
p=/go/src/x/y/z
mkdir -p $p
cp -r * $p
cd $p
# Add vendor to the GOPATH so get will pull it in
export GOPATH=$p/vendor:/go
env

case "$1" in
  vendor)  echo "Vendoring dependencies..."
      go get
      cp -r -u $p/vendor $wd
      chmod -R a+rw $wd/vendor
#      cd $wd
      ;;
  build)  echo  "Building..."
      go "$@"
      cp -r -u * $wd
      ;;
  cross)  echo  "Cross compiling..."
      for GOOS in darwin linux windows; do
        for GOARCH in 386 amd64; do
        echo "Building $GOOS-$GOARCH"
        export GOOS=$GOOS
        export GOARCH=$GOARCH
        go build -o bin/app-$GOOS-$GOARCH
        done
      done
      cp -r bin $wd
      chmod -R a+rw $wd/bin
#      ls -al $wd/bin
      ;;
  static) echo  "Building static binary..."
      CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o static
      cp static $wd
      ;;
  *) echo "Invalid command, see https://github.com/treeder/dockers/tree/master/go for reference."
      ;;
esac
exit 0
