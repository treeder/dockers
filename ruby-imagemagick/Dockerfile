FROM iron/ruby

RUN apk update && apk upgrade

RUN apk add imagemagick

# Clean APK cache
RUN rm -rf /var/cache/apk/*
