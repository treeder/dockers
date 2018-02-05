# build stage
FROM golang:1.9-alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ENV D=/go/src/github.com/treeder/dockers/bump
# If dep ever gets decent enough to use, try `dep ensure --vendor-only` from here: https://medium.com/travis-on-docker/triple-stage-docker-builds-with-go-and-angular-1b7d2006cb88
RUN go get -u github.com/golang/dep/cmd/dep
ADD Gopkg.* $D/
RUN cd $D && dep ensure --vendor-only
ADD . $D
RUN cd $D && go build -o bump && cp bump /tmp/

# final stage
FROM alpine:3.7
RUN apk add --no-cache ca-certificates curl
WORKDIR /app
COPY --from=build-env /tmp/bump /script/bump
ENTRYPOINT ["/script/bump"]
