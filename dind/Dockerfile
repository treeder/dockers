FROM alpine

ADD files.tar /
ENTRYPOINT ["/usr/bin/dockerlaunch", "/usr/bin/docker"]
VOLUME /var/lib/docker
CMD ["-d", "-D"]

# Clean APK cache
RUN rm -rf /var/cache/apk/*
