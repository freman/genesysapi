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
	case 408:
		result := NewGetExternalcontactsOrganizationNotesRequestTimeout()
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

/*
GetExternalcontactsOrganizationNotesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsOrganizationNotesOK struct {
	Payload *models.NoteListing
}

// IsSuccess returns true when this get externalcontacts organization notes o k response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts organization notes o k response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes o k response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization notes o k response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes o k response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsOrganizationNotesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesOK) String() string {
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

/*
GetExternalcontactsOrganizationNotesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsOrganizationNotesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes bad request response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes bad request response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes bad request response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes bad request response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes bad request response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsOrganizationNotesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesBadRequest) String() string {
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

/*
GetExternalcontactsOrganizationNotesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsOrganizationNotesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes unauthorized response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes unauthorized response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes unauthorized response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes unauthorized response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes unauthorized response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsOrganizationNotesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesUnauthorized) String() string {
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

/*
GetExternalcontactsOrganizationNotesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsOrganizationNotesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes forbidden response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes forbidden response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes forbidden response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes forbidden response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes forbidden response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsOrganizationNotesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesForbidden) String() string {
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

/*
GetExternalcontactsOrganizationNotesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsOrganizationNotesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes not found response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes not found response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes not found response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes not found response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes not found response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsOrganizationNotesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesNotFound) String() string {
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

// NewGetExternalcontactsOrganizationNotesRequestTimeout creates a GetExternalcontactsOrganizationNotesRequestTimeout with default headers values
func NewGetExternalcontactsOrganizationNotesRequestTimeout() *GetExternalcontactsOrganizationNotesRequestTimeout {
	return &GetExternalcontactsOrganizationNotesRequestTimeout{}
}

/*
GetExternalcontactsOrganizationNotesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsOrganizationNotesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes request timeout response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes request timeout response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes request timeout response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes request timeout response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes request timeout response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsOrganizationNotesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationNotesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetExternalcontactsOrganizationNotesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetExternalcontactsOrganizationNotesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes request entity too large response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes request entity too large response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes request entity too large response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes request entity too large response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes request entity too large response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesRequestEntityTooLarge) String() string {
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

/*
GetExternalcontactsOrganizationNotesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsOrganizationNotesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes unsupported media type response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes unsupported media type response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes unsupported media type response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes unsupported media type response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes unsupported media type response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesUnsupportedMediaType) String() string {
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

/*
GetExternalcontactsOrganizationNotesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsOrganizationNotesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes too many requests response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes too many requests response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes too many requests response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts organization notes too many requests response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts organization notes too many requests response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsOrganizationNotesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesTooManyRequests) String() string {
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

/*
GetExternalcontactsOrganizationNotesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsOrganizationNotesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes internal server error response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes internal server error response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes internal server error response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization notes internal server error response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organization notes internal server error response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsOrganizationNotesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesInternalServerError) String() string {
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

/*
GetExternalcontactsOrganizationNotesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsOrganizationNotesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes service unavailable response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes service unavailable response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes service unavailable response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization notes service unavailable response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organization notes service unavailable response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesServiceUnavailable) String() string {
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

/*
GetExternalcontactsOrganizationNotesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsOrganizationNotesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts organization notes gateway timeout response has a 2xx status code
func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts organization notes gateway timeout response has a 3xx status code
func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts organization notes gateway timeout response has a 4xx status code
func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts organization notes gateway timeout response has a 5xx status code
func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts organization notes gateway timeout response a status code equal to that given
func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/{externalOrganizationId}/notes][%d] getExternalcontactsOrganizationNotesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationNotesGatewayTimeout) String() string {
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
