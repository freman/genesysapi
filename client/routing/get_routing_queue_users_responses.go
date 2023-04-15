// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingQueueUsersReader is a Reader for the GetRoutingQueueUsers structure.
type GetRoutingQueueUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueueUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueueUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueueUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueueUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueueUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueueUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingQueueUsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueueUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueueUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueueUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueueUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueueUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueueUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueueUsersOK creates a GetRoutingQueueUsersOK with default headers values
func NewGetRoutingQueueUsersOK() *GetRoutingQueueUsersOK {
	return &GetRoutingQueueUsersOK{}
}

/*
GetRoutingQueueUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingQueueUsersOK struct {
	Payload *models.QueueMemberEntityListingV1
}

// IsSuccess returns true when this get routing queue users o k response has a 2xx status code
func (o *GetRoutingQueueUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing queue users o k response has a 3xx status code
func (o *GetRoutingQueueUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users o k response has a 4xx status code
func (o *GetRoutingQueueUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue users o k response has a 5xx status code
func (o *GetRoutingQueueUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users o k response a status code equal to that given
func (o *GetRoutingQueueUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingQueueUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueueUsersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueueUsersOK) GetPayload() *models.QueueMemberEntityListingV1 {
	return o.Payload
}

func (o *GetRoutingQueueUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueueMemberEntityListingV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersBadRequest creates a GetRoutingQueueUsersBadRequest with default headers values
func NewGetRoutingQueueUsersBadRequest() *GetRoutingQueueUsersBadRequest {
	return &GetRoutingQueueUsersBadRequest{}
}

/*
GetRoutingQueueUsersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueueUsersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users bad request response has a 2xx status code
func (o *GetRoutingQueueUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users bad request response has a 3xx status code
func (o *GetRoutingQueueUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users bad request response has a 4xx status code
func (o *GetRoutingQueueUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users bad request response has a 5xx status code
func (o *GetRoutingQueueUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users bad request response a status code equal to that given
func (o *GetRoutingQueueUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingQueueUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueueUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueueUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersUnauthorized creates a GetRoutingQueueUsersUnauthorized with default headers values
func NewGetRoutingQueueUsersUnauthorized() *GetRoutingQueueUsersUnauthorized {
	return &GetRoutingQueueUsersUnauthorized{}
}

/*
GetRoutingQueueUsersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueueUsersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users unauthorized response has a 2xx status code
func (o *GetRoutingQueueUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users unauthorized response has a 3xx status code
func (o *GetRoutingQueueUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users unauthorized response has a 4xx status code
func (o *GetRoutingQueueUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users unauthorized response has a 5xx status code
func (o *GetRoutingQueueUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users unauthorized response a status code equal to that given
func (o *GetRoutingQueueUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingQueueUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueueUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueueUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersForbidden creates a GetRoutingQueueUsersForbidden with default headers values
func NewGetRoutingQueueUsersForbidden() *GetRoutingQueueUsersForbidden {
	return &GetRoutingQueueUsersForbidden{}
}

/*
GetRoutingQueueUsersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueueUsersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users forbidden response has a 2xx status code
func (o *GetRoutingQueueUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users forbidden response has a 3xx status code
func (o *GetRoutingQueueUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users forbidden response has a 4xx status code
func (o *GetRoutingQueueUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users forbidden response has a 5xx status code
func (o *GetRoutingQueueUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users forbidden response a status code equal to that given
func (o *GetRoutingQueueUsersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingQueueUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueueUsersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueueUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersNotFound creates a GetRoutingQueueUsersNotFound with default headers values
func NewGetRoutingQueueUsersNotFound() *GetRoutingQueueUsersNotFound {
	return &GetRoutingQueueUsersNotFound{}
}

/*
GetRoutingQueueUsersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingQueueUsersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users not found response has a 2xx status code
func (o *GetRoutingQueueUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users not found response has a 3xx status code
func (o *GetRoutingQueueUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users not found response has a 4xx status code
func (o *GetRoutingQueueUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users not found response has a 5xx status code
func (o *GetRoutingQueueUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users not found response a status code equal to that given
func (o *GetRoutingQueueUsersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingQueueUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueueUsersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueueUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersRequestTimeout creates a GetRoutingQueueUsersRequestTimeout with default headers values
func NewGetRoutingQueueUsersRequestTimeout() *GetRoutingQueueUsersRequestTimeout {
	return &GetRoutingQueueUsersRequestTimeout{}
}

/*
GetRoutingQueueUsersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingQueueUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users request timeout response has a 2xx status code
func (o *GetRoutingQueueUsersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users request timeout response has a 3xx status code
func (o *GetRoutingQueueUsersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users request timeout response has a 4xx status code
func (o *GetRoutingQueueUsersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users request timeout response has a 5xx status code
func (o *GetRoutingQueueUsersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users request timeout response a status code equal to that given
func (o *GetRoutingQueueUsersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingQueueUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingQueueUsersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingQueueUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersRequestEntityTooLarge creates a GetRoutingQueueUsersRequestEntityTooLarge with default headers values
func NewGetRoutingQueueUsersRequestEntityTooLarge() *GetRoutingQueueUsersRequestEntityTooLarge {
	return &GetRoutingQueueUsersRequestEntityTooLarge{}
}

/*
GetRoutingQueueUsersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingQueueUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users request entity too large response has a 2xx status code
func (o *GetRoutingQueueUsersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users request entity too large response has a 3xx status code
func (o *GetRoutingQueueUsersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users request entity too large response has a 4xx status code
func (o *GetRoutingQueueUsersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users request entity too large response has a 5xx status code
func (o *GetRoutingQueueUsersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users request entity too large response a status code equal to that given
func (o *GetRoutingQueueUsersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingQueueUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueueUsersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueueUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersUnsupportedMediaType creates a GetRoutingQueueUsersUnsupportedMediaType with default headers values
func NewGetRoutingQueueUsersUnsupportedMediaType() *GetRoutingQueueUsersUnsupportedMediaType {
	return &GetRoutingQueueUsersUnsupportedMediaType{}
}

/*
GetRoutingQueueUsersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueueUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users unsupported media type response has a 2xx status code
func (o *GetRoutingQueueUsersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users unsupported media type response has a 3xx status code
func (o *GetRoutingQueueUsersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users unsupported media type response has a 4xx status code
func (o *GetRoutingQueueUsersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users unsupported media type response has a 5xx status code
func (o *GetRoutingQueueUsersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users unsupported media type response a status code equal to that given
func (o *GetRoutingQueueUsersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingQueueUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueueUsersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueueUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersTooManyRequests creates a GetRoutingQueueUsersTooManyRequests with default headers values
func NewGetRoutingQueueUsersTooManyRequests() *GetRoutingQueueUsersTooManyRequests {
	return &GetRoutingQueueUsersTooManyRequests{}
}

/*
GetRoutingQueueUsersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingQueueUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users too many requests response has a 2xx status code
func (o *GetRoutingQueueUsersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users too many requests response has a 3xx status code
func (o *GetRoutingQueueUsersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users too many requests response has a 4xx status code
func (o *GetRoutingQueueUsersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue users too many requests response has a 5xx status code
func (o *GetRoutingQueueUsersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue users too many requests response a status code equal to that given
func (o *GetRoutingQueueUsersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingQueueUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueueUsersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueueUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersInternalServerError creates a GetRoutingQueueUsersInternalServerError with default headers values
func NewGetRoutingQueueUsersInternalServerError() *GetRoutingQueueUsersInternalServerError {
	return &GetRoutingQueueUsersInternalServerError{}
}

/*
GetRoutingQueueUsersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueueUsersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users internal server error response has a 2xx status code
func (o *GetRoutingQueueUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users internal server error response has a 3xx status code
func (o *GetRoutingQueueUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users internal server error response has a 4xx status code
func (o *GetRoutingQueueUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue users internal server error response has a 5xx status code
func (o *GetRoutingQueueUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing queue users internal server error response a status code equal to that given
func (o *GetRoutingQueueUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingQueueUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueueUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueueUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersServiceUnavailable creates a GetRoutingQueueUsersServiceUnavailable with default headers values
func NewGetRoutingQueueUsersServiceUnavailable() *GetRoutingQueueUsersServiceUnavailable {
	return &GetRoutingQueueUsersServiceUnavailable{}
}

/*
GetRoutingQueueUsersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueueUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users service unavailable response has a 2xx status code
func (o *GetRoutingQueueUsersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users service unavailable response has a 3xx status code
func (o *GetRoutingQueueUsersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users service unavailable response has a 4xx status code
func (o *GetRoutingQueueUsersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue users service unavailable response has a 5xx status code
func (o *GetRoutingQueueUsersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing queue users service unavailable response a status code equal to that given
func (o *GetRoutingQueueUsersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingQueueUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueueUsersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueueUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUsersGatewayTimeout creates a GetRoutingQueueUsersGatewayTimeout with default headers values
func NewGetRoutingQueueUsersGatewayTimeout() *GetRoutingQueueUsersGatewayTimeout {
	return &GetRoutingQueueUsersGatewayTimeout{}
}

/*
GetRoutingQueueUsersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingQueueUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue users gateway timeout response has a 2xx status code
func (o *GetRoutingQueueUsersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue users gateway timeout response has a 3xx status code
func (o *GetRoutingQueueUsersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue users gateway timeout response has a 4xx status code
func (o *GetRoutingQueueUsersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue users gateway timeout response has a 5xx status code
func (o *GetRoutingQueueUsersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing queue users gateway timeout response a status code equal to that given
func (o *GetRoutingQueueUsersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingQueueUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueueUsersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/users][%d] getRoutingQueueUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueueUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
