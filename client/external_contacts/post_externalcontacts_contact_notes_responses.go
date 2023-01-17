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

// PostExternalcontactsContactNotesReader is a Reader for the PostExternalcontactsContactNotes structure.
type PostExternalcontactsContactNotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsContactNotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsContactNotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsContactNotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsContactNotesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsContactNotesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsContactNotesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsContactNotesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsContactNotesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsContactNotesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsContactNotesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsContactNotesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsContactNotesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsContactNotesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsContactNotesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsContactNotesOK creates a PostExternalcontactsContactNotesOK with default headers values
func NewPostExternalcontactsContactNotesOK() *PostExternalcontactsContactNotesOK {
	return &PostExternalcontactsContactNotesOK{}
}

/*
PostExternalcontactsContactNotesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsContactNotesOK struct {
	Payload *models.Note
}

// IsSuccess returns true when this post externalcontacts contact notes o k response has a 2xx status code
func (o *PostExternalcontactsContactNotesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts contact notes o k response has a 3xx status code
func (o *PostExternalcontactsContactNotesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes o k response has a 4xx status code
func (o *PostExternalcontactsContactNotesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts contact notes o k response has a 5xx status code
func (o *PostExternalcontactsContactNotesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes o k response a status code equal to that given
func (o *PostExternalcontactsContactNotesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsContactNotesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsContactNotesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsContactNotesOK) GetPayload() *models.Note {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Note)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesBadRequest creates a PostExternalcontactsContactNotesBadRequest with default headers values
func NewPostExternalcontactsContactNotesBadRequest() *PostExternalcontactsContactNotesBadRequest {
	return &PostExternalcontactsContactNotesBadRequest{}
}

/*
PostExternalcontactsContactNotesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsContactNotesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes bad request response has a 2xx status code
func (o *PostExternalcontactsContactNotesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes bad request response has a 3xx status code
func (o *PostExternalcontactsContactNotesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes bad request response has a 4xx status code
func (o *PostExternalcontactsContactNotesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes bad request response has a 5xx status code
func (o *PostExternalcontactsContactNotesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes bad request response a status code equal to that given
func (o *PostExternalcontactsContactNotesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsContactNotesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsContactNotesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsContactNotesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesUnauthorized creates a PostExternalcontactsContactNotesUnauthorized with default headers values
func NewPostExternalcontactsContactNotesUnauthorized() *PostExternalcontactsContactNotesUnauthorized {
	return &PostExternalcontactsContactNotesUnauthorized{}
}

/*
PostExternalcontactsContactNotesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsContactNotesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes unauthorized response has a 2xx status code
func (o *PostExternalcontactsContactNotesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes unauthorized response has a 3xx status code
func (o *PostExternalcontactsContactNotesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes unauthorized response has a 4xx status code
func (o *PostExternalcontactsContactNotesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes unauthorized response has a 5xx status code
func (o *PostExternalcontactsContactNotesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes unauthorized response a status code equal to that given
func (o *PostExternalcontactsContactNotesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsContactNotesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsContactNotesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsContactNotesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesForbidden creates a PostExternalcontactsContactNotesForbidden with default headers values
func NewPostExternalcontactsContactNotesForbidden() *PostExternalcontactsContactNotesForbidden {
	return &PostExternalcontactsContactNotesForbidden{}
}

/*
PostExternalcontactsContactNotesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsContactNotesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes forbidden response has a 2xx status code
func (o *PostExternalcontactsContactNotesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes forbidden response has a 3xx status code
func (o *PostExternalcontactsContactNotesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes forbidden response has a 4xx status code
func (o *PostExternalcontactsContactNotesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes forbidden response has a 5xx status code
func (o *PostExternalcontactsContactNotesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes forbidden response a status code equal to that given
func (o *PostExternalcontactsContactNotesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsContactNotesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsContactNotesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsContactNotesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesNotFound creates a PostExternalcontactsContactNotesNotFound with default headers values
func NewPostExternalcontactsContactNotesNotFound() *PostExternalcontactsContactNotesNotFound {
	return &PostExternalcontactsContactNotesNotFound{}
}

/*
PostExternalcontactsContactNotesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsContactNotesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes not found response has a 2xx status code
func (o *PostExternalcontactsContactNotesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes not found response has a 3xx status code
func (o *PostExternalcontactsContactNotesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes not found response has a 4xx status code
func (o *PostExternalcontactsContactNotesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes not found response has a 5xx status code
func (o *PostExternalcontactsContactNotesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes not found response a status code equal to that given
func (o *PostExternalcontactsContactNotesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsContactNotesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsContactNotesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsContactNotesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesRequestTimeout creates a PostExternalcontactsContactNotesRequestTimeout with default headers values
func NewPostExternalcontactsContactNotesRequestTimeout() *PostExternalcontactsContactNotesRequestTimeout {
	return &PostExternalcontactsContactNotesRequestTimeout{}
}

/*
PostExternalcontactsContactNotesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsContactNotesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes request timeout response has a 2xx status code
func (o *PostExternalcontactsContactNotesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes request timeout response has a 3xx status code
func (o *PostExternalcontactsContactNotesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes request timeout response has a 4xx status code
func (o *PostExternalcontactsContactNotesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes request timeout response has a 5xx status code
func (o *PostExternalcontactsContactNotesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes request timeout response a status code equal to that given
func (o *PostExternalcontactsContactNotesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsContactNotesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsContactNotesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsContactNotesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesRequestEntityTooLarge creates a PostExternalcontactsContactNotesRequestEntityTooLarge with default headers values
func NewPostExternalcontactsContactNotesRequestEntityTooLarge() *PostExternalcontactsContactNotesRequestEntityTooLarge {
	return &PostExternalcontactsContactNotesRequestEntityTooLarge{}
}

/*
PostExternalcontactsContactNotesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostExternalcontactsContactNotesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes request entity too large response has a 2xx status code
func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes request entity too large response has a 3xx status code
func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes request entity too large response has a 4xx status code
func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes request entity too large response has a 5xx status code
func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes request entity too large response a status code equal to that given
func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesUnsupportedMediaType creates a PostExternalcontactsContactNotesUnsupportedMediaType with default headers values
func NewPostExternalcontactsContactNotesUnsupportedMediaType() *PostExternalcontactsContactNotesUnsupportedMediaType {
	return &PostExternalcontactsContactNotesUnsupportedMediaType{}
}

/*
PostExternalcontactsContactNotesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsContactNotesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes unsupported media type response has a 2xx status code
func (o *PostExternalcontactsContactNotesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes unsupported media type response has a 3xx status code
func (o *PostExternalcontactsContactNotesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes unsupported media type response has a 4xx status code
func (o *PostExternalcontactsContactNotesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes unsupported media type response has a 5xx status code
func (o *PostExternalcontactsContactNotesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes unsupported media type response a status code equal to that given
func (o *PostExternalcontactsContactNotesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsContactNotesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsContactNotesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsContactNotesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesUnprocessableEntity creates a PostExternalcontactsContactNotesUnprocessableEntity with default headers values
func NewPostExternalcontactsContactNotesUnprocessableEntity() *PostExternalcontactsContactNotesUnprocessableEntity {
	return &PostExternalcontactsContactNotesUnprocessableEntity{}
}

/*
PostExternalcontactsContactNotesUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsContactNotesUnprocessableEntity post externalcontacts contact notes unprocessable entity
*/
type PostExternalcontactsContactNotesUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsContactNotesUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsContactNotesUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsContactNotesUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsContactNotesUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsContactNotesUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsContactNotesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsContactNotesUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsContactNotesUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesTooManyRequests creates a PostExternalcontactsContactNotesTooManyRequests with default headers values
func NewPostExternalcontactsContactNotesTooManyRequests() *PostExternalcontactsContactNotesTooManyRequests {
	return &PostExternalcontactsContactNotesTooManyRequests{}
}

