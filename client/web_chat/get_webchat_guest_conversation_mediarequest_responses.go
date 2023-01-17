// Code generated by go-swagger; DO NOT EDIT.

package web_chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWebchatGuestConversationMediarequestReader is a Reader for the GetWebchatGuestConversationMediarequest structure.
type GetWebchatGuestConversationMediarequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebchatGuestConversationMediarequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebchatGuestConversationMediarequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebchatGuestConversationMediarequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWebchatGuestConversationMediarequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebchatGuestConversationMediarequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebchatGuestConversationMediarequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWebchatGuestConversationMediarequestRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWebchatGuestConversationMediarequestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWebchatGuestConversationMediarequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWebchatGuestConversationMediarequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebchatGuestConversationMediarequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWebchatGuestConversationMediarequestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWebchatGuestConversationMediarequestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebchatGuestConversationMediarequestOK creates a GetWebchatGuestConversationMediarequestOK with default headers values
func NewGetWebchatGuestConversationMediarequestOK() *GetWebchatGuestConversationMediarequestOK {
	return &GetWebchatGuestConversationMediarequestOK{}
}

/*
GetWebchatGuestConversationMediarequestOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebchatGuestConversationMediarequestOK struct {
	Payload *models.WebChatGuestMediaRequest
}

// IsSuccess returns true when this get webchat guest conversation mediarequest o k response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get webchat guest conversation mediarequest o k response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest o k response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation mediarequest o k response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest o k response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWebchatGuestConversationMediarequestOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestOK  %+v", 200, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestOK) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestOK  %+v", 200, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestOK) GetPayload() *models.WebChatGuestMediaRequest {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatGuestMediaRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestBadRequest creates a GetWebchatGuestConversationMediarequestBadRequest with default headers values
func NewGetWebchatGuestConversationMediarequestBadRequest() *GetWebchatGuestConversationMediarequestBadRequest {
	return &GetWebchatGuestConversationMediarequestBadRequest{}
}

/*
GetWebchatGuestConversationMediarequestBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWebchatGuestConversationMediarequestBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest bad request response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest bad request response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest bad request response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest bad request response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest bad request response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWebchatGuestConversationMediarequestBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestUnauthorized creates a GetWebchatGuestConversationMediarequestUnauthorized with default headers values
func NewGetWebchatGuestConversationMediarequestUnauthorized() *GetWebchatGuestConversationMediarequestUnauthorized {
	return &GetWebchatGuestConversationMediarequestUnauthorized{}
}

/*
GetWebchatGuestConversationMediarequestUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWebchatGuestConversationMediarequestUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest unauthorized response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest unauthorized response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest unauthorized response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest unauthorized response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest unauthorized response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWebchatGuestConversationMediarequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestForbidden creates a GetWebchatGuestConversationMediarequestForbidden with default headers values
func NewGetWebchatGuestConversationMediarequestForbidden() *GetWebchatGuestConversationMediarequestForbidden {
	return &GetWebchatGuestConversationMediarequestForbidden{}
}

/*
GetWebchatGuestConversationMediarequestForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWebchatGuestConversationMediarequestForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest forbidden response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest forbidden response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest forbidden response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest forbidden response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest forbidden response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWebchatGuestConversationMediarequestForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestForbidden  %+v", 403, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestForbidden  %+v", 403, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestNotFound creates a GetWebchatGuestConversationMediarequestNotFound with default headers values
func NewGetWebchatGuestConversationMediarequestNotFound() *GetWebchatGuestConversationMediarequestNotFound {
	return &GetWebchatGuestConversationMediarequestNotFound{}
}

/*
GetWebchatGuestConversationMediarequestNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWebchatGuestConversationMediarequestNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest not found response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest not found response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest not found response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest not found response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest not found response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWebchatGuestConversationMediarequestNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestNotFound  %+v", 404, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestNotFound  %+v", 404, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestRequestTimeout creates a GetWebchatGuestConversationMediarequestRequestTimeout with default headers values
func NewGetWebchatGuestConversationMediarequestRequestTimeout() *GetWebchatGuestConversationMediarequestRequestTimeout {
	return &GetWebchatGuestConversationMediarequestRequestTimeout{}
}

/*
GetWebchatGuestConversationMediarequestRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWebchatGuestConversationMediarequestRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest request timeout response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest request timeout response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest request timeout response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest request timeout response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest request timeout response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWebchatGuestConversationMediarequestRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestRequestEntityTooLarge creates a GetWebchatGuestConversationMediarequestRequestEntityTooLarge with default headers values
func NewGetWebchatGuestConversationMediarequestRequestEntityTooLarge() *GetWebchatGuestConversationMediarequestRequestEntityTooLarge {
	return &GetWebchatGuestConversationMediarequestRequestEntityTooLarge{}
}

/*
GetWebchatGuestConversationMediarequestRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWebchatGuestConversationMediarequestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest request entity too large response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest request entity too large response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest request entity too large response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest request entity too large response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest request entity too large response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestUnsupportedMediaType creates a GetWebchatGuestConversationMediarequestUnsupportedMediaType with default headers values
func NewGetWebchatGuestConversationMediarequestUnsupportedMediaType() *GetWebchatGuestConversationMediarequestUnsupportedMediaType {
	return &GetWebchatGuestConversationMediarequestUnsupportedMediaType{}
}

/*
GetWebchatGuestConversationMediarequestUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWebchatGuestConversationMediarequestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest unsupported media type response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest unsupported media type response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest unsupported media type response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest unsupported media type response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest unsupported media type response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestTooManyRequests creates a GetWebchatGuestConversationMediarequestTooManyRequests with default headers values
func NewGetWebchatGuestConversationMediarequestTooManyRequests() *GetWebchatGuestConversationMediarequestTooManyRequests {
	return &GetWebchatGuestConversationMediarequestTooManyRequests{}
}

/*
GetWebchatGuestConversationMediarequestTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWebchatGuestConversationMediarequestTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest too many requests response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest too many requests response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest too many requests response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webchat guest conversation mediarequest too many requests response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get webchat guest conversation mediarequest too many requests response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWebchatGuestConversationMediarequestTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestInternalServerError creates a GetWebchatGuestConversationMediarequestInternalServerError with default headers values
func NewGetWebchatGuestConversationMediarequestInternalServerError() *GetWebchatGuestConversationMediarequestInternalServerError {
	return &GetWebchatGuestConversationMediarequestInternalServerError{}
}

/*
GetWebchatGuestConversationMediarequestInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWebchatGuestConversationMediarequestInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest internal server error response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest internal server error response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest internal server error response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation mediarequest internal server error response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get webchat guest conversation mediarequest internal server error response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWebchatGuestConversationMediarequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestServiceUnavailable creates a GetWebchatGuestConversationMediarequestServiceUnavailable with default headers values
func NewGetWebchatGuestConversationMediarequestServiceUnavailable() *GetWebchatGuestConversationMediarequestServiceUnavailable {
	return &GetWebchatGuestConversationMediarequestServiceUnavailable{}
}

/*
GetWebchatGuestConversationMediarequestServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWebchatGuestConversationMediarequestServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest service unavailable response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest service unavailable response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest service unavailable response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation mediarequest service unavailable response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get webchat guest conversation mediarequest service unavailable response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebchatGuestConversationMediarequestGatewayTimeout creates a GetWebchatGuestConversationMediarequestGatewayTimeout with default headers values
func NewGetWebchatGuestConversationMediarequestGatewayTimeout() *GetWebchatGuestConversationMediarequestGatewayTimeout {
	return &GetWebchatGuestConversationMediarequestGatewayTimeout{}
}

/*
GetWebchatGuestConversationMediarequestGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWebchatGuestConversationMediarequestGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get webchat guest conversation mediarequest gateway timeout response has a 2xx status code
func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webchat guest conversation mediarequest gateway timeout response has a 3xx status code
func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webchat guest conversation mediarequest gateway timeout response has a 4xx status code
func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webchat guest conversation mediarequest gateway timeout response has a 5xx status code
func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get webchat guest conversation mediarequest gateway timeout response a status code equal to that given
func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/webchat/guest/conversations/{conversationId}/mediarequests/{mediaRequestId}][%d] getWebchatGuestConversationMediarequestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWebchatGuestConversationMediarequestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
