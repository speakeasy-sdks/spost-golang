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