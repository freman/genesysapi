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

// GetExternalcontactsOrganizationNotesReader is a Reader for the GetExternalcontactsOrganizationNotes structure.
type GetExternalcontactsOrganizationNotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsOrganizationNotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsOrganizationNotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsOrganizationNotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsOrganizationNotesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsOrganizationNotesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsOrganizationNotesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsOrganizationNotesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsOrganizationNotesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsOrganizationNotesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsOrganizationNotesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsOrganizationNotesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsOrganizationNotesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsOrganizationNotesOK creates a GetExternalcontactsOrganizationNotesOK with default headers values
func NewGetExternalcontactsOrganizationNotesOK() *GetExternalcontactsOrganizationNotesOK {
	return &GetExternalcontactsOrganizationNotesOK{}
}

/*GetExternalcontactsOrganizationNotesOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsOrganizationNotesOK struct {
	Payload *models.NoteListing
}

func (o *GetExternalcontactsOrganizationNotesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesOK) GetPayload() *models.NoteListing {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NoteListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesBadRequest creates a GetExternalcontactsOrganizationNotesBadRequest with default headers values
func NewGetExternalcontactsOrganizationNotesBadRequest() *GetExternalcontactsOrganizationNotesBadRequest {
	return &GetExternalcontactsOrganizationNotesBadRequest{}
}

/*GetExternalcontactsOrganizationNotesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsOrganizationNotesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesUnauthorized creates a GetExternalcontactsOrganizationNotesUnauthorized with default headers values
func NewGetExternalcontactsOrganizationNotesUnauthorized() *GetExternalcontactsOrganizationNotesUnauthorized {
	return &GetExternalcontactsOrganizationNotesUnauthorized{}
}

/*GetExternalcontactsOrganizationNotesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsOrganizationNotesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesForbidden creates a GetExternalcontactsOrganizationNotesForbidden with default headers values
func NewGetExternalcontactsOrganizationNotesForbidden() *GetExternalcontactsOrganizationNotesForbidden {
	return &GetExternalcontactsOrganizationNotesForbidden{}
}

/*GetExternalcontactsOrganizationNotesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsOrganizationNotesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesNotFound creates a GetExternalcontactsOrganizationNotesNotFound with default headers values
func NewGetExternalcontactsOrganizationNotesNotFound() *GetExternalcontactsOrganizationNotesNotFound {
	return &GetExternalcontactsOrganizationNotesNotFound{}
}

/*GetExternalcontactsOrganizationNotesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsOrganizationNotesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesRequestEntityTooLarge creates a GetExternalcontactsOrganizationNotesRequestEntityTooLarge with default headers values
func NewGetExternalcontactsOrganizationNotesRequestEntityTooLarge() *GetExternalcontactsOrganizationNotesRequestEntityTooLarge {
	return &GetExternalcontactsOrganizationNotesRequestEntityTooLarge{}
}

/*GetExternalcontactsOrganizationNotesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsOrganizationNotesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesUnsupportedMediaType creates a GetExternalcontactsOrganizationNotesUnsupportedMediaType with default headers values
func NewGetExternalcontactsOrganizationNotesUnsupportedMediaType() *GetExternalcontactsOrganizationNotesUnsupportedMediaType {
	return &GetExternalcontactsOrganizationNotesUnsupportedMediaType{}
}

/*GetExternalcontactsOrganizationNotesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsOrganizationNotesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesTooManyRequests creates a GetExternalcontactsOrganizationNotesTooManyRequests with default headers values
func NewGetExternalcontactsOrganizationNotesTooManyRequests() *GetExternalcontactsOrganizationNotesTooManyRequests {
	return &GetExternalcontactsOrganizationNotesTooManyRequests{}
}

/*GetExternalcontactsOrganizationNotesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetExternalcontactsOrganizationNotesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesInternalServerError creates a GetExternalcontactsOrganizationNotesInternalServerError with default headers values
func NewGetExternalcontactsOrganizationNotesInternalServerError() *GetExternalcontactsOrganizationNotesInternalServerError {
	return &GetExternalcontactsOrganizationNotesInternalServerError{}
}

/*GetExternalcontactsOrganizationNotesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsOrganizationNotesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesServiceUnavailable creates a GetExternalcontactsOrganizationNotesServiceUnavailable with default headers values
func NewGetExternalcontactsOrganizationNotesServiceUnavailable() *GetExternalcontactsOrganizationNotesServiceUnavailable {
	return &GetExternalcontactsOrganizationNotesServiceUnavailable{}
}

/*GetExternalcontactsOrganizationNotesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsOrganizationNotesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationNotesGatewayTimeout creates a GetExternalcontactsOrganizationNotesGatewayTimeout with default headers values
func NewGetExternalcontactsOrganizationNotesGatewayTimeout() *GetExternalcontactsOrganizationNotesGatewayTimeout {
	return &GetExternalcontactsOrganizationNotesGatewayTimeout{}
}

/*GetExternalcontactsOrganizationNotesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsOrganizationNotesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}