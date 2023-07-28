<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/spost-golang"
	"github.com/speakeasy-sdks/spost-golang/pkg/models/operations"
	"github.com/speakeasy-sdks/spost-golang/pkg/models/shared"
)

func main() {
    s := sendpost.New()
    requestBody := []byte("corrupti")
    xSubAccountAPIKey := "provident"
    xSendPostMockEmail := false
    xSendPostMockTimeShift := "distinctio"

    ctx := context.Background()
    res, err := s.SubaccountEmail.EmailRouterSendEmail(ctx, requestBody, xSubAccountAPIKey, xSendPostMockEmail, xSendPostMockTimeShift)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->