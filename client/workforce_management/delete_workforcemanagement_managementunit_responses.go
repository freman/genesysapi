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

// DeleteWorkforcemanagementManagementunitReader is a Reader for the DeleteWorkforcemanagementManagementunit structure.
type DeleteWorkforcemanagementManagementunitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementManagementunitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkforcemanagementManagementunitNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkforcemanagementManagementunitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementManagementunitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementManagementunitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementManagementunitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementManagementunitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementManagementunitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementManagementunitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementManagementunitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementManagementunitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementManagementunitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkforcemanagementManagementunitNoContent creates a DeleteWorkforcemanagementManagementunitNoContent with default headers values
func NewDeleteWorkforcemanagementManagementunitNoContent() *DeleteWorkforcemanagementManagementunitNoContent {
	return &DeleteWorkforcemanagementManagementunitNoContent{}
}

/*DeleteWorkforcemanagementManagementunitNoContent handles this case with default header values.

The management unit was successfully deleted
*/
type DeleteWorkforcemanagementManagementunitNoContent struct {
}

func (o *DeleteWorkforcemanagementManagementunitNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitNoContent ", 204)
}

func (o *DeleteWorkforcemanagementManagementunitNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkforcemanagementManagementunitBadRequest creates a DeleteWorkforcemanagementManagementunitBadRequest with default headers values
func NewDeleteWorkforcemanagementManagementunitBadRequest() *DeleteWorkforcemanagementManagementunitBadRequest {
	return &DeleteWorkforcemanagementManagementunitBadRequest{}
}

/*DeleteWorkforcemanagementManagementunitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementManagementunitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitUnauthorized creates a DeleteWorkforcemanagementManagementunitUnauthorized with default headers values
func NewDeleteWorkforcemanagementManagementunitUnauthorized() *DeleteWorkforcemanagementManagementunitUnauthorized {
	return &DeleteWorkforcemanagementManagementunitUnauthorized{}
}

/*DeleteWorkforcemanagementManagementunitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementManagementunitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitForbidden creates a DeleteWorkforcemanagementManagementunitForbidden with default headers values
func NewDeleteWorkforcemanagementManagementunitForbidden() *DeleteWorkforcemanagementManagementunitForbidden {
	return &DeleteWorkforcemanagementManagementunitForbidden{}
}

/*DeleteWorkforcemanagementManagementunitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementManagementunitForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitNotFound creates a DeleteWorkforcemanagementManagementunitNotFound with default headers values
func NewDeleteWorkforcemanagementManagementunitNotFound() *DeleteWorkforcemanagementManagementunitNotFound {
	return &DeleteWorkforcemanagementManagementunitNotFound{}
}

/*DeleteWorkforcemanagementManagementunitNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementManagementunitNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitRequestEntityTooLarge creates a DeleteWorkforcemanagementManagementunitRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementManagementunitRequestEntityTooLarge() *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge {
	return &DeleteWorkforcemanagementManagementunitRequestEntityTooLarge{}
}

/*DeleteWorkforcemanagementManagementunitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWorkforcemanagementManagementunitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitUnsupportedMediaType creates a DeleteWorkforcemanagementManagementunitUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementManagementunitUnsupportedMediaType() *DeleteWorkforcemanagementManagementunitUnsupportedMediaType {
	return &DeleteWorkforcemanagementManagementunitUnsupportedMediaType{}
}

/*DeleteWorkforcemanagementManagementunitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementManagementunitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitTooManyRequests creates a DeleteWorkforcemanagementManagementunitTooManyRequests with default headers values
func NewDeleteWorkforcemanagementManagementunitTooManyRequests() *DeleteWorkforcemanagementManagementunitTooManyRequests {
	return &DeleteWorkforcemanagementManagementunitTooManyRequests{}
}

/*DeleteWorkforcemanagementManagementunitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteWorkforcemanagementManagementunitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitInternalServerError creates a DeleteWorkforcemanagementManagementunitInternalServerError with default headers values
func NewDeleteWorkforcemanagementManagementunitInternalServerError() *DeleteWorkforcemanagementManagementunitInternalServerError {
	return &DeleteWorkforcemanagementManagementunitInternalServerError{}
}

/*DeleteWorkforcemanagementManagementunitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementManagementunitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitServiceUnavailable creates a DeleteWorkforcemanagementManagementunitServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementManagementunitServiceUnavailable() *DeleteWorkforcemanagementManagementunitServiceUnavailable {
	return &DeleteWorkforcemanagementManagementunitServiceUnavailable{}
}

/*DeleteWorkforcemanagementManagementunitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementManagementunitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementManagementunitGatewayTimeout creates a DeleteWorkforcemanagementManagementunitGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementManagementunitGatewayTimeout() *DeleteWorkforcemanagementManagementunitGatewayTimeout {
	return &DeleteWorkforcemanagementManagementunitGatewayTimeout{}
}

/*DeleteWorkforcemanagementManagementunitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementManagementunitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] deleteWorkforcemanagementManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementManagementunitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
