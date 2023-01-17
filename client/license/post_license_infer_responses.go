// Code generated by go-swagger; DO NOT EDIT.

package license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostLicenseInferReader is a Reader for the PostLicenseInfer structure.
type PostLicenseInferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLicenseInferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLicenseInferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLicenseInferBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLicenseInferUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLicenseInferForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLicenseInferNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLicenseInferRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLicenseInferRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLicenseInferUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLicenseInferTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLicenseInferInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLicenseInferServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLicenseInferGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLicenseInferOK creates a PostLicenseInferOK with default headers values
func NewPostLicenseInferOK() *PostLicenseInferOK {
	return &PostLicenseInferOK{}
}

/*
PostLicenseInferOK describes a response with status code 200, with default header values.

successful operation
*/
type PostLicenseInferOK struct {
	Payload []string
}

// IsSuccess returns true when this post license infer o k response has a 2xx status code
func (o *PostLicenseInferOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post license infer o k response has a 3xx status code
func (o *PostLicenseInferOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer o k response has a 4xx status code
func (o *PostLicenseInferOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license infer o k response has a 5xx status code
func (o *PostLicenseInferOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer o k response a status code equal to that given
func (o *PostLicenseInferOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLicenseInferOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferOK  %+v", 200, o.Payload)
}

func (o *PostLicenseInferOK) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferOK  %+v", 200, o.Payload)
}

func (o *PostLicenseInferOK) GetPayload() []string {
	return o.Payload
}

func (o *PostLicenseInferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferBadRequest creates a PostLicenseInferBadRequest with default headers values
func NewPostLicenseInferBadRequest() *PostLicenseInferBadRequest {
	return &PostLicenseInferBadRequest{}
}

/*
PostLicenseInferBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLicenseInferBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer bad request response has a 2xx status code
func (o *PostLicenseInferBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer bad request response has a 3xx status code
func (o *PostLicenseInferBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer bad request response has a 4xx status code
func (o *PostLicenseInferBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer bad request response has a 5xx status code
func (o *PostLicenseInferBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer bad request response a status code equal to that given
func (o *PostLicenseInferBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLicenseInferBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferBadRequest  %+v", 400, o.Payload)
}

func (o *PostLicenseInferBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferBadRequest  %+v", 400, o.Payload)
}

func (o *PostLicenseInferBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferUnauthorized creates a PostLicenseInferUnauthorized with default headers values
func NewPostLicenseInferUnauthorized() *PostLicenseInferUnauthorized {
	return &PostLicenseInferUnauthorized{}
}

/*
PostLicenseInferUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLicenseInferUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer unauthorized response has a 2xx status code
func (o *PostLicenseInferUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer unauthorized response has a 3xx status code
func (o *PostLicenseInferUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer unauthorized response has a 4xx status code
func (o *PostLicenseInferUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer unauthorized response has a 5xx status code
func (o *PostLicenseInferUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer unauthorized response a status code equal to that given
func (o *PostLicenseInferUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLicenseInferUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLicenseInferUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLicenseInferUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferForbidden creates a PostLicenseInferForbidden with default headers values
func NewPostLicenseInferForbidden() *PostLicenseInferForbidden {
	return &PostLicenseInferForbidden{}
}

/*
PostLicenseInferForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLicenseInferForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer forbidden response has a 2xx status code
func (o *PostLicenseInferForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer forbidden response has a 3xx status code
func (o *PostLicenseInferForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer forbidden response has a 4xx status code
func (o *PostLicenseInferForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer forbidden response has a 5xx status code
func (o *PostLicenseInferForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer forbidden response a status code equal to that given
func (o *PostLicenseInferForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLicenseInferForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferForbidden  %+v", 403, o.Payload)
}

func (o *PostLicenseInferForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferForbidden  %+v", 403, o.Payload)
}

func (o *PostLicenseInferForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferNotFound creates a PostLicenseInferNotFound with default headers values
func NewPostLicenseInferNotFound() *PostLicenseInferNotFound {
	return &PostLicenseInferNotFound{}
}

/*
PostLicenseInferNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLicenseInferNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer not found response has a 2xx status code
func (o *PostLicenseInferNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer not found response has a 3xx status code
func (o *PostLicenseInferNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer not found response has a 4xx status code
func (o *PostLicenseInferNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer not found response has a 5xx status code
func (o *PostLicenseInferNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer not found response a status code equal to that given
func (o *PostLicenseInferNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLicenseInferNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferNotFound  %+v", 404, o.Payload)
}

func (o *PostLicenseInferNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferNotFound  %+v", 404, o.Payload)
}

func (o *PostLicenseInferNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferRequestTimeout creates a PostLicenseInferRequestTimeout with default headers values
func NewPostLicenseInferRequestTimeout() *PostLicenseInferRequestTimeout {
	return &PostLicenseInferRequestTimeout{}
}

/*
PostLicenseInferRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLicenseInferRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer request timeout response has a 2xx status code
func (o *PostLicenseInferRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer request timeout response has a 3xx status code
func (o *PostLicenseInferRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer request timeout response has a 4xx status code
func (o *PostLicenseInferRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer request timeout response has a 5xx status code
func (o *PostLicenseInferRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer request timeout response a status code equal to that given
func (o *PostLicenseInferRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLicenseInferRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLicenseInferRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLicenseInferRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferRequestEntityTooLarge creates a PostLicenseInferRequestEntityTooLarge with default headers values
func NewPostLicenseInferRequestEntityTooLarge() *PostLicenseInferRequestEntityTooLarge {
	return &PostLicenseInferRequestEntityTooLarge{}
}

/*
PostLicenseInferRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostLicenseInferRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer request entity too large response has a 2xx status code
func (o *PostLicenseInferRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer request entity too large response has a 3xx status code
func (o *PostLicenseInferRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer request entity too large response has a 4xx status code
func (o *PostLicenseInferRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer request entity too large response has a 5xx status code
func (o *PostLicenseInferRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer request entity too large response a status code equal to that given
func (o *PostLicenseInferRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLicenseInferRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLicenseInferRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLicenseInferRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferUnsupportedMediaType creates a PostLicenseInferUnsupportedMediaType with default headers values
func NewPostLicenseInferUnsupportedMediaType() *PostLicenseInferUnsupportedMediaType {
	return &PostLicenseInferUnsupportedMediaType{}
}

/*
PostLicenseInferUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLicenseInferUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer unsupported media type response has a 2xx status code
func (o *PostLicenseInferUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer unsupported media type response has a 3xx status code
func (o *PostLicenseInferUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer unsupported media type response has a 4xx status code
func (o *PostLicenseInferUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer unsupported media type response has a 5xx status code
func (o *PostLicenseInferUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer unsupported media type response a status code equal to that given
func (o *PostLicenseInferUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLicenseInferUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLicenseInferUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLicenseInferUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferTooManyRequests creates a PostLicenseInferTooManyRequests with default headers values
func NewPostLicenseInferTooManyRequests() *PostLicenseInferTooManyRequests {
	return &PostLicenseInferTooManyRequests{}
}

/*
PostLicenseInferTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLicenseInferTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer too many requests response has a 2xx status code
func (o *PostLicenseInferTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer too many requests response has a 3xx status code
func (o *PostLicenseInferTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer too many requests response has a 4xx status code
func (o *PostLicenseInferTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post license infer too many requests response has a 5xx status code
func (o *PostLicenseInferTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post license infer too many requests response a status code equal to that given
func (o *PostLicenseInferTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLicenseInferTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLicenseInferTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLicenseInferTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferInternalServerError creates a PostLicenseInferInternalServerError with default headers values
func NewPostLicenseInferInternalServerError() *PostLicenseInferInternalServerError {
	return &PostLicenseInferInternalServerError{}
}

/*
PostLicenseInferInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLicenseInferInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer internal server error response has a 2xx status code
func (o *PostLicenseInferInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer internal server error response has a 3xx status code
func (o *PostLicenseInferInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer internal server error response has a 4xx status code
func (o *PostLicenseInferInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license infer internal server error response has a 5xx status code
func (o *PostLicenseInferInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post license infer internal server error response a status code equal to that given
func (o *PostLicenseInferInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLicenseInferInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLicenseInferInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLicenseInferInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferServiceUnavailable creates a PostLicenseInferServiceUnavailable with default headers values
func NewPostLicenseInferServiceUnavailable() *PostLicenseInferServiceUnavailable {
	return &PostLicenseInferServiceUnavailable{}
}

/*
PostLicenseInferServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLicenseInferServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer service unavailable response has a 2xx status code
func (o *PostLicenseInferServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer service unavailable response has a 3xx status code
func (o *PostLicenseInferServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer service unavailable response has a 4xx status code
func (o *PostLicenseInferServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license infer service unavailable response has a 5xx status code
func (o *PostLicenseInferServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post license infer service unavailable response a status code equal to that given
func (o *PostLicenseInferServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLicenseInferServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLicenseInferServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLicenseInferServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLicenseInferGatewayTimeout creates a PostLicenseInferGatewayTimeout with default headers values
func NewPostLicenseInferGatewayTimeout() *PostLicenseInferGatewayTimeout {
	return &PostLicenseInferGatewayTimeout{}
}

/*
PostLicenseInferGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLicenseInferGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post license infer gateway timeout response has a 2xx status code
func (o *PostLicenseInferGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post license infer gateway timeout response has a 3xx status code
func (o *PostLicenseInferGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post license infer gateway timeout response has a 4xx status code
func (o *PostLicenseInferGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post license infer gateway timeout response has a 5xx status code
func (o *PostLicenseInferGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post license infer gateway timeout response a status code equal to that given
func (o *PostLicenseInferGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLicenseInferGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLicenseInferGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/license/infer][%d] postLicenseInferGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLicenseInferGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLicenseInferGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
