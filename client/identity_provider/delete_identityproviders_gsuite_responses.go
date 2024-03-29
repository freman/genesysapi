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

// DeleteIdentityprovidersGsuiteReader is a Reader for the DeleteIdentityprovidersGsuite structure.
type DeleteIdentityprovidersGsuiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIdentityprovidersGsuiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIdentityprovidersGsuiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIdentityprovidersGsuiteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIdentityprovidersGsuiteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIdentityprovidersGsuiteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIdentityprovidersGsuiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteIdentityprovidersGsuiteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIdentityprovidersGsuiteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIdentityprovidersGsuiteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIdentityprovidersGsuiteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIdentityprovidersGsuiteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIdentityprovidersGsuiteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIdentityprovidersGsuiteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIdentityprovidersGsuiteOK creates a DeleteIdentityprovidersGsuiteOK with default headers values
func NewDeleteIdentityprovidersGsuiteOK() *DeleteIdentityprovidersGsuiteOK {
	return &DeleteIdentityprovidersGsuiteOK{}
}

/*
DeleteIdentityprovidersGsuiteOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteIdentityprovidersGsuiteOK struct {
	Payload models.Empty
}

// IsSuccess returns true when this delete identityproviders gsuite o k response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete identityproviders gsuite o k response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite o k response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders gsuite o k response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite o k response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteIdentityprovidersGsuiteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteOK  %+v", 200, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteBadRequest creates a DeleteIdentityprovidersGsuiteBadRequest with default headers values
func NewDeleteIdentityprovidersGsuiteBadRequest() *DeleteIdentityprovidersGsuiteBadRequest {
	return &DeleteIdentityprovidersGsuiteBadRequest{}
}

/*
DeleteIdentityprovidersGsuiteBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIdentityprovidersGsuiteBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite bad request response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite bad request response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite bad request response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite bad request response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite bad request response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteIdentityprovidersGsuiteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteUnauthorized creates a DeleteIdentityprovidersGsuiteUnauthorized with default headers values
func NewDeleteIdentityprovidersGsuiteUnauthorized() *DeleteIdentityprovidersGsuiteUnauthorized {
	return &DeleteIdentityprovidersGsuiteUnauthorized{}
}

/*
DeleteIdentityprovidersGsuiteUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIdentityprovidersGsuiteUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite unauthorized response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite unauthorized response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite unauthorized response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite unauthorized response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite unauthorized response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteIdentityprovidersGsuiteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteForbidden creates a DeleteIdentityprovidersGsuiteForbidden with default headers values
func NewDeleteIdentityprovidersGsuiteForbidden() *DeleteIdentityprovidersGsuiteForbidden {
	return &DeleteIdentityprovidersGsuiteForbidden{}
}

/*
DeleteIdentityprovidersGsuiteForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIdentityprovidersGsuiteForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite forbidden response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite forbidden response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite forbidden response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite forbidden response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite forbidden response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteIdentityprovidersGsuiteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteNotFound creates a DeleteIdentityprovidersGsuiteNotFound with default headers values
func NewDeleteIdentityprovidersGsuiteNotFound() *DeleteIdentityprovidersGsuiteNotFound {
	return &DeleteIdentityprovidersGsuiteNotFound{}
}

/*
DeleteIdentityprovidersGsuiteNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteIdentityprovidersGsuiteNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite not found response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite not found response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite not found response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite not found response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite not found response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteIdentityprovidersGsuiteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteRequestTimeout creates a DeleteIdentityprovidersGsuiteRequestTimeout with default headers values
func NewDeleteIdentityprovidersGsuiteRequestTimeout() *DeleteIdentityprovidersGsuiteRequestTimeout {
	return &DeleteIdentityprovidersGsuiteRequestTimeout{}
}

/*
DeleteIdentityprovidersGsuiteRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteIdentityprovidersGsuiteRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite request timeout response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite request timeout response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite request timeout response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite request timeout response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite request timeout response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteIdentityprovidersGsuiteRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteRequestEntityTooLarge creates a DeleteIdentityprovidersGsuiteRequestEntityTooLarge with default headers values
func NewDeleteIdentityprovidersGsuiteRequestEntityTooLarge() *DeleteIdentityprovidersGsuiteRequestEntityTooLarge {
	return &DeleteIdentityprovidersGsuiteRequestEntityTooLarge{}
}

/*
DeleteIdentityprovidersGsuiteRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteIdentityprovidersGsuiteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite request entity too large response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite request entity too large response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite request entity too large response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite request entity too large response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite request entity too large response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteUnsupportedMediaType creates a DeleteIdentityprovidersGsuiteUnsupportedMediaType with default headers values
func NewDeleteIdentityprovidersGsuiteUnsupportedMediaType() *DeleteIdentityprovidersGsuiteUnsupportedMediaType {
	return &DeleteIdentityprovidersGsuiteUnsupportedMediaType{}
}

/*
DeleteIdentityprovidersGsuiteUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIdentityprovidersGsuiteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite unsupported media type response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite unsupported media type response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite unsupported media type response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite unsupported media type response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite unsupported media type response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteTooManyRequests creates a DeleteIdentityprovidersGsuiteTooManyRequests with default headers values
func NewDeleteIdentityprovidersGsuiteTooManyRequests() *DeleteIdentityprovidersGsuiteTooManyRequests {
	return &DeleteIdentityprovidersGsuiteTooManyRequests{}
}

/*
DeleteIdentityprovidersGsuiteTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteIdentityprovidersGsuiteTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite too many requests response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite too many requests response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite too many requests response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete identityproviders gsuite too many requests response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete identityproviders gsuite too many requests response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteIdentityprovidersGsuiteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteInternalServerError creates a DeleteIdentityprovidersGsuiteInternalServerError with default headers values
func NewDeleteIdentityprovidersGsuiteInternalServerError() *DeleteIdentityprovidersGsuiteInternalServerError {
	return &DeleteIdentityprovidersGsuiteInternalServerError{}
}

/*
DeleteIdentityprovidersGsuiteInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIdentityprovidersGsuiteInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite internal server error response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite internal server error response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite internal server error response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders gsuite internal server error response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders gsuite internal server error response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteIdentityprovidersGsuiteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteServiceUnavailable creates a DeleteIdentityprovidersGsuiteServiceUnavailable with default headers values
func NewDeleteIdentityprovidersGsuiteServiceUnavailable() *DeleteIdentityprovidersGsuiteServiceUnavailable {
	return &DeleteIdentityprovidersGsuiteServiceUnavailable{}
}

/*
DeleteIdentityprovidersGsuiteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIdentityprovidersGsuiteServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite service unavailable response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite service unavailable response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite service unavailable response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders gsuite service unavailable response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders gsuite service unavailable response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIdentityprovidersGsuiteGatewayTimeout creates a DeleteIdentityprovidersGsuiteGatewayTimeout with default headers values
func NewDeleteIdentityprovidersGsuiteGatewayTimeout() *DeleteIdentityprovidersGsuiteGatewayTimeout {
	return &DeleteIdentityprovidersGsuiteGatewayTimeout{}
}

/*
DeleteIdentityprovidersGsuiteGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteIdentityprovidersGsuiteGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete identityproviders gsuite gateway timeout response has a 2xx status code
func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete identityproviders gsuite gateway timeout response has a 3xx status code
func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete identityproviders gsuite gateway timeout response has a 4xx status code
func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete identityproviders gsuite gateway timeout response has a 5xx status code
func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete identityproviders gsuite gateway timeout response a status code equal to that given
func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/identityproviders/gsuite][%d] deleteIdentityprovidersGsuiteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIdentityprovidersGsuiteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
