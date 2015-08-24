FROM alpine

RUN apk update && apk upgrade

RUN apk add curl 
RUN apk add php

# Clean APK cache
RUN rm -rf /var/cache/apk/*
