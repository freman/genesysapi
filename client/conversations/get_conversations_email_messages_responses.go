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

// GetConversationsEmailMessagesReader is a Reader for the GetConversationsEmailMessages structure.
type GetConversationsEmailMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsEmailMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsEmailMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsEmailMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsEmailMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsEmailMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsEmailMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsEmailMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsEmailMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsEmailMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsEmailMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsEmailMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsEmailMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsEmailMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsEmailMessagesOK creates a GetConversationsEmailMessagesOK with default headers values
func NewGetConversationsEmailMessagesOK() *GetConversationsEmailMessagesOK {
	return &GetConversationsEmailMessagesOK{}
}

/*
GetConversationsEmailMessagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsEmailMessagesOK struct {
	Payload *models.EmailMessageListing
}

// IsSuccess returns true when this get conversations email messages o k response has a 2xx status code
func (o *GetConversationsEmailMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations email messages o k response has a 3xx status code
func (o *GetConversationsEmailMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages o k response has a 4xx status code
func (o *GetConversationsEmailMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations email messages o k response has a 5xx status code
func (o *GetConversationsEmailMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages o k response a status code equal to that given
func (o *GetConversationsEmailMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsEmailMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsEmailMessagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsEmailMessagesOK) GetPayload() *models.EmailMessageListing {
	return o.Payload
}

func (o *GetConversationsEmailMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailMessageListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesBadRequest creates a GetConversationsEmailMessagesBadRequest with default headers values
func NewGetConversationsEmailMessagesBadRequest() *GetConversationsEmailMessagesBadRequest {
	return &GetConversationsEmailMessagesBadRequest{}
}

/*
GetConversationsEmailMessagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsEmailMessagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages bad request response has a 2xx status code
func (o *GetConversationsEmailMessagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages bad request response has a 3xx status code
func (o *GetConversationsEmailMessagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages bad request response has a 4xx status code
func (o *GetConversationsEmailMessagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages bad request response has a 5xx status code
func (o *GetConversationsEmailMessagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages bad request response a status code equal to that given
func (o *GetConversationsEmailMessagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsEmailMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsEmailMessagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsEmailMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesUnauthorized creates a GetConversationsEmailMessagesUnauthorized with default headers values
func NewGetConversationsEmailMessagesUnauthorized() *GetConversationsEmailMessagesUnauthorized {
	return &GetConversationsEmailMessagesUnauthorized{}
}

/*
GetConversationsEmailMessagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsEmailMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages unauthorized response has a 2xx status code
func (o *GetConversationsEmailMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages unauthorized response has a 3xx status code
func (o *GetConversationsEmailMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages unauthorized response has a 4xx status code
func (o *GetConversationsEmailMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages unauthorized response has a 5xx status code
func (o *GetConversationsEmailMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages unauthorized response a status code equal to that given
func (o *GetConversationsEmailMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsEmailMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsEmailMessagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsEmailMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesForbidden creates a GetConversationsEmailMessagesForbidden with default headers values
func NewGetConversationsEmailMessagesForbidden() *GetConversationsEmailMessagesForbidden {
	return &GetConversationsEmailMessagesForbidden{}
}

/*
GetConversationsEmailMessagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsEmailMessagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages forbidden response has a 2xx status code
func (o *GetConversationsEmailMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages forbidden response has a 3xx status code
func (o *GetConversationsEmailMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages forbidden response has a 4xx status code
func (o *GetConversationsEmailMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages forbidden response has a 5xx status code
func (o *GetConversationsEmailMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages forbidden response a status code equal to that given
func (o *GetConversationsEmailMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsEmailMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsEmailMessagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsEmailMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesNotFound creates a GetConversationsEmailMessagesNotFound with default headers values
func NewGetConversationsEmailMessagesNotFound() *GetConversationsEmailMessagesNotFound {
	return &GetConversationsEmailMessagesNotFound{}
}

/*
GetConversationsEmailMessagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsEmailMessagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages not found response has a 2xx status code
func (o *GetConversationsEmailMessagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages not found response has a 3xx status code
func (o *GetConversationsEmailMessagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages not found response has a 4xx status code
func (o *GetConversationsEmailMessagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages not found response has a 5xx status code
func (o *GetConversationsEmailMessagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages not found response a status code equal to that given
func (o *GetConversationsEmailMessagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsEmailMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsEmailMessagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsEmailMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesRequestTimeout creates a GetConversationsEmailMessagesRequestTimeout with default headers values
func NewGetConversationsEmailMessagesRequestTimeout() *GetConversationsEmailMessagesRequestTimeout {
	return &GetConversationsEmailMessagesRequestTimeout{}
}

/*
GetConversationsEmailMessagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsEmailMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages request timeout response has a 2xx status code
func (o *GetConversationsEmailMessagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages request timeout response has a 3xx status code
func (o *GetConversationsEmailMessagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages request timeout response has a 4xx status code
func (o *GetConversationsEmailMessagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages request timeout response has a 5xx status code
func (o *GetConversationsEmailMessagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages request timeout response a status code equal to that given
func (o *GetConversationsEmailMessagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsEmailMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsEmailMessagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsEmailMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesRequestEntityTooLarge creates a GetConversationsEmailMessagesRequestEntityTooLarge with default headers values
func NewGetConversationsEmailMessagesRequestEntityTooLarge() *GetConversationsEmailMessagesRequestEntityTooLarge {
	return &GetConversationsEmailMessagesRequestEntityTooLarge{}
}

/*
GetConversationsEmailMessagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsEmailMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages request entity too large response has a 2xx status code
func (o *GetConversationsEmailMessagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages request entity too large response has a 3xx status code
func (o *GetConversationsEmailMessagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages request entity too large response has a 4xx status code
func (o *GetConversationsEmailMessagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages request entity too large response has a 5xx status code
func (o *GetConversationsEmailMessagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages request entity too large response a status code equal to that given
func (o *GetConversationsEmailMessagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsEmailMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsEmailMessagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsEmailMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesUnsupportedMediaType creates a GetConversationsEmailMessagesUnsupportedMediaType with default headers values
func NewGetConversationsEmailMessagesUnsupportedMediaType() *GetConversationsEmailMessagesUnsupportedMediaType {
	return &GetConversationsEmailMessagesUnsupportedMediaType{}
}

/*
GetConversationsEmailMessagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsEmailMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages unsupported media type response has a 2xx status code
func (o *GetConversationsEmailMessagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages unsupported media type response has a 3xx status code
func (o *GetConversationsEmailMessagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages unsupported media type response has a 4xx status code
func (o *GetConversationsEmailMessagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages unsupported media type response has a 5xx status code
func (o *GetConversationsEmailMessagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages unsupported media type response a status code equal to that given
func (o *GetConversationsEmailMessagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsEmailMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsEmailMessagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsEmailMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesTooManyRequests creates a GetConversationsEmailMessagesTooManyRequests with default headers values
func NewGetConversationsEmailMessagesTooManyRequests() *GetConversationsEmailMessagesTooManyRequests {
	return &GetConversationsEmailMessagesTooManyRequests{}
}

/*
GetConversationsEmailMessagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsEmailMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages too many requests response has a 2xx status code
func (o *GetConversationsEmailMessagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages too many requests response has a 3xx status code
func (o *GetConversationsEmailMessagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages too many requests response has a 4xx status code
func (o *GetConversationsEmailMessagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations email messages too many requests response has a 5xx status code
func (o *GetConversationsEmailMessagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations email messages too many requests response a status code equal to that given
func (o *GetConversationsEmailMessagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsEmailMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsEmailMessagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsEmailMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesInternalServerError creates a GetConversationsEmailMessagesInternalServerError with default headers values
func NewGetConversationsEmailMessagesInternalServerError() *GetConversationsEmailMessagesInternalServerError {
	return &GetConversationsEmailMessagesInternalServerError{}
}

/*
GetConversationsEmailMessagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsEmailMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages internal server error response has a 2xx status code
func (o *GetConversationsEmailMessagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages internal server error response has a 3xx status code
func (o *GetConversationsEmailMessagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages internal server error response has a 4xx status code
func (o *GetConversationsEmailMessagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations email messages internal server error response has a 5xx status code
func (o *GetConversationsEmailMessagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations email messages internal server error response a status code equal to that given
func (o *GetConversationsEmailMessagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsEmailMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsEmailMessagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsEmailMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesServiceUnavailable creates a GetConversationsEmailMessagesServiceUnavailable with default headers values
func NewGetConversationsEmailMessagesServiceUnavailable() *GetConversationsEmailMessagesServiceUnavailable {
	return &GetConversationsEmailMessagesServiceUnavailable{}
}

/*
GetConversationsEmailMessagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsEmailMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages service unavailable response has a 2xx status code
func (o *GetConversationsEmailMessagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages service unavailable response has a 3xx status code
func (o *GetConversationsEmailMessagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages service unavailable response has a 4xx status code
func (o *GetConversationsEmailMessagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations email messages service unavailable response has a 5xx status code
func (o *GetConversationsEmailMessagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations email messages service unavailable response a status code equal to that given
func (o *GetConversationsEmailMessagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsEmailMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsEmailMessagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsEmailMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsEmailMessagesGatewayTimeout creates a GetConversationsEmailMessagesGatewayTimeout with default headers values
func NewGetConversationsEmailMessagesGatewayTimeout() *GetConversationsEmailMessagesGatewayTimeout {
	return &GetConversationsEmailMessagesGatewayTimeout{}
}

/*
GetConversationsEmailMessagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsEmailMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations email messages gateway timeout response has a 2xx status code
func (o *GetConversationsEmailMessagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations email messages gateway timeout response has a 3xx status code
func (o *GetConversationsEmailMessagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations email messages gateway timeout response has a 4xx status code
func (o *GetConversationsEmailMessagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations email messages gateway timeout response has a 5xx status code
func (o *GetConversationsEmailMessagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations email messages gateway timeout response a status code equal to that given
func (o *GetConversationsEmailMessagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsEmailMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsEmailMessagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/emails/{conversationId}/messages][%d] getConversationsEmailMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsEmailMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsEmailMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
