# jq

jq tool in a container.

Eg:

```sh
docker run --rm -p 8080 --name sinatra treeder/hello-sinatra
```

To pull out the assigned port:

```sh
docker inspect --format="{{json .}}" sinatra | docker run --rm -i treeder/jq '.NetworkSettings.Ports["8080/tcp"][0].HostPort'
```

Equivalent with jq installed:

```sh
docker inspect --format="{{json .}}" sinatra | jq '.NetworkSettings.Ports["8080/tcp"][0].HostPort'
```
