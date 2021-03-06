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

// DeleteWorkforcemanagementBusinessunitReader is a Reader for the DeleteWorkforcemanagementBusinessunit structure.
type DeleteWorkforcemanagementBusinessunitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementBusinessunitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkforcemanagementBusinessunitNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkforcemanagementBusinessunitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementBusinessunitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementBusinessunitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementBusinessunitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementBusinessunitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementBusinessunitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementBusinessunitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementBusinessunitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementBusinessunitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementBusinessunitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkforcemanagementBusinessunitNoContent creates a DeleteWorkforcemanagementBusinessunitNoContent with default headers values
func NewDeleteWorkforcemanagementBusinessunitNoContent() *DeleteWorkforcemanagementBusinessunitNoContent {
	return &DeleteWorkforcemanagementBusinessunitNoContent{}
}

/*DeleteWorkforcemanagementBusinessunitNoContent handles this case with default header values.

The business unit was successfully deleted
*/
type DeleteWorkforcemanagementBusinessunitNoContent struct {
}

func (o *DeleteWorkforcemanagementBusinessunitNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitNoContent ", 204)
}

func (o *DeleteWorkforcemanagementBusinessunitNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitBadRequest creates a DeleteWorkforcemanagementBusinessunitBadRequest with default headers values
func NewDeleteWorkforcemanagementBusinessunitBadRequest() *DeleteWorkforcemanagementBusinessunitBadRequest {
	return &DeleteWorkforcemanagementBusinessunitBadRequest{}
}

/*DeleteWorkforcemanagementBusinessunitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementBusinessunitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitUnauthorized creates a DeleteWorkforcemanagementBusinessunitUnauthorized with default headers values
func NewDeleteWorkforcemanagementBusinessunitUnauthorized() *DeleteWorkforcemanagementBusinessunitUnauthorized {
	return &DeleteWorkforcemanagementBusinessunitUnauthorized{}
}

/*DeleteWorkforcemanagementBusinessunitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementBusinessunitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitForbidden creates a DeleteWorkforcemanagementBusinessunitForbidden with default headers values
func NewDeleteWorkforcemanagementBusinessunitForbidden() *DeleteWorkforcemanagementBusinessunitForbidden {
	return &DeleteWorkforcemanagementBusinessunitForbidden{}
}

/*DeleteWorkforcemanagementBusinessunitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementBusinessunitForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitNotFound creates a DeleteWorkforcemanagementBusinessunitNotFound with default headers values
func NewDeleteWorkforcemanagementBusinessunitNotFound() *DeleteWorkforcemanagementBusinessunitNotFound {
	return &DeleteWorkforcemanagementBusinessunitNotFound{}
}

/*DeleteWorkforcemanagementBusinessunitNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementBusinessunitNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitRequestEntityTooLarge creates a DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementBusinessunitRequestEntityTooLarge() *DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge {
	return &DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge{}
}

/*DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitUnsupportedMediaType creates a DeleteWorkforcemanagementBusinessunitUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementBusinessunitUnsupportedMediaType() *DeleteWorkforcemanagementBusinessunitUnsupportedMediaType {
	return &DeleteWorkforcemanagementBusinessunitUnsupportedMediaType{}
}

/*DeleteWorkforcemanagementBusinessunitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementBusinessunitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitTooManyRequests creates a DeleteWorkforcemanagementBusinessunitTooManyRequests with default headers values
func NewDeleteWorkforcemanagementBusinessunitTooManyRequests() *DeleteWorkforcemanagementBusinessunitTooManyRequests {
	return &DeleteWorkforcemanagementBusinessunitTooManyRequests{}
}

/*DeleteWorkforcemanagementBusinessunitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteWorkforcemanagementBusinessunitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitInternalServerError creates a DeleteWorkforcemanagementBusinessunitInternalServerError with default headers values
func NewDeleteWorkforcemanagementBusinessunitInternalServerError() *DeleteWorkforcemanagementBusinessunitInternalServerError {
	return &DeleteWorkforcemanagementBusinessunitInternalServerError{}
}

/*DeleteWorkforcemanagementBusinessunitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementBusinessunitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitServiceUnavailable creates a DeleteWorkforcemanagementBusinessunitServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementBusinessunitServiceUnavailable() *DeleteWorkforcemanagementBusinessunitServiceUnavailable {
	return &DeleteWorkforcemanagementBusinessunitServiceUnavailable{}
}

/*DeleteWorkforcemanagementBusinessunitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementBusinessunitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitGatewayTimeout creates a DeleteWorkforcemanagementBusinessunitGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementBusinessunitGatewayTimeout() *DeleteWorkforcemanagementBusinessunitGatewayTimeout {
	return &DeleteWorkforcemanagementBusinessunitGatewayTimeout{}
}

/*DeleteWorkforcemanagementBusinessunitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementBusinessunitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}][%d] deleteWorkforcemanagementBusinessunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
