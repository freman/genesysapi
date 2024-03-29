// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserCallforwardingReader is a Reader for the GetUserCallforwarding structure.
type GetUserCallforwardingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserCallforwardingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserCallforwardingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserCallforwardingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserCallforwardingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserCallforwardingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserCallforwardingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserCallforwardingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserCallforwardingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserCallforwardingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 424:
		result := NewGetUserCallforwardingFailedDependency()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserCallforwardingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserCallforwardingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserCallforwardingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserCallforwardingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserCallforwardingOK creates a GetUserCallforwardingOK with default headers values
func NewGetUserCallforwardingOK() *GetUserCallforwardingOK {
	return &GetUserCallforwardingOK{}
}

/*
GetUserCallforwardingOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserCallforwardingOK struct {
	Payload *models.CallForwarding
}

// IsSuccess returns true when this get user callforwarding o k response has a 2xx status code
func (o *GetUserCallforwardingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user callforwarding o k response has a 3xx status code
func (o *GetUserCallforwardingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding o k response has a 4xx status code
func (o *GetUserCallforwardingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user callforwarding o k response has a 5xx status code
func (o *GetUserCallforwardingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding o k response a status code equal to that given
func (o *GetUserCallforwardingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserCallforwardingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingOK  %+v", 200, o.Payload)
}

func (o *GetUserCallforwardingOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingOK  %+v", 200, o.Payload)
}

func (o *GetUserCallforwardingOK) GetPayload() *models.CallForwarding {
	return o.Payload
}

func (o *GetUserCallforwardingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallForwarding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingBadRequest creates a GetUserCallforwardingBadRequest with default headers values
func NewGetUserCallforwardingBadRequest() *GetUserCallforwardingBadRequest {
	return &GetUserCallforwardingBadRequest{}
}

/*
GetUserCallforwardingBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserCallforwardingBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding bad request response has a 2xx status code
func (o *GetUserCallforwardingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding bad request response has a 3xx status code
func (o *GetUserCallforwardingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding bad request response has a 4xx status code
func (o *GetUserCallforwardingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding bad request response has a 5xx status code
func (o *GetUserCallforwardingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding bad request response a status code equal to that given
func (o *GetUserCallforwardingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserCallforwardingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserCallforwardingBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserCallforwardingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingUnauthorized creates a GetUserCallforwardingUnauthorized with default headers values
func NewGetUserCallforwardingUnauthorized() *GetUserCallforwardingUnauthorized {
	return &GetUserCallforwardingUnauthorized{}
}

/*
GetUserCallforwardingUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserCallforwardingUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding unauthorized response has a 2xx status code
func (o *GetUserCallforwardingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding unauthorized response has a 3xx status code
func (o *GetUserCallforwardingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding unauthorized response has a 4xx status code
func (o *GetUserCallforwardingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding unauthorized response has a 5xx status code
func (o *GetUserCallforwardingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding unauthorized response a status code equal to that given
func (o *GetUserCallforwardingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserCallforwardingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserCallforwardingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserCallforwardingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingForbidden creates a GetUserCallforwardingForbidden with default headers values
func NewGetUserCallforwardingForbidden() *GetUserCallforwardingForbidden {
	return &GetUserCallforwardingForbidden{}
}

/*
GetUserCallforwardingForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserCallforwardingForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding forbidden response has a 2xx status code
func (o *GetUserCallforwardingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding forbidden response has a 3xx status code
func (o *GetUserCallforwardingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding forbidden response has a 4xx status code
func (o *GetUserCallforwardingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding forbidden response has a 5xx status code
func (o *GetUserCallforwardingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding forbidden response a status code equal to that given
func (o *GetUserCallforwardingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserCallforwardingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingForbidden  %+v", 403, o.Payload)
}

func (o *GetUserCallforwardingForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingForbidden  %+v", 403, o.Payload)
}

func (o *GetUserCallforwardingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingNotFound creates a GetUserCallforwardingNotFound with default headers values
func NewGetUserCallforwardingNotFound() *GetUserCallforwardingNotFound {
	return &GetUserCallforwardingNotFound{}
}

/*
GetUserCallforwardingNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserCallforwardingNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding not found response has a 2xx status code
func (o *GetUserCallforwardingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding not found response has a 3xx status code
func (o *GetUserCallforwardingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding not found response has a 4xx status code
func (o *GetUserCallforwardingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding not found response has a 5xx status code
func (o *GetUserCallforwardingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding not found response a status code equal to that given
func (o *GetUserCallforwardingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserCallforwardingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingNotFound  %+v", 404, o.Payload)
}

func (o *GetUserCallforwardingNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingNotFound  %+v", 404, o.Payload)
}

func (o *GetUserCallforwardingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingRequestTimeout creates a GetUserCallforwardingRequestTimeout with default headers values
func NewGetUserCallforwardingRequestTimeout() *GetUserCallforwardingRequestTimeout {
	return &GetUserCallforwardingRequestTimeout{}
}

/*
GetUserCallforwardingRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserCallforwardingRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding request timeout response has a 2xx status code
func (o *GetUserCallforwardingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding request timeout response has a 3xx status code
func (o *GetUserCallforwardingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding request timeout response has a 4xx status code
func (o *GetUserCallforwardingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding request timeout response has a 5xx status code
func (o *GetUserCallforwardingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding request timeout response a status code equal to that given
func (o *GetUserCallforwardingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserCallforwardingRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserCallforwardingRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserCallforwardingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingRequestEntityTooLarge creates a GetUserCallforwardingRequestEntityTooLarge with default headers values
func NewGetUserCallforwardingRequestEntityTooLarge() *GetUserCallforwardingRequestEntityTooLarge {
	return &GetUserCallforwardingRequestEntityTooLarge{}
}

/*
GetUserCallforwardingRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUserCallforwardingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding request entity too large response has a 2xx status code
func (o *GetUserCallforwardingRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding request entity too large response has a 3xx status code
func (o *GetUserCallforwardingRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding request entity too large response has a 4xx status code
func (o *GetUserCallforwardingRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding request entity too large response has a 5xx status code
func (o *GetUserCallforwardingRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding request entity too large response a status code equal to that given
func (o *GetUserCallforwardingRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserCallforwardingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserCallforwardingRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserCallforwardingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingUnsupportedMediaType creates a GetUserCallforwardingUnsupportedMediaType with default headers values
func NewGetUserCallforwardingUnsupportedMediaType() *GetUserCallforwardingUnsupportedMediaType {
	return &GetUserCallforwardingUnsupportedMediaType{}
}

/*
GetUserCallforwardingUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserCallforwardingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding unsupported media type response has a 2xx status code
func (o *GetUserCallforwardingUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding unsupported media type response has a 3xx status code
func (o *GetUserCallforwardingUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding unsupported media type response has a 4xx status code
func (o *GetUserCallforwardingUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding unsupported media type response has a 5xx status code
func (o *GetUserCallforwardingUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding unsupported media type response a status code equal to that given
func (o *GetUserCallforwardingUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserCallforwardingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserCallforwardingUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserCallforwardingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingFailedDependency creates a GetUserCallforwardingFailedDependency with default headers values
func NewGetUserCallforwardingFailedDependency() *GetUserCallforwardingFailedDependency {
	return &GetUserCallforwardingFailedDependency{}
}

/*
GetUserCallforwardingFailedDependency describes a response with status code 424, with default header values.

GetUserCallforwardingFailedDependency get user callforwarding failed dependency
*/
type GetUserCallforwardingFailedDependency struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding failed dependency response has a 2xx status code
func (o *GetUserCallforwardingFailedDependency) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding failed dependency response has a 3xx status code
func (o *GetUserCallforwardingFailedDependency) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding failed dependency response has a 4xx status code
func (o *GetUserCallforwardingFailedDependency) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding failed dependency response has a 5xx status code
func (o *GetUserCallforwardingFailedDependency) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding failed dependency response a status code equal to that given
func (o *GetUserCallforwardingFailedDependency) IsCode(code int) bool {
	return code == 424
}

