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

// DeleteExternalcontactsContactsSchemaReader is a Reader for the DeleteExternalcontactsContactsSchema structure.
type DeleteExternalcontactsContactsSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalcontactsContactsSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteExternalcontactsContactsSchemaNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalcontactsContactsSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalcontactsContactsSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalcontactsContactsSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalcontactsContactsSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteExternalcontactsContactsSchemaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteExternalcontactsContactsSchemaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalcontactsContactsSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteExternalcontactsContactsSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteExternalcontactsContactsSchemaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteExternalcontactsContactsSchemaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExternalcontactsContactsSchemaNoContent creates a DeleteExternalcontactsContactsSchemaNoContent with default headers values
func NewDeleteExternalcontactsContactsSchemaNoContent() *DeleteExternalcontactsContactsSchemaNoContent {
	return &DeleteExternalcontactsContactsSchemaNoContent{}
}

/*DeleteExternalcontactsContactsSchemaNoContent handles this case with default header values.

Operation was successful.
*/
type DeleteExternalcontactsContactsSchemaNoContent struct {
}

func (o *DeleteExternalcontactsContactsSchemaNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaNoContent ", 204)
}

func (o *DeleteExternalcontactsContactsSchemaNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalcontactsContactsSchemaBadRequest creates a DeleteExternalcontactsContactsSchemaBadRequest with default headers values
func NewDeleteExternalcontactsContactsSchemaBadRequest() *DeleteExternalcontactsContactsSchemaBadRequest {
	return &DeleteExternalcontactsContactsSchemaBadRequest{}
}

/*DeleteExternalcontactsContactsSchemaBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteExternalcontactsContactsSchemaBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaUnauthorized creates a DeleteExternalcontactsContactsSchemaUnauthorized with default headers values
func NewDeleteExternalcontactsContactsSchemaUnauthorized() *DeleteExternalcontactsContactsSchemaUnauthorized {
	return &DeleteExternalcontactsContactsSchemaUnauthorized{}
}

/*DeleteExternalcontactsContactsSchemaUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteExternalcontactsContactsSchemaUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaForbidden creates a DeleteExternalcontactsContactsSchemaForbidden with default headers values
func NewDeleteExternalcontactsContactsSchemaForbidden() *DeleteExternalcontactsContactsSchemaForbidden {
	return &DeleteExternalcontactsContactsSchemaForbidden{}
}

/*DeleteExternalcontactsContactsSchemaForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteExternalcontactsContactsSchemaForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaNotFound creates a DeleteExternalcontactsContactsSchemaNotFound with default headers values
func NewDeleteExternalcontactsContactsSchemaNotFound() *DeleteExternalcontactsContactsSchemaNotFound {
	return &DeleteExternalcontactsContactsSchemaNotFound{}
}

/*DeleteExternalcontactsContactsSchemaNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteExternalcontactsContactsSchemaNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaRequestEntityTooLarge creates a DeleteExternalcontactsContactsSchemaRequestEntityTooLarge with default headers values
func NewDeleteExternalcontactsContactsSchemaRequestEntityTooLarge() *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge {
	return &DeleteExternalcontactsContactsSchemaRequestEntityTooLarge{}
}

/*DeleteExternalcontactsContactsSchemaRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteExternalcontactsContactsSchemaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaUnsupportedMediaType creates a DeleteExternalcontactsContactsSchemaUnsupportedMediaType with default headers values
func NewDeleteExternalcontactsContactsSchemaUnsupportedMediaType() *DeleteExternalcontactsContactsSchemaUnsupportedMediaType {
	return &DeleteExternalcontactsContactsSchemaUnsupportedMediaType{}
}

/*DeleteExternalcontactsContactsSchemaUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteExternalcontactsContactsSchemaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaTooManyRequests creates a DeleteExternalcontactsContactsSchemaTooManyRequests with default headers values
func NewDeleteExternalcontactsContactsSchemaTooManyRequests() *DeleteExternalcontactsContactsSchemaTooManyRequests {
	return &DeleteExternalcontactsContactsSchemaTooManyRequests{}
}

/*DeleteExternalcontactsContactsSchemaTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteExternalcontactsContactsSchemaTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaInternalServerError creates a DeleteExternalcontactsContactsSchemaInternalServerError with default headers values
func NewDeleteExternalcontactsContactsSchemaInternalServerError() *DeleteExternalcontactsContactsSchemaInternalServerError {
	return &DeleteExternalcontactsContactsSchemaInternalServerError{}
}

/*DeleteExternalcontactsContactsSchemaInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteExternalcontactsContactsSchemaInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaServiceUnavailable creates a DeleteExternalcontactsContactsSchemaServiceUnavailable with default headers values
func NewDeleteExternalcontactsContactsSchemaServiceUnavailable() *DeleteExternalcontactsContactsSchemaServiceUnavailable {
	return &DeleteExternalcontactsContactsSchemaServiceUnavailable{}
}

/*DeleteExternalcontactsContactsSchemaServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteExternalcontactsContactsSchemaServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaGatewayTimeout creates a DeleteExternalcontactsContactsSchemaGatewayTimeout with default headers values
func NewDeleteExternalcontactsContactsSchemaGatewayTimeout() *DeleteExternalcontactsContactsSchemaGatewayTimeout {
	return &DeleteExternalcontactsContactsSchemaGatewayTimeout{}
}

/*DeleteExternalcontactsContactsSchemaGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteExternalcontactsContactsSchemaGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
