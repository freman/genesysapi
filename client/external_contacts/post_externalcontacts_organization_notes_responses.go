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

// PostExternalcontactsOrganizationNotesReader is a Reader for the PostExternalcontactsOrganizationNotes structure.
type PostExternalcontactsOrganizationNotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsOrganizationNotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsOrganizationNotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsOrganizationNotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsOrganizationNotesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsOrganizationNotesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsOrganizationNotesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsOrganizationNotesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsOrganizationNotesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsOrganizationNotesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsOrganizationNotesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsOrganizationNotesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsOrganizationNotesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsOrganizationNotesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsOrganizationNotesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsOrganizationNotesOK creates a PostExternalcontactsOrganizationNotesOK with default headers values
func NewPostExternalcontactsOrganizationNotesOK() *PostExternalcontactsOrganizationNotesOK {
	return &PostExternalcontactsOrganizationNotesOK{}
}

/*PostExternalcontactsOrganizationNotesOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsOrganizationNotesOK struct {
	Payload *models.Note
}

func (o *PostExternalcontactsOrganizationNotesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesBadRequest creates a PostExternalcontactsOrganizationNotesBadRequest with default headers values
func NewPostExternalcontactsOrganizationNotesBadRequest() *PostExternalcontactsOrganizationNotesBadRequest {
	return &PostExternalcontactsOrganizationNotesBadRequest{}
}

/*PostExternalcontactsOrganizationNotesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsOrganizationNotesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesUnauthorized creates a PostExternalcontactsOrganizationNotesUnauthorized with default headers values
func NewPostExternalcontactsOrganizationNotesUnauthorized() *PostExternalcontactsOrganizationNotesUnauthorized {
	return &PostExternalcontactsOrganizationNotesUnauthorized{}
}

/*PostExternalcontactsOrganizationNotesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsOrganizationNotesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesForbidden creates a PostExternalcontactsOrganizationNotesForbidden with default headers values
func NewPostExternalcontactsOrganizationNotesForbidden() *PostExternalcontactsOrganizationNotesForbidden {
	return &PostExternalcontactsOrganizationNotesForbidden{}
}

/*PostExternalcontactsOrganizationNotesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsOrganizationNotesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesNotFound creates a PostExternalcontactsOrganizationNotesNotFound with default headers values
func NewPostExternalcontactsOrganizationNotesNotFound() *PostExternalcontactsOrganizationNotesNotFound {
	return &PostExternalcontactsOrganizationNotesNotFound{}
}

/*PostExternalcontactsOrganizationNotesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsOrganizationNotesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesRequestTimeout creates a PostExternalcontactsOrganizationNotesRequestTimeout with default headers values
func NewPostExternalcontactsOrganizationNotesRequestTimeout() *PostExternalcontactsOrganizationNotesRequestTimeout {
	return &PostExternalcontactsOrganizationNotesRequestTimeout{}
}

/*PostExternalcontactsOrganizationNotesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsOrganizationNotesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesRequestEntityTooLarge creates a PostExternalcontactsOrganizationNotesRequestEntityTooLarge with default headers values
func NewPostExternalcontactsOrganizationNotesRequestEntityTooLarge() *PostExternalcontactsOrganizationNotesRequestEntityTooLarge {
	return &PostExternalcontactsOrganizationNotesRequestEntityTooLarge{}
}

/*PostExternalcontactsOrganizationNotesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostExternalcontactsOrganizationNotesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesUnsupportedMediaType creates a PostExternalcontactsOrganizationNotesUnsupportedMediaType with default headers values
func NewPostExternalcontactsOrganizationNotesUnsupportedMediaType() *PostExternalcontactsOrganizationNotesUnsupportedMediaType {
	return &PostExternalcontactsOrganizationNotesUnsupportedMediaType{}
}

/*PostExternalcontactsOrganizationNotesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsOrganizationNotesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesUnprocessableEntity creates a PostExternalcontactsOrganizationNotesUnprocessableEntity with default headers values
func NewPostExternalcontactsOrganizationNotesUnprocessableEntity() *PostExternalcontactsOrganizationNotesUnprocessableEntity {
	return &PostExternalcontactsOrganizationNotesUnprocessableEntity{}
}

/*PostExternalcontactsOrganizationNotesUnprocessableEntity handles this case with default header values.

PostExternalcontactsOrganizationNotesUnprocessableEntity post externalcontacts organization notes unprocessable entity
*/
type PostExternalcontactsOrganizationNotesUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesTooManyRequests creates a PostExternalcontactsOrganizationNotesTooManyRequests with default headers values
func NewPostExternalcontactsOrganizationNotesTooManyRequests() *PostExternalcontactsOrganizationNotesTooManyRequests {
	return &PostExternalcontactsOrganizationNotesTooManyRequests{}
}

/*PostExternalcontactsOrganizationNotesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsOrganizationNotesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesInternalServerError creates a PostExternalcontactsOrganizationNotesInternalServerError with default headers values
func NewPostExternalcontactsOrganizationNotesInternalServerError() *PostExternalcontactsOrganizationNotesInternalServerError {
	return &PostExternalcontactsOrganizationNotesInternalServerError{}
}

/*PostExternalcontactsOrganizationNotesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsOrganizationNotesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesServiceUnavailable creates a PostExternalcontactsOrganizationNotesServiceUnavailable with default headers values
func NewPostExternalcontactsOrganizationNotesServiceUnavailable() *PostExternalcontactsOrganizationNotesServiceUnavailable {
	return &PostExternalcontactsOrganizationNotesServiceUnavailable{}
}

/*PostExternalcontactsOrganizationNotesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsOrganizationNotesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsOrganizationNotesGatewayTimeout creates a PostExternalcontactsOrganizationNotesGatewayTimeout with default headers values
func NewPostExternalcontactsOrganizationNotesGatewayTimeout() *PostExternalcontactsOrganizationNotesGatewayTimeout {
	return &PostExternalcontactsOrganizationNotesGatewayTimeout{}
}

/*PostExternalcontactsOrganizationNotesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsOrganizationNotesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsOrganizationNotesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] postExternalcontactsOrganizationNotesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsOrganizationNotesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsOrganizationNotesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
