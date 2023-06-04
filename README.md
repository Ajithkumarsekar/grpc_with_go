# grpc_with_go

[//]: # ()
[//]: # (# FAQ&#40;by myself 😛&#41;)

[//]: # ()
[//]: # (- [ What is gRPC?]&#40;#What-is-gRPC?&#41;)

### What is gRPC?
Developed by google (2015). It uses `HTTP/2` for transport, `Protocol Buffers` as the default `interface description language`
### what is interface description language?
IDL is a generic term for a language that lets a program or object written in one language communicate with another program written in an unknown language. IDLs describe an interface in a language-independent way, enabling communication between software components that do not share one language, for example, between those written in C++ and those written in Java.
Popular IDL used with gRPC `Apache Thrift` (by Facebook), `Apache Avro`, `Flatbuffers` (by Google)
### Why do you mean by `HTTP/2` for transport? gRPC doesn't have any components of HTTP like url, method, headers so HTTP?
That's not entirely true. gRPC use HTTP components like URL path, request/response headers (content-type, timeout, custom-metadata) behind the scenes to make the rpc call!
Most importantly it uses HTTP 2 streams to make multiple RPC requests within a single TCP connection with support for stream cancellation etc. refer [PROTOCOL-HTTP2](https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md).
### What is channel and subchannels?

### What are all the states of a channel?

### What happens to the channel when the server goes down? Does it reresolve the DNS to obtain new address and retry the rpc?

### Does gRPC server usually sent GO AWAY when it shuts down?

### Unary RPC flow?
![unary_flow](/screenshots/unary_flow.png)
### can gRPC method be called without generated stub (client code)?
yes, using gRPC reflections. gRPC Server Reflection provides information about publicly-accessible gRPC services on a server, and assists clients at runtime to construct RPC requests and responses without precompiled service information.
### What is evans?
[evans](https://github.com/ktr0731/evans) is a generic gRPC command line client which uses either gRPC server reflection or proto files to construct, decode RPC request, response
![evans_client](/screenshots/evans_client.png)

# Reference
- [gRPC getting started with Go](https://grpc.io/docs/languages/go/quickstart)
- [gRPC core concepts](https://grpc.io/docs/what-is-grpc/core-concepts/)
- [gRPC metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md)
- [gRPC server reflection protocol](https://github.com/grpc/grpc/blob/master/doc/server-reflection.md)
- [Deadlines](https://grpc.io/blog/deadlines/)