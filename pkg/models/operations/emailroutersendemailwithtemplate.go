// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EmailRouterSendEmailWithTemplateRequest struct {
	// The Email Message
	RequestBody []byte `request:"mediaType=*/*"`
	// Sub-Account API Key
	XSubAccountAPIKey string `header:"style=simple,explode=false,name=X-SubAccount-ApiKey"`
}

func (o *EmailRouterSendEmailWithTemplateRequest) GetRequestBody() []byte {
	if o == nil {
		return []byte{}
	}
	return o.RequestBody
}

func (o *EmailRouterSendEmailWithTemplateRequest) GetXSubAccountAPIKey() string {
	if o == nil {
		return ""
	}
	return o.XSubAccountAPIKey
}

type EmailRouterSendEmailWithTemplateResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *EmailRouterSendEmailWithTemplateResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *EmailRouterSendEmailWithTemplateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EmailRouterSendEmailWithTemplateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EmailRouterSendEmailWithTemplateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
