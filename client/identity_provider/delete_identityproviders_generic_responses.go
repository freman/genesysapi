// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteIdentityprovidersGenericReader is a Reader for the DeleteIdentityprovidersGeneric structure.
type DeleteIdentityprovidersGenericReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityprovidersGenericReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityprovidersGenericOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityprovidersGenericBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIdentityprovidersGenericUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIdentityprovidersGenericForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIdentityprovidersGenericNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteIdentityprovidersGenericRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIdentityprovidersGenericRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIdentityprovidersGenericUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIdentityprovidersGenericTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIdentityprovidersGenericInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIdentityprovidersGenericServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIdentityprovidersGenericGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIdentityprovidersGenericOK creates a DeleteIdentityprovidersGenericOK with default headers values
func NewDeleteIdentityprovidersGenericOK() *DeleteIdentityprovidersGenericOK {
	return &DeleteIdentityprovidersGenericOK{}
}

/*
DeleteIdentityprovidersGenericOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteIdentityprovidersGenericOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete identityproviders generic o k response has a 2xx status code
func (o *DeleteIdentityprovidersGenericOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete identityproviders generic o k response has a 3xx status code
func (o *DeleteIdentityprovidersGenericOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic o k response has a 4xx status code
func (o *DeleteIdentityprovidersGenericOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders generic o k response has a 5xx status code
func (o *DeleteIdentityprovidersGenericOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic o k response a status code equal to that given
func (o *DeleteIdentityprovidersGenericOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteIdentityprovidersGenericOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersGenericOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersGenericOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericBadRequest creates a DeleteIdentityprovidersGenericBadRequest with default headers values
func NewDeleteIdentityprovidersGenericBadRequest() *DeleteIdentityprovidersGenericBadRequest {
	return &DeleteIdentityprovidersGenericBadRequest{}
}

/*
DeleteIdentityprovidersGenericBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIdentityprovidersGenericBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic bad request response has a 2xx status code
func (o *DeleteIdentityprovidersGenericBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic bad request response has a 3xx status code
func (o *DeleteIdentityprovidersGenericBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic bad request response has a 4xx status code
func (o *DeleteIdentityprovidersGenericBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic bad request response has a 5xx status code
func (o *DeleteIdentityprovidersGenericBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic bad request response a status code equal to that given
func (o *DeleteIdentityprovidersGenericBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteIdentityprovidersGenericBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersGenericBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersGenericBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericUnauthorized creates a DeleteIdentityprovidersGenericUnauthorized with default headers values
func NewDeleteIdentityprovidersGenericUnauthorized() *DeleteIdentityprovidersGenericUnauthorized {
	return &DeleteIdentityprovidersGenericUnauthorized{}
}

/*
DeleteIdentityprovidersGenericUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIdentityprovidersGenericUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic unauthorized response has a 2xx status code
func (o *DeleteIdentityprovidersGenericUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic unauthorized response has a 3xx status code
func (o *DeleteIdentityprovidersGenericUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic unauthorized response has a 4xx status code
func (o *DeleteIdentityprovidersGenericUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic unauthorized response has a 5xx status code
func (o *DeleteIdentityprovidersGenericUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic unauthorized response a status code equal to that given
func (o *DeleteIdentityprovidersGenericUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteIdentityprovidersGenericUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersGenericUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersGenericUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericForbidden creates a DeleteIdentityprovidersGenericForbidden with default headers values
func NewDeleteIdentityprovidersGenericForbidden() *DeleteIdentityprovidersGenericForbidden {
	return &DeleteIdentityprovidersGenericForbidden{}
}

/*
DeleteIdentityprovidersGenericForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIdentityprovidersGenericForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic forbidden response has a 2xx status code
func (o *DeleteIdentityprovidersGenericForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic forbidden response has a 3xx status code
func (o *DeleteIdentityprovidersGenericForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic forbidden response has a 4xx status code
func (o *DeleteIdentityprovidersGenericForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic forbidden response has a 5xx status code
func (o *DeleteIdentityprovidersGenericForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic forbidden response a status code equal to that given
func (o *DeleteIdentityprovidersGenericForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteIdentityprovidersGenericForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersGenericForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersGenericForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericNotFound creates a DeleteIdentityprovidersGenericNotFound with default headers values
func NewDeleteIdentityprovidersGenericNotFound() *DeleteIdentityprovidersGenericNotFound {
	return &DeleteIdentityprovidersGenericNotFound{}
}

/*
DeleteIdentityprovidersGenericNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteIdentityprovidersGenericNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic not found response has a 2xx status code
func (o *DeleteIdentityprovidersGenericNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic not found response has a 3xx status code
func (o *DeleteIdentityprovidersGenericNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic not found response has a 4xx status code
func (o *DeleteIdentityprovidersGenericNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic not found response has a 5xx status code
func (o *DeleteIdentityprovidersGenericNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic not found response a status code equal to that given
func (o *DeleteIdentityprovidersGenericNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteIdentityprovidersGenericNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersGenericNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersGenericNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericRequestTimeout creates a DeleteIdentityprovidersGenericRequestTimeout with default headers values
func NewDeleteIdentityprovidersGenericRequestTimeout() *DeleteIdentityprovidersGenericRequestTimeout {
	return &DeleteIdentityprovidersGenericRequestTimeout{}
}

/*
DeleteIdentityprovidersGenericRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteIdentityprovidersGenericRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic request timeout response has a 2xx status code
func (o *DeleteIdentityprovidersGenericRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic request timeout response has a 3xx status code
func (o *DeleteIdentityprovidersGenericRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic request timeout response has a 4xx status code
func (o *DeleteIdentityprovidersGenericRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic request timeout response has a 5xx status code
func (o *DeleteIdentityprovidersGenericRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic request timeout response a status code equal to that given
func (o *DeleteIdentityprovidersGenericRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteIdentityprovidersGenericRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersGenericRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersGenericRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericRequestEntityTooLarge creates a DeleteIdentityprovidersGenericRequestEntityTooLarge with default headers values
func NewDeleteIdentityprovidersGenericRequestEntityTooLarge() *DeleteIdentityprovidersGenericRequestEntityTooLarge {
	return &DeleteIdentityprovidersGenericRequestEntityTooLarge{}
}

/*
DeleteIdentityprovidersGenericRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteIdentityprovidersGenericRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic request entity too large response has a 2xx status code
func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic request entity too large response has a 3xx status code
func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic request entity too large response has a 4xx status code
func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic request entity too large response has a 5xx status code
func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic request entity too large response a status code equal to that given
func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericUnsupportedMediaType creates a DeleteIdentityprovidersGenericUnsupportedMediaType with default headers values
func NewDeleteIdentityprovidersGenericUnsupportedMediaType() *DeleteIdentityprovidersGenericUnsupportedMediaType {
	return &DeleteIdentityprovidersGenericUnsupportedMediaType{}
}

/*
DeleteIdentityprovidersGenericUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIdentityprovidersGenericUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic unsupported media type response has a 2xx status code
func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic unsupported media type response has a 3xx status code
func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic unsupported media type response has a 4xx status code
func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic unsupported media type response has a 5xx status code
func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic unsupported media type response a status code equal to that given
func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericTooManyRequests creates a DeleteIdentityprovidersGenericTooManyRequests with default headers values
func NewDeleteIdentityprovidersGenericTooManyRequests() *DeleteIdentityprovidersGenericTooManyRequests {
	return &DeleteIdentityprovidersGenericTooManyRequests{}
}

/*
DeleteIdentityprovidersGenericTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteIdentityprovidersGenericTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic too many requests response has a 2xx status code
func (o *DeleteIdentityprovidersGenericTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic too many requests response has a 3xx status code
func (o *DeleteIdentityprovidersGenericTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic too many requests response has a 4xx status code
func (o *DeleteIdentityprovidersGenericTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders generic too many requests response has a 5xx status code
func (o *DeleteIdentityprovidersGenericTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders generic too many requests response a status code equal to that given
func (o *DeleteIdentityprovidersGenericTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteIdentityprovidersGenericTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersGenericTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersGenericTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericInternalServerError creates a DeleteIdentityprovidersGenericInternalServerError with default headers values
func NewDeleteIdentityprovidersGenericInternalServerError() *DeleteIdentityprovidersGenericInternalServerError {
	return &DeleteIdentityprovidersGenericInternalServerError{}
}

/*
DeleteIdentityprovidersGenericInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIdentityprovidersGenericInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic internal server error response has a 2xx status code
func (o *DeleteIdentityprovidersGenericInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic internal server error response has a 3xx status code
func (o *DeleteIdentityprovidersGenericInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic internal server error response has a 4xx status code
func (o *DeleteIdentityprovidersGenericInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders generic internal server error response has a 5xx status code
func (o *DeleteIdentityprovidersGenericInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders generic internal server error response a status code equal to that given
func (o *DeleteIdentityprovidersGenericInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteIdentityprovidersGenericInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersGenericInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersGenericInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericServiceUnavailable creates a DeleteIdentityprovidersGenericServiceUnavailable with default headers values
func NewDeleteIdentityprovidersGenericServiceUnavailable() *DeleteIdentityprovidersGenericServiceUnavailable {
	return &DeleteIdentityprovidersGenericServiceUnavailable{}
}

/*
DeleteIdentityprovidersGenericServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIdentityprovidersGenericServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic service unavailable response has a 2xx status code
func (o *DeleteIdentityprovidersGenericServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic service unavailable response has a 3xx status code
func (o *DeleteIdentityprovidersGenericServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic service unavailable response has a 4xx status code
func (o *DeleteIdentityprovidersGenericServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders generic service unavailable response has a 5xx status code
func (o *DeleteIdentityprovidersGenericServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders generic service unavailable response a status code equal to that given
func (o *DeleteIdentityprovidersGenericServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteIdentityprovidersGenericServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersGenericServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersGenericServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGenericGatewayTimeout creates a DeleteIdentityprovidersGenericGatewayTimeout with default headers values
func NewDeleteIdentityprovidersGenericGatewayTimeout() *DeleteIdentityprovidersGenericGatewayTimeout {
	return &DeleteIdentityprovidersGenericGatewayTimeout{}
}

/*
DeleteIdentityprovidersGenericGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteIdentityprovidersGenericGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders generic gateway timeout response has a 2xx status code
func (o *DeleteIdentityprovidersGenericGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders generic gateway timeout response has a 3xx status code
func (o *DeleteIdentityprovidersGenericGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders generic gateway timeout response has a 4xx status code
func (o *DeleteIdentityprovidersGenericGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders generic gateway timeout response has a 5xx status code
func (o *DeleteIdentityprovidersGenericGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders generic gateway timeout response a status code equal to that given
func (o *DeleteIdentityprovidersGenericGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteIdentityprovidersGenericGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersGenericGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/generic][%d] deleteIdentityprovidersGenericGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersGenericGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGenericGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
