FROM golang:1.4.2-cross

ADD https://raw.githubusercontent.com/treeder/dockers/master/go/go.sh /scripts/

ADD https://raw.githubusercontent.com/treeder/dockers/master/go/app.go /app/app.go
WORKDIR /app

# Default is to update and install dependencies
ENTRYPOINT ["sh", "/scripts/go.sh"]
