FROM alpine

RUN apk update && apk upgrade
RUN apk add git
RUN rm -rf /var/cache/apk/*

ENTRYPOINT ["git"]
