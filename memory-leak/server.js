var http = require("http");

const fillmem = [];

http.createServer(function (request, response) {
   // Send the HTTP header 
   // HTTP Status: 200 : OK
   // Content Type: text/plain
   response.writeHead(200, { 'Content-Type': 'text/plain' });

   for (let i = 0; i < 10000; i++) {
      fillmem.push(i);
   }

   // Send the response body as "Hello World"
   response.end('This is version 4 of the application\n');
}).listen(80);

// Console will print the message
console.log('Server running at http://127.0.0.1:80/');
