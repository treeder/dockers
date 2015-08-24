FROM alpine

RUN apk update && apk upgrade

RUN apk add curl git make
RUN apk add openssl-dev build-base
RUN apk add ruby-dev ruby ruby-io-console ruby-bundler
# For pg gem:
#RUN apk add postgresql-dev

# Clean APK cache
RUN rm -rf /var/cache/apk/*

ADD update-and-install.sh /scripts/

# Default is to update and install dependencies
CMD ["sh", "-c", "/scripts/update-and-install.sh"]
