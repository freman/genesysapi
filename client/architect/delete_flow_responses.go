// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteFlowReader is a Reader for the DeleteFlow structure.
type DeleteFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteFlowMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteFlowConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteFlowGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteFlowRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteFlowUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFlowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFlowServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFlowGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteFlowDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteFlowBadRequest creates a DeleteFlowBadRequest with default headers values
func NewDeleteFlowBadRequest() *DeleteFlowBadRequest {
	return &DeleteFlowBadRequest{}
}

/*DeleteFlowBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteFlowBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowUnauthorized creates a DeleteFlowUnauthorized with default headers values
func NewDeleteFlowUnauthorized() *DeleteFlowUnauthorized {
	return &DeleteFlowUnauthorized{}
}

/*DeleteFlowUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteFlowUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowForbidden creates a DeleteFlowForbidden with default headers values
func NewDeleteFlowForbidden() *DeleteFlowForbidden {
	return &DeleteFlowForbidden{}
}

/*DeleteFlowForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteFlowForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowNotFound creates a DeleteFlowNotFound with default headers values
func NewDeleteFlowNotFound() *DeleteFlowNotFound {
	return &DeleteFlowNotFound{}
}

/*DeleteFlowNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteFlowNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowMethodNotAllowed creates a DeleteFlowMethodNotAllowed with default headers values
func NewDeleteFlowMethodNotAllowed() *DeleteFlowMethodNotAllowed {
	return &DeleteFlowMethodNotAllowed{}
}

/*DeleteFlowMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type DeleteFlowMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeleteFlowMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowConflict creates a DeleteFlowConflict with default headers values
func NewDeleteFlowConflict() *DeleteFlowConflict {
	return &DeleteFlowConflict{}
}

/*DeleteFlowConflict handles this case with default header values.

Conflict
*/
type DeleteFlowConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowConflict  %+v", 409, o.Payload)
}

func (o *DeleteFlowConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowGone creates a DeleteFlowGone with default headers values
func NewDeleteFlowGone() *DeleteFlowGone {
	return &DeleteFlowGone{}
}

/*DeleteFlowGone handles this case with default header values.

Gone
*/
type DeleteFlowGone struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowGone) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowGone  %+v", 410, o.Payload)
}

func (o *DeleteFlowGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowRequestEntityTooLarge creates a DeleteFlowRequestEntityTooLarge with default headers values
func NewDeleteFlowRequestEntityTooLarge() *DeleteFlowRequestEntityTooLarge {
	return &DeleteFlowRequestEntityTooLarge{}
}

/*DeleteFlowRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteFlowRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowUnsupportedMediaType creates a DeleteFlowUnsupportedMediaType with default headers values
func NewDeleteFlowUnsupportedMediaType() *DeleteFlowUnsupportedMediaType {
	return &DeleteFlowUnsupportedMediaType{}
}

/*DeleteFlowUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteFlowUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowTooManyRequests creates a DeleteFlowTooManyRequests with default headers values
func NewDeleteFlowTooManyRequests() *DeleteFlowTooManyRequests {
	return &DeleteFlowTooManyRequests{}
}

/*DeleteFlowTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteFlowTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowInternalServerError creates a DeleteFlowInternalServerError with default headers values
func NewDeleteFlowInternalServerError() *DeleteFlowInternalServerError {
	return &DeleteFlowInternalServerError{}
}

/*DeleteFlowInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteFlowInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowServiceUnavailable creates a DeleteFlowServiceUnavailable with default headers values
func NewDeleteFlowServiceUnavailable() *DeleteFlowServiceUnavailable {
	return &DeleteFlowServiceUnavailable{}
}

/*DeleteFlowServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteFlowServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowGatewayTimeout creates a DeleteFlowGatewayTimeout with default headers values
func NewDeleteFlowGatewayTimeout() *DeleteFlowGatewayTimeout {
	return &DeleteFlowGatewayTimeout{}
}

/*DeleteFlowGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteFlowGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowDefault creates a DeleteFlowDefault with default headers values
func NewDeleteFlowDefault(code int) *DeleteFlowDefault {
	return &DeleteFlowDefault{
		_statusCode: code,
	}
}

/*DeleteFlowDefault handles this case with default header values.

successful operation
*/
type DeleteFlowDefault struct {
	_statusCode int
}

// Code gets the status code for the delete flow default response
func (o *DeleteFlowDefault) Code() int {
	return o._statusCode
}

func (o *DeleteFlowDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/{flowId}][%d] deleteFlow default ", o._statusCode)
}

func (o *DeleteFlowDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
