// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetExternalcontactsContactReader is a Reader for the GetExternalcontactsContact structure.
type GetExternalcontactsContactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsContactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsContactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsContactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsContactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsContactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsContactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsContactRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsContactRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsContactUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsContactTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsContactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsContactServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsContactGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsContactOK creates a GetExternalcontactsContactOK with default headers values
func NewGetExternalcontactsContactOK() *GetExternalcontactsContactOK {
	return &GetExternalcontactsContactOK{}
}

/*GetExternalcontactsContactOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsContactOK struct {
	Payload *models.ExternalContact
}

func (o *GetExternalcontactsContactOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsContactOK) GetPayload() *models.ExternalContact {
	return o.Payload
}

func (o *GetExternalcontactsContactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactBadRequest creates a GetExternalcontactsContactBadRequest with default headers values
func NewGetExternalcontactsContactBadRequest() *GetExternalcontactsContactBadRequest {
	return &GetExternalcontactsContactBadRequest{}
}

/*GetExternalcontactsContactBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsContactBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsContactBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactUnauthorized creates a GetExternalcontactsContactUnauthorized with default headers values
func NewGetExternalcontactsContactUnauthorized() *GetExternalcontactsContactUnauthorized {
	return &GetExternalcontactsContactUnauthorized{}
}

/*GetExternalcontactsContactUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsContactUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsContactUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactForbidden creates a GetExternalcontactsContactForbidden with default headers values
func NewGetExternalcontactsContactForbidden() *GetExternalcontactsContactForbidden {
	return &GetExternalcontactsContactForbidden{}
}

/*GetExternalcontactsContactForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsContactForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsContactForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactNotFound creates a GetExternalcontactsContactNotFound with default headers values
func NewGetExternalcontactsContactNotFound() *GetExternalcontactsContactNotFound {
	return &GetExternalcontactsContactNotFound{}
}

/*GetExternalcontactsContactNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsContactNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsContactNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactRequestTimeout creates a GetExternalcontactsContactRequestTimeout with default headers values
func NewGetExternalcontactsContactRequestTimeout() *GetExternalcontactsContactRequestTimeout {
	return &GetExternalcontactsContactRequestTimeout{}
}

/*GetExternalcontactsContactRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsContactRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsContactRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactRequestEntityTooLarge creates a GetExternalcontactsContactRequestEntityTooLarge with default headers values
func NewGetExternalcontactsContactRequestEntityTooLarge() *GetExternalcontactsContactRequestEntityTooLarge {
	return &GetExternalcontactsContactRequestEntityTooLarge{}
}

/*GetExternalcontactsContactRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsContactRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsContactRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactUnsupportedMediaType creates a GetExternalcontactsContactUnsupportedMediaType with default headers values
func NewGetExternalcontactsContactUnsupportedMediaType() *GetExternalcontactsContactUnsupportedMediaType {
	return &GetExternalcontactsContactUnsupportedMediaType{}
}

/*GetExternalcontactsContactUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsContactUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsContactUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactTooManyRequests creates a GetExternalcontactsContactTooManyRequests with default headers values
func NewGetExternalcontactsContactTooManyRequests() *GetExternalcontactsContactTooManyRequests {
	return &GetExternalcontactsContactTooManyRequests{}
}

/*GetExternalcontactsContactTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsContactTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsContactTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactInternalServerError creates a GetExternalcontactsContactInternalServerError with default headers values
func NewGetExternalcontactsContactInternalServerError() *GetExternalcontactsContactInternalServerError {
	return &GetExternalcontactsContactInternalServerError{}
}

/*GetExternalcontactsContactInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsContactInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsContactInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactServiceUnavailable creates a GetExternalcontactsContactServiceUnavailable with default headers values
func NewGetExternalcontactsContactServiceUnavailable() *GetExternalcontactsContactServiceUnavailable {
	return &GetExternalcontactsContactServiceUnavailable{}
}

/*GetExternalcontactsContactServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsContactServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsContactServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactGatewayTimeout creates a GetExternalcontactsContactGatewayTimeout with default headers values
func NewGetExternalcontactsContactGatewayTimeout() *GetExternalcontactsContactGatewayTimeout {
	return &GetExternalcontactsContactGatewayTimeout{}
}

/*GetExternalcontactsContactGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsContactGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}][%d] getExternalcontactsContactGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsContactGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
