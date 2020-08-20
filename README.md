# http-echo

Echos header information back to receiver. 

Dockerhub: https://hub.docker.com/repository/docker/keaaa/http-echo

```
docker run -p 8080:8080 keaaa/http-echo
curl -X PUT -d arg=val -d arg2=val2 -H "Authorization: Bearer 123.123.123" localhost:8080
General
Request URL: "/"
Request Method: "PUT"
Remote Address: "172.17.0.1:33634"

Request Headers
Header field "Accept", Value ["*/*"]
Header field "Authorization", Value ["Bearer 123.123.123"]
Header field "Content-Length", Value ["17"]
Header field "Content-Type", Value ["application/x-www-form-urlencoded"]
Header field "User-Agent", Value ["curl/7.64.1"]

Body
arg=val&arg2=val2
```

Can also return host information by setting environment variable INCLUDE_HOST_INFORMATION
