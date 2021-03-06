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

// GetExternalcontactsContactsSchemaReader is a Reader for the GetExternalcontactsContactsSchema structure.
type GetExternalcontactsContactsSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsContactsSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsContactsSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsContactsSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsContactsSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsContactsSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsContactsSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsContactsSchemaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsContactsSchemaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsContactsSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsContactsSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsContactsSchemaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsContactsSchemaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsContactsSchemaOK creates a GetExternalcontactsContactsSchemaOK with default headers values
func NewGetExternalcontactsContactsSchemaOK() *GetExternalcontactsContactsSchemaOK {
	return &GetExternalcontactsContactsSchemaOK{}
}

/*GetExternalcontactsContactsSchemaOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsContactsSchemaOK struct {
	Payload *models.DataSchema
}

func (o *GetExternalcontactsContactsSchemaOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaOK) GetPayload() *models.DataSchema {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaBadRequest creates a GetExternalcontactsContactsSchemaBadRequest with default headers values
func NewGetExternalcontactsContactsSchemaBadRequest() *GetExternalcontactsContactsSchemaBadRequest {
	return &GetExternalcontactsContactsSchemaBadRequest{}
}

/*GetExternalcontactsContactsSchemaBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsContactsSchemaBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaUnauthorized creates a GetExternalcontactsContactsSchemaUnauthorized with default headers values
func NewGetExternalcontactsContactsSchemaUnauthorized() *GetExternalcontactsContactsSchemaUnauthorized {
	return &GetExternalcontactsContactsSchemaUnauthorized{}
}

/*GetExternalcontactsContactsSchemaUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsContactsSchemaUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaForbidden creates a GetExternalcontactsContactsSchemaForbidden with default headers values
func NewGetExternalcontactsContactsSchemaForbidden() *GetExternalcontactsContactsSchemaForbidden {
	return &GetExternalcontactsContactsSchemaForbidden{}
}

/*GetExternalcontactsContactsSchemaForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsContactsSchemaForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaNotFound creates a GetExternalcontactsContactsSchemaNotFound with default headers values
func NewGetExternalcontactsContactsSchemaNotFound() *GetExternalcontactsContactsSchemaNotFound {
	return &GetExternalcontactsContactsSchemaNotFound{}
}

/*GetExternalcontactsContactsSchemaNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsContactsSchemaNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaRequestEntityTooLarge creates a GetExternalcontactsContactsSchemaRequestEntityTooLarge with default headers values
func NewGetExternalcontactsContactsSchemaRequestEntityTooLarge() *GetExternalcontactsContactsSchemaRequestEntityTooLarge {
	return &GetExternalcontactsContactsSchemaRequestEntityTooLarge{}
}

/*GetExternalcontactsContactsSchemaRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsContactsSchemaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaUnsupportedMediaType creates a GetExternalcontactsContactsSchemaUnsupportedMediaType with default headers values
func NewGetExternalcontactsContactsSchemaUnsupportedMediaType() *GetExternalcontactsContactsSchemaUnsupportedMediaType {
	return &GetExternalcontactsContactsSchemaUnsupportedMediaType{}
}

/*GetExternalcontactsContactsSchemaUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsContactsSchemaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaTooManyRequests creates a GetExternalcontactsContactsSchemaTooManyRequests with default headers values
func NewGetExternalcontactsContactsSchemaTooManyRequests() *GetExternalcontactsContactsSchemaTooManyRequests {
	return &GetExternalcontactsContactsSchemaTooManyRequests{}
}

/*GetExternalcontactsContactsSchemaTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetExternalcontactsContactsSchemaTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaInternalServerError creates a GetExternalcontactsContactsSchemaInternalServerError with default headers values
func NewGetExternalcontactsContactsSchemaInternalServerError() *GetExternalcontactsContactsSchemaInternalServerError {
	return &GetExternalcontactsContactsSchemaInternalServerError{}
}

/*GetExternalcontactsContactsSchemaInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsContactsSchemaInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaServiceUnavailable creates a GetExternalcontactsContactsSchemaServiceUnavailable with default headers values
func NewGetExternalcontactsContactsSchemaServiceUnavailable() *GetExternalcontactsContactsSchemaServiceUnavailable {
	return &GetExternalcontactsContactsSchemaServiceUnavailable{}
}

/*GetExternalcontactsContactsSchemaServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsContactsSchemaServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaGatewayTimeout creates a GetExternalcontactsContactsSchemaGatewayTimeout with default headers values
func NewGetExternalcontactsContactsSchemaGatewayTimeout() *GetExternalcontactsContactsSchemaGatewayTimeout {
	return &GetExternalcontactsContactsSchemaGatewayTimeout{}
}

/*GetExternalcontactsContactsSchemaGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsContactsSchemaGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] getExternalcontactsContactsSchemaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
