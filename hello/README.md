Simple hello world as shell script (fast).

## Building this image

```sh
docker build -t treeder/hello:sh .
```

Push:

```sh
docker push treeder/hello:sh
```

## To fire this up on a fresh Ubuntu box

Fire up a box on Digital Ocean or wherever and ssh in. Then:

```sh
# Install go:
wget https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.7.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

# Install docker:
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
sudo apt-get install linux-image-extra-$(uname -r) linux-image-extra-virtual
echo "deb https://apt.dockerproject.org/repo ubuntu-xenial main" > /etc/apt/sources.list.d/docker.list
sudo apt-get update
apt-cache policy docker-engine
sudo apt-get install docker-engine
sudo docker run hello-world

# Get timer program:
wget https://raw.githubusercontent.com/treeder/dockers/master/hello/time.go
wget https://raw.githubusercontent.com/treeder/dockers/master/hello/hello.sh
chmod a+x hello.sh
./hello.sh
go run time.go
```
