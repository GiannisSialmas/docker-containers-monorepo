# Description
memoryLeak is a nodejs webserver that keeps leaking memory with each request. It does that by pushing numbers to a global array in each request. The array is not picked up by the garbage collector as it is defined in the global scope.
This proved useful when studing about kubernetes resource limits.