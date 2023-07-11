# SubaccountEmail

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
	"sendpost"
	"sendpost/pkg/models/operations"
	"sendpost/pkg/models/shared"
)

func main() {
    s := sendpost.New()

    ctx := context.Background()
    res, err := s.SubaccountEmail.EmailRouterSendEmail(ctx, operations.EmailRouterSendEmailRequest{
        RequestBody: []byte("quibusdam"),
        XSendPostMockEmail: sendpost.Bool(false),
        XSendPostMockTimeShift: sendpost.String("unde"),
        XSubAccountAPIKey: "nulla",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.EmailRouterSendEmailRequest](../../models/operations/emailroutersendemailrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


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
	"sendpost"
	"sendpost/pkg/models/operations"
	"sendpost/pkg/models/shared"
)

func main() {
    s := sendpost.New()

    ctx := context.Background()
    res, err := s.SubaccountEmail.EmailRouterSendEmailWithTemplate(ctx, operations.EmailRouterSendEmailWithTemplateRequest{
        RequestBody: []byte("corrupti"),
        XSubAccountAPIKey: "illum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.EmailRouterSendEmailWithTemplateRequest](../../models/operations/emailroutersendemailwithtemplaterequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.EmailRouterSendEmailWithTemplateResponse](../../models/operations/emailroutersendemailwithtemplateresponse.md), error**

