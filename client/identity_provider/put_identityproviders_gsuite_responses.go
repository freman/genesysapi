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

// PutIdentityprovidersGsuiteReader is a Reader for the PutIdentityprovidersGsuite structure.
type PutIdentityprovidersGsuiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersGsuiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersGsuiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersGsuiteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersGsuiteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersGsuiteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersGsuiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutIdentityprovidersGsuiteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersGsuiteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersGsuiteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersGsuiteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersGsuiteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersGsuiteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersGsuiteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersGsuiteOK creates a PutIdentityprovidersGsuiteOK with default headers values
func NewPutIdentityprovidersGsuiteOK() *PutIdentityprovidersGsuiteOK {
	return &PutIdentityprovidersGsuiteOK{}
}

/*
PutIdentityprovidersGsuiteOK describes a response with status code 200, with default header values.

successful operation
*/
type PutIdentityprovidersGsuiteOK struct {
	Payload *models.OAuthProvider
}

// IsSuccess returns true when this put identityproviders gsuite o k response has a 2xx status code
func (o *PutIdentityprovidersGsuiteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put identityproviders gsuite o k response has a 3xx status code
func (o *PutIdentityprovidersGsuiteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite o k response has a 4xx status code
func (o *PutIdentityprovidersGsuiteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders gsuite o k response has a 5xx status code
func (o *PutIdentityprovidersGsuiteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite o k response a status code equal to that given
func (o *PutIdentityprovidersGsuiteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutIdentityprovidersGsuiteOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersGsuiteOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersGsuiteOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteBadRequest creates a PutIdentityprovidersGsuiteBadRequest with default headers values
func NewPutIdentityprovidersGsuiteBadRequest() *PutIdentityprovidersGsuiteBadRequest {
	return &PutIdentityprovidersGsuiteBadRequest{}
}

/*
PutIdentityprovidersGsuiteBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersGsuiteBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite bad request response has a 2xx status code
func (o *PutIdentityprovidersGsuiteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite bad request response has a 3xx status code
func (o *PutIdentityprovidersGsuiteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite bad request response has a 4xx status code
func (o *PutIdentityprovidersGsuiteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite bad request response has a 5xx status code
func (o *PutIdentityprovidersGsuiteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite bad request response a status code equal to that given
func (o *PutIdentityprovidersGsuiteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutIdentityprovidersGsuiteBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersGsuiteBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersGsuiteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteUnauthorized creates a PutIdentityprovidersGsuiteUnauthorized with default headers values
func NewPutIdentityprovidersGsuiteUnauthorized() *PutIdentityprovidersGsuiteUnauthorized {
	return &PutIdentityprovidersGsuiteUnauthorized{}
}

/*
PutIdentityprovidersGsuiteUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersGsuiteUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite unauthorized response has a 2xx status code
func (o *PutIdentityprovidersGsuiteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite unauthorized response has a 3xx status code
func (o *PutIdentityprovidersGsuiteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite unauthorized response has a 4xx status code
func (o *PutIdentityprovidersGsuiteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite unauthorized response has a 5xx status code
func (o *PutIdentityprovidersGsuiteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite unauthorized response a status code equal to that given
func (o *PutIdentityprovidersGsuiteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutIdentityprovidersGsuiteUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersGsuiteUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersGsuiteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteForbidden creates a PutIdentityprovidersGsuiteForbidden with default headers values
func NewPutIdentityprovidersGsuiteForbidden() *PutIdentityprovidersGsuiteForbidden {
	return &PutIdentityprovidersGsuiteForbidden{}
}

/*
PutIdentityprovidersGsuiteForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersGsuiteForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite forbidden response has a 2xx status code
func (o *PutIdentityprovidersGsuiteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite forbidden response has a 3xx status code
func (o *PutIdentityprovidersGsuiteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite forbidden response has a 4xx status code
func (o *PutIdentityprovidersGsuiteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite forbidden response has a 5xx status code
func (o *PutIdentityprovidersGsuiteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite forbidden response a status code equal to that given
func (o *PutIdentityprovidersGsuiteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutIdentityprovidersGsuiteForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersGsuiteForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersGsuiteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteNotFound creates a PutIdentityprovidersGsuiteNotFound with default headers values
func NewPutIdentityprovidersGsuiteNotFound() *PutIdentityprovidersGsuiteNotFound {
	return &PutIdentityprovidersGsuiteNotFound{}
}

/*
PutIdentityprovidersGsuiteNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersGsuiteNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite not found response has a 2xx status code
func (o *PutIdentityprovidersGsuiteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite not found response has a 3xx status code
func (o *PutIdentityprovidersGsuiteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite not found response has a 4xx status code
func (o *PutIdentityprovidersGsuiteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite not found response has a 5xx status code
func (o *PutIdentityprovidersGsuiteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite not found response a status code equal to that given
func (o *PutIdentityprovidersGsuiteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutIdentityprovidersGsuiteNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersGsuiteNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersGsuiteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteRequestTimeout creates a PutIdentityprovidersGsuiteRequestTimeout with default headers values
func NewPutIdentityprovidersGsuiteRequestTimeout() *PutIdentityprovidersGsuiteRequestTimeout {
	return &PutIdentityprovidersGsuiteRequestTimeout{}
}

/*
PutIdentityprovidersGsuiteRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutIdentityprovidersGsuiteRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite request timeout response has a 2xx status code
func (o *PutIdentityprovidersGsuiteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite request timeout response has a 3xx status code
func (o *PutIdentityprovidersGsuiteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite request timeout response has a 4xx status code
func (o *PutIdentityprovidersGsuiteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite request timeout response has a 5xx status code
func (o *PutIdentityprovidersGsuiteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite request timeout response a status code equal to that given
func (o *PutIdentityprovidersGsuiteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutIdentityprovidersGsuiteRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersGsuiteRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersGsuiteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteRequestEntityTooLarge creates a PutIdentityprovidersGsuiteRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersGsuiteRequestEntityTooLarge() *PutIdentityprovidersGsuiteRequestEntityTooLarge {
	return &PutIdentityprovidersGsuiteRequestEntityTooLarge{}
}

/*
PutIdentityprovidersGsuiteRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutIdentityprovidersGsuiteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite request entity too large response has a 2xx status code
func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite request entity too large response has a 3xx status code
func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite request entity too large response has a 4xx status code
func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite request entity too large response has a 5xx status code
func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite request entity too large response a status code equal to that given
func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteUnsupportedMediaType creates a PutIdentityprovidersGsuiteUnsupportedMediaType with default headers values
func NewPutIdentityprovidersGsuiteUnsupportedMediaType() *PutIdentityprovidersGsuiteUnsupportedMediaType {
	return &PutIdentityprovidersGsuiteUnsupportedMediaType{}
}

/*
PutIdentityprovidersGsuiteUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersGsuiteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite unsupported media type response has a 2xx status code
func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite unsupported media type response has a 3xx status code
func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite unsupported media type response has a 4xx status code
func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite unsupported media type response has a 5xx status code
func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite unsupported media type response a status code equal to that given
func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteTooManyRequests creates a PutIdentityprovidersGsuiteTooManyRequests with default headers values
func NewPutIdentityprovidersGsuiteTooManyRequests() *PutIdentityprovidersGsuiteTooManyRequests {
	return &PutIdentityprovidersGsuiteTooManyRequests{}
}

/*
PutIdentityprovidersGsuiteTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutIdentityprovidersGsuiteTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite too many requests response has a 2xx status code
func (o *PutIdentityprovidersGsuiteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite too many requests response has a 3xx status code
func (o *PutIdentityprovidersGsuiteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite too many requests response has a 4xx status code
func (o *PutIdentityprovidersGsuiteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders gsuite too many requests response has a 5xx status code
func (o *PutIdentityprovidersGsuiteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders gsuite too many requests response a status code equal to that given
func (o *PutIdentityprovidersGsuiteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutIdentityprovidersGsuiteTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersGsuiteTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersGsuiteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteInternalServerError creates a PutIdentityprovidersGsuiteInternalServerError with default headers values
func NewPutIdentityprovidersGsuiteInternalServerError() *PutIdentityprovidersGsuiteInternalServerError {
	return &PutIdentityprovidersGsuiteInternalServerError{}
}

/*
PutIdentityprovidersGsuiteInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersGsuiteInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite internal server error response has a 2xx status code
func (o *PutIdentityprovidersGsuiteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite internal server error response has a 3xx status code
func (o *PutIdentityprovidersGsuiteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite internal server error response has a 4xx status code
func (o *PutIdentityprovidersGsuiteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders gsuite internal server error response has a 5xx status code
func (o *PutIdentityprovidersGsuiteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders gsuite internal server error response a status code equal to that given
func (o *PutIdentityprovidersGsuiteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutIdentityprovidersGsuiteInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersGsuiteInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersGsuiteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteServiceUnavailable creates a PutIdentityprovidersGsuiteServiceUnavailable with default headers values
func NewPutIdentityprovidersGsuiteServiceUnavailable() *PutIdentityprovidersGsuiteServiceUnavailable {
	return &PutIdentityprovidersGsuiteServiceUnavailable{}
}

/*
PutIdentityprovidersGsuiteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersGsuiteServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite service unavailable response has a 2xx status code
func (o *PutIdentityprovidersGsuiteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite service unavailable response has a 3xx status code
func (o *PutIdentityprovidersGsuiteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite service unavailable response has a 4xx status code
func (o *PutIdentityprovidersGsuiteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders gsuite service unavailable response has a 5xx status code
func (o *PutIdentityprovidersGsuiteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders gsuite service unavailable response a status code equal to that given
func (o *PutIdentityprovidersGsuiteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutIdentityprovidersGsuiteServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersGsuiteServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersGsuiteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteGatewayTimeout creates a PutIdentityprovidersGsuiteGatewayTimeout with default headers values
func NewPutIdentityprovidersGsuiteGatewayTimeout() *PutIdentityprovidersGsuiteGatewayTimeout {
	return &PutIdentityprovidersGsuiteGatewayTimeout{}
}

/*
PutIdentityprovidersGsuiteGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutIdentityprovidersGsuiteGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders gsuite gateway timeout response has a 2xx status code
func (o *PutIdentityprovidersGsuiteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders gsuite gateway timeout response has a 3xx status code
func (o *PutIdentityprovidersGsuiteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders gsuite gateway timeout response has a 4xx status code
func (o *PutIdentityprovidersGsuiteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders gsuite gateway timeout response has a 5xx status code
func (o *PutIdentityprovidersGsuiteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders gsuite gateway timeout response a status code equal to that given
func (o *PutIdentityprovidersGsuiteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutIdentityprovidersGsuiteGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersGsuiteGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersGsuiteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
