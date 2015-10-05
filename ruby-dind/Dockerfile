FROM iron/ruby-bundle

RUN echo '@community http://nl.alpinelinux.org/alpine/edge/community' >> /etc/apk/repositories
#RUN sed -i -e 's/v3\.2/edge/g' /etc/apk/repositories; apk update;

RUN apk update && apk upgrade

RUN apk add docker@community
# puts file in /etc/runlevels/...
RUN rc-update add docker default
#RUN rc default

# run this to get files.tar
# docker export $(docker create rancher/docker) > files.tar

#ADD files.tar /
#ENTRYPOINT ["/usr/bin/dockerlaunch", "/usr/bin/docker -d -D", "/bin/sh"]
#ENTRYPOINT ["/usr/bin/dockerlaunch", ]
#ENTRYPOINT ["rc default", "/bin/sh"]
#VOLUME /var/lib/docker
#CMD ["-d", "-D"]

CMD ["/sbin/rc", "default"]

# Clean APK cache
RUN rm -rf /var/cache/apk/*
