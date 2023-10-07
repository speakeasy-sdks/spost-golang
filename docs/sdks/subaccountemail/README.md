# SubaccountEmail
(*SubaccountEmail*)

### Available Operations

* [EmailRouterSendEmail](#emailroutersendemail) - Send Email To Contacts
* [EmailRouterSendEmailWithTemplate](#emailroutersendemailwithtemplate) - Send Email To Contacts With Template

## EmailRouterSendEmail

Send Email To Contacts

### Example Usage

```go
package main

import(
	"context"
	"log"
	spostgolang "github.com/speakeasy-sdks/spost-golang"
	"github.com/speakeasy-sdks/spost-golang/pkg/models/shared"
)

func main() {
    s := spostgolang.New()
    var requestBody []byte = []byte(":k13|`asY9")
    var xSubAccountAPIKey string = "Northeast"
    var xSendPostMockEmail *bool = false
    var xSendPostMockTimeShift *string = "primary"

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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `requestBody`                                         | *[]byte*                                              | :heavy_check_mark:                                    | The Email Message                                     |
| `xSubAccountAPIKey`                                   | *string*                                              | :heavy_check_mark:                                    | Sub-Account API Key                                   |
| `xSendPostMockEmail`                                  | **bool*                                               | :heavy_minus_sign:                                    | Mock email header                                     |
| `xSendPostMockTimeShift`                              | **string*                                             | :heavy_minus_sign:                                    | Mock email time shift                                 |


### Response

**[*operations.EmailRouterSendEmailResponse](../../models/operations/emailroutersendemailresponse.md), error**


## EmailRouterSendEmailWithTemplate

Send Email To Contacts With Template

### Example Usage

```go
package main

import(
	"context"
	"log"
	spostgolang "github.com/speakeasy-sdks/spost-golang"
	"github.com/speakeasy-sdks/spost-golang/pkg/models/shared"
)

func main() {
    s := spostgolang.New()
    var requestBody []byte = []byte("9hY_GIO^\M")
    var xSubAccountAPIKey string = "Kina"

    ctx := context.Background()
    res, err := s.SubaccountEmail.EmailRouterSendEmailWithTemplate(ctx, requestBody, xSubAccountAPIKey)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `requestBody`                                         | *[]byte*                                              | :heavy_check_mark:                                    | The Email Message                                     |
| `xSubAccountAPIKey`                                   | *string*                                              | :heavy_check_mark:                                    | Sub-Account API Key                                   |


### Response

**[*operations.EmailRouterSendEmailWithTemplateResponse](../../models/operations/emailroutersendemailwithtemplateresponse.md), error**

