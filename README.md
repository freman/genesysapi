# GenesysAPI

An unofficial go SDK client library for the mypurecloud api generated from the [official](https://github.com/MyPureCloud/platform-client-sdk-go) swagger.json

## Get SDK Package

Retrieve the package from https://github.com/freman/genesysapi using go get:

```bash
go get github.com/freman/genesysapi
```

## Use the SDK

### Importing the package

```go
import (
    "github.com/freman/geneysapi/client"
)
```

### Configuring the SDK

The SDK can be configured by setting properties on a Configuration instance. Applications using this library really can only
use client credentials as I have not implemented any of the other auth schemes.

```
    uri, err := url.Parse("https://api.mypurecloud.com.au")
    if err != nil {
        panic(err)
    }

	config := client.Config{
		URL: uri,
	}

	err := config.AuthorizeClientCredentials(os.Getenv("GENESYS_CLOUD_CLIENT_ID"), os.Getenv("GENESYS_CLOUD_CLIENT_SECRET"))
	if err != nil {
		panic(err)
	}

    api := client.New(config)
```

### Transport debugging

Enabling debug will trace out all http requests

```go
api.SetDebug(true)
```

## Generating

### Prerequisites

You need to have [goswagger](https://goswagger.io/) installed but I have provided a shell script to re-generate this package from the [official](https://github.com/MyPureCloud/platform-client-sdk-go) swagger.json

```bash
./generate.sh
```

## Versioning

None yet

## Authors

* **Shannon Wynter** - *Initial work* - [Freman](https://github.com/freman)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

* Genesys  / MyPureCloud, couldn't have done this without their swagger document.
* [Stratoscale](https://github.com/Stratoscale) for their [swagger](https://github.com/Stratoscale/swagger) template from which [swagger-template/client](swagger-template/client) is derived from
* etc
