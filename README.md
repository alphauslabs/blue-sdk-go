![main](https://github.com/alphauslabs/blue-sdk-go/workflows/main/badge.svg)

**(work-in-progress)**

The main branch can be broken. Make sure to use tagged releases.

By default, this library will look for the following environment variables for [authentication](https://alphauslabs.github.io/blueapi/authentication/apikey.html).

```bash
First, it will look for:
ALPHAUS_CLIENT_ID
ALPHAUS_CLIENT_SECRET

# If those are not set, it will then look for:
ALPHAUS_RIPPLE_CLIENT_ID
ALPHAUS_RIPPLE_CLIENT_SECRET

# If those are not set, it will then look for:
ALPHAUS_WAVE_CLIENT_ID
ALPHAUS_WAVE_CLIENT_SECRET
```

If you're mainly a Ripple user, we recommend you to set the following:
```bash
ALPHAUS_CLIENT_ID
ALPHAUS_CLIENT_SECRET
```

If you're mainly a Wave/WavePro user, we recommend you to set the following:

```bash
ALPHAUS_CLIENT_ID=myclientid
ALPHAUS_CLIENT_SECRET=myclientsecret
ALPHAUS_AUTH_URL=https://login.alphaus.cloud/access_token
```

If you're using both Ripple and Wave, we recommend you to set the following:

```bash
ALPHAUS_RIPPLE_CLIENT_ID=myrippleclientid
ALPHAUS_RIPPLE_CLIENT_SECRET=myrippleclientsecret
ALPHAUS_WAVE_CLIENT_ID=mywaveclientid
ALPHAUS_WAVE_CLIENT_SECRET=mywaveclientsecret
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
		ClientId:     "my-client-id",
		ClientSecret: "my-client-secret",
	}),
))

conn, _ := grpc.DialContext(context.Background(), session.BlueEndpoint, opts...)
defer conn.Close()
client := blue.NewBlueClient(conn)
...
```
