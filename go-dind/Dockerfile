FROM iron/go

RUN echo '@community http://nl.alpinelinux.org/alpine/edge/community' >> /etc/apk/repositories
RUN apk update && apk upgrade

RUN apk add docker@community
# puts file in /etc/runlevels/...
RUN rc-update add docker default

CMD ["/sbin/rc", "default"]

# Clean APK cache
RUN rm -rf /var/cache/apk/*
