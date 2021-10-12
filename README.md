![main](https://github.com/alphauslabs/blue-sdk-go/workflows/main/badge.svg)

The main branch can be broken. Make sure to use tagged releases.

For authentication, check out this [document](https://alphauslabs.github.io/blueapi/authentication/apikey.html).

You can check out the documentation [here](https://pkg.go.dev/github.com/alphauslabs/blue-sdk-go).

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
