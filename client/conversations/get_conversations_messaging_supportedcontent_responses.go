// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetConversationsMessagingSupportedcontentReader is a Reader for the GetConversationsMessagingSupportedcontent structure.
type GetConversationsMessagingSupportedcontentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingSupportedcontentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingSupportedcontentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingSupportedcontentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingSupportedcontentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingSupportedcontentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingSupportedcontentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingSupportedcontentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingSupportedcontentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingSupportedcontentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingSupportedcontentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingSupportedcontentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingSupportedcontentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingSupportedcontentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingSupportedcontentOK creates a GetConversationsMessagingSupportedcontentOK with default headers values
func NewGetConversationsMessagingSupportedcontentOK() *GetConversationsMessagingSupportedcontentOK {
	return &GetConversationsMessagingSupportedcontentOK{}
}

/*
GetConversationsMessagingSupportedcontentOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessagingSupportedcontentOK struct {
	Payload *models.SupportedContentListing
}

// IsSuccess returns true when this get conversations messaging supportedcontent o k response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations messaging supportedcontent o k response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent o k response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging supportedcontent o k response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent o k response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessagingSupportedcontentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentOK) GetPayload() *models.SupportedContentListing {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedContentListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentBadRequest creates a GetConversationsMessagingSupportedcontentBadRequest with default headers values
func NewGetConversationsMessagingSupportedcontentBadRequest() *GetConversationsMessagingSupportedcontentBadRequest {
	return &GetConversationsMessagingSupportedcontentBadRequest{}
}

/*
GetConversationsMessagingSupportedcontentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingSupportedcontentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent bad request response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent bad request response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent bad request response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent bad request response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent bad request response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessagingSupportedcontentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentUnauthorized creates a GetConversationsMessagingSupportedcontentUnauthorized with default headers values
func NewGetConversationsMessagingSupportedcontentUnauthorized() *GetConversationsMessagingSupportedcontentUnauthorized {
	return &GetConversationsMessagingSupportedcontentUnauthorized{}
}

/*
GetConversationsMessagingSupportedcontentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingSupportedcontentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent unauthorized response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent unauthorized response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent unauthorized response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent unauthorized response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent unauthorized response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessagingSupportedcontentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentForbidden creates a GetConversationsMessagingSupportedcontentForbidden with default headers values
func NewGetConversationsMessagingSupportedcontentForbidden() *GetConversationsMessagingSupportedcontentForbidden {
	return &GetConversationsMessagingSupportedcontentForbidden{}
}

/*
GetConversationsMessagingSupportedcontentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingSupportedcontentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent forbidden response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent forbidden response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent forbidden response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent forbidden response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent forbidden response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessagingSupportedcontentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentNotFound creates a GetConversationsMessagingSupportedcontentNotFound with default headers values
func NewGetConversationsMessagingSupportedcontentNotFound() *GetConversationsMessagingSupportedcontentNotFound {
	return &GetConversationsMessagingSupportedcontentNotFound{}
}

/*
GetConversationsMessagingSupportedcontentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingSupportedcontentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent not found response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent not found response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent not found response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent not found response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent not found response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessagingSupportedcontentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentRequestTimeout creates a GetConversationsMessagingSupportedcontentRequestTimeout with default headers values
func NewGetConversationsMessagingSupportedcontentRequestTimeout() *GetConversationsMessagingSupportedcontentRequestTimeout {
	return &GetConversationsMessagingSupportedcontentRequestTimeout{}
}

/*
GetConversationsMessagingSupportedcontentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingSupportedcontentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent request timeout response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent request timeout response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent request timeout response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent request timeout response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent request timeout response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessagingSupportedcontentRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentRequestEntityTooLarge creates a GetConversationsMessagingSupportedcontentRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingSupportedcontentRequestEntityTooLarge() *GetConversationsMessagingSupportedcontentRequestEntityTooLarge {
	return &GetConversationsMessagingSupportedcontentRequestEntityTooLarge{}
}

/*
GetConversationsMessagingSupportedcontentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsMessagingSupportedcontentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent request entity too large response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent request entity too large response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent request entity too large response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent request entity too large response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent request entity too large response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentUnsupportedMediaType creates a GetConversationsMessagingSupportedcontentUnsupportedMediaType with default headers values
func NewGetConversationsMessagingSupportedcontentUnsupportedMediaType() *GetConversationsMessagingSupportedcontentUnsupportedMediaType {
	return &GetConversationsMessagingSupportedcontentUnsupportedMediaType{}
}

/*
GetConversationsMessagingSupportedcontentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingSupportedcontentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent unsupported media type response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent unsupported media type response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent unsupported media type response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent unsupported media type response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent unsupported media type response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentTooManyRequests creates a GetConversationsMessagingSupportedcontentTooManyRequests with default headers values
func NewGetConversationsMessagingSupportedcontentTooManyRequests() *GetConversationsMessagingSupportedcontentTooManyRequests {
	return &GetConversationsMessagingSupportedcontentTooManyRequests{}
}

/*
GetConversationsMessagingSupportedcontentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingSupportedcontentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent too many requests response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent too many requests response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent too many requests response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations messaging supportedcontent too many requests response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations messaging supportedcontent too many requests response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessagingSupportedcontentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentInternalServerError creates a GetConversationsMessagingSupportedcontentInternalServerError with default headers values
func NewGetConversationsMessagingSupportedcontentInternalServerError() *GetConversationsMessagingSupportedcontentInternalServerError {
	return &GetConversationsMessagingSupportedcontentInternalServerError{}
}

/*
GetConversationsMessagingSupportedcontentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingSupportedcontentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent internal server error response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent internal server error response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent internal server error response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging supportedcontent internal server error response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging supportedcontent internal server error response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessagingSupportedcontentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentServiceUnavailable creates a GetConversationsMessagingSupportedcontentServiceUnavailable with default headers values
func NewGetConversationsMessagingSupportedcontentServiceUnavailable() *GetConversationsMessagingSupportedcontentServiceUnavailable {
	return &GetConversationsMessagingSupportedcontentServiceUnavailable{}
}

/*
GetConversationsMessagingSupportedcontentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingSupportedcontentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent service unavailable response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent service unavailable response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent service unavailable response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging supportedcontent service unavailable response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging supportedcontent service unavailable response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingSupportedcontentGatewayTimeout creates a GetConversationsMessagingSupportedcontentGatewayTimeout with default headers values
func NewGetConversationsMessagingSupportedcontentGatewayTimeout() *GetConversationsMessagingSupportedcontentGatewayTimeout {
	return &GetConversationsMessagingSupportedcontentGatewayTimeout{}
}

/*
GetConversationsMessagingSupportedcontentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessagingSupportedcontentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations messaging supportedcontent gateway timeout response has a 2xx status code
func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations messaging supportedcontent gateway timeout response has a 3xx status code
func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations messaging supportedcontent gateway timeout response has a 4xx status code
func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations messaging supportedcontent gateway timeout response has a 5xx status code
func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations messaging supportedcontent gateway timeout response a status code equal to that given
func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/supportedcontent][%d] getConversationsMessagingSupportedcontentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingSupportedcontentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
