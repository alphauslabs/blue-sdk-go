![main](https://github.com/alphauslabs/blue-sdk-go/workflows/main/badge.svg)

**(work-in-progress)**

The main branch can be broken. Make sure to use tagged releases.

For authentication, check out this [document](https://alphauslabs.github.io/blueapi/authentication/apikey.html).

To use the default client(s), you can try something like:

```go
ctx := context.Background()
client, _ := iam.NewClient(ctx)
defer client.Close()
out, err := client.WhoAmI(ctx, &iam.WhoAmIRequest{})
log.Println(out, err)
...
```
