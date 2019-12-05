# http-echo

Echos header information back to receiver

```
docker run -p 8080:8080 keaaa/http-echo
curl localhost:8080
Header field "User-Agent", Value ["curl/7.64.1"]
Header field "Accept", Value ["*/*"]
```