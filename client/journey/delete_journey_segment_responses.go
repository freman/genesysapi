// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteJourneySegmentReader is a Reader for the DeleteJourneySegment structure.
type DeleteJourneySegmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJourneySegmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteJourneySegmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteJourneySegmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteJourneySegmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteJourneySegmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteJourneySegmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteJourneySegmentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteJourneySegmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteJourneySegmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteJourneySegmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteJourneySegmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteJourneySegmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteJourneySegmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteJourneySegmentNoContent creates a DeleteJourneySegmentNoContent with default headers values
func NewDeleteJourneySegmentNoContent() *DeleteJourneySegmentNoContent {
	return &DeleteJourneySegmentNoContent{}
}

/*
DeleteJourneySegmentNoContent describes a response with status code 204, with default header values.

Segment deleted.
*/
type DeleteJourneySegmentNoContent struct {
}

// IsSuccess returns true when this delete journey segment no content response has a 2xx status code
func (o *DeleteJourneySegmentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete journey segment no content response has a 3xx status code
func (o *DeleteJourneySegmentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment no content response has a 4xx status code
func (o *DeleteJourneySegmentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete journey segment no content response has a 5xx status code
func (o *DeleteJourneySegmentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment no content response a status code equal to that given
func (o *DeleteJourneySegmentNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteJourneySegmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentNoContent ", 204)
}

func (o *DeleteJourneySegmentNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentNoContent ", 204)
}

func (o *DeleteJourneySegmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJourneySegmentBadRequest creates a DeleteJourneySegmentBadRequest with default headers values
func NewDeleteJourneySegmentBadRequest() *DeleteJourneySegmentBadRequest {
	return &DeleteJourneySegmentBadRequest{}
}

/*
DeleteJourneySegmentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteJourneySegmentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment bad request response has a 2xx status code
func (o *DeleteJourneySegmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment bad request response has a 3xx status code
func (o *DeleteJourneySegmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment bad request response has a 4xx status code
func (o *DeleteJourneySegmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment bad request response has a 5xx status code
func (o *DeleteJourneySegmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment bad request response a status code equal to that given
func (o *DeleteJourneySegmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteJourneySegmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteJourneySegmentBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteJourneySegmentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentUnauthorized creates a DeleteJourneySegmentUnauthorized with default headers values
func NewDeleteJourneySegmentUnauthorized() *DeleteJourneySegmentUnauthorized {
	return &DeleteJourneySegmentUnauthorized{}
}

/*
DeleteJourneySegmentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteJourneySegmentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment unauthorized response has a 2xx status code
func (o *DeleteJourneySegmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment unauthorized response has a 3xx status code
func (o *DeleteJourneySegmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment unauthorized response has a 4xx status code
func (o *DeleteJourneySegmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment unauthorized response has a 5xx status code
func (o *DeleteJourneySegmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment unauthorized response a status code equal to that given
func (o *DeleteJourneySegmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteJourneySegmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteJourneySegmentUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteJourneySegmentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentForbidden creates a DeleteJourneySegmentForbidden with default headers values
func NewDeleteJourneySegmentForbidden() *DeleteJourneySegmentForbidden {
	return &DeleteJourneySegmentForbidden{}
}

/*
DeleteJourneySegmentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteJourneySegmentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment forbidden response has a 2xx status code
func (o *DeleteJourneySegmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment forbidden response has a 3xx status code
func (o *DeleteJourneySegmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment forbidden response has a 4xx status code
func (o *DeleteJourneySegmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment forbidden response has a 5xx status code
func (o *DeleteJourneySegmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment forbidden response a status code equal to that given
func (o *DeleteJourneySegmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteJourneySegmentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteJourneySegmentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteJourneySegmentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentNotFound creates a DeleteJourneySegmentNotFound with default headers values
func NewDeleteJourneySegmentNotFound() *DeleteJourneySegmentNotFound {
	return &DeleteJourneySegmentNotFound{}
}

/*
DeleteJourneySegmentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteJourneySegmentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment not found response has a 2xx status code
func (o *DeleteJourneySegmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment not found response has a 3xx status code
func (o *DeleteJourneySegmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment not found response has a 4xx status code
func (o *DeleteJourneySegmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment not found response has a 5xx status code
func (o *DeleteJourneySegmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment not found response a status code equal to that given
func (o *DeleteJourneySegmentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteJourneySegmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteJourneySegmentNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteJourneySegmentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentRequestTimeout creates a DeleteJourneySegmentRequestTimeout with default headers values
func NewDeleteJourneySegmentRequestTimeout() *DeleteJourneySegmentRequestTimeout {
	return &DeleteJourneySegmentRequestTimeout{}
}

/*
DeleteJourneySegmentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteJourneySegmentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment request timeout response has a 2xx status code
func (o *DeleteJourneySegmentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment request timeout response has a 3xx status code
func (o *DeleteJourneySegmentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment request timeout response has a 4xx status code
func (o *DeleteJourneySegmentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment request timeout response has a 5xx status code
func (o *DeleteJourneySegmentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment request timeout response a status code equal to that given
func (o *DeleteJourneySegmentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteJourneySegmentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteJourneySegmentRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteJourneySegmentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentRequestEntityTooLarge creates a DeleteJourneySegmentRequestEntityTooLarge with default headers values
func NewDeleteJourneySegmentRequestEntityTooLarge() *DeleteJourneySegmentRequestEntityTooLarge {
	return &DeleteJourneySegmentRequestEntityTooLarge{}
}

/*
DeleteJourneySegmentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteJourneySegmentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment request entity too large response has a 2xx status code
func (o *DeleteJourneySegmentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment request entity too large response has a 3xx status code
func (o *DeleteJourneySegmentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment request entity too large response has a 4xx status code
func (o *DeleteJourneySegmentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment request entity too large response has a 5xx status code
func (o *DeleteJourneySegmentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment request entity too large response a status code equal to that given
func (o *DeleteJourneySegmentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteJourneySegmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteJourneySegmentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteJourneySegmentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentUnsupportedMediaType creates a DeleteJourneySegmentUnsupportedMediaType with default headers values
func NewDeleteJourneySegmentUnsupportedMediaType() *DeleteJourneySegmentUnsupportedMediaType {
	return &DeleteJourneySegmentUnsupportedMediaType{}
}

/*
DeleteJourneySegmentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteJourneySegmentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment unsupported media type response has a 2xx status code
func (o *DeleteJourneySegmentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment unsupported media type response has a 3xx status code
func (o *DeleteJourneySegmentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment unsupported media type response has a 4xx status code
func (o *DeleteJourneySegmentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment unsupported media type response has a 5xx status code
func (o *DeleteJourneySegmentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment unsupported media type response a status code equal to that given
func (o *DeleteJourneySegmentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteJourneySegmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteJourneySegmentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteJourneySegmentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentTooManyRequests creates a DeleteJourneySegmentTooManyRequests with default headers values
func NewDeleteJourneySegmentTooManyRequests() *DeleteJourneySegmentTooManyRequests {
	return &DeleteJourneySegmentTooManyRequests{}
}

/*
DeleteJourneySegmentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteJourneySegmentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment too many requests response has a 2xx status code
func (o *DeleteJourneySegmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment too many requests response has a 3xx status code
func (o *DeleteJourneySegmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment too many requests response has a 4xx status code
func (o *DeleteJourneySegmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete journey segment too many requests response has a 5xx status code
func (o *DeleteJourneySegmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete journey segment too many requests response a status code equal to that given
func (o *DeleteJourneySegmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteJourneySegmentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteJourneySegmentTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteJourneySegmentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentInternalServerError creates a DeleteJourneySegmentInternalServerError with default headers values
func NewDeleteJourneySegmentInternalServerError() *DeleteJourneySegmentInternalServerError {
	return &DeleteJourneySegmentInternalServerError{}
}

/*
DeleteJourneySegmentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteJourneySegmentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment internal server error response has a 2xx status code
func (o *DeleteJourneySegmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment internal server error response has a 3xx status code
func (o *DeleteJourneySegmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment internal server error response has a 4xx status code
func (o *DeleteJourneySegmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete journey segment internal server error response has a 5xx status code
func (o *DeleteJourneySegmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete journey segment internal server error response a status code equal to that given
func (o *DeleteJourneySegmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteJourneySegmentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteJourneySegmentInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteJourneySegmentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentServiceUnavailable creates a DeleteJourneySegmentServiceUnavailable with default headers values
func NewDeleteJourneySegmentServiceUnavailable() *DeleteJourneySegmentServiceUnavailable {
	return &DeleteJourneySegmentServiceUnavailable{}
}

/*
DeleteJourneySegmentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteJourneySegmentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment service unavailable response has a 2xx status code
func (o *DeleteJourneySegmentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment service unavailable response has a 3xx status code
func (o *DeleteJourneySegmentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment service unavailable response has a 4xx status code
func (o *DeleteJourneySegmentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete journey segment service unavailable response has a 5xx status code
func (o *DeleteJourneySegmentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete journey segment service unavailable response a status code equal to that given
func (o *DeleteJourneySegmentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteJourneySegmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteJourneySegmentServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteJourneySegmentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneySegmentGatewayTimeout creates a DeleteJourneySegmentGatewayTimeout with default headers values
func NewDeleteJourneySegmentGatewayTimeout() *DeleteJourneySegmentGatewayTimeout {
	return &DeleteJourneySegmentGatewayTimeout{}
}

/*
DeleteJourneySegmentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteJourneySegmentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete journey segment gateway timeout response has a 2xx status code
func (o *DeleteJourneySegmentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete journey segment gateway timeout response has a 3xx status code
func (o *DeleteJourneySegmentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete journey segment gateway timeout response has a 4xx status code
func (o *DeleteJourneySegmentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete journey segment gateway timeout response has a 5xx status code
func (o *DeleteJourneySegmentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete journey segment gateway timeout response a status code equal to that given
func (o *DeleteJourneySegmentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteJourneySegmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteJourneySegmentGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/segments/{segmentId}][%d] deleteJourneySegmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteJourneySegmentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneySegmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
