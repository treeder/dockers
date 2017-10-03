# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/treeder/dockers/bump
# If dep ever gets decent enough to use, try `dep ensure --vendor-only` from here: https://medium.com/travis-on-docker/triple-stage-docker-builds-with-go-and-angular-1b7d2006cb88
RUN go get -u github.com/Masterminds/glide
ADD glide.* $D/
RUN cd $D && glide install -v
ADD . $D
RUN cd $D && go build -o bump && cp bump /tmp/

# final stage
FROM alpine
RUN apk add --no-cache ca-certificates curl
WORKDIR /script
COPY --from=build-env /tmp/bump /script/bump
ENTRYPOINT ["/script/bump"]
