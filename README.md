![main](https://github.com/alphauslabs/blue-sdk-go/workflows/main/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/alphauslabs/blue-sdk-go.svg)](https://pkg.go.dev/github.com/alphauslabs/blue-sdk-go)

> [!IMPORTANT]
> The main branch can be broken. Make sure to use tagged releases.

This repo contains Blue's (mostly auto-generated) client library for the Go programming language.

For authentication, check out this [document](https://alphauslabs.github.io/docs/blueapi/authentication/).

To use the default client(s), you can try something like:

```go
ctx := context.Background()
client, _ := iam.NewClient(ctx)
defer client.Close()
out, err := client.WhoAmI(ctx, &iam.WhoAmIRequest{})
log.Println(out, err)
...
```

You can also check the examples [here](./examples/).

The generated [gRPC](https://grpc.io/) documentation is [here](https://labs.alphaus.cloud/blue-sdk-go/).
