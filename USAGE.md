<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	spostgolang "github.com/speakeasy-sdks/spost-golang/v3"
	"log"
)

func main() {
	s := spostgolang.New()

	var requestBody []byte = []byte("0x6B34FffDd5")

	var xSubAccountAPIKey string = "<value>"

	var xSendPostMockEmail *bool = spostgolang.Bool(false)

	var xSendPostMockTimeShift *string = spostgolang.String("<value>")

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
<!-- End SDK Example Usage [usage] -->