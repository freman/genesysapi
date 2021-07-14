// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAuthorizationSubjectDivisionRoleReader is a Reader for the PostAuthorizationSubjectDivisionRole structure.
type PostAuthorizationSubjectDivisionRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthorizationSubjectDivisionRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostAuthorizationSubjectDivisionRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuthorizationSubjectDivisionRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuthorizationSubjectDivisionRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuthorizationSubjectDivisionRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAuthorizationSubjectDivisionRoleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuthorizationSubjectDivisionRoleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuthorizationSubjectDivisionRoleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuthorizationSubjectDivisionRoleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthorizationSubjectDivisionRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuthorizationSubjectDivisionRoleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuthorizationSubjectDivisionRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostAuthorizationSubjectDivisionRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostAuthorizationSubjectDivisionRoleBadRequest creates a PostAuthorizationSubjectDivisionRoleBadRequest with default headers values
func NewPostAuthorizationSubjectDivisionRoleBadRequest() *PostAuthorizationSubjectDivisionRoleBadRequest {
	return &PostAuthorizationSubjectDivisionRoleBadRequest{}
}

/*PostAuthorizationSubjectDivisionRoleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuthorizationSubjectDivisionRoleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleUnauthorized creates a PostAuthorizationSubjectDivisionRoleUnauthorized with default headers values
func NewPostAuthorizationSubjectDivisionRoleUnauthorized() *PostAuthorizationSubjectDivisionRoleUnauthorized {
	return &PostAuthorizationSubjectDivisionRoleUnauthorized{}
}

/*PostAuthorizationSubjectDivisionRoleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuthorizationSubjectDivisionRoleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleForbidden creates a PostAuthorizationSubjectDivisionRoleForbidden with default headers values
func NewPostAuthorizationSubjectDivisionRoleForbidden() *PostAuthorizationSubjectDivisionRoleForbidden {
	return &PostAuthorizationSubjectDivisionRoleForbidden{}
}

/*PostAuthorizationSubjectDivisionRoleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAuthorizationSubjectDivisionRoleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleForbidden  %+v", 403, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleNotFound creates a PostAuthorizationSubjectDivisionRoleNotFound with default headers values
func NewPostAuthorizationSubjectDivisionRoleNotFound() *PostAuthorizationSubjectDivisionRoleNotFound {
	return &PostAuthorizationSubjectDivisionRoleNotFound{}
}

/*PostAuthorizationSubjectDivisionRoleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAuthorizationSubjectDivisionRoleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleNotFound  %+v", 404, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleRequestTimeout creates a PostAuthorizationSubjectDivisionRoleRequestTimeout with default headers values
func NewPostAuthorizationSubjectDivisionRoleRequestTimeout() *PostAuthorizationSubjectDivisionRoleRequestTimeout {
	return &PostAuthorizationSubjectDivisionRoleRequestTimeout{}
}

/*PostAuthorizationSubjectDivisionRoleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAuthorizationSubjectDivisionRoleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleRequestEntityTooLarge creates a PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge with default headers values
func NewPostAuthorizationSubjectDivisionRoleRequestEntityTooLarge() *PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge {
	return &PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge{}
}

/*PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleUnsupportedMediaType creates a PostAuthorizationSubjectDivisionRoleUnsupportedMediaType with default headers values
func NewPostAuthorizationSubjectDivisionRoleUnsupportedMediaType() *PostAuthorizationSubjectDivisionRoleUnsupportedMediaType {
	return &PostAuthorizationSubjectDivisionRoleUnsupportedMediaType{}
}

/*PostAuthorizationSubjectDivisionRoleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuthorizationSubjectDivisionRoleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleTooManyRequests creates a PostAuthorizationSubjectDivisionRoleTooManyRequests with default headers values
func NewPostAuthorizationSubjectDivisionRoleTooManyRequests() *PostAuthorizationSubjectDivisionRoleTooManyRequests {
	return &PostAuthorizationSubjectDivisionRoleTooManyRequests{}
}

/*PostAuthorizationSubjectDivisionRoleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAuthorizationSubjectDivisionRoleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleInternalServerError creates a PostAuthorizationSubjectDivisionRoleInternalServerError with default headers values
func NewPostAuthorizationSubjectDivisionRoleInternalServerError() *PostAuthorizationSubjectDivisionRoleInternalServerError {
	return &PostAuthorizationSubjectDivisionRoleInternalServerError{}
}

/*PostAuthorizationSubjectDivisionRoleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuthorizationSubjectDivisionRoleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleServiceUnavailable creates a PostAuthorizationSubjectDivisionRoleServiceUnavailable with default headers values
func NewPostAuthorizationSubjectDivisionRoleServiceUnavailable() *PostAuthorizationSubjectDivisionRoleServiceUnavailable {
	return &PostAuthorizationSubjectDivisionRoleServiceUnavailable{}
}

/*PostAuthorizationSubjectDivisionRoleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuthorizationSubjectDivisionRoleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleGatewayTimeout creates a PostAuthorizationSubjectDivisionRoleGatewayTimeout with default headers values
func NewPostAuthorizationSubjectDivisionRoleGatewayTimeout() *PostAuthorizationSubjectDivisionRoleGatewayTimeout {
	return &PostAuthorizationSubjectDivisionRoleGatewayTimeout{}
}

/*PostAuthorizationSubjectDivisionRoleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAuthorizationSubjectDivisionRoleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectDivisionRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuthorizationSubjectDivisionRoleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectDivisionRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectDivisionRoleDefault creates a PostAuthorizationSubjectDivisionRoleDefault with default headers values
func NewPostAuthorizationSubjectDivisionRoleDefault(code int) *PostAuthorizationSubjectDivisionRoleDefault {
	return &PostAuthorizationSubjectDivisionRoleDefault{
		_statusCode: code,
	}
}

/*PostAuthorizationSubjectDivisionRoleDefault handles this case with default header values.

successful operation
*/
type PostAuthorizationSubjectDivisionRoleDefault struct {
	_statusCode int
}

// Code gets the status code for the post authorization subject division role default response
func (o *PostAuthorizationSubjectDivisionRoleDefault) Code() int {
	return o._statusCode
}

func (o *PostAuthorizationSubjectDivisionRoleDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] postAuthorizationSubjectDivisionRole default ", o._statusCode)
}

func (o *PostAuthorizationSubjectDivisionRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
