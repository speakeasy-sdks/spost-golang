<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	spostgolang "github.com/speakeasy-sdks/spost-golang"
	"github.com/speakeasy-sdks/spost-golang/pkg/models/operations"
	"github.com/speakeasy-sdks/spost-golang/pkg/models/shared"
)

func main() {
    s := spostgolang.New()
    requestBody := []byte(":k13|`asY9")
    xSubAccountAPIKey := "Northeast"
    xSendPostMockEmail := false
    xSendPostMockTimeShift := "primary"

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