# sendpost

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/spost-golang
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"sendpost"
	"sendpost/pkg/models/operations"
	"sendpost/pkg/models/shared"
)

func main() {
    s := sendpost.New()

    ctx := context.Background()
    res, err := s.SubaccountEmail.EmailRouterSendEmail(ctx, operations.EmailRouterSendEmailRequest{
        RequestBody: []byte("corrupti"),
        XSendPostMockEmail: sendpost.Bool(false),
        XSendPostMockTimeShift: sendpost.String("provident"),
        XSubAccountAPIKey: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [SubaccountEmail](docs/sdks/subaccountemail/README.md)

* [EmailRouterSendEmail](docs/sdks/subaccountemail/README.md#emailroutersendemail) - Send Email To Contacts
* [EmailRouterSendEmailWithTemplate](docs/sdks/subaccountemail/README.md#emailroutersendemailwithtemplate) - Send Email To Contacts With Template
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
