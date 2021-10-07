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

// GetFlowsDatatablesDivisionviewReader is a Reader for the GetFlowsDatatablesDivisionview structure.
type GetFlowsDatatablesDivisionviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsDatatablesDivisionviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsDatatablesDivisionviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsDatatablesDivisionviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsDatatablesDivisionviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsDatatablesDivisionviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsDatatablesDivisionviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowsDatatablesDivisionviewRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsDatatablesDivisionviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsDatatablesDivisionviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsDatatablesDivisionviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsDatatablesDivisionviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsDatatablesDivisionviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsDatatablesDivisionviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsDatatablesDivisionviewOK creates a GetFlowsDatatablesDivisionviewOK with default headers values
func NewGetFlowsDatatablesDivisionviewOK() *GetFlowsDatatablesDivisionviewOK {
	return &GetFlowsDatatablesDivisionviewOK{}
}

/*GetFlowsDatatablesDivisionviewOK handles this case with default header values.

successful operation
*/
type GetFlowsDatatablesDivisionviewOK struct {
	Payload *models.DataTable
}

func (o *GetFlowsDatatablesDivisionviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewOK  %+v", 200, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewOK) GetPayload() *models.DataTable {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataTable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewBadRequest creates a GetFlowsDatatablesDivisionviewBadRequest with default headers values
func NewGetFlowsDatatablesDivisionviewBadRequest() *GetFlowsDatatablesDivisionviewBadRequest {
	return &GetFlowsDatatablesDivisionviewBadRequest{}
}

/*GetFlowsDatatablesDivisionviewBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsDatatablesDivisionviewBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewUnauthorized creates a GetFlowsDatatablesDivisionviewUnauthorized with default headers values
func NewGetFlowsDatatablesDivisionviewUnauthorized() *GetFlowsDatatablesDivisionviewUnauthorized {
	return &GetFlowsDatatablesDivisionviewUnauthorized{}
}

/*GetFlowsDatatablesDivisionviewUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsDatatablesDivisionviewUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewForbidden creates a GetFlowsDatatablesDivisionviewForbidden with default headers values
func NewGetFlowsDatatablesDivisionviewForbidden() *GetFlowsDatatablesDivisionviewForbidden {
	return &GetFlowsDatatablesDivisionviewForbidden{}
}

/*GetFlowsDatatablesDivisionviewForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsDatatablesDivisionviewForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewNotFound creates a GetFlowsDatatablesDivisionviewNotFound with default headers values
func NewGetFlowsDatatablesDivisionviewNotFound() *GetFlowsDatatablesDivisionviewNotFound {
	return &GetFlowsDatatablesDivisionviewNotFound{}
}

/*GetFlowsDatatablesDivisionviewNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsDatatablesDivisionviewNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewRequestTimeout creates a GetFlowsDatatablesDivisionviewRequestTimeout with default headers values
func NewGetFlowsDatatablesDivisionviewRequestTimeout() *GetFlowsDatatablesDivisionviewRequestTimeout {
	return &GetFlowsDatatablesDivisionviewRequestTimeout{}
}

/*GetFlowsDatatablesDivisionviewRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowsDatatablesDivisionviewRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewRequestEntityTooLarge creates a GetFlowsDatatablesDivisionviewRequestEntityTooLarge with default headers values
func NewGetFlowsDatatablesDivisionviewRequestEntityTooLarge() *GetFlowsDatatablesDivisionviewRequestEntityTooLarge {
	return &GetFlowsDatatablesDivisionviewRequestEntityTooLarge{}
}

/*GetFlowsDatatablesDivisionviewRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowsDatatablesDivisionviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewUnsupportedMediaType creates a GetFlowsDatatablesDivisionviewUnsupportedMediaType with default headers values
func NewGetFlowsDatatablesDivisionviewUnsupportedMediaType() *GetFlowsDatatablesDivisionviewUnsupportedMediaType {
	return &GetFlowsDatatablesDivisionviewUnsupportedMediaType{}
}

/*GetFlowsDatatablesDivisionviewUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsDatatablesDivisionviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewTooManyRequests creates a GetFlowsDatatablesDivisionviewTooManyRequests with default headers values
func NewGetFlowsDatatablesDivisionviewTooManyRequests() *GetFlowsDatatablesDivisionviewTooManyRequests {
	return &GetFlowsDatatablesDivisionviewTooManyRequests{}
}

/*GetFlowsDatatablesDivisionviewTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowsDatatablesDivisionviewTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewInternalServerError creates a GetFlowsDatatablesDivisionviewInternalServerError with default headers values
func NewGetFlowsDatatablesDivisionviewInternalServerError() *GetFlowsDatatablesDivisionviewInternalServerError {
	return &GetFlowsDatatablesDivisionviewInternalServerError{}
}

/*GetFlowsDatatablesDivisionviewInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsDatatablesDivisionviewInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewServiceUnavailable creates a GetFlowsDatatablesDivisionviewServiceUnavailable with default headers values
func NewGetFlowsDatatablesDivisionviewServiceUnavailable() *GetFlowsDatatablesDivisionviewServiceUnavailable {
	return &GetFlowsDatatablesDivisionviewServiceUnavailable{}
}

/*GetFlowsDatatablesDivisionviewServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsDatatablesDivisionviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsDatatablesDivisionviewGatewayTimeout creates a GetFlowsDatatablesDivisionviewGatewayTimeout with default headers values
func NewGetFlowsDatatablesDivisionviewGatewayTimeout() *GetFlowsDatatablesDivisionviewGatewayTimeout {
	return &GetFlowsDatatablesDivisionviewGatewayTimeout{}
}

/*GetFlowsDatatablesDivisionviewGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsDatatablesDivisionviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsDatatablesDivisionviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/datatables/divisionviews/{datatableId}][%d] getFlowsDatatablesDivisionviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsDatatablesDivisionviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsDatatablesDivisionviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
