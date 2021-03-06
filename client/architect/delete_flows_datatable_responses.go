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

// DeleteFlowsDatatableReader is a Reader for the DeleteFlowsDatatable structure.
type DeleteFlowsDatatableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlowsDatatableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFlowsDatatableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFlowsDatatableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFlowsDatatableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFlowsDatatableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFlowsDatatableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteFlowsDatatableConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteFlowsDatatableRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteFlowsDatatableUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFlowsDatatableTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFlowsDatatableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFlowsDatatableServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFlowsDatatableGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFlowsDatatableNoContent creates a DeleteFlowsDatatableNoContent with default headers values
func NewDeleteFlowsDatatableNoContent() *DeleteFlowsDatatableNoContent {
	return &DeleteFlowsDatatableNoContent{}
}

/*DeleteFlowsDatatableNoContent handles this case with default header values.

The datatable was deleted successfully
*/
type DeleteFlowsDatatableNoContent struct {
}

func (o *DeleteFlowsDatatableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableNoContent ", 204)
}

func (o *DeleteFlowsDatatableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowsDatatableBadRequest creates a DeleteFlowsDatatableBadRequest with default headers values
func NewDeleteFlowsDatatableBadRequest() *DeleteFlowsDatatableBadRequest {
	return &DeleteFlowsDatatableBadRequest{}
}

/*DeleteFlowsDatatableBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteFlowsDatatableBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowsDatatableBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableUnauthorized creates a DeleteFlowsDatatableUnauthorized with default headers values
func NewDeleteFlowsDatatableUnauthorized() *DeleteFlowsDatatableUnauthorized {
	return &DeleteFlowsDatatableUnauthorized{}
}

/*DeleteFlowsDatatableUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteFlowsDatatableUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowsDatatableUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableForbidden creates a DeleteFlowsDatatableForbidden with default headers values
func NewDeleteFlowsDatatableForbidden() *DeleteFlowsDatatableForbidden {
	return &DeleteFlowsDatatableForbidden{}
}

/*DeleteFlowsDatatableForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteFlowsDatatableForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowsDatatableForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableNotFound creates a DeleteFlowsDatatableNotFound with default headers values
func NewDeleteFlowsDatatableNotFound() *DeleteFlowsDatatableNotFound {
	return &DeleteFlowsDatatableNotFound{}
}

/*DeleteFlowsDatatableNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteFlowsDatatableNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowsDatatableNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableConflict creates a DeleteFlowsDatatableConflict with default headers values
func NewDeleteFlowsDatatableConflict() *DeleteFlowsDatatableConflict {
	return &DeleteFlowsDatatableConflict{}
}

/*DeleteFlowsDatatableConflict handles this case with default header values.

Conflict
*/
type DeleteFlowsDatatableConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableConflict  %+v", 409, o.Payload)
}

func (o *DeleteFlowsDatatableConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRequestEntityTooLarge creates a DeleteFlowsDatatableRequestEntityTooLarge with default headers values
func NewDeleteFlowsDatatableRequestEntityTooLarge() *DeleteFlowsDatatableRequestEntityTooLarge {
	return &DeleteFlowsDatatableRequestEntityTooLarge{}
}

/*DeleteFlowsDatatableRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteFlowsDatatableRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowsDatatableRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableUnsupportedMediaType creates a DeleteFlowsDatatableUnsupportedMediaType with default headers values
func NewDeleteFlowsDatatableUnsupportedMediaType() *DeleteFlowsDatatableUnsupportedMediaType {
	return &DeleteFlowsDatatableUnsupportedMediaType{}
}

/*DeleteFlowsDatatableUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteFlowsDatatableUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowsDatatableUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableTooManyRequests creates a DeleteFlowsDatatableTooManyRequests with default headers values
func NewDeleteFlowsDatatableTooManyRequests() *DeleteFlowsDatatableTooManyRequests {
	return &DeleteFlowsDatatableTooManyRequests{}
}

/*DeleteFlowsDatatableTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteFlowsDatatableTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowsDatatableTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableInternalServerError creates a DeleteFlowsDatatableInternalServerError with default headers values
func NewDeleteFlowsDatatableInternalServerError() *DeleteFlowsDatatableInternalServerError {
	return &DeleteFlowsDatatableInternalServerError{}
}

/*DeleteFlowsDatatableInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteFlowsDatatableInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowsDatatableInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableServiceUnavailable creates a DeleteFlowsDatatableServiceUnavailable with default headers values
func NewDeleteFlowsDatatableServiceUnavailable() *DeleteFlowsDatatableServiceUnavailable {
	return &DeleteFlowsDatatableServiceUnavailable{}
}

/*DeleteFlowsDatatableServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteFlowsDatatableServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowsDatatableServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableGatewayTimeout creates a DeleteFlowsDatatableGatewayTimeout with default headers values
func NewDeleteFlowsDatatableGatewayTimeout() *DeleteFlowsDatatableGatewayTimeout {
	return &DeleteFlowsDatatableGatewayTimeout{}
}

/*DeleteFlowsDatatableGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteFlowsDatatableGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}][%d] deleteFlowsDatatableGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowsDatatableGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
