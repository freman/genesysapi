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

// GetFlowsDatatableReader is a Reader for the GetFlowsDatatable structure.
type GetFlowsDatatableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDatatableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDatatableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDatatableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDatatableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDatatableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDatatableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDatatableRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDatatableUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDatatableTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDatatableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDatatableServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDatatableGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDatatableOK creates a GetFlowsDatatableOK with default headers values
func NewGetFlowsDatatableOK() *GetFlowsDatatableOK {
	return &GetFlowsDatatableOK{}
}

/*GetFlowsDatatableOK handles this case with default header values.

successful operation
*/
type GetFlowsDatatableOK struct {
	Payload *models.DataTable
}

func (o *GetFlowsDatatableOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatableOK) GetPayload() *models.DataTable {
	return o.Payload
}

func (o *GetFlowsDatatableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableBadRequest creates a GetFlowsDatatableBadRequest with default headers values
func NewGetFlowsDatatableBadRequest() *GetFlowsDatatableBadRequest {
	return &GetFlowsDatatableBadRequest{}
}

/*GetFlowsDatatableBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDatatableBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatableBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableUnauthorized creates a GetFlowsDatatableUnauthorized with default headers values
func NewGetFlowsDatatableUnauthorized() *GetFlowsDatatableUnauthorized {
	return &GetFlowsDatatableUnauthorized{}
}

/*GetFlowsDatatableUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDatatableUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatableUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableForbidden creates a GetFlowsDatatableForbidden with default headers values
func NewGetFlowsDatatableForbidden() *GetFlowsDatatableForbidden {
	return &GetFlowsDatatableForbidden{}
}

/*GetFlowsDatatableForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDatatableForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatableForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableNotFound creates a GetFlowsDatatableNotFound with default headers values
func NewGetFlowsDatatableNotFound() *GetFlowsDatatableNotFound {
	return &GetFlowsDatatableNotFound{}
}

/*GetFlowsDatatableNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsDatatableNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatableNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRequestEntityTooLarge creates a GetFlowsDatatableRequestEntityTooLarge with default headers values
func NewGetFlowsDatatableRequestEntityTooLarge() *GetFlowsDatatableRequestEntityTooLarge {
	return &GetFlowsDatatableRequestEntityTooLarge{}
}

/*GetFlowsDatatableRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowsDatatableRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatableRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableUnsupportedMediaType creates a GetFlowsDatatableUnsupportedMediaType with default headers values
func NewGetFlowsDatatableUnsupportedMediaType() *GetFlowsDatatableUnsupportedMediaType {
	return &GetFlowsDatatableUnsupportedMediaType{}
}

/*GetFlowsDatatableUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDatatableUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatableUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableTooManyRequests creates a GetFlowsDatatableTooManyRequests with default headers values
func NewGetFlowsDatatableTooManyRequests() *GetFlowsDatatableTooManyRequests {
	return &GetFlowsDatatableTooManyRequests{}
}

/*GetFlowsDatatableTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetFlowsDatatableTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatableTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableInternalServerError creates a GetFlowsDatatableInternalServerError with default headers values
func NewGetFlowsDatatableInternalServerError() *GetFlowsDatatableInternalServerError {
	return &GetFlowsDatatableInternalServerError{}
}

/*GetFlowsDatatableInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDatatableInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatableInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableServiceUnavailable creates a GetFlowsDatatableServiceUnavailable with default headers values
func NewGetFlowsDatatableServiceUnavailable() *GetFlowsDatatableServiceUnavailable {
	return &GetFlowsDatatableServiceUnavailable{}
}

/*GetFlowsDatatableServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDatatableServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatableServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableGatewayTimeout creates a GetFlowsDatatableGatewayTimeout with default headers values
func NewGetFlowsDatatableGatewayTimeout() *GetFlowsDatatableGatewayTimeout {
	return &GetFlowsDatatableGatewayTimeout{}
}

/*GetFlowsDatatableGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsDatatableGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}][%d] getFlowsDatatableGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatableGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
