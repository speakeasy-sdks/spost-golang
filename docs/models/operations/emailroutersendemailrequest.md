# EmailRouterSendEmailRequest


## Fields

| Field                    | Type                     | Required                 | Description              |
| ------------------------ | ------------------------ | ------------------------ | ------------------------ |
| `RequestBody`            | *[]byte*                 | :heavy_check_mark:       | The Email Message        |
| `XSendPostMockEmail`     | **bool*                  | :heavy_minus_sign:       | Mock email header        |
| `XSendPostMockTimeShift` | **string*                | :heavy_minus_sign:       | Mock email time shift    |
| `XSubAccountAPIKey`      | *string*                 | :heavy_check_mark:       | Sub-Account API Key      |