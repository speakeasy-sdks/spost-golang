<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	spostgolang "github.com/speakeasy-sdks/spost-golang"
	"github.com/speakeasy-sdks/spost-golang/pkg/models/shared"
	"log"
)

func main() {
	s := spostgolang.New()

	var requestBody []byte = []byte(":k13|`asY9")

	var xSubAccountAPIKey string = "string"

	var xSendPostMockEmail *bool = false

	var xSendPostMockTimeShift *string = "string"

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