FROM iron/go

#RUN apk update && apk upgrade
# Clean APK cache
#RUN rm -rf /var/cache/apk/*

# run this to get files.tar
# docker export $(docker create rancher/docker) > files.tar

ADD files.tar /
ENTRYPOINT ["/usr/bin/dockerlaunch", "/usr/bin/docker"]
VOLUME /var/lib/docker
CMD ["-d", "-D"]
