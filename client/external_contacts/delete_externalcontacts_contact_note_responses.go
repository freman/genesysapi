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

// DeleteExternalcontactsContactNoteReader is a Reader for the DeleteExternalcontactsContactNote structure.
type DeleteExternalcontactsContactNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalcontactsContactNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExternalcontactsContactNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalcontactsContactNoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalcontactsContactNoteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalcontactsContactNoteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalcontactsContactNoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteExternalcontactsContactNoteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteExternalcontactsContactNoteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteExternalcontactsContactNoteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalcontactsContactNoteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteExternalcontactsContactNoteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteExternalcontactsContactNoteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteExternalcontactsContactNoteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExternalcontactsContactNoteOK creates a DeleteExternalcontactsContactNoteOK with default headers values
func NewDeleteExternalcontactsContactNoteOK() *DeleteExternalcontactsContactNoteOK {
	return &DeleteExternalcontactsContactNoteOK{}
}

/*DeleteExternalcontactsContactNoteOK handles this case with default header values.

successful operation
*/
type DeleteExternalcontactsContactNoteOK struct {
	Payload models.Empty
}

func (o *DeleteExternalcontactsContactNoteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteOK  %+v", 200, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteBadRequest creates a DeleteExternalcontactsContactNoteBadRequest with default headers values
func NewDeleteExternalcontactsContactNoteBadRequest() *DeleteExternalcontactsContactNoteBadRequest {
	return &DeleteExternalcontactsContactNoteBadRequest{}
}

/*DeleteExternalcontactsContactNoteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteExternalcontactsContactNoteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteUnauthorized creates a DeleteExternalcontactsContactNoteUnauthorized with default headers values
func NewDeleteExternalcontactsContactNoteUnauthorized() *DeleteExternalcontactsContactNoteUnauthorized {
	return &DeleteExternalcontactsContactNoteUnauthorized{}
}

/*DeleteExternalcontactsContactNoteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteExternalcontactsContactNoteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteForbidden creates a DeleteExternalcontactsContactNoteForbidden with default headers values
func NewDeleteExternalcontactsContactNoteForbidden() *DeleteExternalcontactsContactNoteForbidden {
	return &DeleteExternalcontactsContactNoteForbidden{}
}

/*DeleteExternalcontactsContactNoteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteExternalcontactsContactNoteForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteNotFound creates a DeleteExternalcontactsContactNoteNotFound with default headers values
func NewDeleteExternalcontactsContactNoteNotFound() *DeleteExternalcontactsContactNoteNotFound {
	return &DeleteExternalcontactsContactNoteNotFound{}
}

/*DeleteExternalcontactsContactNoteNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteExternalcontactsContactNoteNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteRequestTimeout creates a DeleteExternalcontactsContactNoteRequestTimeout with default headers values
func NewDeleteExternalcontactsContactNoteRequestTimeout() *DeleteExternalcontactsContactNoteRequestTimeout {
	return &DeleteExternalcontactsContactNoteRequestTimeout{}
}

/*DeleteExternalcontactsContactNoteRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteExternalcontactsContactNoteRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteRequestEntityTooLarge creates a DeleteExternalcontactsContactNoteRequestEntityTooLarge with default headers values
func NewDeleteExternalcontactsContactNoteRequestEntityTooLarge() *DeleteExternalcontactsContactNoteRequestEntityTooLarge {
	return &DeleteExternalcontactsContactNoteRequestEntityTooLarge{}
}

/*DeleteExternalcontactsContactNoteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteExternalcontactsContactNoteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteUnsupportedMediaType creates a DeleteExternalcontactsContactNoteUnsupportedMediaType with default headers values
func NewDeleteExternalcontactsContactNoteUnsupportedMediaType() *DeleteExternalcontactsContactNoteUnsupportedMediaType {
	return &DeleteExternalcontactsContactNoteUnsupportedMediaType{}
}

/*DeleteExternalcontactsContactNoteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteExternalcontactsContactNoteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteTooManyRequests creates a DeleteExternalcontactsContactNoteTooManyRequests with default headers values
func NewDeleteExternalcontactsContactNoteTooManyRequests() *DeleteExternalcontactsContactNoteTooManyRequests {
	return &DeleteExternalcontactsContactNoteTooManyRequests{}
}

/*DeleteExternalcontactsContactNoteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteExternalcontactsContactNoteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteInternalServerError creates a DeleteExternalcontactsContactNoteInternalServerError with default headers values
func NewDeleteExternalcontactsContactNoteInternalServerError() *DeleteExternalcontactsContactNoteInternalServerError {
	return &DeleteExternalcontactsContactNoteInternalServerError{}
}

/*DeleteExternalcontactsContactNoteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteExternalcontactsContactNoteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteServiceUnavailable creates a DeleteExternalcontactsContactNoteServiceUnavailable with default headers values
func NewDeleteExternalcontactsContactNoteServiceUnavailable() *DeleteExternalcontactsContactNoteServiceUnavailable {
	return &DeleteExternalcontactsContactNoteServiceUnavailable{}
}

/*DeleteExternalcontactsContactNoteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteExternalcontactsContactNoteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactNoteGatewayTimeout creates a DeleteExternalcontactsContactNoteGatewayTimeout with default headers values
func NewDeleteExternalcontactsContactNoteGatewayTimeout() *DeleteExternalcontactsContactNoteGatewayTimeout {
	return &DeleteExternalcontactsContactNoteGatewayTimeout{}
}

/*DeleteExternalcontactsContactNoteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteExternalcontactsContactNoteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteExternalcontactsContactNoteGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/{contactId}/notes/{noteId}][%d] deleteExternalcontactsContactNoteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsContactNoteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactNoteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
