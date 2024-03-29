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

// DeleteRoutingQueueMemberReader is a Reader for the DeleteRoutingQueueMember structure.
type DeleteRoutingQueueMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingQueueMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingQueueMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingQueueMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingQueueMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingQueueMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingQueueMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingQueueMemberRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingQueueMemberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingQueueMemberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingQueueMemberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingQueueMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingQueueMemberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingQueueMemberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingQueueMemberNoContent creates a DeleteRoutingQueueMemberNoContent with default headers values
func NewDeleteRoutingQueueMemberNoContent() *DeleteRoutingQueueMemberNoContent {
	return &DeleteRoutingQueueMemberNoContent{}
}

/*
DeleteRoutingQueueMemberNoContent describes a response with status code 204, with default header values.

Deleted.
*/
type DeleteRoutingQueueMemberNoContent struct {
}

// IsSuccess returns true when this delete routing queue member no content response has a 2xx status code
func (o *DeleteRoutingQueueMemberNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing queue member no content response has a 3xx status code
func (o *DeleteRoutingQueueMemberNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member no content response has a 4xx status code
func (o *DeleteRoutingQueueMemberNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing queue member no content response has a 5xx status code
func (o *DeleteRoutingQueueMemberNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member no content response a status code equal to that given
func (o *DeleteRoutingQueueMemberNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteRoutingQueueMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberNoContent ", 204)
}

func (o *DeleteRoutingQueueMemberNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberNoContent ", 204)
}

func (o *DeleteRoutingQueueMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingQueueMemberBadRequest creates a DeleteRoutingQueueMemberBadRequest with default headers values
func NewDeleteRoutingQueueMemberBadRequest() *DeleteRoutingQueueMemberBadRequest {
	return &DeleteRoutingQueueMemberBadRequest{}
}

/*
DeleteRoutingQueueMemberBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingQueueMemberBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member bad request response has a 2xx status code
func (o *DeleteRoutingQueueMemberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member bad request response has a 3xx status code
func (o *DeleteRoutingQueueMemberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member bad request response has a 4xx status code
func (o *DeleteRoutingQueueMemberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member bad request response has a 5xx status code
func (o *DeleteRoutingQueueMemberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member bad request response a status code equal to that given
func (o *DeleteRoutingQueueMemberBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRoutingQueueMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingQueueMemberBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingQueueMemberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberUnauthorized creates a DeleteRoutingQueueMemberUnauthorized with default headers values
func NewDeleteRoutingQueueMemberUnauthorized() *DeleteRoutingQueueMemberUnauthorized {
	return &DeleteRoutingQueueMemberUnauthorized{}
}

/*
DeleteRoutingQueueMemberUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingQueueMemberUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member unauthorized response has a 2xx status code
func (o *DeleteRoutingQueueMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member unauthorized response has a 3xx status code
func (o *DeleteRoutingQueueMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member unauthorized response has a 4xx status code
func (o *DeleteRoutingQueueMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member unauthorized response has a 5xx status code
func (o *DeleteRoutingQueueMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member unauthorized response a status code equal to that given
func (o *DeleteRoutingQueueMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoutingQueueMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingQueueMemberUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingQueueMemberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberForbidden creates a DeleteRoutingQueueMemberForbidden with default headers values
func NewDeleteRoutingQueueMemberForbidden() *DeleteRoutingQueueMemberForbidden {
	return &DeleteRoutingQueueMemberForbidden{}
}

/*
DeleteRoutingQueueMemberForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingQueueMemberForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member forbidden response has a 2xx status code
func (o *DeleteRoutingQueueMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member forbidden response has a 3xx status code
func (o *DeleteRoutingQueueMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member forbidden response has a 4xx status code
func (o *DeleteRoutingQueueMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member forbidden response has a 5xx status code
func (o *DeleteRoutingQueueMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member forbidden response a status code equal to that given
func (o *DeleteRoutingQueueMemberForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoutingQueueMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingQueueMemberForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingQueueMemberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberNotFound creates a DeleteRoutingQueueMemberNotFound with default headers values
func NewDeleteRoutingQueueMemberNotFound() *DeleteRoutingQueueMemberNotFound {
	return &DeleteRoutingQueueMemberNotFound{}
}

/*
DeleteRoutingQueueMemberNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteRoutingQueueMemberNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member not found response has a 2xx status code
func (o *DeleteRoutingQueueMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member not found response has a 3xx status code
func (o *DeleteRoutingQueueMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member not found response has a 4xx status code
func (o *DeleteRoutingQueueMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member not found response has a 5xx status code
func (o *DeleteRoutingQueueMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member not found response a status code equal to that given
func (o *DeleteRoutingQueueMemberNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoutingQueueMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingQueueMemberNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingQueueMemberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberRequestTimeout creates a DeleteRoutingQueueMemberRequestTimeout with default headers values
func NewDeleteRoutingQueueMemberRequestTimeout() *DeleteRoutingQueueMemberRequestTimeout {
	return &DeleteRoutingQueueMemberRequestTimeout{}
}

/*
DeleteRoutingQueueMemberRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingQueueMemberRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member request timeout response has a 2xx status code
func (o *DeleteRoutingQueueMemberRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member request timeout response has a 3xx status code
func (o *DeleteRoutingQueueMemberRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member request timeout response has a 4xx status code
func (o *DeleteRoutingQueueMemberRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member request timeout response has a 5xx status code
func (o *DeleteRoutingQueueMemberRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member request timeout response a status code equal to that given
func (o *DeleteRoutingQueueMemberRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteRoutingQueueMemberRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingQueueMemberRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingQueueMemberRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberRequestEntityTooLarge creates a DeleteRoutingQueueMemberRequestEntityTooLarge with default headers values
func NewDeleteRoutingQueueMemberRequestEntityTooLarge() *DeleteRoutingQueueMemberRequestEntityTooLarge {
	return &DeleteRoutingQueueMemberRequestEntityTooLarge{}
}

/*
DeleteRoutingQueueMemberRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteRoutingQueueMemberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member request entity too large response has a 2xx status code
func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member request entity too large response has a 3xx status code
func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member request entity too large response has a 4xx status code
func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member request entity too large response has a 5xx status code
func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member request entity too large response a status code equal to that given
func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberUnsupportedMediaType creates a DeleteRoutingQueueMemberUnsupportedMediaType with default headers values
func NewDeleteRoutingQueueMemberUnsupportedMediaType() *DeleteRoutingQueueMemberUnsupportedMediaType {
	return &DeleteRoutingQueueMemberUnsupportedMediaType{}
}

/*
DeleteRoutingQueueMemberUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingQueueMemberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member unsupported media type response has a 2xx status code
func (o *DeleteRoutingQueueMemberUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member unsupported media type response has a 3xx status code
func (o *DeleteRoutingQueueMemberUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member unsupported media type response has a 4xx status code
func (o *DeleteRoutingQueueMemberUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member unsupported media type response has a 5xx status code
func (o *DeleteRoutingQueueMemberUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member unsupported media type response a status code equal to that given
func (o *DeleteRoutingQueueMemberUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteRoutingQueueMemberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingQueueMemberUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingQueueMemberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberTooManyRequests creates a DeleteRoutingQueueMemberTooManyRequests with default headers values
func NewDeleteRoutingQueueMemberTooManyRequests() *DeleteRoutingQueueMemberTooManyRequests {
	return &DeleteRoutingQueueMemberTooManyRequests{}
}

/*
DeleteRoutingQueueMemberTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingQueueMemberTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member too many requests response has a 2xx status code
func (o *DeleteRoutingQueueMemberTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member too many requests response has a 3xx status code
func (o *DeleteRoutingQueueMemberTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member too many requests response has a 4xx status code
func (o *DeleteRoutingQueueMemberTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing queue member too many requests response has a 5xx status code
func (o *DeleteRoutingQueueMemberTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing queue member too many requests response a status code equal to that given
func (o *DeleteRoutingQueueMemberTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRoutingQueueMemberTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingQueueMemberTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingQueueMemberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberInternalServerError creates a DeleteRoutingQueueMemberInternalServerError with default headers values
func NewDeleteRoutingQueueMemberInternalServerError() *DeleteRoutingQueueMemberInternalServerError {
	return &DeleteRoutingQueueMemberInternalServerError{}
}

/*
DeleteRoutingQueueMemberInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingQueueMemberInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member internal server error response has a 2xx status code
func (o *DeleteRoutingQueueMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member internal server error response has a 3xx status code
func (o *DeleteRoutingQueueMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member internal server error response has a 4xx status code
func (o *DeleteRoutingQueueMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing queue member internal server error response has a 5xx status code
func (o *DeleteRoutingQueueMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing queue member internal server error response a status code equal to that given
func (o *DeleteRoutingQueueMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoutingQueueMemberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingQueueMemberInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingQueueMemberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberServiceUnavailable creates a DeleteRoutingQueueMemberServiceUnavailable with default headers values
func NewDeleteRoutingQueueMemberServiceUnavailable() *DeleteRoutingQueueMemberServiceUnavailable {
	return &DeleteRoutingQueueMemberServiceUnavailable{}
}

/*
DeleteRoutingQueueMemberServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingQueueMemberServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member service unavailable response has a 2xx status code
func (o *DeleteRoutingQueueMemberServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member service unavailable response has a 3xx status code
func (o *DeleteRoutingQueueMemberServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member service unavailable response has a 4xx status code
func (o *DeleteRoutingQueueMemberServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing queue member service unavailable response has a 5xx status code
func (o *DeleteRoutingQueueMemberServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing queue member service unavailable response a status code equal to that given
func (o *DeleteRoutingQueueMemberServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteRoutingQueueMemberServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingQueueMemberServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingQueueMemberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberGatewayTimeout creates a DeleteRoutingQueueMemberGatewayTimeout with default headers values
func NewDeleteRoutingQueueMemberGatewayTimeout() *DeleteRoutingQueueMemberGatewayTimeout {
	return &DeleteRoutingQueueMemberGatewayTimeout{}
}

/*
DeleteRoutingQueueMemberGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteRoutingQueueMemberGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing queue member gateway timeout response has a 2xx status code
func (o *DeleteRoutingQueueMemberGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing queue member gateway timeout response has a 3xx status code
func (o *DeleteRoutingQueueMemberGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing queue member gateway timeout response has a 4xx status code
func (o *DeleteRoutingQueueMemberGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing queue member gateway timeout response has a 5xx status code
func (o *DeleteRoutingQueueMemberGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing queue member gateway timeout response a status code equal to that given
func (o *DeleteRoutingQueueMemberGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteRoutingQueueMemberGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingQueueMemberGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingQueueMemberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
