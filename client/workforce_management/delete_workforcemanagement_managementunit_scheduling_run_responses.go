// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteWorkforcemanagementManagementunitSchedulingRunReader is a Reader for the DeleteWorkforcemanagementManagementunitSchedulingRun structure.
type DeleteWorkforcemanagementManagementunitSchedulingRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementManagementunitSchedulingRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteWorkforcemanagementManagementunitSchedulingRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunBadRequest creates a DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunBadRequest() *DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized creates a DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized() *DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunForbidden creates a DeleteWorkforcemanagementManagementunitSchedulingRunForbidden with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunForbidden() *DeleteWorkforcemanagementManagementunitSchedulingRunForbidden {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunForbidden{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunNotFound creates a DeleteWorkforcemanagementManagementunitSchedulingRunNotFound with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunNotFound() *DeleteWorkforcemanagementManagementunitSchedulingRunNotFound {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunNotFound{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunGone creates a DeleteWorkforcemanagementManagementunitSchedulingRunGone with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunGone() *DeleteWorkforcemanagementManagementunitSchedulingRunGone {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunGone{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunGone handles this case with default header values.

Gone
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunGone struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunGone) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunGone  %+v", 410, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge creates a DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge() *DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType creates a DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType() *DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests creates a DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests() *DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError creates a DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError() *DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable creates a DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable() *DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout creates a DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout() *DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout{}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitSchedulingRunDefault creates a DeleteWorkforcemanagementManagementunitSchedulingRunDefault with default headers values
func NewDeleteWorkforcemanagementManagementunitSchedulingRunDefault(code int) *DeleteWorkforcemanagementManagementunitSchedulingRunDefault {
	return &DeleteWorkforcemanagementManagementunitSchedulingRunDefault{
		_statusCode: code,
	}
}

/*DeleteWorkforcemanagementManagementunitSchedulingRunDefault handles this case with default header values.

successful operation
*/
type DeleteWorkforcemanagementManagementunitSchedulingRunDefault struct {
	_statusCode int
}

// Code gets the status code for the delete workforcemanagement managementunit scheduling run default response
func (o *DeleteWorkforcemanagementManagementunitSchedulingRunDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}/scheduling/runs/{runId}][%d] deleteWorkforcemanagementManagementunitSchedulingRun default ", o._statusCode)
}

func (o *DeleteWorkforcemanagementManagementunitSchedulingRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
