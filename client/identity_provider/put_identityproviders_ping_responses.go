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

// PutIdentityprovidersPingReader is a Reader for the PutIdentityprovidersPing structure.
type PutIdentityprovidersPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersPingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersPingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersPingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutIdentityprovidersPingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersPingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersPingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersPingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersPingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersPingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersPingOK creates a PutIdentityprovidersPingOK with default headers values
func NewPutIdentityprovidersPingOK() *PutIdentityprovidersPingOK {
	return &PutIdentityprovidersPingOK{}
}

/*
PutIdentityprovidersPingOK describes a response with status code 200, with default header values.

successful operation
*/
type PutIdentityprovidersPingOK struct {
	Payload *models.OAuthProvider
}

// IsSuccess returns true when this put identityproviders ping o k response has a 2xx status code
func (o *PutIdentityprovidersPingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put identityproviders ping o k response has a 3xx status code
func (o *PutIdentityprovidersPingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping o k response has a 4xx status code
func (o *PutIdentityprovidersPingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders ping o k response has a 5xx status code
func (o *PutIdentityprovidersPingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping o k response a status code equal to that given
func (o *PutIdentityprovidersPingOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutIdentityprovidersPingOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersPingOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersPingOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingBadRequest creates a PutIdentityprovidersPingBadRequest with default headers values
func NewPutIdentityprovidersPingBadRequest() *PutIdentityprovidersPingBadRequest {
	return &PutIdentityprovidersPingBadRequest{}
}

/*
PutIdentityprovidersPingBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersPingBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping bad request response has a 2xx status code
func (o *PutIdentityprovidersPingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping bad request response has a 3xx status code
func (o *PutIdentityprovidersPingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping bad request response has a 4xx status code
func (o *PutIdentityprovidersPingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping bad request response has a 5xx status code
func (o *PutIdentityprovidersPingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping bad request response a status code equal to that given
func (o *PutIdentityprovidersPingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutIdentityprovidersPingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersPingBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersPingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingUnauthorized creates a PutIdentityprovidersPingUnauthorized with default headers values
func NewPutIdentityprovidersPingUnauthorized() *PutIdentityprovidersPingUnauthorized {
	return &PutIdentityprovidersPingUnauthorized{}
}

/*
PutIdentityprovidersPingUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersPingUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping unauthorized response has a 2xx status code
func (o *PutIdentityprovidersPingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping unauthorized response has a 3xx status code
func (o *PutIdentityprovidersPingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping unauthorized response has a 4xx status code
func (o *PutIdentityprovidersPingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping unauthorized response has a 5xx status code
func (o *PutIdentityprovidersPingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping unauthorized response a status code equal to that given
func (o *PutIdentityprovidersPingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutIdentityprovidersPingUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersPingUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersPingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingForbidden creates a PutIdentityprovidersPingForbidden with default headers values
func NewPutIdentityprovidersPingForbidden() *PutIdentityprovidersPingForbidden {
	return &PutIdentityprovidersPingForbidden{}
}

/*
PutIdentityprovidersPingForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersPingForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping forbidden response has a 2xx status code
func (o *PutIdentityprovidersPingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping forbidden response has a 3xx status code
func (o *PutIdentityprovidersPingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping forbidden response has a 4xx status code
func (o *PutIdentityprovidersPingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping forbidden response has a 5xx status code
func (o *PutIdentityprovidersPingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping forbidden response a status code equal to that given
func (o *PutIdentityprovidersPingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutIdentityprovidersPingForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersPingForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersPingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingNotFound creates a PutIdentityprovidersPingNotFound with default headers values
func NewPutIdentityprovidersPingNotFound() *PutIdentityprovidersPingNotFound {
	return &PutIdentityprovidersPingNotFound{}
}

/*
PutIdentityprovidersPingNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersPingNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping not found response has a 2xx status code
func (o *PutIdentityprovidersPingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping not found response has a 3xx status code
func (o *PutIdentityprovidersPingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping not found response has a 4xx status code
func (o *PutIdentityprovidersPingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping not found response has a 5xx status code
func (o *PutIdentityprovidersPingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping not found response a status code equal to that given
func (o *PutIdentityprovidersPingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutIdentityprovidersPingNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersPingNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersPingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingRequestTimeout creates a PutIdentityprovidersPingRequestTimeout with default headers values
func NewPutIdentityprovidersPingRequestTimeout() *PutIdentityprovidersPingRequestTimeout {
	return &PutIdentityprovidersPingRequestTimeout{}
}

/*
PutIdentityprovidersPingRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutIdentityprovidersPingRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping request timeout response has a 2xx status code
func (o *PutIdentityprovidersPingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping request timeout response has a 3xx status code
func (o *PutIdentityprovidersPingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping request timeout response has a 4xx status code
func (o *PutIdentityprovidersPingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping request timeout response has a 5xx status code
func (o *PutIdentityprovidersPingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping request timeout response a status code equal to that given
func (o *PutIdentityprovidersPingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutIdentityprovidersPingRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersPingRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersPingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingRequestEntityTooLarge creates a PutIdentityprovidersPingRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersPingRequestEntityTooLarge() *PutIdentityprovidersPingRequestEntityTooLarge {
	return &PutIdentityprovidersPingRequestEntityTooLarge{}
}

/*
PutIdentityprovidersPingRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutIdentityprovidersPingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping request entity too large response has a 2xx status code
func (o *PutIdentityprovidersPingRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping request entity too large response has a 3xx status code
func (o *PutIdentityprovidersPingRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping request entity too large response has a 4xx status code
func (o *PutIdentityprovidersPingRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping request entity too large response has a 5xx status code
func (o *PutIdentityprovidersPingRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping request entity too large response a status code equal to that given
func (o *PutIdentityprovidersPingRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutIdentityprovidersPingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersPingRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersPingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingUnsupportedMediaType creates a PutIdentityprovidersPingUnsupportedMediaType with default headers values
func NewPutIdentityprovidersPingUnsupportedMediaType() *PutIdentityprovidersPingUnsupportedMediaType {
	return &PutIdentityprovidersPingUnsupportedMediaType{}
}

/*
PutIdentityprovidersPingUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersPingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping unsupported media type response has a 2xx status code
func (o *PutIdentityprovidersPingUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping unsupported media type response has a 3xx status code
func (o *PutIdentityprovidersPingUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping unsupported media type response has a 4xx status code
func (o *PutIdentityprovidersPingUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping unsupported media type response has a 5xx status code
func (o *PutIdentityprovidersPingUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping unsupported media type response a status code equal to that given
func (o *PutIdentityprovidersPingUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutIdentityprovidersPingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersPingUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersPingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingTooManyRequests creates a PutIdentityprovidersPingTooManyRequests with default headers values
func NewPutIdentityprovidersPingTooManyRequests() *PutIdentityprovidersPingTooManyRequests {
	return &PutIdentityprovidersPingTooManyRequests{}
}

/*
PutIdentityprovidersPingTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutIdentityprovidersPingTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping too many requests response has a 2xx status code
func (o *PutIdentityprovidersPingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping too many requests response has a 3xx status code
func (o *PutIdentityprovidersPingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping too many requests response has a 4xx status code
func (o *PutIdentityprovidersPingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put identityproviders ping too many requests response has a 5xx status code
func (o *PutIdentityprovidersPingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put identityproviders ping too many requests response a status code equal to that given
func (o *PutIdentityprovidersPingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutIdentityprovidersPingTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersPingTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersPingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingInternalServerError creates a PutIdentityprovidersPingInternalServerError with default headers values
func NewPutIdentityprovidersPingInternalServerError() *PutIdentityprovidersPingInternalServerError {
	return &PutIdentityprovidersPingInternalServerError{}
}

/*
PutIdentityprovidersPingInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersPingInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping internal server error response has a 2xx status code
func (o *PutIdentityprovidersPingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping internal server error response has a 3xx status code
func (o *PutIdentityprovidersPingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping internal server error response has a 4xx status code
func (o *PutIdentityprovidersPingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders ping internal server error response has a 5xx status code
func (o *PutIdentityprovidersPingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders ping internal server error response a status code equal to that given
func (o *PutIdentityprovidersPingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutIdentityprovidersPingInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersPingInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersPingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingServiceUnavailable creates a PutIdentityprovidersPingServiceUnavailable with default headers values
func NewPutIdentityprovidersPingServiceUnavailable() *PutIdentityprovidersPingServiceUnavailable {
	return &PutIdentityprovidersPingServiceUnavailable{}
}

/*
PutIdentityprovidersPingServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersPingServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping service unavailable response has a 2xx status code
func (o *PutIdentityprovidersPingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping service unavailable response has a 3xx status code
func (o *PutIdentityprovidersPingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping service unavailable response has a 4xx status code
func (o *PutIdentityprovidersPingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders ping service unavailable response has a 5xx status code
func (o *PutIdentityprovidersPingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders ping service unavailable response a status code equal to that given
func (o *PutIdentityprovidersPingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutIdentityprovidersPingServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersPingServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersPingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPingGatewayTimeout creates a PutIdentityprovidersPingGatewayTimeout with default headers values
func NewPutIdentityprovidersPingGatewayTimeout() *PutIdentityprovidersPingGatewayTimeout {
	return &PutIdentityprovidersPingGatewayTimeout{}
}

/*
PutIdentityprovidersPingGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutIdentityprovidersPingGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put identityproviders ping gateway timeout response has a 2xx status code
func (o *PutIdentityprovidersPingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put identityproviders ping gateway timeout response has a 3xx status code
func (o *PutIdentityprovidersPingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put identityproviders ping gateway timeout response has a 4xx status code
func (o *PutIdentityprovidersPingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put identityproviders ping gateway timeout response has a 5xx status code
func (o *PutIdentityprovidersPingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put identityproviders ping gateway timeout response a status code equal to that given
func (o *PutIdentityprovidersPingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutIdentityprovidersPingGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersPingGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/ping][%d] putIdentityprovidersPingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersPingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
