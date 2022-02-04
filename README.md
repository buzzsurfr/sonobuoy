# sonobuoy <img src="https://github.com/buzzsurfr/sonobuoy/raw/main/docs/img/dot-circle-regular.svg" alt="fa-dot-circle-regular" style="width: 40px; height: 40px;">

Simple, multi-protocol echo proxy

A sonobuoy (a portmanteau of sonar and buoy) is a relatively small buoy - typically 13 cm (5 in) diameter and 91 cm (3 ft) long - expendable sonar system that is dropped/ejected from aircraft or ships conducting anti-submarine warfare or underwater acoustic research.

These sonobuoys can be deployed as a means of passing health checks and are a great example of a passthrough health check (to ensure an upstream is healthy).

## Usage

Use sonobuoy in order to pass health checks of different protocol types, including TCP, HTTP, and gRPC.

```
Usage:
  sonobuoy [flags]

Flags:
  -p, --port string         Port (default ":2869")
  -P, --protocol protocol   Protocol: tcp | http | grpc (default tcp)
```

### TCP (default)

The TCP sonobuoy will gracefully close the TCP connection whenever one is successfully established and a payload sent.

Test using [netcat](http://netcat.sourceforge.net/):
```
$ nc -v localhost 2869
Connection to localhost port 2869 [tcp/icslap] succeeded!
```

### HTTP

The HTTP sonobuoy sends an empty HTTP response with the 200 status.

Test using [curl](https://curl.se/):
```
$ curl -v localhost:2869
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 2869 (#0)
> GET / HTTP/1.1
> Host: localhost:2869
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Fri, 04 Feb 2022 13:51:59 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
* Closing connection 0
```

### gRPC

The gRPC sonobuoy uses a [signal protobuf](https://github.com/buzzsurfr/sonobuoy/blob/main/proto/signal.proto) that accepts and returns an empty set.

Test using [gRPCurl](https://github.com/fullstorydev/grpcurl):
```
$ grpcurl -plaintext -proto proto/signal.proto localhost:2869 signal.v1alpha.Echo/Signal
{

}
```