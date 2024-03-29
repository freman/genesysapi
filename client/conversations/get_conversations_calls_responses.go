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

// GetConversationsCallsReader is a Reader for the GetConversationsCalls structure.
type GetConversationsCallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallsOK creates a GetConversationsCallsOK with default headers values
func NewGetConversationsCallsOK() *GetConversationsCallsOK {
	return &GetConversationsCallsOK{}
}

/*
GetConversationsCallsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsCallsOK struct {
	Payload *models.CallConversationEntityListing
}

// IsSuccess returns true when this get conversations calls o k response has a 2xx status code
func (o *GetConversationsCallsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations calls o k response has a 3xx status code
func (o *GetConversationsCallsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls o k response has a 4xx status code
func (o *GetConversationsCallsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations calls o k response has a 5xx status code
func (o *GetConversationsCallsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls o k response a status code equal to that given
func (o *GetConversationsCallsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsCallsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallsOK) GetPayload() *models.CallConversationEntityListing {
	return o.Payload
}

func (o *GetConversationsCallsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallConversationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsBadRequest creates a GetConversationsCallsBadRequest with default headers values
func NewGetConversationsCallsBadRequest() *GetConversationsCallsBadRequest {
	return &GetConversationsCallsBadRequest{}
}

/*
GetConversationsCallsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls bad request response has a 2xx status code
func (o *GetConversationsCallsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls bad request response has a 3xx status code
func (o *GetConversationsCallsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls bad request response has a 4xx status code
func (o *GetConversationsCallsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls bad request response has a 5xx status code
func (o *GetConversationsCallsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls bad request response a status code equal to that given
func (o *GetConversationsCallsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsCallsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsUnauthorized creates a GetConversationsCallsUnauthorized with default headers values
func NewGetConversationsCallsUnauthorized() *GetConversationsCallsUnauthorized {
	return &GetConversationsCallsUnauthorized{}
}

/*
GetConversationsCallsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls unauthorized response has a 2xx status code
func (o *GetConversationsCallsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls unauthorized response has a 3xx status code
func (o *GetConversationsCallsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls unauthorized response has a 4xx status code
func (o *GetConversationsCallsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls unauthorized response has a 5xx status code
func (o *GetConversationsCallsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls unauthorized response a status code equal to that given
func (o *GetConversationsCallsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsCallsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsForbidden creates a GetConversationsCallsForbidden with default headers values
func NewGetConversationsCallsForbidden() *GetConversationsCallsForbidden {
	return &GetConversationsCallsForbidden{}
}

/*
GetConversationsCallsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls forbidden response has a 2xx status code
func (o *GetConversationsCallsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls forbidden response has a 3xx status code
func (o *GetConversationsCallsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls forbidden response has a 4xx status code
func (o *GetConversationsCallsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls forbidden response has a 5xx status code
func (o *GetConversationsCallsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls forbidden response a status code equal to that given
func (o *GetConversationsCallsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsCallsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsNotFound creates a GetConversationsCallsNotFound with default headers values
func NewGetConversationsCallsNotFound() *GetConversationsCallsNotFound {
	return &GetConversationsCallsNotFound{}
}

/*
GetConversationsCallsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsCallsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls not found response has a 2xx status code
func (o *GetConversationsCallsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls not found response has a 3xx status code
func (o *GetConversationsCallsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls not found response has a 4xx status code
func (o *GetConversationsCallsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls not found response has a 5xx status code
func (o *GetConversationsCallsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls not found response a status code equal to that given
func (o *GetConversationsCallsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsCallsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsRequestTimeout creates a GetConversationsCallsRequestTimeout with default headers values
func NewGetConversationsCallsRequestTimeout() *GetConversationsCallsRequestTimeout {
	return &GetConversationsCallsRequestTimeout{}
}

/*
GetConversationsCallsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls request timeout response has a 2xx status code
func (o *GetConversationsCallsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls request timeout response has a 3xx status code
func (o *GetConversationsCallsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls request timeout response has a 4xx status code
func (o *GetConversationsCallsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls request timeout response has a 5xx status code
func (o *GetConversationsCallsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls request timeout response a status code equal to that given
func (o *GetConversationsCallsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsCallsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsRequestEntityTooLarge creates a GetConversationsCallsRequestEntityTooLarge with default headers values
func NewGetConversationsCallsRequestEntityTooLarge() *GetConversationsCallsRequestEntityTooLarge {
	return &GetConversationsCallsRequestEntityTooLarge{}
}

/*
GetConversationsCallsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsCallsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls request entity too large response has a 2xx status code
func (o *GetConversationsCallsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls request entity too large response has a 3xx status code
func (o *GetConversationsCallsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls request entity too large response has a 4xx status code
func (o *GetConversationsCallsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls request entity too large response has a 5xx status code
func (o *GetConversationsCallsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls request entity too large response a status code equal to that given
func (o *GetConversationsCallsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsCallsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsUnsupportedMediaType creates a GetConversationsCallsUnsupportedMediaType with default headers values
func NewGetConversationsCallsUnsupportedMediaType() *GetConversationsCallsUnsupportedMediaType {
	return &GetConversationsCallsUnsupportedMediaType{}
}

/*
GetConversationsCallsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls unsupported media type response has a 2xx status code
func (o *GetConversationsCallsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls unsupported media type response has a 3xx status code
func (o *GetConversationsCallsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls unsupported media type response has a 4xx status code
func (o *GetConversationsCallsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls unsupported media type response has a 5xx status code
func (o *GetConversationsCallsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls unsupported media type response a status code equal to that given
func (o *GetConversationsCallsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsCallsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsTooManyRequests creates a GetConversationsCallsTooManyRequests with default headers values
func NewGetConversationsCallsTooManyRequests() *GetConversationsCallsTooManyRequests {
	return &GetConversationsCallsTooManyRequests{}
}

/*
GetConversationsCallsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls too many requests response has a 2xx status code
func (o *GetConversationsCallsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls too many requests response has a 3xx status code
func (o *GetConversationsCallsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls too many requests response has a 4xx status code
func (o *GetConversationsCallsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations calls too many requests response has a 5xx status code
func (o *GetConversationsCallsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations calls too many requests response a status code equal to that given
func (o *GetConversationsCallsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsCallsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsInternalServerError creates a GetConversationsCallsInternalServerError with default headers values
func NewGetConversationsCallsInternalServerError() *GetConversationsCallsInternalServerError {
	return &GetConversationsCallsInternalServerError{}
}

/*
GetConversationsCallsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls internal server error response has a 2xx status code
func (o *GetConversationsCallsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls internal server error response has a 3xx status code
func (o *GetConversationsCallsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls internal server error response has a 4xx status code
func (o *GetConversationsCallsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations calls internal server error response has a 5xx status code
func (o *GetConversationsCallsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations calls internal server error response a status code equal to that given
func (o *GetConversationsCallsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsCallsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsServiceUnavailable creates a GetConversationsCallsServiceUnavailable with default headers values
func NewGetConversationsCallsServiceUnavailable() *GetConversationsCallsServiceUnavailable {
	return &GetConversationsCallsServiceUnavailable{}
}

/*
GetConversationsCallsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls service unavailable response has a 2xx status code
func (o *GetConversationsCallsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls service unavailable response has a 3xx status code
func (o *GetConversationsCallsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls service unavailable response has a 4xx status code
func (o *GetConversationsCallsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations calls service unavailable response has a 5xx status code
func (o *GetConversationsCallsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations calls service unavailable response a status code equal to that given
func (o *GetConversationsCallsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsCallsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallsGatewayTimeout creates a GetConversationsCallsGatewayTimeout with default headers values
func NewGetConversationsCallsGatewayTimeout() *GetConversationsCallsGatewayTimeout {
	return &GetConversationsCallsGatewayTimeout{}
}

/*
GetConversationsCallsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsCallsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations calls gateway timeout response has a 2xx status code
func (o *GetConversationsCallsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations calls gateway timeout response has a 3xx status code
func (o *GetConversationsCallsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations calls gateway timeout response has a 4xx status code
func (o *GetConversationsCallsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations calls gateway timeout response has a 5xx status code
func (o *GetConversationsCallsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations calls gateway timeout response a status code equal to that given
func (o *GetConversationsCallsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsCallsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls][%d] getConversationsCallsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
