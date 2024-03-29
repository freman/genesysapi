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

// GetConversationsMessageDetailsReader is a Reader for the GetConversationsMessageDetails structure.
type GetConversationsMessageDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessageDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessageDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessageDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessageDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessageDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessageDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessageDetailsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessageDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessageDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessageDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessageDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessageDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessageDetailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessageDetailsOK creates a GetConversationsMessageDetailsOK with default headers values
func NewGetConversationsMessageDetailsOK() *GetConversationsMessageDetailsOK {
	return &GetConversationsMessageDetailsOK{}
}

/*
GetConversationsMessageDetailsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsMessageDetailsOK struct {
	Payload *models.MessageData
}

// IsSuccess returns true when this get conversations message details o k response has a 2xx status code
func (o *GetConversationsMessageDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations message details o k response has a 3xx status code
func (o *GetConversationsMessageDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details o k response has a 4xx status code
func (o *GetConversationsMessageDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message details o k response has a 5xx status code
func (o *GetConversationsMessageDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details o k response a status code equal to that given
func (o *GetConversationsMessageDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsMessageDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessageDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessageDetailsOK) GetPayload() *models.MessageData {
	return o.Payload
}

func (o *GetConversationsMessageDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsBadRequest creates a GetConversationsMessageDetailsBadRequest with default headers values
func NewGetConversationsMessageDetailsBadRequest() *GetConversationsMessageDetailsBadRequest {
	return &GetConversationsMessageDetailsBadRequest{}
}

/*
GetConversationsMessageDetailsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessageDetailsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details bad request response has a 2xx status code
func (o *GetConversationsMessageDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details bad request response has a 3xx status code
func (o *GetConversationsMessageDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details bad request response has a 4xx status code
func (o *GetConversationsMessageDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details bad request response has a 5xx status code
func (o *GetConversationsMessageDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details bad request response a status code equal to that given
func (o *GetConversationsMessageDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsMessageDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessageDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessageDetailsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsUnauthorized creates a GetConversationsMessageDetailsUnauthorized with default headers values
func NewGetConversationsMessageDetailsUnauthorized() *GetConversationsMessageDetailsUnauthorized {
	return &GetConversationsMessageDetailsUnauthorized{}
}

/*
GetConversationsMessageDetailsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessageDetailsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details unauthorized response has a 2xx status code
func (o *GetConversationsMessageDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details unauthorized response has a 3xx status code
func (o *GetConversationsMessageDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details unauthorized response has a 4xx status code
func (o *GetConversationsMessageDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details unauthorized response has a 5xx status code
func (o *GetConversationsMessageDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details unauthorized response a status code equal to that given
func (o *GetConversationsMessageDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsMessageDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessageDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessageDetailsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsForbidden creates a GetConversationsMessageDetailsForbidden with default headers values
func NewGetConversationsMessageDetailsForbidden() *GetConversationsMessageDetailsForbidden {
	return &GetConversationsMessageDetailsForbidden{}
}

/*
GetConversationsMessageDetailsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessageDetailsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details forbidden response has a 2xx status code
func (o *GetConversationsMessageDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details forbidden response has a 3xx status code
func (o *GetConversationsMessageDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details forbidden response has a 4xx status code
func (o *GetConversationsMessageDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details forbidden response has a 5xx status code
func (o *GetConversationsMessageDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details forbidden response a status code equal to that given
func (o *GetConversationsMessageDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsMessageDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessageDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessageDetailsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsNotFound creates a GetConversationsMessageDetailsNotFound with default headers values
func NewGetConversationsMessageDetailsNotFound() *GetConversationsMessageDetailsNotFound {
	return &GetConversationsMessageDetailsNotFound{}
}

/*
GetConversationsMessageDetailsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsMessageDetailsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details not found response has a 2xx status code
func (o *GetConversationsMessageDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details not found response has a 3xx status code
func (o *GetConversationsMessageDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details not found response has a 4xx status code
func (o *GetConversationsMessageDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details not found response has a 5xx status code
func (o *GetConversationsMessageDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details not found response a status code equal to that given
func (o *GetConversationsMessageDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsMessageDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessageDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessageDetailsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsRequestTimeout creates a GetConversationsMessageDetailsRequestTimeout with default headers values
func NewGetConversationsMessageDetailsRequestTimeout() *GetConversationsMessageDetailsRequestTimeout {
	return &GetConversationsMessageDetailsRequestTimeout{}
}

/*
GetConversationsMessageDetailsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessageDetailsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details request timeout response has a 2xx status code
func (o *GetConversationsMessageDetailsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details request timeout response has a 3xx status code
func (o *GetConversationsMessageDetailsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details request timeout response has a 4xx status code
func (o *GetConversationsMessageDetailsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details request timeout response has a 5xx status code
func (o *GetConversationsMessageDetailsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details request timeout response a status code equal to that given
func (o *GetConversationsMessageDetailsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsMessageDetailsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessageDetailsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessageDetailsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsRequestEntityTooLarge creates a GetConversationsMessageDetailsRequestEntityTooLarge with default headers values
func NewGetConversationsMessageDetailsRequestEntityTooLarge() *GetConversationsMessageDetailsRequestEntityTooLarge {
	return &GetConversationsMessageDetailsRequestEntityTooLarge{}
}

/*
GetConversationsMessageDetailsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsMessageDetailsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details request entity too large response has a 2xx status code
func (o *GetConversationsMessageDetailsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details request entity too large response has a 3xx status code
func (o *GetConversationsMessageDetailsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details request entity too large response has a 4xx status code
func (o *GetConversationsMessageDetailsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details request entity too large response has a 5xx status code
func (o *GetConversationsMessageDetailsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details request entity too large response a status code equal to that given
func (o *GetConversationsMessageDetailsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsMessageDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessageDetailsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessageDetailsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsUnsupportedMediaType creates a GetConversationsMessageDetailsUnsupportedMediaType with default headers values
func NewGetConversationsMessageDetailsUnsupportedMediaType() *GetConversationsMessageDetailsUnsupportedMediaType {
	return &GetConversationsMessageDetailsUnsupportedMediaType{}
}

/*
GetConversationsMessageDetailsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessageDetailsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details unsupported media type response has a 2xx status code
func (o *GetConversationsMessageDetailsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details unsupported media type response has a 3xx status code
func (o *GetConversationsMessageDetailsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details unsupported media type response has a 4xx status code
func (o *GetConversationsMessageDetailsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details unsupported media type response has a 5xx status code
func (o *GetConversationsMessageDetailsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details unsupported media type response a status code equal to that given
func (o *GetConversationsMessageDetailsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsMessageDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessageDetailsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessageDetailsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsTooManyRequests creates a GetConversationsMessageDetailsTooManyRequests with default headers values
func NewGetConversationsMessageDetailsTooManyRequests() *GetConversationsMessageDetailsTooManyRequests {
	return &GetConversationsMessageDetailsTooManyRequests{}
}

/*
GetConversationsMessageDetailsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessageDetailsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details too many requests response has a 2xx status code
func (o *GetConversationsMessageDetailsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details too many requests response has a 3xx status code
func (o *GetConversationsMessageDetailsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details too many requests response has a 4xx status code
func (o *GetConversationsMessageDetailsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations message details too many requests response has a 5xx status code
func (o *GetConversationsMessageDetailsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations message details too many requests response a status code equal to that given
func (o *GetConversationsMessageDetailsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsMessageDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessageDetailsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessageDetailsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsInternalServerError creates a GetConversationsMessageDetailsInternalServerError with default headers values
func NewGetConversationsMessageDetailsInternalServerError() *GetConversationsMessageDetailsInternalServerError {
	return &GetConversationsMessageDetailsInternalServerError{}
}

/*
GetConversationsMessageDetailsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessageDetailsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details internal server error response has a 2xx status code
func (o *GetConversationsMessageDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details internal server error response has a 3xx status code
func (o *GetConversationsMessageDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details internal server error response has a 4xx status code
func (o *GetConversationsMessageDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message details internal server error response has a 5xx status code
func (o *GetConversationsMessageDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations message details internal server error response a status code equal to that given
func (o *GetConversationsMessageDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsMessageDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessageDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessageDetailsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsServiceUnavailable creates a GetConversationsMessageDetailsServiceUnavailable with default headers values
func NewGetConversationsMessageDetailsServiceUnavailable() *GetConversationsMessageDetailsServiceUnavailable {
	return &GetConversationsMessageDetailsServiceUnavailable{}
}

/*
GetConversationsMessageDetailsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessageDetailsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details service unavailable response has a 2xx status code
func (o *GetConversationsMessageDetailsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details service unavailable response has a 3xx status code
func (o *GetConversationsMessageDetailsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details service unavailable response has a 4xx status code
func (o *GetConversationsMessageDetailsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message details service unavailable response has a 5xx status code
func (o *GetConversationsMessageDetailsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations message details service unavailable response a status code equal to that given
func (o *GetConversationsMessageDetailsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsMessageDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessageDetailsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessageDetailsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageDetailsGatewayTimeout creates a GetConversationsMessageDetailsGatewayTimeout with default headers values
func NewGetConversationsMessageDetailsGatewayTimeout() *GetConversationsMessageDetailsGatewayTimeout {
	return &GetConversationsMessageDetailsGatewayTimeout{}
}

/*
GetConversationsMessageDetailsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsMessageDetailsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations message details gateway timeout response has a 2xx status code
func (o *GetConversationsMessageDetailsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations message details gateway timeout response has a 3xx status code
func (o *GetConversationsMessageDetailsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations message details gateway timeout response has a 4xx status code
func (o *GetConversationsMessageDetailsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations message details gateway timeout response has a 5xx status code
func (o *GetConversationsMessageDetailsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations message details gateway timeout response a status code equal to that given
func (o *GetConversationsMessageDetailsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsMessageDetailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessageDetailsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{messageId}/details][%d] getConversationsMessageDetailsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessageDetailsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageDetailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
