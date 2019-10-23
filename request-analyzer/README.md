# Description
requestAnalyzer is just a webserver that returns some info about the request. Pretty usefull for debugging networking and dns.


```
{
    "hostname": "second.first.battlesable.local",
    "originalUrl": "/hello?name=battlesable&image=request-analyzer",
    "clientIp": "172.17.0.1",
    "path": "/hello",
    "protocol": "http",
    "query": {
        "name": "battlesable",
        "image": "request-analyzer"
    },
    "subdomains": [
        "first",
        "second"
    ],
    "serverInterfaces": {
        "eth0": "172.17.0.2"
    }
}
```