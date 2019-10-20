# Description
requestAnalyzer is just a webserver that returns some info about the request. Pretty usefull for debugging networking and dns.


```
{
hostname: "second.first.battlesable.local",
originalUrl: "/hello?name=giannis",
forwardedFor: "X.X.X.X",
clientIp: "Y.Y.Y.Y"
path: "/hello",
protocol: "http",
query: {
name: "giannis"
    },
subdomains: [
        "first",
        "second"
    ],
serverInterfaces: {
eth0: "172.17.0.2"
    }
}
```