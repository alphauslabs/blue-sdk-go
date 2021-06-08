![main](https://github.com/alphauslabs/blue-sdk-go/workflows/main/badge.svg)

**(work-in-progress)**

The main branch can be broken. Make sure to use tagged releases.

By default, this library will look for the following environment variables for [authentication](https://alphauslabs.github.io/blueapi/authentication/apikey.html).

```bash
# First, it will look for:
ALPHAUS_CLIENT_ID
ALPHAUS_CLIENT_SECRET

# If those are not set, it will then look for:
ALPHAUS_RIPPLE_CLIENT_ID
ALPHAUS_RIPPLE_CLIENT_SECRET

# If those are not set, it will then look for:
ALPHAUS_WAVE_CLIENT_ID
ALPHAUS_WAVE_CLIENT_SECRET

# If you're mainly a Ripple user, we recommend you to set the following:
ALPHAUS_CLIENT_ID=myclientid
ALPHAUS_CLIENT_SECRET=myclientsecret

# If you're mainly a Wave/WavePro user, we recommend you to set the following:
ALPHAUS_CLIENT_ID=myclientid
ALPHAUS_CLIENT_SECRET=myclientsecret
ALPHAUS_AUTH_URL=https://login.alphaus.cloud/access_token
```

To use the default client(s), you can try something like:

```go
ctx := context.Background()
client, _ := iam.NewClient(ctx)
defer client.Close()
out, err := client.WhoAmI(ctx, &iam.WhoAmIRequest{})
log.Println(out, err)
...
```
