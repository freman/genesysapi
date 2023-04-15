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

// GetAuthorizationPermissionsReader is a Reader for the GetAuthorizationPermissions structure.
type GetAuthorizationPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationPermissionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationPermissionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationPermissionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationPermissionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationPermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationPermissionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationPermissionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationPermissionsOK creates a GetAuthorizationPermissionsOK with default headers values
func NewGetAuthorizationPermissionsOK() *GetAuthorizationPermissionsOK {
	return &GetAuthorizationPermissionsOK{}
}

/*
GetAuthorizationPermissionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAuthorizationPermissionsOK struct {
	Payload *models.PermissionCollectionEntityListing
}

// IsSuccess returns true when this get authorization permissions o k response has a 2xx status code
func (o *GetAuthorizationPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get authorization permissions o k response has a 3xx status code
func (o *GetAuthorizationPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions o k response has a 4xx status code
func (o *GetAuthorizationPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization permissions o k response has a 5xx status code
func (o *GetAuthorizationPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions o k response a status code equal to that given
func (o *GetAuthorizationPermissionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAuthorizationPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationPermissionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationPermissionsOK) GetPayload() *models.PermissionCollectionEntityListing {
	return o.Payload
}

func (o *GetAuthorizationPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionCollectionEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsBadRequest creates a GetAuthorizationPermissionsBadRequest with default headers values
func NewGetAuthorizationPermissionsBadRequest() *GetAuthorizationPermissionsBadRequest {
	return &GetAuthorizationPermissionsBadRequest{}
}

/*
GetAuthorizationPermissionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationPermissionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions bad request response has a 2xx status code
func (o *GetAuthorizationPermissionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions bad request response has a 3xx status code
func (o *GetAuthorizationPermissionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions bad request response has a 4xx status code
func (o *GetAuthorizationPermissionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions bad request response has a 5xx status code
func (o *GetAuthorizationPermissionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions bad request response a status code equal to that given
func (o *GetAuthorizationPermissionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAuthorizationPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationPermissionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationPermissionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsUnauthorized creates a GetAuthorizationPermissionsUnauthorized with default headers values
func NewGetAuthorizationPermissionsUnauthorized() *GetAuthorizationPermissionsUnauthorized {
	return &GetAuthorizationPermissionsUnauthorized{}
}

/*
GetAuthorizationPermissionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationPermissionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions unauthorized response has a 2xx status code
func (o *GetAuthorizationPermissionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions unauthorized response has a 3xx status code
func (o *GetAuthorizationPermissionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions unauthorized response has a 4xx status code
func (o *GetAuthorizationPermissionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions unauthorized response has a 5xx status code
func (o *GetAuthorizationPermissionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions unauthorized response a status code equal to that given
func (o *GetAuthorizationPermissionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAuthorizationPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationPermissionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationPermissionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsForbidden creates a GetAuthorizationPermissionsForbidden with default headers values
func NewGetAuthorizationPermissionsForbidden() *GetAuthorizationPermissionsForbidden {
	return &GetAuthorizationPermissionsForbidden{}
}

/*
GetAuthorizationPermissionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationPermissionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions forbidden response has a 2xx status code
func (o *GetAuthorizationPermissionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions forbidden response has a 3xx status code
func (o *GetAuthorizationPermissionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions forbidden response has a 4xx status code
func (o *GetAuthorizationPermissionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions forbidden response has a 5xx status code
func (o *GetAuthorizationPermissionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions forbidden response a status code equal to that given
func (o *GetAuthorizationPermissionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAuthorizationPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationPermissionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationPermissionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsNotFound creates a GetAuthorizationPermissionsNotFound with default headers values
func NewGetAuthorizationPermissionsNotFound() *GetAuthorizationPermissionsNotFound {
	return &GetAuthorizationPermissionsNotFound{}
}

/*
GetAuthorizationPermissionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAuthorizationPermissionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions not found response has a 2xx status code
func (o *GetAuthorizationPermissionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions not found response has a 3xx status code
func (o *GetAuthorizationPermissionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions not found response has a 4xx status code
func (o *GetAuthorizationPermissionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions not found response has a 5xx status code
func (o *GetAuthorizationPermissionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions not found response a status code equal to that given
func (o *GetAuthorizationPermissionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAuthorizationPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationPermissionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationPermissionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsRequestTimeout creates a GetAuthorizationPermissionsRequestTimeout with default headers values
func NewGetAuthorizationPermissionsRequestTimeout() *GetAuthorizationPermissionsRequestTimeout {
	return &GetAuthorizationPermissionsRequestTimeout{}
}

/*
GetAuthorizationPermissionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationPermissionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions request timeout response has a 2xx status code
func (o *GetAuthorizationPermissionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions request timeout response has a 3xx status code
func (o *GetAuthorizationPermissionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions request timeout response has a 4xx status code
func (o *GetAuthorizationPermissionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions request timeout response has a 5xx status code
func (o *GetAuthorizationPermissionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions request timeout response a status code equal to that given
func (o *GetAuthorizationPermissionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAuthorizationPermissionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationPermissionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationPermissionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsRequestEntityTooLarge creates a GetAuthorizationPermissionsRequestEntityTooLarge with default headers values
func NewGetAuthorizationPermissionsRequestEntityTooLarge() *GetAuthorizationPermissionsRequestEntityTooLarge {
	return &GetAuthorizationPermissionsRequestEntityTooLarge{}
}

/*
GetAuthorizationPermissionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAuthorizationPermissionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions request entity too large response has a 2xx status code
func (o *GetAuthorizationPermissionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions request entity too large response has a 3xx status code
func (o *GetAuthorizationPermissionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions request entity too large response has a 4xx status code
func (o *GetAuthorizationPermissionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions request entity too large response has a 5xx status code
func (o *GetAuthorizationPermissionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions request entity too large response a status code equal to that given
func (o *GetAuthorizationPermissionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAuthorizationPermissionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationPermissionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationPermissionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsUnsupportedMediaType creates a GetAuthorizationPermissionsUnsupportedMediaType with default headers values
func NewGetAuthorizationPermissionsUnsupportedMediaType() *GetAuthorizationPermissionsUnsupportedMediaType {
	return &GetAuthorizationPermissionsUnsupportedMediaType{}
}

/*
GetAuthorizationPermissionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationPermissionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions unsupported media type response has a 2xx status code
func (o *GetAuthorizationPermissionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions unsupported media type response has a 3xx status code
func (o *GetAuthorizationPermissionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions unsupported media type response has a 4xx status code
func (o *GetAuthorizationPermissionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions unsupported media type response has a 5xx status code
func (o *GetAuthorizationPermissionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions unsupported media type response a status code equal to that given
func (o *GetAuthorizationPermissionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAuthorizationPermissionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationPermissionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationPermissionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsTooManyRequests creates a GetAuthorizationPermissionsTooManyRequests with default headers values
func NewGetAuthorizationPermissionsTooManyRequests() *GetAuthorizationPermissionsTooManyRequests {
	return &GetAuthorizationPermissionsTooManyRequests{}
}

/*
GetAuthorizationPermissionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationPermissionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions too many requests response has a 2xx status code
func (o *GetAuthorizationPermissionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions too many requests response has a 3xx status code
func (o *GetAuthorizationPermissionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions too many requests response has a 4xx status code
func (o *GetAuthorizationPermissionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization permissions too many requests response has a 5xx status code
func (o *GetAuthorizationPermissionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization permissions too many requests response a status code equal to that given
func (o *GetAuthorizationPermissionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAuthorizationPermissionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationPermissionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationPermissionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsInternalServerError creates a GetAuthorizationPermissionsInternalServerError with default headers values
func NewGetAuthorizationPermissionsInternalServerError() *GetAuthorizationPermissionsInternalServerError {
	return &GetAuthorizationPermissionsInternalServerError{}
}

/*
GetAuthorizationPermissionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationPermissionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions internal server error response has a 2xx status code
func (o *GetAuthorizationPermissionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions internal server error response has a 3xx status code
func (o *GetAuthorizationPermissionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions internal server error response has a 4xx status code
func (o *GetAuthorizationPermissionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization permissions internal server error response has a 5xx status code
func (o *GetAuthorizationPermissionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization permissions internal server error response a status code equal to that given
func (o *GetAuthorizationPermissionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAuthorizationPermissionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationPermissionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationPermissionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsServiceUnavailable creates a GetAuthorizationPermissionsServiceUnavailable with default headers values
func NewGetAuthorizationPermissionsServiceUnavailable() *GetAuthorizationPermissionsServiceUnavailable {
	return &GetAuthorizationPermissionsServiceUnavailable{}
}

/*
GetAuthorizationPermissionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationPermissionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions service unavailable response has a 2xx status code
func (o *GetAuthorizationPermissionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions service unavailable response has a 3xx status code
func (o *GetAuthorizationPermissionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions service unavailable response has a 4xx status code
func (o *GetAuthorizationPermissionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization permissions service unavailable response has a 5xx status code
func (o *GetAuthorizationPermissionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization permissions service unavailable response a status code equal to that given
func (o *GetAuthorizationPermissionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAuthorizationPermissionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationPermissionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationPermissionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationPermissionsGatewayTimeout creates a GetAuthorizationPermissionsGatewayTimeout with default headers values
func NewGetAuthorizationPermissionsGatewayTimeout() *GetAuthorizationPermissionsGatewayTimeout {
	return &GetAuthorizationPermissionsGatewayTimeout{}
}

/*
GetAuthorizationPermissionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAuthorizationPermissionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization permissions gateway timeout response has a 2xx status code
func (o *GetAuthorizationPermissionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization permissions gateway timeout response has a 3xx status code
func (o *GetAuthorizationPermissionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization permissions gateway timeout response has a 4xx status code
func (o *GetAuthorizationPermissionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization permissions gateway timeout response has a 5xx status code
func (o *GetAuthorizationPermissionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization permissions gateway timeout response a status code equal to that given
func (o *GetAuthorizationPermissionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAuthorizationPermissionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationPermissionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/permissions][%d] getAuthorizationPermissionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationPermissionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationPermissionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
