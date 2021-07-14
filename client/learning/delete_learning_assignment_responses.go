// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteLearningAssignmentReader is a Reader for the DeleteLearningAssignment structure.
type DeleteLearningAssignmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearningAssignmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearningAssignmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearningAssignmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLearningAssignmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearningAssignmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearningAssignmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteLearningAssignmentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteLearningAssignmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteLearningAssignmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteLearningAssignmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteLearningAssignmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLearningAssignmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteLearningAssignmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteLearningAssignmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearningAssignmentNoContent creates a DeleteLearningAssignmentNoContent with default headers values
func NewDeleteLearningAssignmentNoContent() *DeleteLearningAssignmentNoContent {
	return &DeleteLearningAssignmentNoContent{}
}

/*DeleteLearningAssignmentNoContent handles this case with default header values.

The learning assignment was deleted successfully
*/
type DeleteLearningAssignmentNoContent struct {
}

func (o *DeleteLearningAssignmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentNoContent ", 204)
}

func (o *DeleteLearningAssignmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearningAssignmentBadRequest creates a DeleteLearningAssignmentBadRequest with default headers values
func NewDeleteLearningAssignmentBadRequest() *DeleteLearningAssignmentBadRequest {
	return &DeleteLearningAssignmentBadRequest{}
}

/*DeleteLearningAssignmentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteLearningAssignmentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearningAssignmentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentUnauthorized creates a DeleteLearningAssignmentUnauthorized with default headers values
func NewDeleteLearningAssignmentUnauthorized() *DeleteLearningAssignmentUnauthorized {
	return &DeleteLearningAssignmentUnauthorized{}
}

/*DeleteLearningAssignmentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteLearningAssignmentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLearningAssignmentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentForbidden creates a DeleteLearningAssignmentForbidden with default headers values
func NewDeleteLearningAssignmentForbidden() *DeleteLearningAssignmentForbidden {
	return &DeleteLearningAssignmentForbidden{}
}

/*DeleteLearningAssignmentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteLearningAssignmentForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearningAssignmentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentNotFound creates a DeleteLearningAssignmentNotFound with default headers values
func NewDeleteLearningAssignmentNotFound() *DeleteLearningAssignmentNotFound {
	return &DeleteLearningAssignmentNotFound{}
}

/*DeleteLearningAssignmentNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteLearningAssignmentNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearningAssignmentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentRequestTimeout creates a DeleteLearningAssignmentRequestTimeout with default headers values
func NewDeleteLearningAssignmentRequestTimeout() *DeleteLearningAssignmentRequestTimeout {
	return &DeleteLearningAssignmentRequestTimeout{}
}

/*DeleteLearningAssignmentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteLearningAssignmentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteLearningAssignmentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentConflict creates a DeleteLearningAssignmentConflict with default headers values
func NewDeleteLearningAssignmentConflict() *DeleteLearningAssignmentConflict {
	return &DeleteLearningAssignmentConflict{}
}

/*DeleteLearningAssignmentConflict handles this case with default header values.

Conflict
*/
type DeleteLearningAssignmentConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentConflict  %+v", 409, o.Payload)
}

func (o *DeleteLearningAssignmentConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentRequestEntityTooLarge creates a DeleteLearningAssignmentRequestEntityTooLarge with default headers values
func NewDeleteLearningAssignmentRequestEntityTooLarge() *DeleteLearningAssignmentRequestEntityTooLarge {
	return &DeleteLearningAssignmentRequestEntityTooLarge{}
}

/*DeleteLearningAssignmentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteLearningAssignmentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteLearningAssignmentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentUnsupportedMediaType creates a DeleteLearningAssignmentUnsupportedMediaType with default headers values
func NewDeleteLearningAssignmentUnsupportedMediaType() *DeleteLearningAssignmentUnsupportedMediaType {
	return &DeleteLearningAssignmentUnsupportedMediaType{}
}

/*DeleteLearningAssignmentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteLearningAssignmentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteLearningAssignmentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentTooManyRequests creates a DeleteLearningAssignmentTooManyRequests with default headers values
func NewDeleteLearningAssignmentTooManyRequests() *DeleteLearningAssignmentTooManyRequests {
	return &DeleteLearningAssignmentTooManyRequests{}
}

/*DeleteLearningAssignmentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteLearningAssignmentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteLearningAssignmentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentInternalServerError creates a DeleteLearningAssignmentInternalServerError with default headers values
func NewDeleteLearningAssignmentInternalServerError() *DeleteLearningAssignmentInternalServerError {
	return &DeleteLearningAssignmentInternalServerError{}
}

/*DeleteLearningAssignmentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteLearningAssignmentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLearningAssignmentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentServiceUnavailable creates a DeleteLearningAssignmentServiceUnavailable with default headers values
func NewDeleteLearningAssignmentServiceUnavailable() *DeleteLearningAssignmentServiceUnavailable {
	return &DeleteLearningAssignmentServiceUnavailable{}
}

/*DeleteLearningAssignmentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteLearningAssignmentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLearningAssignmentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningAssignmentGatewayTimeout creates a DeleteLearningAssignmentGatewayTimeout with default headers values
func NewDeleteLearningAssignmentGatewayTimeout() *DeleteLearningAssignmentGatewayTimeout {
	return &DeleteLearningAssignmentGatewayTimeout{}
}

/*DeleteLearningAssignmentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteLearningAssignmentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningAssignmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/assignments/{assignmentId}][%d] deleteLearningAssignmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteLearningAssignmentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningAssignmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
