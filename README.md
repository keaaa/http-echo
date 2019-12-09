# http-echo

Echos header information back to receiver. 

Dockerhub: https://hub.docker.com/repository/docker/keaaa/http-echo

```
docker run -p 8080:8080 keaaa/http-echo
curl localhost:8080
Method "GET", Request "/asdfas/adsfasf"
Header field "User-Agent", Value ["curl/7.64.1"]
Header field "Accept", Value ["*/*"]
```
