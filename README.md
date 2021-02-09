![main](https://github.com/alphauslabs/blue-sdk-go/workflows/main/badge.svg)

The main branch can be broken. Make sure to use tagged releases.

By default, this library will look for the following environment variables for [authentication](https://alphauslabs.github.io/blueapi/authentication/apikey.html):

```bash
ALPHAUS_CLIENT_ID
ALPHAUS_CLIENT_SECRET
```

To use the default client, you can try something like:

```go
ctx := context.Background()
client, _ := awscost.NewClient(ctx)
defer client.Close()
...
```

You can also provide your own gRPC connection:

```go
var opts []grpc.DialOption
creds := credentials.NewTLS(&tls.Config{})
opts = append(opts, grpc.WithTransportCredentials(creds))
opts = append(opts, grpc.WithBlock())
opts = append(opts, grpc.WithPerRPCCredentials(
	session.NewRpcCredentials(session.RpcCredentialsInput{
		LoginUrl:     session.LoginUrlRipple,
		ClientId:     params.ClientId,
		ClientSecret: params.ClientSecret,
	}),
))

conn, _ := grpc.DialContext(context.Background(), session.BlueAwsEndpoint, opts...)
defer conn.Close()
client := blueaws.NewBlueAwsClient(conn)
...
```
