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

// DeleteAuthorizationSubjectDivisionRoleReader is a Reader for the DeleteAuthorizationSubjectDivisionRole structure.
type DeleteAuthorizationSubjectDivisionRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthorizationSubjectDivisionRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteAuthorizationSubjectDivisionRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAuthorizationSubjectDivisionRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAuthorizationSubjectDivisionRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAuthorizationSubjectDivisionRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteAuthorizationSubjectDivisionRoleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAuthorizationSubjectDivisionRoleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAuthorizationSubjectDivisionRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAuthorizationSubjectDivisionRoleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteAuthorizationSubjectDivisionRoleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAuthorizationSubjectDivisionRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAuthorizationSubjectDivisionRoleBadRequest creates a DeleteAuthorizationSubjectDivisionRoleBadRequest with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleBadRequest() *DeleteAuthorizationSubjectDivisionRoleBadRequest {
	return &DeleteAuthorizationSubjectDivisionRoleBadRequest{}
}

/*DeleteAuthorizationSubjectDivisionRoleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteAuthorizationSubjectDivisionRoleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleUnauthorized creates a DeleteAuthorizationSubjectDivisionRoleUnauthorized with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleUnauthorized() *DeleteAuthorizationSubjectDivisionRoleUnauthorized {
	return &DeleteAuthorizationSubjectDivisionRoleUnauthorized{}
}

/*DeleteAuthorizationSubjectDivisionRoleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteAuthorizationSubjectDivisionRoleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleForbidden creates a DeleteAuthorizationSubjectDivisionRoleForbidden with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleForbidden() *DeleteAuthorizationSubjectDivisionRoleForbidden {
	return &DeleteAuthorizationSubjectDivisionRoleForbidden{}
}

/*DeleteAuthorizationSubjectDivisionRoleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteAuthorizationSubjectDivisionRoleForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleNotFound creates a DeleteAuthorizationSubjectDivisionRoleNotFound with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleNotFound() *DeleteAuthorizationSubjectDivisionRoleNotFound {
	return &DeleteAuthorizationSubjectDivisionRoleNotFound{}
}

/*DeleteAuthorizationSubjectDivisionRoleNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteAuthorizationSubjectDivisionRoleNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleRequestTimeout creates a DeleteAuthorizationSubjectDivisionRoleRequestTimeout with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleRequestTimeout() *DeleteAuthorizationSubjectDivisionRoleRequestTimeout {
	return &DeleteAuthorizationSubjectDivisionRoleRequestTimeout{}
}

/*DeleteAuthorizationSubjectDivisionRoleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteAuthorizationSubjectDivisionRoleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge creates a DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge() *DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge {
	return &DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge{}
}

/*DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType creates a DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType() *DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType {
	return &DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType{}
}

/*DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleTooManyRequests creates a DeleteAuthorizationSubjectDivisionRoleTooManyRequests with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleTooManyRequests() *DeleteAuthorizationSubjectDivisionRoleTooManyRequests {
	return &DeleteAuthorizationSubjectDivisionRoleTooManyRequests{}
}

/*DeleteAuthorizationSubjectDivisionRoleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteAuthorizationSubjectDivisionRoleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleInternalServerError creates a DeleteAuthorizationSubjectDivisionRoleInternalServerError with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleInternalServerError() *DeleteAuthorizationSubjectDivisionRoleInternalServerError {
	return &DeleteAuthorizationSubjectDivisionRoleInternalServerError{}
}

/*DeleteAuthorizationSubjectDivisionRoleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteAuthorizationSubjectDivisionRoleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleServiceUnavailable creates a DeleteAuthorizationSubjectDivisionRoleServiceUnavailable with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleServiceUnavailable() *DeleteAuthorizationSubjectDivisionRoleServiceUnavailable {
	return &DeleteAuthorizationSubjectDivisionRoleServiceUnavailable{}
}

/*DeleteAuthorizationSubjectDivisionRoleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteAuthorizationSubjectDivisionRoleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleGatewayTimeout creates a DeleteAuthorizationSubjectDivisionRoleGatewayTimeout with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleGatewayTimeout() *DeleteAuthorizationSubjectDivisionRoleGatewayTimeout {
	return &DeleteAuthorizationSubjectDivisionRoleGatewayTimeout{}
}

/*DeleteAuthorizationSubjectDivisionRoleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteAuthorizationSubjectDivisionRoleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAuthorizationSubjectDivisionRoleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRoleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteAuthorizationSubjectDivisionRoleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAuthorizationSubjectDivisionRoleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAuthorizationSubjectDivisionRoleDefault creates a DeleteAuthorizationSubjectDivisionRoleDefault with default headers values
func NewDeleteAuthorizationSubjectDivisionRoleDefault(code int) *DeleteAuthorizationSubjectDivisionRoleDefault {
	return &DeleteAuthorizationSubjectDivisionRoleDefault{
		_statusCode: code,
	}
}

/*DeleteAuthorizationSubjectDivisionRoleDefault handles this case with default header values.

successful operation
*/
type DeleteAuthorizationSubjectDivisionRoleDefault struct {
	_statusCode int
}

// Code gets the status code for the delete authorization subject division role default response
func (o *DeleteAuthorizationSubjectDivisionRoleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAuthorizationSubjectDivisionRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/authorization/subjects/{subjectId}/divisions/{divisionId}/roles/{roleId}][%d] deleteAuthorizationSubjectDivisionRole default ", o._statusCode)
}

func (o *DeleteAuthorizationSubjectDivisionRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
