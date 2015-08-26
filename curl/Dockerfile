FROM alpine

RUN apk update && apk upgrade
RUN apk add curl
RUN rm -rf /var/cache/apk/*

ENTRYPOINT ["curl"]
