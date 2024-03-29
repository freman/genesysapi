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

// PutIdentityprovidersOneloginReader is a Reader for the PutIdentityprovidersOnelogin structure.
type PutIdentityprovidersOneloginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersOneloginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersOneloginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersOneloginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersOneloginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersOneloginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersOneloginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutIdentityprovidersOneloginRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersOneloginRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersOneloginUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersOneloginTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersOneloginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersOneloginServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersOneloginGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersOneloginOK creates a PutIdentityprovidersOneloginOK with default headers values
func NewPutIdentityprovidersOneloginOK() *PutIdentityprovidersOneloginOK {
	return &PutIdentityprovidersOneloginOK{}
}

/*
PutIdentityprovidersOneloginOK describes a response with status code 200, with default header values.

successful operation
*/
type PutIdentityprovidersOneloginOK struct {
	Payload *models.OAuthProvider
}

// IsSuccess returns true when this put identityproviders onelogin o k response has a 2xx status code
func (o *PutIdentityprovidersOneloginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put identityproviders onelogin o k response has a 3xx status code
func (o *PutIdentityprovidersOneloginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin o k response has a 4xx status code
func (o *PutIdentityprovidersOneloginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders onelogin o k response has a 5xx status code
func (o *PutIdentityprovidersOneloginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin o k response a status code equal to that given
func (o *PutIdentityprovidersOneloginOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutIdentityprovidersOneloginOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersOneloginOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersOneloginOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginBadRequest creates a PutIdentityprovidersOneloginBadRequest with default headers values
func NewPutIdentityprovidersOneloginBadRequest() *PutIdentityprovidersOneloginBadRequest {
	return &PutIdentityprovidersOneloginBadRequest{}
}

/*
PutIdentityprovidersOneloginBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersOneloginBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin bad request response has a 2xx status code
func (o *PutIdentityprovidersOneloginBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin bad request response has a 3xx status code
func (o *PutIdentityprovidersOneloginBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin bad request response has a 4xx status code
func (o *PutIdentityprovidersOneloginBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin bad request response has a 5xx status code
func (o *PutIdentityprovidersOneloginBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin bad request response a status code equal to that given
func (o *PutIdentityprovidersOneloginBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutIdentityprovidersOneloginBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersOneloginBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersOneloginBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginUnauthorized creates a PutIdentityprovidersOneloginUnauthorized with default headers values
func NewPutIdentityprovidersOneloginUnauthorized() *PutIdentityprovidersOneloginUnauthorized {
	return &PutIdentityprovidersOneloginUnauthorized{}
}

/*
PutIdentityprovidersOneloginUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersOneloginUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin unauthorized response has a 2xx status code
func (o *PutIdentityprovidersOneloginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin unauthorized response has a 3xx status code
func (o *PutIdentityprovidersOneloginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin unauthorized response has a 4xx status code
func (o *PutIdentityprovidersOneloginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin unauthorized response has a 5xx status code
func (o *PutIdentityprovidersOneloginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin unauthorized response a status code equal to that given
func (o *PutIdentityprovidersOneloginUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutIdentityprovidersOneloginUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersOneloginUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersOneloginUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginForbidden creates a PutIdentityprovidersOneloginForbidden with default headers values
func NewPutIdentityprovidersOneloginForbidden() *PutIdentityprovidersOneloginForbidden {
	return &PutIdentityprovidersOneloginForbidden{}
}

/*
PutIdentityprovidersOneloginForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersOneloginForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin forbidden response has a 2xx status code
func (o *PutIdentityprovidersOneloginForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin forbidden response has a 3xx status code
func (o *PutIdentityprovidersOneloginForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin forbidden response has a 4xx status code
func (o *PutIdentityprovidersOneloginForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin forbidden response has a 5xx status code
func (o *PutIdentityprovidersOneloginForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin forbidden response a status code equal to that given
func (o *PutIdentityprovidersOneloginForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutIdentityprovidersOneloginForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersOneloginForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersOneloginForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginNotFound creates a PutIdentityprovidersOneloginNotFound with default headers values
func NewPutIdentityprovidersOneloginNotFound() *PutIdentityprovidersOneloginNotFound {
	return &PutIdentityprovidersOneloginNotFound{}
}

/*
PutIdentityprovidersOneloginNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersOneloginNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin not found response has a 2xx status code
func (o *PutIdentityprovidersOneloginNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin not found response has a 3xx status code
func (o *PutIdentityprovidersOneloginNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin not found response has a 4xx status code
func (o *PutIdentityprovidersOneloginNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin not found response has a 5xx status code
func (o *PutIdentityprovidersOneloginNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin not found response a status code equal to that given
func (o *PutIdentityprovidersOneloginNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutIdentityprovidersOneloginNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersOneloginNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersOneloginNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginRequestTimeout creates a PutIdentityprovidersOneloginRequestTimeout with default headers values
func NewPutIdentityprovidersOneloginRequestTimeout() *PutIdentityprovidersOneloginRequestTimeout {
	return &PutIdentityprovidersOneloginRequestTimeout{}
}

/*
PutIdentityprovidersOneloginRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutIdentityprovidersOneloginRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin request timeout response has a 2xx status code
func (o *PutIdentityprovidersOneloginRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin request timeout response has a 3xx status code
func (o *PutIdentityprovidersOneloginRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin request timeout response has a 4xx status code
func (o *PutIdentityprovidersOneloginRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin request timeout response has a 5xx status code
func (o *PutIdentityprovidersOneloginRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin request timeout response a status code equal to that given
func (o *PutIdentityprovidersOneloginRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutIdentityprovidersOneloginRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersOneloginRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersOneloginRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginRequestEntityTooLarge creates a PutIdentityprovidersOneloginRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersOneloginRequestEntityTooLarge() *PutIdentityprovidersOneloginRequestEntityTooLarge {
	return &PutIdentityprovidersOneloginRequestEntityTooLarge{}
}

/*
PutIdentityprovidersOneloginRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutIdentityprovidersOneloginRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin request entity too large response has a 2xx status code
func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin request entity too large response has a 3xx status code
func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin request entity too large response has a 4xx status code
func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin request entity too large response has a 5xx status code
func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin request entity too large response a status code equal to that given
func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginUnsupportedMediaType creates a PutIdentityprovidersOneloginUnsupportedMediaType with default headers values
func NewPutIdentityprovidersOneloginUnsupportedMediaType() *PutIdentityprovidersOneloginUnsupportedMediaType {
	return &PutIdentityprovidersOneloginUnsupportedMediaType{}
}

/*
PutIdentityprovidersOneloginUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersOneloginUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin unsupported media type response has a 2xx status code
func (o *PutIdentityprovidersOneloginUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin unsupported media type response has a 3xx status code
func (o *PutIdentityprovidersOneloginUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin unsupported media type response has a 4xx status code
func (o *PutIdentityprovidersOneloginUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin unsupported media type response has a 5xx status code
func (o *PutIdentityprovidersOneloginUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin unsupported media type response a status code equal to that given
func (o *PutIdentityprovidersOneloginUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutIdentityprovidersOneloginUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersOneloginUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersOneloginUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginTooManyRequests creates a PutIdentityprovidersOneloginTooManyRequests with default headers values
func NewPutIdentityprovidersOneloginTooManyRequests() *PutIdentityprovidersOneloginTooManyRequests {
	return &PutIdentityprovidersOneloginTooManyRequests{}
}

/*
PutIdentityprovidersOneloginTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutIdentityprovidersOneloginTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin too many requests response has a 2xx status code
func (o *PutIdentityprovidersOneloginTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin too many requests response has a 3xx status code
func (o *PutIdentityprovidersOneloginTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin too many requests response has a 4xx status code
func (o *PutIdentityprovidersOneloginTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders onelogin too many requests response has a 5xx status code
func (o *PutIdentityprovidersOneloginTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders onelogin too many requests response a status code equal to that given
func (o *PutIdentityprovidersOneloginTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutIdentityprovidersOneloginTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersOneloginTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersOneloginTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginInternalServerError creates a PutIdentityprovidersOneloginInternalServerError with default headers values
func NewPutIdentityprovidersOneloginInternalServerError() *PutIdentityprovidersOneloginInternalServerError {
	return &PutIdentityprovidersOneloginInternalServerError{}
}

/*
PutIdentityprovidersOneloginInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersOneloginInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin internal server error response has a 2xx status code
func (o *PutIdentityprovidersOneloginInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin internal server error response has a 3xx status code
func (o *PutIdentityprovidersOneloginInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin internal server error response has a 4xx status code
func (o *PutIdentityprovidersOneloginInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders onelogin internal server error response has a 5xx status code
func (o *PutIdentityprovidersOneloginInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders onelogin internal server error response a status code equal to that given
func (o *PutIdentityprovidersOneloginInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutIdentityprovidersOneloginInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersOneloginInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersOneloginInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginServiceUnavailable creates a PutIdentityprovidersOneloginServiceUnavailable with default headers values
func NewPutIdentityprovidersOneloginServiceUnavailable() *PutIdentityprovidersOneloginServiceUnavailable {
	return &PutIdentityprovidersOneloginServiceUnavailable{}
}

/*
PutIdentityprovidersOneloginServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersOneloginServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin service unavailable response has a 2xx status code
func (o *PutIdentityprovidersOneloginServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin service unavailable response has a 3xx status code
func (o *PutIdentityprovidersOneloginServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin service unavailable response has a 4xx status code
func (o *PutIdentityprovidersOneloginServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders onelogin service unavailable response has a 5xx status code
func (o *PutIdentityprovidersOneloginServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders onelogin service unavailable response a status code equal to that given
func (o *PutIdentityprovidersOneloginServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutIdentityprovidersOneloginServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersOneloginServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersOneloginServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersOneloginGatewayTimeout creates a PutIdentityprovidersOneloginGatewayTimeout with default headers values
func NewPutIdentityprovidersOneloginGatewayTimeout() *PutIdentityprovidersOneloginGatewayTimeout {
	return &PutIdentityprovidersOneloginGatewayTimeout{}
}

/*
PutIdentityprovidersOneloginGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutIdentityprovidersOneloginGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders onelogin gateway timeout response has a 2xx status code
func (o *PutIdentityprovidersOneloginGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders onelogin gateway timeout response has a 3xx status code
func (o *PutIdentityprovidersOneloginGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders onelogin gateway timeout response has a 4xx status code
func (o *PutIdentityprovidersOneloginGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders onelogin gateway timeout response has a 5xx status code
func (o *PutIdentityprovidersOneloginGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders onelogin gateway timeout response a status code equal to that given
func (o *PutIdentityprovidersOneloginGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutIdentityprovidersOneloginGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersOneloginGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/onelogin][%d] putIdentityprovidersOneloginGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersOneloginGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersOneloginGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
