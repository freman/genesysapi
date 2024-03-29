// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutRoutingWrapupcodeReader is a Reader for the PutRoutingWrapupcode structure.
type PutRoutingWrapupcodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRoutingWrapupcodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRoutingWrapupcodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRoutingWrapupcodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRoutingWrapupcodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRoutingWrapupcodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRoutingWrapupcodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutRoutingWrapupcodeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutRoutingWrapupcodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutRoutingWrapupcodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutRoutingWrapupcodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRoutingWrapupcodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutRoutingWrapupcodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutRoutingWrapupcodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRoutingWrapupcodeOK creates a PutRoutingWrapupcodeOK with default headers values
func NewPutRoutingWrapupcodeOK() *PutRoutingWrapupcodeOK {
	return &PutRoutingWrapupcodeOK{}
}

/*
PutRoutingWrapupcodeOK describes a response with status code 200, with default header values.

successful operation
*/
type PutRoutingWrapupcodeOK struct {
	Payload *models.WrapupCode
}

// IsSuccess returns true when this put routing wrapupcode o k response has a 2xx status code
func (o *PutRoutingWrapupcodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put routing wrapupcode o k response has a 3xx status code
func (o *PutRoutingWrapupcodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode o k response has a 4xx status code
func (o *PutRoutingWrapupcodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put routing wrapupcode o k response has a 5xx status code
func (o *PutRoutingWrapupcodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode o k response a status code equal to that given
func (o *PutRoutingWrapupcodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutRoutingWrapupcodeOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeOK  %+v", 200, o.Payload)
}

func (o *PutRoutingWrapupcodeOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeOK  %+v", 200, o.Payload)
}

func (o *PutRoutingWrapupcodeOK) GetPayload() *models.WrapupCode {
	return o.Payload
}

func (o *PutRoutingWrapupcodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WrapupCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeBadRequest creates a PutRoutingWrapupcodeBadRequest with default headers values
func NewPutRoutingWrapupcodeBadRequest() *PutRoutingWrapupcodeBadRequest {
	return &PutRoutingWrapupcodeBadRequest{}
}

/*
PutRoutingWrapupcodeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutRoutingWrapupcodeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode bad request response has a 2xx status code
func (o *PutRoutingWrapupcodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode bad request response has a 3xx status code
func (o *PutRoutingWrapupcodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode bad request response has a 4xx status code
func (o *PutRoutingWrapupcodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode bad request response has a 5xx status code
func (o *PutRoutingWrapupcodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode bad request response a status code equal to that given
func (o *PutRoutingWrapupcodeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutRoutingWrapupcodeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeBadRequest  %+v", 400, o.Payload)
}

func (o *PutRoutingWrapupcodeBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeBadRequest  %+v", 400, o.Payload)
}

func (o *PutRoutingWrapupcodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeUnauthorized creates a PutRoutingWrapupcodeUnauthorized with default headers values
func NewPutRoutingWrapupcodeUnauthorized() *PutRoutingWrapupcodeUnauthorized {
	return &PutRoutingWrapupcodeUnauthorized{}
}

/*
PutRoutingWrapupcodeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutRoutingWrapupcodeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode unauthorized response has a 2xx status code
func (o *PutRoutingWrapupcodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode unauthorized response has a 3xx status code
func (o *PutRoutingWrapupcodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode unauthorized response has a 4xx status code
func (o *PutRoutingWrapupcodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode unauthorized response has a 5xx status code
func (o *PutRoutingWrapupcodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode unauthorized response a status code equal to that given
func (o *PutRoutingWrapupcodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutRoutingWrapupcodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRoutingWrapupcodeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRoutingWrapupcodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeForbidden creates a PutRoutingWrapupcodeForbidden with default headers values
func NewPutRoutingWrapupcodeForbidden() *PutRoutingWrapupcodeForbidden {
	return &PutRoutingWrapupcodeForbidden{}
}

/*
PutRoutingWrapupcodeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutRoutingWrapupcodeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode forbidden response has a 2xx status code
func (o *PutRoutingWrapupcodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode forbidden response has a 3xx status code
func (o *PutRoutingWrapupcodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode forbidden response has a 4xx status code
func (o *PutRoutingWrapupcodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode forbidden response has a 5xx status code
func (o *PutRoutingWrapupcodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode forbidden response a status code equal to that given
func (o *PutRoutingWrapupcodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutRoutingWrapupcodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeForbidden  %+v", 403, o.Payload)
}

func (o *PutRoutingWrapupcodeForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeForbidden  %+v", 403, o.Payload)
}

func (o *PutRoutingWrapupcodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeNotFound creates a PutRoutingWrapupcodeNotFound with default headers values
func NewPutRoutingWrapupcodeNotFound() *PutRoutingWrapupcodeNotFound {
	return &PutRoutingWrapupcodeNotFound{}
}

/*
PutRoutingWrapupcodeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutRoutingWrapupcodeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode not found response has a 2xx status code
func (o *PutRoutingWrapupcodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode not found response has a 3xx status code
func (o *PutRoutingWrapupcodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode not found response has a 4xx status code
func (o *PutRoutingWrapupcodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode not found response has a 5xx status code
func (o *PutRoutingWrapupcodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode not found response a status code equal to that given
func (o *PutRoutingWrapupcodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutRoutingWrapupcodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeNotFound  %+v", 404, o.Payload)
}

func (o *PutRoutingWrapupcodeNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeNotFound  %+v", 404, o.Payload)
}

func (o *PutRoutingWrapupcodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeRequestTimeout creates a PutRoutingWrapupcodeRequestTimeout with default headers values
func NewPutRoutingWrapupcodeRequestTimeout() *PutRoutingWrapupcodeRequestTimeout {
	return &PutRoutingWrapupcodeRequestTimeout{}
}

/*
PutRoutingWrapupcodeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutRoutingWrapupcodeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode request timeout response has a 2xx status code
func (o *PutRoutingWrapupcodeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode request timeout response has a 3xx status code
func (o *PutRoutingWrapupcodeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode request timeout response has a 4xx status code
func (o *PutRoutingWrapupcodeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode request timeout response has a 5xx status code
func (o *PutRoutingWrapupcodeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode request timeout response a status code equal to that given
func (o *PutRoutingWrapupcodeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutRoutingWrapupcodeRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRoutingWrapupcodeRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRoutingWrapupcodeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeRequestEntityTooLarge creates a PutRoutingWrapupcodeRequestEntityTooLarge with default headers values
func NewPutRoutingWrapupcodeRequestEntityTooLarge() *PutRoutingWrapupcodeRequestEntityTooLarge {
	return &PutRoutingWrapupcodeRequestEntityTooLarge{}
}

/*
PutRoutingWrapupcodeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutRoutingWrapupcodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode request entity too large response has a 2xx status code
func (o *PutRoutingWrapupcodeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode request entity too large response has a 3xx status code
func (o *PutRoutingWrapupcodeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode request entity too large response has a 4xx status code
func (o *PutRoutingWrapupcodeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode request entity too large response has a 5xx status code
func (o *PutRoutingWrapupcodeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode request entity too large response a status code equal to that given
func (o *PutRoutingWrapupcodeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutRoutingWrapupcodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRoutingWrapupcodeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRoutingWrapupcodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeUnsupportedMediaType creates a PutRoutingWrapupcodeUnsupportedMediaType with default headers values
func NewPutRoutingWrapupcodeUnsupportedMediaType() *PutRoutingWrapupcodeUnsupportedMediaType {
	return &PutRoutingWrapupcodeUnsupportedMediaType{}
}

/*
PutRoutingWrapupcodeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutRoutingWrapupcodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode unsupported media type response has a 2xx status code
func (o *PutRoutingWrapupcodeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode unsupported media type response has a 3xx status code
func (o *PutRoutingWrapupcodeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode unsupported media type response has a 4xx status code
func (o *PutRoutingWrapupcodeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode unsupported media type response has a 5xx status code
func (o *PutRoutingWrapupcodeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode unsupported media type response a status code equal to that given
func (o *PutRoutingWrapupcodeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutRoutingWrapupcodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRoutingWrapupcodeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRoutingWrapupcodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeTooManyRequests creates a PutRoutingWrapupcodeTooManyRequests with default headers values
func NewPutRoutingWrapupcodeTooManyRequests() *PutRoutingWrapupcodeTooManyRequests {
	return &PutRoutingWrapupcodeTooManyRequests{}
}

/*
PutRoutingWrapupcodeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutRoutingWrapupcodeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode too many requests response has a 2xx status code
func (o *PutRoutingWrapupcodeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode too many requests response has a 3xx status code
func (o *PutRoutingWrapupcodeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode too many requests response has a 4xx status code
func (o *PutRoutingWrapupcodeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put routing wrapupcode too many requests response has a 5xx status code
func (o *PutRoutingWrapupcodeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put routing wrapupcode too many requests response a status code equal to that given
func (o *PutRoutingWrapupcodeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutRoutingWrapupcodeTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRoutingWrapupcodeTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRoutingWrapupcodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeInternalServerError creates a PutRoutingWrapupcodeInternalServerError with default headers values
func NewPutRoutingWrapupcodeInternalServerError() *PutRoutingWrapupcodeInternalServerError {
	return &PutRoutingWrapupcodeInternalServerError{}
}

/*
PutRoutingWrapupcodeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutRoutingWrapupcodeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode internal server error response has a 2xx status code
func (o *PutRoutingWrapupcodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode internal server error response has a 3xx status code
func (o *PutRoutingWrapupcodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode internal server error response has a 4xx status code
func (o *PutRoutingWrapupcodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put routing wrapupcode internal server error response has a 5xx status code
func (o *PutRoutingWrapupcodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put routing wrapupcode internal server error response a status code equal to that given
func (o *PutRoutingWrapupcodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutRoutingWrapupcodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRoutingWrapupcodeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRoutingWrapupcodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeServiceUnavailable creates a PutRoutingWrapupcodeServiceUnavailable with default headers values
func NewPutRoutingWrapupcodeServiceUnavailable() *PutRoutingWrapupcodeServiceUnavailable {
	return &PutRoutingWrapupcodeServiceUnavailable{}
}

/*
PutRoutingWrapupcodeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutRoutingWrapupcodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode service unavailable response has a 2xx status code
func (o *PutRoutingWrapupcodeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode service unavailable response has a 3xx status code
func (o *PutRoutingWrapupcodeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode service unavailable response has a 4xx status code
func (o *PutRoutingWrapupcodeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put routing wrapupcode service unavailable response has a 5xx status code
func (o *PutRoutingWrapupcodeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put routing wrapupcode service unavailable response a status code equal to that given
func (o *PutRoutingWrapupcodeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutRoutingWrapupcodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRoutingWrapupcodeServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRoutingWrapupcodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingWrapupcodeGatewayTimeout creates a PutRoutingWrapupcodeGatewayTimeout with default headers values
func NewPutRoutingWrapupcodeGatewayTimeout() *PutRoutingWrapupcodeGatewayTimeout {
	return &PutRoutingWrapupcodeGatewayTimeout{}
}

/*
PutRoutingWrapupcodeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutRoutingWrapupcodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put routing wrapupcode gateway timeout response has a 2xx status code
func (o *PutRoutingWrapupcodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put routing wrapupcode gateway timeout response has a 3xx status code
func (o *PutRoutingWrapupcodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put routing wrapupcode gateway timeout response has a 4xx status code
func (o *PutRoutingWrapupcodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put routing wrapupcode gateway timeout response has a 5xx status code
func (o *PutRoutingWrapupcodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put routing wrapupcode gateway timeout response a status code equal to that given
func (o *PutRoutingWrapupcodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutRoutingWrapupcodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRoutingWrapupcodeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/routing/wrapupcodes/{codeId}][%d] putRoutingWrapupcodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRoutingWrapupcodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingWrapupcodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
