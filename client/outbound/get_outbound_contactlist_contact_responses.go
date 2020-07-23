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

// GetOutboundContactlistContactReader is a Reader for the GetOutboundContactlistContact structure.
type GetOutboundContactlistContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundContactlistContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundContactlistContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundContactlistContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundContactlistContactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundContactlistContactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundContactlistContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundContactlistContactRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundContactlistContactUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundContactlistContactTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundContactlistContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundContactlistContactServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundContactlistContactGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundContactlistContactOK creates a GetOutboundContactlistContactOK with default headers values
func NewGetOutboundContactlistContactOK() *GetOutboundContactlistContactOK {
	return &GetOutboundContactlistContactOK{}
}

/*GetOutboundContactlistContactOK handles this case with default header values.

successful operation
*/
type GetOutboundContactlistContactOK struct {
	Payload *models.DialerContact
}

func (o *GetOutboundContactlistContactOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistContactOK) GetPayload() *models.DialerContact {
	return o.Payload
}

func (o *GetOutboundContactlistContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DialerContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactBadRequest creates a GetOutboundContactlistContactBadRequest with default headers values
func NewGetOutboundContactlistContactBadRequest() *GetOutboundContactlistContactBadRequest {
	return &GetOutboundContactlistContactBadRequest{}
}

/*GetOutboundContactlistContactBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundContactlistContactBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistContactBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactUnauthorized creates a GetOutboundContactlistContactUnauthorized with default headers values
func NewGetOutboundContactlistContactUnauthorized() *GetOutboundContactlistContactUnauthorized {
	return &GetOutboundContactlistContactUnauthorized{}
}

/*GetOutboundContactlistContactUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundContactlistContactUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistContactUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactForbidden creates a GetOutboundContactlistContactForbidden with default headers values
func NewGetOutboundContactlistContactForbidden() *GetOutboundContactlistContactForbidden {
	return &GetOutboundContactlistContactForbidden{}
}

/*GetOutboundContactlistContactForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundContactlistContactForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistContactForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactNotFound creates a GetOutboundContactlistContactNotFound with default headers values
func NewGetOutboundContactlistContactNotFound() *GetOutboundContactlistContactNotFound {
	return &GetOutboundContactlistContactNotFound{}
}

/*GetOutboundContactlistContactNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundContactlistContactNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistContactNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactRequestEntityTooLarge creates a GetOutboundContactlistContactRequestEntityTooLarge with default headers values
func NewGetOutboundContactlistContactRequestEntityTooLarge() *GetOutboundContactlistContactRequestEntityTooLarge {
	return &GetOutboundContactlistContactRequestEntityTooLarge{}
}

/*GetOutboundContactlistContactRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundContactlistContactRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistContactRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactUnsupportedMediaType creates a GetOutboundContactlistContactUnsupportedMediaType with default headers values
func NewGetOutboundContactlistContactUnsupportedMediaType() *GetOutboundContactlistContactUnsupportedMediaType {
	return &GetOutboundContactlistContactUnsupportedMediaType{}
}

/*GetOutboundContactlistContactUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundContactlistContactUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistContactUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactTooManyRequests creates a GetOutboundContactlistContactTooManyRequests with default headers values
func NewGetOutboundContactlistContactTooManyRequests() *GetOutboundContactlistContactTooManyRequests {
	return &GetOutboundContactlistContactTooManyRequests{}
}

/*GetOutboundContactlistContactTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundContactlistContactTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistContactTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactInternalServerError creates a GetOutboundContactlistContactInternalServerError with default headers values
func NewGetOutboundContactlistContactInternalServerError() *GetOutboundContactlistContactInternalServerError {
	return &GetOutboundContactlistContactInternalServerError{}
}

/*GetOutboundContactlistContactInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundContactlistContactInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistContactInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactServiceUnavailable creates a GetOutboundContactlistContactServiceUnavailable with default headers values
func NewGetOutboundContactlistContactServiceUnavailable() *GetOutboundContactlistContactServiceUnavailable {
	return &GetOutboundContactlistContactServiceUnavailable{}
}

/*GetOutboundContactlistContactServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundContactlistContactServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistContactServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistContactGatewayTimeout creates a GetOutboundContactlistContactGatewayTimeout with default headers values
func NewGetOutboundContactlistContactGatewayTimeout() *GetOutboundContactlistContactGatewayTimeout {
	return &GetOutboundContactlistContactGatewayTimeout{}
}

/*GetOutboundContactlistContactGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundContactlistContactGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistContactGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/contacts/{contactId}][%d] getOutboundContactlistContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistContactGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistContactGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
