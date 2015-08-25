# To test this script before building an image:
# docker run --rm -v $PWD:/app -w /app golang:1.4.2-cross sh go.sh version

# Original command to vendor
# docker run --rm -it -v "$PWD":/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang go get
# Original command to build
# docker run --rm -it -v $PWD:/go/src/x/y/z -w /go/src/x/y/z -e "GOPATH=/go/src/x/y/z/vendor:/go" golang go build -o hello
# Original command to cross compile:

cmd="$1"
# echo "Args: $*"
if [ "$#" -lt 1 ]
then
    echo "No command provided."
    exit 1
fi

wd=$PWD
p=/go/src/x/y/z
mkdir -p $p
cp -r * $p
cd $p
# Add vendor to the GOPATH so get will pull it in the right spot
export GOPATH=$p/vendor:/go
# env

vendor () {
  go get
  cp -r -u $1/vendor $2
  chmod -R a+rw $2/vendor
  #      cd $wd
}
build () {
  echo "build $1 $2"
  go $1
  ls -al
  cp -r -u * $2
}

case "$1" in
  vendor)  echo "Vendoring dependencies..."
      vendor $p $wd
      ;;
  build)  echo  "Building..."
      build $@ $wd
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
  remote) echo  "Building binary from $2"
      pwd
      userwd=$wd
      cd
      git clone $2 repo
      cd repo
      ls -al
      wd=$PWD
      # Need to redo some initial setup here:
      cp -r * $p
      cd $p
      vendor $p $wd
      build "build -o app" $wd
      ls -al $wd
      cp $wd/app $userwd
      echo "userwd $userwd"
      ls -al $userwd
      chmod a+rwx $userwd/app
      ;;
  version)
      go version
      ;;
  *) echo "Invalid command, see https://github.com/treeder/dockers/tree/master/go for reference."
      ;;
esac
exit 0
