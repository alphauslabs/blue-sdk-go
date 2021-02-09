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