func (o *GetUserCallforwardingFailedDependency) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingFailedDependency  %+v", 424, o.Payload)
}

func (o *GetUserCallforwardingFailedDependency) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingFailedDependency  %+v", 424, o.Payload)
}

func (o *GetUserCallforwardingFailedDependency) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingFailedDependency) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingTooManyRequests creates a GetUserCallforwardingTooManyRequests with default headers values
func NewGetUserCallforwardingTooManyRequests() *GetUserCallforwardingTooManyRequests {
	return &GetUserCallforwardingTooManyRequests{}
}

/*
GetUserCallforwardingTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserCallforwardingTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding too many requests response has a 2xx status code
func (o *GetUserCallforwardingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding too many requests response has a 3xx status code
func (o *GetUserCallforwardingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding too many requests response has a 4xx status code
func (o *GetUserCallforwardingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user callforwarding too many requests response has a 5xx status code
func (o *GetUserCallforwardingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user callforwarding too many requests response a status code equal to that given
func (o *GetUserCallforwardingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserCallforwardingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserCallforwardingTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserCallforwardingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingInternalServerError creates a GetUserCallforwardingInternalServerError with default headers values
func NewGetUserCallforwardingInternalServerError() *GetUserCallforwardingInternalServerError {
	return &GetUserCallforwardingInternalServerError{}
}

/*
GetUserCallforwardingInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserCallforwardingInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding internal server error response has a 2xx status code
func (o *GetUserCallforwardingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding internal server error response has a 3xx status code
func (o *GetUserCallforwardingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding internal server error response has a 4xx status code
func (o *GetUserCallforwardingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user callforwarding internal server error response has a 5xx status code
func (o *GetUserCallforwardingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user callforwarding internal server error response a status code equal to that given
func (o *GetUserCallforwardingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserCallforwardingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserCallforwardingInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserCallforwardingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingServiceUnavailable creates a GetUserCallforwardingServiceUnavailable with default headers values
func NewGetUserCallforwardingServiceUnavailable() *GetUserCallforwardingServiceUnavailable {
	return &GetUserCallforwardingServiceUnavailable{}
}

/*
GetUserCallforwardingServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserCallforwardingServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding service unavailable response has a 2xx status code
func (o *GetUserCallforwardingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding service unavailable response has a 3xx status code
func (o *GetUserCallforwardingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding service unavailable response has a 4xx status code
func (o *GetUserCallforwardingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user callforwarding service unavailable response has a 5xx status code
func (o *GetUserCallforwardingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user callforwarding service unavailable response a status code equal to that given
func (o *GetUserCallforwardingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserCallforwardingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserCallforwardingServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserCallforwardingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserCallforwardingGatewayTimeout creates a GetUserCallforwardingGatewayTimeout with default headers values
func NewGetUserCallforwardingGatewayTimeout() *GetUserCallforwardingGatewayTimeout {
	return &GetUserCallforwardingGatewayTimeout{}
}

/*
GetUserCallforwardingGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserCallforwardingGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user callforwarding gateway timeout response has a 2xx status code
func (o *GetUserCallforwardingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user callforwarding gateway timeout response has a 3xx status code
func (o *GetUserCallforwardingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user callforwarding gateway timeout response has a 4xx status code
func (o *GetUserCallforwardingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user callforwarding gateway timeout response has a 5xx status code
func (o *GetUserCallforwardingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user callforwarding gateway timeout response a status code equal to that given
func (o *GetUserCallforwardingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserCallforwardingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserCallforwardingGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/callforwarding][%d] getUserCallforwardingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserCallforwardingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserCallforwardingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