/*
PostExternalcontactsContactNotesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsContactNotesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes too many requests response has a 2xx status code
func (o *PostExternalcontactsContactNotesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes too many requests response has a 3xx status code
func (o *PostExternalcontactsContactNotesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes too many requests response has a 4xx status code
func (o *PostExternalcontactsContactNotesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts contact notes too many requests response has a 5xx status code
func (o *PostExternalcontactsContactNotesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts contact notes too many requests response a status code equal to that given
func (o *PostExternalcontactsContactNotesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsContactNotesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsContactNotesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsContactNotesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesInternalServerError creates a PostExternalcontactsContactNotesInternalServerError with default headers values
func NewPostExternalcontactsContactNotesInternalServerError() *PostExternalcontactsContactNotesInternalServerError {
	return &PostExternalcontactsContactNotesInternalServerError{}
}

/*
PostExternalcontactsContactNotesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsContactNotesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes internal server error response has a 2xx status code
func (o *PostExternalcontactsContactNotesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes internal server error response has a 3xx status code
func (o *PostExternalcontactsContactNotesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes internal server error response has a 4xx status code
func (o *PostExternalcontactsContactNotesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts contact notes internal server error response has a 5xx status code
func (o *PostExternalcontactsContactNotesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts contact notes internal server error response a status code equal to that given
func (o *PostExternalcontactsContactNotesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsContactNotesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsContactNotesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsContactNotesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesServiceUnavailable creates a PostExternalcontactsContactNotesServiceUnavailable with default headers values
func NewPostExternalcontactsContactNotesServiceUnavailable() *PostExternalcontactsContactNotesServiceUnavailable {
	return &PostExternalcontactsContactNotesServiceUnavailable{}
}

/*
PostExternalcontactsContactNotesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsContactNotesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes service unavailable response has a 2xx status code
func (o *PostExternalcontactsContactNotesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes service unavailable response has a 3xx status code
func (o *PostExternalcontactsContactNotesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes service unavailable response has a 4xx status code
func (o *PostExternalcontactsContactNotesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts contact notes service unavailable response has a 5xx status code
func (o *PostExternalcontactsContactNotesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts contact notes service unavailable response a status code equal to that given
func (o *PostExternalcontactsContactNotesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsContactNotesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsContactNotesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsContactNotesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactNotesGatewayTimeout creates a PostExternalcontactsContactNotesGatewayTimeout with default headers values
func NewPostExternalcontactsContactNotesGatewayTimeout() *PostExternalcontactsContactNotesGatewayTimeout {
	return &PostExternalcontactsContactNotesGatewayTimeout{}
}

/*
PostExternalcontactsContactNotesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsContactNotesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts contact notes gateway timeout response has a 2xx status code
func (o *PostExternalcontactsContactNotesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts contact notes gateway timeout response has a 3xx status code
func (o *PostExternalcontactsContactNotesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts contact notes gateway timeout response has a 4xx status code
func (o *PostExternalcontactsContactNotesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts contact notes gateway timeout response has a 5xx status code
func (o *PostExternalcontactsContactNotesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts contact notes gateway timeout response a status code equal to that given
func (o *PostExternalcontactsContactNotesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsContactNotesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsContactNotesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts/{contactId}/notes][%d] postExternalcontactsContactNotesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsContactNotesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactNotesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
