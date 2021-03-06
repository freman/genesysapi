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

// DeleteFlowsDatatableRowReader is a Reader for the DeleteFlowsDatatableRow structure.
type DeleteFlowsDatatableRowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlowsDatatableRowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteFlowsDatatableRowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFlowsDatatableRowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFlowsDatatableRowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFlowsDatatableRowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFlowsDatatableRowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteFlowsDatatableRowRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteFlowsDatatableRowUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFlowsDatatableRowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFlowsDatatableRowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFlowsDatatableRowServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFlowsDatatableRowGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFlowsDatatableRowNoContent creates a DeleteFlowsDatatableRowNoContent with default headers values
func NewDeleteFlowsDatatableRowNoContent() *DeleteFlowsDatatableRowNoContent {
	return &DeleteFlowsDatatableRowNoContent{}
}

/*DeleteFlowsDatatableRowNoContent handles this case with default header values.

The row was deleted successfully
*/
type DeleteFlowsDatatableRowNoContent struct {
}

func (o *DeleteFlowsDatatableRowNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowNoContent ", 204)
}

func (o *DeleteFlowsDatatableRowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowsDatatableRowBadRequest creates a DeleteFlowsDatatableRowBadRequest with default headers values
func NewDeleteFlowsDatatableRowBadRequest() *DeleteFlowsDatatableRowBadRequest {
	return &DeleteFlowsDatatableRowBadRequest{}
}

/*DeleteFlowsDatatableRowBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteFlowsDatatableRowBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowsDatatableRowBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowUnauthorized creates a DeleteFlowsDatatableRowUnauthorized with default headers values
func NewDeleteFlowsDatatableRowUnauthorized() *DeleteFlowsDatatableRowUnauthorized {
	return &DeleteFlowsDatatableRowUnauthorized{}
}

/*DeleteFlowsDatatableRowUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteFlowsDatatableRowUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowsDatatableRowUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowForbidden creates a DeleteFlowsDatatableRowForbidden with default headers values
func NewDeleteFlowsDatatableRowForbidden() *DeleteFlowsDatatableRowForbidden {
	return &DeleteFlowsDatatableRowForbidden{}
}

/*DeleteFlowsDatatableRowForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteFlowsDatatableRowForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowsDatatableRowForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowNotFound creates a DeleteFlowsDatatableRowNotFound with default headers values
func NewDeleteFlowsDatatableRowNotFound() *DeleteFlowsDatatableRowNotFound {
	return &DeleteFlowsDatatableRowNotFound{}
}

/*DeleteFlowsDatatableRowNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteFlowsDatatableRowNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowsDatatableRowNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowRequestEntityTooLarge creates a DeleteFlowsDatatableRowRequestEntityTooLarge with default headers values
func NewDeleteFlowsDatatableRowRequestEntityTooLarge() *DeleteFlowsDatatableRowRequestEntityTooLarge {
	return &DeleteFlowsDatatableRowRequestEntityTooLarge{}
}

/*DeleteFlowsDatatableRowRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteFlowsDatatableRowRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowsDatatableRowRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowUnsupportedMediaType creates a DeleteFlowsDatatableRowUnsupportedMediaType with default headers values
func NewDeleteFlowsDatatableRowUnsupportedMediaType() *DeleteFlowsDatatableRowUnsupportedMediaType {
	return &DeleteFlowsDatatableRowUnsupportedMediaType{}
}

/*DeleteFlowsDatatableRowUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteFlowsDatatableRowUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowsDatatableRowUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowTooManyRequests creates a DeleteFlowsDatatableRowTooManyRequests with default headers values
func NewDeleteFlowsDatatableRowTooManyRequests() *DeleteFlowsDatatableRowTooManyRequests {
	return &DeleteFlowsDatatableRowTooManyRequests{}
}

/*DeleteFlowsDatatableRowTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteFlowsDatatableRowTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowsDatatableRowTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowInternalServerError creates a DeleteFlowsDatatableRowInternalServerError with default headers values
func NewDeleteFlowsDatatableRowInternalServerError() *DeleteFlowsDatatableRowInternalServerError {
	return &DeleteFlowsDatatableRowInternalServerError{}
}

/*DeleteFlowsDatatableRowInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteFlowsDatatableRowInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowsDatatableRowInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowServiceUnavailable creates a DeleteFlowsDatatableRowServiceUnavailable with default headers values
func NewDeleteFlowsDatatableRowServiceUnavailable() *DeleteFlowsDatatableRowServiceUnavailable {
	return &DeleteFlowsDatatableRowServiceUnavailable{}
}

/*DeleteFlowsDatatableRowServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteFlowsDatatableRowServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowsDatatableRowServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsDatatableRowGatewayTimeout creates a DeleteFlowsDatatableRowGatewayTimeout with default headers values
func NewDeleteFlowsDatatableRowGatewayTimeout() *DeleteFlowsDatatableRowGatewayTimeout {
	return &DeleteFlowsDatatableRowGatewayTimeout{}
}

/*DeleteFlowsDatatableRowGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteFlowsDatatableRowGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsDatatableRowGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] deleteFlowsDatatableRowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowsDatatableRowGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsDatatableRowGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
