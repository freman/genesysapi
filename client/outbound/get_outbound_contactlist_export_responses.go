// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundContactlistExportReader is a Reader for the GetOutboundContactlistExport structure.
type GetOutboundContactlistExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundContactlistExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundContactlistExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundContactlistExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundContactlistExportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundContactlistExportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundContactlistExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundContactlistExportRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundContactlistExportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundContactlistExportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundContactlistExportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundContactlistExportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundContactlistExportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundContactlistExportOK creates a GetOutboundContactlistExportOK with default headers values
func NewGetOutboundContactlistExportOK() *GetOutboundContactlistExportOK {
	return &GetOutboundContactlistExportOK{}
}

/*GetOutboundContactlistExportOK handles this case with default header values.

successful operation
*/
type GetOutboundContactlistExportOK struct {
	Payload *models.ExportURI
}

func (o *GetOutboundContactlistExportOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistExportOK) GetPayload() *models.ExportURI {
	return o.Payload
}

func (o *GetOutboundContactlistExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportURI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportBadRequest creates a GetOutboundContactlistExportBadRequest with default headers values
func NewGetOutboundContactlistExportBadRequest() *GetOutboundContactlistExportBadRequest {
	return &GetOutboundContactlistExportBadRequest{}
}

/*GetOutboundContactlistExportBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundContactlistExportBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistExportBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportUnauthorized creates a GetOutboundContactlistExportUnauthorized with default headers values
func NewGetOutboundContactlistExportUnauthorized() *GetOutboundContactlistExportUnauthorized {
	return &GetOutboundContactlistExportUnauthorized{}
}

/*GetOutboundContactlistExportUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundContactlistExportUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistExportUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportForbidden creates a GetOutboundContactlistExportForbidden with default headers values
func NewGetOutboundContactlistExportForbidden() *GetOutboundContactlistExportForbidden {
	return &GetOutboundContactlistExportForbidden{}
}

/*GetOutboundContactlistExportForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundContactlistExportForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistExportForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportNotFound creates a GetOutboundContactlistExportNotFound with default headers values
func NewGetOutboundContactlistExportNotFound() *GetOutboundContactlistExportNotFound {
	return &GetOutboundContactlistExportNotFound{}
}

/*GetOutboundContactlistExportNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundContactlistExportNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistExportNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportRequestEntityTooLarge creates a GetOutboundContactlistExportRequestEntityTooLarge with default headers values
func NewGetOutboundContactlistExportRequestEntityTooLarge() *GetOutboundContactlistExportRequestEntityTooLarge {
	return &GetOutboundContactlistExportRequestEntityTooLarge{}
}

/*GetOutboundContactlistExportRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundContactlistExportRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistExportRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportUnsupportedMediaType creates a GetOutboundContactlistExportUnsupportedMediaType with default headers values
func NewGetOutboundContactlistExportUnsupportedMediaType() *GetOutboundContactlistExportUnsupportedMediaType {
	return &GetOutboundContactlistExportUnsupportedMediaType{}
}

/*GetOutboundContactlistExportUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundContactlistExportUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistExportUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportTooManyRequests creates a GetOutboundContactlistExportTooManyRequests with default headers values
func NewGetOutboundContactlistExportTooManyRequests() *GetOutboundContactlistExportTooManyRequests {
	return &GetOutboundContactlistExportTooManyRequests{}
}

/*GetOutboundContactlistExportTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundContactlistExportTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistExportTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportInternalServerError creates a GetOutboundContactlistExportInternalServerError with default headers values
func NewGetOutboundContactlistExportInternalServerError() *GetOutboundContactlistExportInternalServerError {
	return &GetOutboundContactlistExportInternalServerError{}
}

/*GetOutboundContactlistExportInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundContactlistExportInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistExportInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportServiceUnavailable creates a GetOutboundContactlistExportServiceUnavailable with default headers values
func NewGetOutboundContactlistExportServiceUnavailable() *GetOutboundContactlistExportServiceUnavailable {
	return &GetOutboundContactlistExportServiceUnavailable{}
}

/*GetOutboundContactlistExportServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundContactlistExportServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistExportServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistExportGatewayTimeout creates a GetOutboundContactlistExportGatewayTimeout with default headers values
func NewGetOutboundContactlistExportGatewayTimeout() *GetOutboundContactlistExportGatewayTimeout {
	return &GetOutboundContactlistExportGatewayTimeout{}
}

/*GetOutboundContactlistExportGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundContactlistExportGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistExportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/export][%d] getOutboundContactlistExportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistExportGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistExportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}