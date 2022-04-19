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

// GetFlowsDatatablesReader is a Reader for the GetFlowsDatatables structure.
type GetFlowsDatatablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDatatablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDatatablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDatatablesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDatatablesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDatatablesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDatatablesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsDatatablesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDatatablesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDatatablesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDatatablesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDatatablesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDatatablesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDatatablesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDatatablesOK creates a GetFlowsDatatablesOK with default headers values
func NewGetFlowsDatatablesOK() *GetFlowsDatatablesOK {
	return &GetFlowsDatatablesOK{}
}

/*GetFlowsDatatablesOK handles this case with default header values.

successful operation
*/
type GetFlowsDatatablesOK struct {
	Payload *models.DataTablesDomainEntityListing
}

func (o *GetFlowsDatatablesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatablesOK) GetPayload() *models.DataTablesDomainEntityListing {
	return o.Payload
}

func (o *GetFlowsDatatablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTablesDomainEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesBadRequest creates a GetFlowsDatatablesBadRequest with default headers values
func NewGetFlowsDatatablesBadRequest() *GetFlowsDatatablesBadRequest {
	return &GetFlowsDatatablesBadRequest{}
}

/*GetFlowsDatatablesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDatatablesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatablesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesUnauthorized creates a GetFlowsDatatablesUnauthorized with default headers values
func NewGetFlowsDatatablesUnauthorized() *GetFlowsDatatablesUnauthorized {
	return &GetFlowsDatatablesUnauthorized{}
}

/*GetFlowsDatatablesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDatatablesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatablesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesForbidden creates a GetFlowsDatatablesForbidden with default headers values
func NewGetFlowsDatatablesForbidden() *GetFlowsDatatablesForbidden {
	return &GetFlowsDatatablesForbidden{}
}

/*GetFlowsDatatablesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDatatablesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatablesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesNotFound creates a GetFlowsDatatablesNotFound with default headers values
func NewGetFlowsDatatablesNotFound() *GetFlowsDatatablesNotFound {
	return &GetFlowsDatatablesNotFound{}
}

/*GetFlowsDatatablesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsDatatablesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatablesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesRequestTimeout creates a GetFlowsDatatablesRequestTimeout with default headers values
func NewGetFlowsDatatablesRequestTimeout() *GetFlowsDatatablesRequestTimeout {
	return &GetFlowsDatatablesRequestTimeout{}
}

/*GetFlowsDatatablesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsDatatablesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatablesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesRequestEntityTooLarge creates a GetFlowsDatatablesRequestEntityTooLarge with default headers values
func NewGetFlowsDatatablesRequestEntityTooLarge() *GetFlowsDatatablesRequestEntityTooLarge {
	return &GetFlowsDatatablesRequestEntityTooLarge{}
}

/*GetFlowsDatatablesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetFlowsDatatablesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatablesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesUnsupportedMediaType creates a GetFlowsDatatablesUnsupportedMediaType with default headers values
func NewGetFlowsDatatablesUnsupportedMediaType() *GetFlowsDatatablesUnsupportedMediaType {
	return &GetFlowsDatatablesUnsupportedMediaType{}
}

/*GetFlowsDatatablesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDatatablesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatablesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesTooManyRequests creates a GetFlowsDatatablesTooManyRequests with default headers values
func NewGetFlowsDatatablesTooManyRequests() *GetFlowsDatatablesTooManyRequests {
	return &GetFlowsDatatablesTooManyRequests{}
}

/*GetFlowsDatatablesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsDatatablesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatablesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesInternalServerError creates a GetFlowsDatatablesInternalServerError with default headers values
func NewGetFlowsDatatablesInternalServerError() *GetFlowsDatatablesInternalServerError {
	return &GetFlowsDatatablesInternalServerError{}
}

/*GetFlowsDatatablesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDatatablesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatablesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesServiceUnavailable creates a GetFlowsDatatablesServiceUnavailable with default headers values
func NewGetFlowsDatatablesServiceUnavailable() *GetFlowsDatatablesServiceUnavailable {
	return &GetFlowsDatatablesServiceUnavailable{}
}

/*GetFlowsDatatablesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDatatablesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatablesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesGatewayTimeout creates a GetFlowsDatatablesGatewayTimeout with default headers values
func NewGetFlowsDatatablesGatewayTimeout() *GetFlowsDatatablesGatewayTimeout {
	return &GetFlowsDatatablesGatewayTimeout{}
}

/*GetFlowsDatatablesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsDatatablesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables][%d] getFlowsDatatablesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatablesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
