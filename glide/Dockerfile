FROM iron/go:dev

ENV VERSION=0.12.3

WORKDIR /app

RUN mkdir -p /go/src/github.com/Masterminds
ENV GOPATH=/go

# RUN cd /go/src/github.com/Masterminds && git clone https://github.com/Masterminds/glide.git && cd glide && go build

RUN apk add --no-cache wget tar
RUN cd /go/src/github.com/Masterminds && wget https://github.com/Masterminds/glide/archive/v$VERSION.tar.gz && tar -zxvf v$VERSION.tar.gz && mv glide-$VERSION glide && cd glide && go build -o glide
RUN cp /go/src/github.com/Masterminds/glide/glide /bin

ENTRYPOINT ["glide"]
