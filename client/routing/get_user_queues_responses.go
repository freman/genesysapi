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

// GetUserQueuesReader is a Reader for the GetUserQueues structure.
type GetUserQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserQueuesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserQueuesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserQueuesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserQueuesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserQueuesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserQueuesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserQueuesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserQueuesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserQueuesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserQueuesOK creates a GetUserQueuesOK with default headers values
func NewGetUserQueuesOK() *GetUserQueuesOK {
	return &GetUserQueuesOK{}
}

/*
GetUserQueuesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserQueuesOK struct {
	Payload *models.UserQueueEntityListing
}

// IsSuccess returns true when this get user queues o k response has a 2xx status code
func (o *GetUserQueuesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user queues o k response has a 3xx status code
func (o *GetUserQueuesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues o k response has a 4xx status code
func (o *GetUserQueuesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user queues o k response has a 5xx status code
func (o *GetUserQueuesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues o k response a status code equal to that given
func (o *GetUserQueuesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserQueuesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesOK  %+v", 200, o.Payload)
}

func (o *GetUserQueuesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesOK  %+v", 200, o.Payload)
}

func (o *GetUserQueuesOK) GetPayload() *models.UserQueueEntityListing {
	return o.Payload
}

func (o *GetUserQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserQueueEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesBadRequest creates a GetUserQueuesBadRequest with default headers values
func NewGetUserQueuesBadRequest() *GetUserQueuesBadRequest {
	return &GetUserQueuesBadRequest{}
}

/*
GetUserQueuesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserQueuesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues bad request response has a 2xx status code
func (o *GetUserQueuesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues bad request response has a 3xx status code
func (o *GetUserQueuesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues bad request response has a 4xx status code
func (o *GetUserQueuesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues bad request response has a 5xx status code
func (o *GetUserQueuesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues bad request response a status code equal to that given
func (o *GetUserQueuesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserQueuesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserQueuesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserQueuesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesUnauthorized creates a GetUserQueuesUnauthorized with default headers values
func NewGetUserQueuesUnauthorized() *GetUserQueuesUnauthorized {
	return &GetUserQueuesUnauthorized{}
}

/*
GetUserQueuesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserQueuesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues unauthorized response has a 2xx status code
func (o *GetUserQueuesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues unauthorized response has a 3xx status code
func (o *GetUserQueuesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues unauthorized response has a 4xx status code
func (o *GetUserQueuesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues unauthorized response has a 5xx status code
func (o *GetUserQueuesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues unauthorized response a status code equal to that given
func (o *GetUserQueuesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserQueuesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserQueuesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserQueuesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesForbidden creates a GetUserQueuesForbidden with default headers values
func NewGetUserQueuesForbidden() *GetUserQueuesForbidden {
	return &GetUserQueuesForbidden{}
}

/*
GetUserQueuesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserQueuesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues forbidden response has a 2xx status code
func (o *GetUserQueuesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues forbidden response has a 3xx status code
func (o *GetUserQueuesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues forbidden response has a 4xx status code
func (o *GetUserQueuesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues forbidden response has a 5xx status code
func (o *GetUserQueuesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues forbidden response a status code equal to that given
func (o *GetUserQueuesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserQueuesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserQueuesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserQueuesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesNotFound creates a GetUserQueuesNotFound with default headers values
func NewGetUserQueuesNotFound() *GetUserQueuesNotFound {
	return &GetUserQueuesNotFound{}
}

/*
GetUserQueuesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserQueuesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues not found response has a 2xx status code
func (o *GetUserQueuesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues not found response has a 3xx status code
func (o *GetUserQueuesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues not found response has a 4xx status code
func (o *GetUserQueuesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues not found response has a 5xx status code
func (o *GetUserQueuesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues not found response a status code equal to that given
func (o *GetUserQueuesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserQueuesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserQueuesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserQueuesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesRequestTimeout creates a GetUserQueuesRequestTimeout with default headers values
func NewGetUserQueuesRequestTimeout() *GetUserQueuesRequestTimeout {
	return &GetUserQueuesRequestTimeout{}
}

/*
GetUserQueuesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserQueuesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues request timeout response has a 2xx status code
func (o *GetUserQueuesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues request timeout response has a 3xx status code
func (o *GetUserQueuesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues request timeout response has a 4xx status code
func (o *GetUserQueuesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues request timeout response has a 5xx status code
func (o *GetUserQueuesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues request timeout response a status code equal to that given
func (o *GetUserQueuesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserQueuesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserQueuesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserQueuesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesRequestEntityTooLarge creates a GetUserQueuesRequestEntityTooLarge with default headers values
func NewGetUserQueuesRequestEntityTooLarge() *GetUserQueuesRequestEntityTooLarge {
	return &GetUserQueuesRequestEntityTooLarge{}
}

/*
GetUserQueuesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUserQueuesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues request entity too large response has a 2xx status code
func (o *GetUserQueuesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues request entity too large response has a 3xx status code
func (o *GetUserQueuesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues request entity too large response has a 4xx status code
func (o *GetUserQueuesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues request entity too large response has a 5xx status code
func (o *GetUserQueuesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues request entity too large response a status code equal to that given
func (o *GetUserQueuesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserQueuesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserQueuesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserQueuesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesUnsupportedMediaType creates a GetUserQueuesUnsupportedMediaType with default headers values
func NewGetUserQueuesUnsupportedMediaType() *GetUserQueuesUnsupportedMediaType {
	return &GetUserQueuesUnsupportedMediaType{}
}

/*
GetUserQueuesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserQueuesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues unsupported media type response has a 2xx status code
func (o *GetUserQueuesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues unsupported media type response has a 3xx status code
func (o *GetUserQueuesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues unsupported media type response has a 4xx status code
func (o *GetUserQueuesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues unsupported media type response has a 5xx status code
func (o *GetUserQueuesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues unsupported media type response a status code equal to that given
func (o *GetUserQueuesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserQueuesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserQueuesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserQueuesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesTooManyRequests creates a GetUserQueuesTooManyRequests with default headers values
func NewGetUserQueuesTooManyRequests() *GetUserQueuesTooManyRequests {
	return &GetUserQueuesTooManyRequests{}
}

/*
GetUserQueuesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserQueuesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues too many requests response has a 2xx status code
func (o *GetUserQueuesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues too many requests response has a 3xx status code
func (o *GetUserQueuesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues too many requests response has a 4xx status code
func (o *GetUserQueuesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user queues too many requests response has a 5xx status code
func (o *GetUserQueuesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user queues too many requests response a status code equal to that given
func (o *GetUserQueuesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserQueuesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserQueuesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserQueuesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesInternalServerError creates a GetUserQueuesInternalServerError with default headers values
func NewGetUserQueuesInternalServerError() *GetUserQueuesInternalServerError {
	return &GetUserQueuesInternalServerError{}
}

/*
GetUserQueuesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserQueuesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues internal server error response has a 2xx status code
func (o *GetUserQueuesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues internal server error response has a 3xx status code
func (o *GetUserQueuesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues internal server error response has a 4xx status code
func (o *GetUserQueuesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user queues internal server error response has a 5xx status code
func (o *GetUserQueuesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user queues internal server error response a status code equal to that given
func (o *GetUserQueuesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserQueuesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserQueuesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserQueuesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesServiceUnavailable creates a GetUserQueuesServiceUnavailable with default headers values
func NewGetUserQueuesServiceUnavailable() *GetUserQueuesServiceUnavailable {
	return &GetUserQueuesServiceUnavailable{}
}

/*
GetUserQueuesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserQueuesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues service unavailable response has a 2xx status code
func (o *GetUserQueuesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues service unavailable response has a 3xx status code
func (o *GetUserQueuesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues service unavailable response has a 4xx status code
func (o *GetUserQueuesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user queues service unavailable response has a 5xx status code
func (o *GetUserQueuesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user queues service unavailable response a status code equal to that given
func (o *GetUserQueuesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserQueuesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserQueuesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserQueuesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesGatewayTimeout creates a GetUserQueuesGatewayTimeout with default headers values
func NewGetUserQueuesGatewayTimeout() *GetUserQueuesGatewayTimeout {
	return &GetUserQueuesGatewayTimeout{}
}

/*
GetUserQueuesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserQueuesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user queues gateway timeout response has a 2xx status code
func (o *GetUserQueuesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user queues gateway timeout response has a 3xx status code
func (o *GetUserQueuesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user queues gateway timeout response has a 4xx status code
func (o *GetUserQueuesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user queues gateway timeout response has a 5xx status code
func (o *GetUserQueuesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user queues gateway timeout response a status code equal to that given
func (o *GetUserQueuesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserQueuesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserQueuesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserQueuesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
