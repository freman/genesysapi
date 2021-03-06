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

// GetExternalcontactsOrganizationNoteReader is a Reader for the GetExternalcontactsOrganizationNote structure.
type GetExternalcontactsOrganizationNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsOrganizationNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsOrganizationNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsOrganizationNoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsOrganizationNoteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsOrganizationNoteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsOrganizationNoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsOrganizationNoteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsOrganizationNoteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsOrganizationNoteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsOrganizationNoteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsOrganizationNoteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsOrganizationNoteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsOrganizationNoteOK creates a GetExternalcontactsOrganizationNoteOK with default headers values
func NewGetExternalcontactsOrganizationNoteOK() *GetExternalcontactsOrganizationNoteOK {
	return &GetExternalcontactsOrganizationNoteOK{}
}

/*GetExternalcontactsOrganizationNoteOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsOrganizationNoteOK struct {
	Payload *models.Note
}

func (o *GetExternalcontactsOrganizationNoteOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteBadRequest creates a GetExternalcontactsOrganizationNoteBadRequest with default headers values
func NewGetExternalcontactsOrganizationNoteBadRequest() *GetExternalcontactsOrganizationNoteBadRequest {
	return &GetExternalcontactsOrganizationNoteBadRequest{}
}

/*GetExternalcontactsOrganizationNoteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsOrganizationNoteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteUnauthorized creates a GetExternalcontactsOrganizationNoteUnauthorized with default headers values
func NewGetExternalcontactsOrganizationNoteUnauthorized() *GetExternalcontactsOrganizationNoteUnauthorized {
	return &GetExternalcontactsOrganizationNoteUnauthorized{}
}

/*GetExternalcontactsOrganizationNoteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsOrganizationNoteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteForbidden creates a GetExternalcontactsOrganizationNoteForbidden with default headers values
func NewGetExternalcontactsOrganizationNoteForbidden() *GetExternalcontactsOrganizationNoteForbidden {
	return &GetExternalcontactsOrganizationNoteForbidden{}
}

/*GetExternalcontactsOrganizationNoteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsOrganizationNoteForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteNotFound creates a GetExternalcontactsOrganizationNoteNotFound with default headers values
func NewGetExternalcontactsOrganizationNoteNotFound() *GetExternalcontactsOrganizationNoteNotFound {
	return &GetExternalcontactsOrganizationNoteNotFound{}
}

/*GetExternalcontactsOrganizationNoteNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsOrganizationNoteNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteRequestEntityTooLarge creates a GetExternalcontactsOrganizationNoteRequestEntityTooLarge with default headers values
func NewGetExternalcontactsOrganizationNoteRequestEntityTooLarge() *GetExternalcontactsOrganizationNoteRequestEntityTooLarge {
	return &GetExternalcontactsOrganizationNoteRequestEntityTooLarge{}
}

/*GetExternalcontactsOrganizationNoteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsOrganizationNoteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteUnsupportedMediaType creates a GetExternalcontactsOrganizationNoteUnsupportedMediaType with default headers values
func NewGetExternalcontactsOrganizationNoteUnsupportedMediaType() *GetExternalcontactsOrganizationNoteUnsupportedMediaType {
	return &GetExternalcontactsOrganizationNoteUnsupportedMediaType{}
}

/*GetExternalcontactsOrganizationNoteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsOrganizationNoteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteTooManyRequests creates a GetExternalcontactsOrganizationNoteTooManyRequests with default headers values
func NewGetExternalcontactsOrganizationNoteTooManyRequests() *GetExternalcontactsOrganizationNoteTooManyRequests {
	return &GetExternalcontactsOrganizationNoteTooManyRequests{}
}

/*GetExternalcontactsOrganizationNoteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetExternalcontactsOrganizationNoteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteInternalServerError creates a GetExternalcontactsOrganizationNoteInternalServerError with default headers values
func NewGetExternalcontactsOrganizationNoteInternalServerError() *GetExternalcontactsOrganizationNoteInternalServerError {
	return &GetExternalcontactsOrganizationNoteInternalServerError{}
}

/*GetExternalcontactsOrganizationNoteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsOrganizationNoteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteServiceUnavailable creates a GetExternalcontactsOrganizationNoteServiceUnavailable with default headers values
func NewGetExternalcontactsOrganizationNoteServiceUnavailable() *GetExternalcontactsOrganizationNoteServiceUnavailable {
	return &GetExternalcontactsOrganizationNoteServiceUnavailable{}
}

/*GetExternalcontactsOrganizationNoteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsOrganizationNoteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNoteGatewayTimeout creates a GetExternalcontactsOrganizationNoteGatewayTimeout with default headers values
func NewGetExternalcontactsOrganizationNoteGatewayTimeout() *GetExternalcontactsOrganizationNoteGatewayTimeout {
	return &GetExternalcontactsOrganizationNoteGatewayTimeout{}
}

/*GetExternalcontactsOrganizationNoteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsOrganizationNoteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNoteGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes/{noteId}][%d] getExternalcontactsOrganizationNoteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationNoteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNoteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
