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

// GetFlowsDatatableRowReader is a Reader for the GetFlowsDatatableRow structure.
type GetFlowsDatatableRowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDatatableRowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDatatableRowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDatatableRowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDatatableRowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDatatableRowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDatatableRowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDatatableRowRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDatatableRowUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDatatableRowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDatatableRowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDatatableRowServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDatatableRowGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDatatableRowOK creates a GetFlowsDatatableRowOK with default headers values
func NewGetFlowsDatatableRowOK() *GetFlowsDatatableRowOK {
	return &GetFlowsDatatableRowOK{}
}

/*GetFlowsDatatableRowOK handles this case with default header values.

successful operation
*/
type GetFlowsDatatableRowOK struct {
	Payload map[string]interface{}
}

func (o *GetFlowsDatatableRowOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatableRowOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetFlowsDatatableRowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowBadRequest creates a GetFlowsDatatableRowBadRequest with default headers values
func NewGetFlowsDatatableRowBadRequest() *GetFlowsDatatableRowBadRequest {
	return &GetFlowsDatatableRowBadRequest{}
}

/*GetFlowsDatatableRowBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDatatableRowBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatableRowBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowUnauthorized creates a GetFlowsDatatableRowUnauthorized with default headers values
func NewGetFlowsDatatableRowUnauthorized() *GetFlowsDatatableRowUnauthorized {
	return &GetFlowsDatatableRowUnauthorized{}
}

/*GetFlowsDatatableRowUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDatatableRowUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatableRowUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowForbidden creates a GetFlowsDatatableRowForbidden with default headers values
func NewGetFlowsDatatableRowForbidden() *GetFlowsDatatableRowForbidden {
	return &GetFlowsDatatableRowForbidden{}
}

/*GetFlowsDatatableRowForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDatatableRowForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatableRowForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowNotFound creates a GetFlowsDatatableRowNotFound with default headers values
func NewGetFlowsDatatableRowNotFound() *GetFlowsDatatableRowNotFound {
	return &GetFlowsDatatableRowNotFound{}
}

/*GetFlowsDatatableRowNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsDatatableRowNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatableRowNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowRequestEntityTooLarge creates a GetFlowsDatatableRowRequestEntityTooLarge with default headers values
func NewGetFlowsDatatableRowRequestEntityTooLarge() *GetFlowsDatatableRowRequestEntityTooLarge {
	return &GetFlowsDatatableRowRequestEntityTooLarge{}
}

/*GetFlowsDatatableRowRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowsDatatableRowRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatableRowRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowUnsupportedMediaType creates a GetFlowsDatatableRowUnsupportedMediaType with default headers values
func NewGetFlowsDatatableRowUnsupportedMediaType() *GetFlowsDatatableRowUnsupportedMediaType {
	return &GetFlowsDatatableRowUnsupportedMediaType{}
}

/*GetFlowsDatatableRowUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDatatableRowUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatableRowUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowTooManyRequests creates a GetFlowsDatatableRowTooManyRequests with default headers values
func NewGetFlowsDatatableRowTooManyRequests() *GetFlowsDatatableRowTooManyRequests {
	return &GetFlowsDatatableRowTooManyRequests{}
}

/*GetFlowsDatatableRowTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetFlowsDatatableRowTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatableRowTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowInternalServerError creates a GetFlowsDatatableRowInternalServerError with default headers values
func NewGetFlowsDatatableRowInternalServerError() *GetFlowsDatatableRowInternalServerError {
	return &GetFlowsDatatableRowInternalServerError{}
}

/*GetFlowsDatatableRowInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDatatableRowInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatableRowInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowServiceUnavailable creates a GetFlowsDatatableRowServiceUnavailable with default headers values
func NewGetFlowsDatatableRowServiceUnavailable() *GetFlowsDatatableRowServiceUnavailable {
	return &GetFlowsDatatableRowServiceUnavailable{}
}

/*GetFlowsDatatableRowServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDatatableRowServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatableRowServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatableRowGatewayTimeout creates a GetFlowsDatatableRowGatewayTimeout with default headers values
func NewGetFlowsDatatableRowGatewayTimeout() *GetFlowsDatatableRowGatewayTimeout {
	return &GetFlowsDatatableRowGatewayTimeout{}
}

/*GetFlowsDatatableRowGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsDatatableRowGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatableRowGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/{datatableId}/rows/{rowId}][%d] getFlowsDatatableRowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatableRowGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatableRowGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
