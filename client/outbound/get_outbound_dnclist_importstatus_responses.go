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

// GetOutboundDnclistImportstatusReader is a Reader for the GetOutboundDnclistImportstatus structure.
type GetOutboundDnclistImportstatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundDnclistImportstatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundDnclistImportstatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundDnclistImportstatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundDnclistImportstatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundDnclistImportstatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundDnclistImportstatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundDnclistImportstatusRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundDnclistImportstatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundDnclistImportstatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundDnclistImportstatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundDnclistImportstatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundDnclistImportstatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundDnclistImportstatusOK creates a GetOutboundDnclistImportstatusOK with default headers values
func NewGetOutboundDnclistImportstatusOK() *GetOutboundDnclistImportstatusOK {
	return &GetOutboundDnclistImportstatusOK{}
}

/*GetOutboundDnclistImportstatusOK handles this case with default header values.

successful operation
*/
type GetOutboundDnclistImportstatusOK struct {
	Payload *models.ImportStatus
}

func (o *GetOutboundDnclistImportstatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusOK  %+v", 200, o.Payload)
}

func (o *GetOutboundDnclistImportstatusOK) GetPayload() *models.ImportStatus {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusBadRequest creates a GetOutboundDnclistImportstatusBadRequest with default headers values
func NewGetOutboundDnclistImportstatusBadRequest() *GetOutboundDnclistImportstatusBadRequest {
	return &GetOutboundDnclistImportstatusBadRequest{}
}

/*GetOutboundDnclistImportstatusBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundDnclistImportstatusBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundDnclistImportstatusBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusUnauthorized creates a GetOutboundDnclistImportstatusUnauthorized with default headers values
func NewGetOutboundDnclistImportstatusUnauthorized() *GetOutboundDnclistImportstatusUnauthorized {
	return &GetOutboundDnclistImportstatusUnauthorized{}
}

/*GetOutboundDnclistImportstatusUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundDnclistImportstatusUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundDnclistImportstatusUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusForbidden creates a GetOutboundDnclistImportstatusForbidden with default headers values
func NewGetOutboundDnclistImportstatusForbidden() *GetOutboundDnclistImportstatusForbidden {
	return &GetOutboundDnclistImportstatusForbidden{}
}

/*GetOutboundDnclistImportstatusForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundDnclistImportstatusForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundDnclistImportstatusForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusNotFound creates a GetOutboundDnclistImportstatusNotFound with default headers values
func NewGetOutboundDnclistImportstatusNotFound() *GetOutboundDnclistImportstatusNotFound {
	return &GetOutboundDnclistImportstatusNotFound{}
}

/*GetOutboundDnclistImportstatusNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundDnclistImportstatusNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundDnclistImportstatusNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusRequestEntityTooLarge creates a GetOutboundDnclistImportstatusRequestEntityTooLarge with default headers values
func NewGetOutboundDnclistImportstatusRequestEntityTooLarge() *GetOutboundDnclistImportstatusRequestEntityTooLarge {
	return &GetOutboundDnclistImportstatusRequestEntityTooLarge{}
}

/*GetOutboundDnclistImportstatusRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundDnclistImportstatusRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusUnsupportedMediaType creates a GetOutboundDnclistImportstatusUnsupportedMediaType with default headers values
func NewGetOutboundDnclistImportstatusUnsupportedMediaType() *GetOutboundDnclistImportstatusUnsupportedMediaType {
	return &GetOutboundDnclistImportstatusUnsupportedMediaType{}
}

/*GetOutboundDnclistImportstatusUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundDnclistImportstatusUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusTooManyRequests creates a GetOutboundDnclistImportstatusTooManyRequests with default headers values
func NewGetOutboundDnclistImportstatusTooManyRequests() *GetOutboundDnclistImportstatusTooManyRequests {
	return &GetOutboundDnclistImportstatusTooManyRequests{}
}

/*GetOutboundDnclistImportstatusTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundDnclistImportstatusTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundDnclistImportstatusTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusInternalServerError creates a GetOutboundDnclistImportstatusInternalServerError with default headers values
func NewGetOutboundDnclistImportstatusInternalServerError() *GetOutboundDnclistImportstatusInternalServerError {
	return &GetOutboundDnclistImportstatusInternalServerError{}
}

/*GetOutboundDnclistImportstatusInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundDnclistImportstatusInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundDnclistImportstatusInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusServiceUnavailable creates a GetOutboundDnclistImportstatusServiceUnavailable with default headers values
func NewGetOutboundDnclistImportstatusServiceUnavailable() *GetOutboundDnclistImportstatusServiceUnavailable {
	return &GetOutboundDnclistImportstatusServiceUnavailable{}
}

/*GetOutboundDnclistImportstatusServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundDnclistImportstatusServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundDnclistImportstatusServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistImportstatusGatewayTimeout creates a GetOutboundDnclistImportstatusGatewayTimeout with default headers values
func NewGetOutboundDnclistImportstatusGatewayTimeout() *GetOutboundDnclistImportstatusGatewayTimeout {
	return &GetOutboundDnclistImportstatusGatewayTimeout{}
}

/*GetOutboundDnclistImportstatusGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundDnclistImportstatusGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistImportstatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/{dncListId}/importstatus][%d] getOutboundDnclistImportstatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundDnclistImportstatusGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistImportstatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
