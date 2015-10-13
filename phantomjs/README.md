


I attempted to build on alpine with the following installed:

apks: build-base, libjpeg-turbo-dev
libpng-dev
flex
bison
libssl1.0
ruby
perl
gperf
sqlite-dev
python
fontconfig-dev
libxext-dev
icu-dev
openssl-dev

but it errored out with:

```
In file included from thread/qmutex.cpp:654:0:
thread/qmutex_linux.cpp:51:25: fatal error: linux/futex.h: No such file or directory
 #include <linux/futex.h>
```

Using this guys info/package instead: http://fabiorehm.com/blog/2015/07/22/building-a-minimum-viable-phantomjs-2-docker-image/
