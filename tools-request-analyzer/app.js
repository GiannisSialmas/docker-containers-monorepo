var express = require('express');
var os = require('os');

var app = express();
var ifaces = os.networkInterfaces();

const serverInterfaces = {};
Object.keys(ifaces).forEach(function (ifname) {
    var alias = 0;
    ifaces[ifname].forEach(function (iface) {
        if ('IPv4' !== iface.family || iface.internal !== false) {
            // skip over internal (i.e. 127.0.0.1) and non-ipv4 addresses
            return;
        }

        if (alias >= 1) {
            // this single interface has multiple ipv4 addresses
            serverInterfaces[ifname] = iface.address;
        } else {
            // this interface has only one ipv4 adress
            serverInterfaces[ifname] =  iface.address;
        }
        ++alias;
    });
});



var server = app.listen(80, '0.0.0.0');

app.get('/*', function (request, response) {

    const obj = {
        hostname: request.hostname,
        originalUrl: request.originalUrl,
        forwardedFor: request.header('x-forwarded-for'),
        clientIp: request.connection.remoteAddress,
        path: request.path,
        protocol: request.protocol,
        query: request.query,
        subdomains: request.subdomains,
        serverInterfaces
    }

    response.send(obj);
})










// var http = require('http');

// //Lets define a port we want to listen to
// const PORT = 80;

// function timestamp() {

//     var date = new Date();

//     var hour = date.getHours();
//     hour = (hour < 10 ? "0" : "") + hour;

//     var min = date.getMinutes();
//     min = (min < 10 ? "0" : "") + min;

//     var sec = date.getSeconds();
//     sec = (sec < 10 ? "0" : "") + sec;

//     return hour + ":" + min + ":" + sec;

// }



// var server = http.createServer((request, response) => {

//     const obj = {
//         app: request.app,
//         baseUrl: request.baseUrl,
//         hostname: request.hostname,
//         ip: request.ip,
//         originalUrl: request.originalUrl,
//         path: request.path,
//         protocol: request.protocol,
//         query: request.query,
//         route: request.route,
//         subdomains: request.subdomains
//     }
//     console.log(Object.keys(request));
//     console.log(timestamp() + ': ' + JSON.stringify(obj));
//     response.end('Path Hit: ' + JSON.stringify(obj));

// });

// //Lets start our server
// server.listen(PORT, function () {
//     //Callback triggered when server is successfully listening. Hurray!
//     console.log("Server listening on: http://localhost:%s", PORT);
// });