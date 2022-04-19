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

// PostExternalcontactsBulkNotesReader is a Reader for the PostExternalcontactsBulkNotes structure.
type PostExternalcontactsBulkNotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkNotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkNotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkNotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkNotesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkNotesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkNotesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkNotesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkNotesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkNotesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkNotesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkNotesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkNotesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkNotesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkNotesOK creates a PostExternalcontactsBulkNotesOK with default headers values
func NewPostExternalcontactsBulkNotesOK() *PostExternalcontactsBulkNotesOK {
	return &PostExternalcontactsBulkNotesOK{}
}

/*PostExternalcontactsBulkNotesOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsBulkNotesOK struct {
	Payload *models.BulkFetchNotesResponse
}

func (o *PostExternalcontactsBulkNotesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkNotesOK) GetPayload() *models.BulkFetchNotesResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkFetchNotesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesBadRequest creates a PostExternalcontactsBulkNotesBadRequest with default headers values
func NewPostExternalcontactsBulkNotesBadRequest() *PostExternalcontactsBulkNotesBadRequest {
	return &PostExternalcontactsBulkNotesBadRequest{}
}

/*PostExternalcontactsBulkNotesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkNotesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkNotesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUnauthorized creates a PostExternalcontactsBulkNotesUnauthorized with default headers values
func NewPostExternalcontactsBulkNotesUnauthorized() *PostExternalcontactsBulkNotesUnauthorized {
	return &PostExternalcontactsBulkNotesUnauthorized{}
}

/*PostExternalcontactsBulkNotesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkNotesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesForbidden creates a PostExternalcontactsBulkNotesForbidden with default headers values
func NewPostExternalcontactsBulkNotesForbidden() *PostExternalcontactsBulkNotesForbidden {
	return &PostExternalcontactsBulkNotesForbidden{}
}

/*PostExternalcontactsBulkNotesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkNotesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkNotesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesNotFound creates a PostExternalcontactsBulkNotesNotFound with default headers values
func NewPostExternalcontactsBulkNotesNotFound() *PostExternalcontactsBulkNotesNotFound {
	return &PostExternalcontactsBulkNotesNotFound{}
}

/*PostExternalcontactsBulkNotesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkNotesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkNotesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRequestTimeout creates a PostExternalcontactsBulkNotesRequestTimeout with default headers values
func NewPostExternalcontactsBulkNotesRequestTimeout() *PostExternalcontactsBulkNotesRequestTimeout {
	return &PostExternalcontactsBulkNotesRequestTimeout{}
}

/*PostExternalcontactsBulkNotesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkNotesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRequestEntityTooLarge creates a PostExternalcontactsBulkNotesRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkNotesRequestEntityTooLarge() *PostExternalcontactsBulkNotesRequestEntityTooLarge {
	return &PostExternalcontactsBulkNotesRequestEntityTooLarge{}
}

/*PostExternalcontactsBulkNotesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostExternalcontactsBulkNotesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUnsupportedMediaType creates a PostExternalcontactsBulkNotesUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkNotesUnsupportedMediaType() *PostExternalcontactsBulkNotesUnsupportedMediaType {
	return &PostExternalcontactsBulkNotesUnsupportedMediaType{}
}

/*PostExternalcontactsBulkNotesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkNotesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesTooManyRequests creates a PostExternalcontactsBulkNotesTooManyRequests with default headers values
func NewPostExternalcontactsBulkNotesTooManyRequests() *PostExternalcontactsBulkNotesTooManyRequests {
	return &PostExternalcontactsBulkNotesTooManyRequests{}
}

/*PostExternalcontactsBulkNotesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkNotesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkNotesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesInternalServerError creates a PostExternalcontactsBulkNotesInternalServerError with default headers values
func NewPostExternalcontactsBulkNotesInternalServerError() *PostExternalcontactsBulkNotesInternalServerError {
	return &PostExternalcontactsBulkNotesInternalServerError{}
}

/*PostExternalcontactsBulkNotesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkNotesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkNotesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesServiceUnavailable creates a PostExternalcontactsBulkNotesServiceUnavailable with default headers values
func NewPostExternalcontactsBulkNotesServiceUnavailable() *PostExternalcontactsBulkNotesServiceUnavailable {
	return &PostExternalcontactsBulkNotesServiceUnavailable{}
}

/*PostExternalcontactsBulkNotesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkNotesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkNotesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesGatewayTimeout creates a PostExternalcontactsBulkNotesGatewayTimeout with default headers values
func NewPostExternalcontactsBulkNotesGatewayTimeout() *PostExternalcontactsBulkNotesGatewayTimeout {
	return &PostExternalcontactsBulkNotesGatewayTimeout{}
}

/*PostExternalcontactsBulkNotesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsBulkNotesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes][%d] postExternalcontactsBulkNotesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkNotesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
