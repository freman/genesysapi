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

// GetUserRolesReader is a Reader for the GetUserRoles structure.
type GetUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserRolesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserRolesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserRolesOK creates a GetUserRolesOK with default headers values
func NewGetUserRolesOK() *GetUserRolesOK {
	return &GetUserRolesOK{}
}

/*
GetUserRolesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserRolesOK struct {
	Payload *models.UserAuthorization
}

// IsSuccess returns true when this get user roles o k response has a 2xx status code
func (o *GetUserRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user roles o k response has a 3xx status code
func (o *GetUserRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles o k response has a 4xx status code
func (o *GetUserRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user roles o k response has a 5xx status code
func (o *GetUserRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles o k response a status code equal to that given
func (o *GetUserRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesOK  %+v", 200, o.Payload)
}

func (o *GetUserRolesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesOK  %+v", 200, o.Payload)
}

func (o *GetUserRolesOK) GetPayload() *models.UserAuthorization {
	return o.Payload
}

func (o *GetUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAuthorization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesBadRequest creates a GetUserRolesBadRequest with default headers values
func NewGetUserRolesBadRequest() *GetUserRolesBadRequest {
	return &GetUserRolesBadRequest{}
}

/*
GetUserRolesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserRolesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles bad request response has a 2xx status code
func (o *GetUserRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles bad request response has a 3xx status code
func (o *GetUserRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles bad request response has a 4xx status code
func (o *GetUserRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles bad request response has a 5xx status code
func (o *GetUserRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles bad request response a status code equal to that given
func (o *GetUserRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserRolesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesUnauthorized creates a GetUserRolesUnauthorized with default headers values
func NewGetUserRolesUnauthorized() *GetUserRolesUnauthorized {
	return &GetUserRolesUnauthorized{}
}

/*
GetUserRolesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserRolesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles unauthorized response has a 2xx status code
func (o *GetUserRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles unauthorized response has a 3xx status code
func (o *GetUserRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles unauthorized response has a 4xx status code
func (o *GetUserRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles unauthorized response has a 5xx status code
func (o *GetUserRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles unauthorized response a status code equal to that given
func (o *GetUserRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserRolesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesForbidden creates a GetUserRolesForbidden with default headers values
func NewGetUserRolesForbidden() *GetUserRolesForbidden {
	return &GetUserRolesForbidden{}
}

/*
GetUserRolesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserRolesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles forbidden response has a 2xx status code
func (o *GetUserRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles forbidden response has a 3xx status code
func (o *GetUserRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles forbidden response has a 4xx status code
func (o *GetUserRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles forbidden response has a 5xx status code
func (o *GetUserRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles forbidden response a status code equal to that given
func (o *GetUserRolesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserRolesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserRolesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesNotFound creates a GetUserRolesNotFound with default headers values
func NewGetUserRolesNotFound() *GetUserRolesNotFound {
	return &GetUserRolesNotFound{}
}

/*
GetUserRolesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserRolesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles not found response has a 2xx status code
func (o *GetUserRolesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles not found response has a 3xx status code
func (o *GetUserRolesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles not found response has a 4xx status code
func (o *GetUserRolesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles not found response has a 5xx status code
func (o *GetUserRolesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles not found response a status code equal to that given
func (o *GetUserRolesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserRolesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserRolesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesRequestTimeout creates a GetUserRolesRequestTimeout with default headers values
func NewGetUserRolesRequestTimeout() *GetUserRolesRequestTimeout {
	return &GetUserRolesRequestTimeout{}
}

/*
GetUserRolesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserRolesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles request timeout response has a 2xx status code
func (o *GetUserRolesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles request timeout response has a 3xx status code
func (o *GetUserRolesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles request timeout response has a 4xx status code
func (o *GetUserRolesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles request timeout response has a 5xx status code
func (o *GetUserRolesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles request timeout response a status code equal to that given
func (o *GetUserRolesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserRolesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserRolesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserRolesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesRequestEntityTooLarge creates a GetUserRolesRequestEntityTooLarge with default headers values
func NewGetUserRolesRequestEntityTooLarge() *GetUserRolesRequestEntityTooLarge {
	return &GetUserRolesRequestEntityTooLarge{}
}

/*
GetUserRolesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUserRolesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles request entity too large response has a 2xx status code
func (o *GetUserRolesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles request entity too large response has a 3xx status code
func (o *GetUserRolesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles request entity too large response has a 4xx status code
func (o *GetUserRolesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles request entity too large response has a 5xx status code
func (o *GetUserRolesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles request entity too large response a status code equal to that given
func (o *GetUserRolesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserRolesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserRolesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserRolesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesUnsupportedMediaType creates a GetUserRolesUnsupportedMediaType with default headers values
func NewGetUserRolesUnsupportedMediaType() *GetUserRolesUnsupportedMediaType {
	return &GetUserRolesUnsupportedMediaType{}
}

/*
GetUserRolesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserRolesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles unsupported media type response has a 2xx status code
func (o *GetUserRolesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles unsupported media type response has a 3xx status code
func (o *GetUserRolesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles unsupported media type response has a 4xx status code
func (o *GetUserRolesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles unsupported media type response has a 5xx status code
func (o *GetUserRolesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles unsupported media type response a status code equal to that given
func (o *GetUserRolesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserRolesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserRolesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesTooManyRequests creates a GetUserRolesTooManyRequests with default headers values
func NewGetUserRolesTooManyRequests() *GetUserRolesTooManyRequests {
	return &GetUserRolesTooManyRequests{}
}

/*
GetUserRolesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserRolesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles too many requests response has a 2xx status code
func (o *GetUserRolesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles too many requests response has a 3xx status code
func (o *GetUserRolesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles too many requests response has a 4xx status code
func (o *GetUserRolesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user roles too many requests response has a 5xx status code
func (o *GetUserRolesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user roles too many requests response a status code equal to that given
func (o *GetUserRolesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserRolesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserRolesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesInternalServerError creates a GetUserRolesInternalServerError with default headers values
func NewGetUserRolesInternalServerError() *GetUserRolesInternalServerError {
	return &GetUserRolesInternalServerError{}
}

/*
GetUserRolesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserRolesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles internal server error response has a 2xx status code
func (o *GetUserRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles internal server error response has a 3xx status code
func (o *GetUserRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles internal server error response has a 4xx status code
func (o *GetUserRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user roles internal server error response has a 5xx status code
func (o *GetUserRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user roles internal server error response a status code equal to that given
func (o *GetUserRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserRolesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserRolesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesServiceUnavailable creates a GetUserRolesServiceUnavailable with default headers values
func NewGetUserRolesServiceUnavailable() *GetUserRolesServiceUnavailable {
	return &GetUserRolesServiceUnavailable{}
}

/*
GetUserRolesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserRolesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles service unavailable response has a 2xx status code
func (o *GetUserRolesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles service unavailable response has a 3xx status code
func (o *GetUserRolesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles service unavailable response has a 4xx status code
func (o *GetUserRolesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user roles service unavailable response has a 5xx status code
func (o *GetUserRolesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user roles service unavailable response a status code equal to that given
func (o *GetUserRolesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserRolesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserRolesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesGatewayTimeout creates a GetUserRolesGatewayTimeout with default headers values
func NewGetUserRolesGatewayTimeout() *GetUserRolesGatewayTimeout {
	return &GetUserRolesGatewayTimeout{}
}

/*
GetUserRolesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserRolesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user roles gateway timeout response has a 2xx status code
func (o *GetUserRolesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user roles gateway timeout response has a 3xx status code
func (o *GetUserRolesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user roles gateway timeout response has a 4xx status code
func (o *GetUserRolesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user roles gateway timeout response has a 5xx status code
func (o *GetUserRolesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user roles gateway timeout response a status code equal to that given
func (o *GetUserRolesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserRolesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{subjectId}/roles][%d] getUserRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserRolesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
