// Code generated by go-swagger; DO NOT EDIT.

package greetings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutGroupGreetingsDefaultsReader is a Reader for the PutGroupGreetingsDefaults structure.
type PutGroupGreetingsDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGroupGreetingsDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutGroupGreetingsDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGroupGreetingsDefaultsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutGroupGreetingsDefaultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutGroupGreetingsDefaultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutGroupGreetingsDefaultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutGroupGreetingsDefaultsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutGroupGreetingsDefaultsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutGroupGreetingsDefaultsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutGroupGreetingsDefaultsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutGroupGreetingsDefaultsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutGroupGreetingsDefaultsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutGroupGreetingsDefaultsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutGroupGreetingsDefaultsOK creates a PutGroupGreetingsDefaultsOK with default headers values
func NewPutGroupGreetingsDefaultsOK() *PutGroupGreetingsDefaultsOK {
	return &PutGroupGreetingsDefaultsOK{}
}

/*
PutGroupGreetingsDefaultsOK describes a response with status code 200, with default header values.

successful operation
*/
type PutGroupGreetingsDefaultsOK struct {
	Payload *models.DefaultGreetingList
}

// IsSuccess returns true when this put group greetings defaults o k response has a 2xx status code
func (o *PutGroupGreetingsDefaultsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put group greetings defaults o k response has a 3xx status code
func (o *PutGroupGreetingsDefaultsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults o k response has a 4xx status code
func (o *PutGroupGreetingsDefaultsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put group greetings defaults o k response has a 5xx status code
func (o *PutGroupGreetingsDefaultsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults o k response a status code equal to that given
func (o *PutGroupGreetingsDefaultsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutGroupGreetingsDefaultsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsOK  %+v", 200, o.Payload)
}

func (o *PutGroupGreetingsDefaultsOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsOK  %+v", 200, o.Payload)
}

func (o *PutGroupGreetingsDefaultsOK) GetPayload() *models.DefaultGreetingList {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultGreetingList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsBadRequest creates a PutGroupGreetingsDefaultsBadRequest with default headers values
func NewPutGroupGreetingsDefaultsBadRequest() *PutGroupGreetingsDefaultsBadRequest {
	return &PutGroupGreetingsDefaultsBadRequest{}
}

/*
PutGroupGreetingsDefaultsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutGroupGreetingsDefaultsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults bad request response has a 2xx status code
func (o *PutGroupGreetingsDefaultsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults bad request response has a 3xx status code
func (o *PutGroupGreetingsDefaultsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults bad request response has a 4xx status code
func (o *PutGroupGreetingsDefaultsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults bad request response has a 5xx status code
func (o *PutGroupGreetingsDefaultsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults bad request response a status code equal to that given
func (o *PutGroupGreetingsDefaultsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutGroupGreetingsDefaultsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *PutGroupGreetingsDefaultsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsBadRequest  %+v", 400, o.Payload)
}

func (o *PutGroupGreetingsDefaultsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsUnauthorized creates a PutGroupGreetingsDefaultsUnauthorized with default headers values
func NewPutGroupGreetingsDefaultsUnauthorized() *PutGroupGreetingsDefaultsUnauthorized {
	return &PutGroupGreetingsDefaultsUnauthorized{}
}

/*
PutGroupGreetingsDefaultsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutGroupGreetingsDefaultsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults unauthorized response has a 2xx status code
func (o *PutGroupGreetingsDefaultsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults unauthorized response has a 3xx status code
func (o *PutGroupGreetingsDefaultsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults unauthorized response has a 4xx status code
func (o *PutGroupGreetingsDefaultsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults unauthorized response has a 5xx status code
func (o *PutGroupGreetingsDefaultsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults unauthorized response a status code equal to that given
func (o *PutGroupGreetingsDefaultsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutGroupGreetingsDefaultsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGroupGreetingsDefaultsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGroupGreetingsDefaultsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsForbidden creates a PutGroupGreetingsDefaultsForbidden with default headers values
func NewPutGroupGreetingsDefaultsForbidden() *PutGroupGreetingsDefaultsForbidden {
	return &PutGroupGreetingsDefaultsForbidden{}
}

/*
PutGroupGreetingsDefaultsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutGroupGreetingsDefaultsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults forbidden response has a 2xx status code
func (o *PutGroupGreetingsDefaultsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults forbidden response has a 3xx status code
func (o *PutGroupGreetingsDefaultsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults forbidden response has a 4xx status code
func (o *PutGroupGreetingsDefaultsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults forbidden response has a 5xx status code
func (o *PutGroupGreetingsDefaultsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults forbidden response a status code equal to that given
func (o *PutGroupGreetingsDefaultsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutGroupGreetingsDefaultsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsForbidden  %+v", 403, o.Payload)
}

func (o *PutGroupGreetingsDefaultsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsForbidden  %+v", 403, o.Payload)
}

func (o *PutGroupGreetingsDefaultsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsNotFound creates a PutGroupGreetingsDefaultsNotFound with default headers values
func NewPutGroupGreetingsDefaultsNotFound() *PutGroupGreetingsDefaultsNotFound {
	return &PutGroupGreetingsDefaultsNotFound{}
}

/*
PutGroupGreetingsDefaultsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutGroupGreetingsDefaultsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults not found response has a 2xx status code
func (o *PutGroupGreetingsDefaultsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults not found response has a 3xx status code
func (o *PutGroupGreetingsDefaultsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults not found response has a 4xx status code
func (o *PutGroupGreetingsDefaultsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults not found response has a 5xx status code
func (o *PutGroupGreetingsDefaultsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults not found response a status code equal to that given
func (o *PutGroupGreetingsDefaultsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutGroupGreetingsDefaultsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsNotFound  %+v", 404, o.Payload)
}

func (o *PutGroupGreetingsDefaultsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsNotFound  %+v", 404, o.Payload)
}

func (o *PutGroupGreetingsDefaultsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsRequestTimeout creates a PutGroupGreetingsDefaultsRequestTimeout with default headers values
func NewPutGroupGreetingsDefaultsRequestTimeout() *PutGroupGreetingsDefaultsRequestTimeout {
	return &PutGroupGreetingsDefaultsRequestTimeout{}
}

/*
PutGroupGreetingsDefaultsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutGroupGreetingsDefaultsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults request timeout response has a 2xx status code
func (o *PutGroupGreetingsDefaultsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults request timeout response has a 3xx status code
func (o *PutGroupGreetingsDefaultsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults request timeout response has a 4xx status code
func (o *PutGroupGreetingsDefaultsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults request timeout response has a 5xx status code
func (o *PutGroupGreetingsDefaultsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults request timeout response a status code equal to that given
func (o *PutGroupGreetingsDefaultsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutGroupGreetingsDefaultsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutGroupGreetingsDefaultsRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutGroupGreetingsDefaultsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsRequestEntityTooLarge creates a PutGroupGreetingsDefaultsRequestEntityTooLarge with default headers values
func NewPutGroupGreetingsDefaultsRequestEntityTooLarge() *PutGroupGreetingsDefaultsRequestEntityTooLarge {
	return &PutGroupGreetingsDefaultsRequestEntityTooLarge{}
}

/*
PutGroupGreetingsDefaultsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutGroupGreetingsDefaultsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults request entity too large response has a 2xx status code
func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults request entity too large response has a 3xx status code
func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults request entity too large response has a 4xx status code
func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults request entity too large response has a 5xx status code
func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults request entity too large response a status code equal to that given
func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsUnsupportedMediaType creates a PutGroupGreetingsDefaultsUnsupportedMediaType with default headers values
func NewPutGroupGreetingsDefaultsUnsupportedMediaType() *PutGroupGreetingsDefaultsUnsupportedMediaType {
	return &PutGroupGreetingsDefaultsUnsupportedMediaType{}
}

/*
PutGroupGreetingsDefaultsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutGroupGreetingsDefaultsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults unsupported media type response has a 2xx status code
func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults unsupported media type response has a 3xx status code
func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults unsupported media type response has a 4xx status code
func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults unsupported media type response has a 5xx status code
func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults unsupported media type response a status code equal to that given
func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsTooManyRequests creates a PutGroupGreetingsDefaultsTooManyRequests with default headers values
func NewPutGroupGreetingsDefaultsTooManyRequests() *PutGroupGreetingsDefaultsTooManyRequests {
	return &PutGroupGreetingsDefaultsTooManyRequests{}
}

/*
PutGroupGreetingsDefaultsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutGroupGreetingsDefaultsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults too many requests response has a 2xx status code
func (o *PutGroupGreetingsDefaultsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults too many requests response has a 3xx status code
func (o *PutGroupGreetingsDefaultsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults too many requests response has a 4xx status code
func (o *PutGroupGreetingsDefaultsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put group greetings defaults too many requests response has a 5xx status code
func (o *PutGroupGreetingsDefaultsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put group greetings defaults too many requests response a status code equal to that given
func (o *PutGroupGreetingsDefaultsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutGroupGreetingsDefaultsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGroupGreetingsDefaultsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGroupGreetingsDefaultsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsInternalServerError creates a PutGroupGreetingsDefaultsInternalServerError with default headers values
func NewPutGroupGreetingsDefaultsInternalServerError() *PutGroupGreetingsDefaultsInternalServerError {
	return &PutGroupGreetingsDefaultsInternalServerError{}
}

/*
PutGroupGreetingsDefaultsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutGroupGreetingsDefaultsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults internal server error response has a 2xx status code
func (o *PutGroupGreetingsDefaultsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults internal server error response has a 3xx status code
func (o *PutGroupGreetingsDefaultsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults internal server error response has a 4xx status code
func (o *PutGroupGreetingsDefaultsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put group greetings defaults internal server error response has a 5xx status code
func (o *PutGroupGreetingsDefaultsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put group greetings defaults internal server error response a status code equal to that given
func (o *PutGroupGreetingsDefaultsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutGroupGreetingsDefaultsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGroupGreetingsDefaultsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGroupGreetingsDefaultsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsServiceUnavailable creates a PutGroupGreetingsDefaultsServiceUnavailable with default headers values
func NewPutGroupGreetingsDefaultsServiceUnavailable() *PutGroupGreetingsDefaultsServiceUnavailable {
	return &PutGroupGreetingsDefaultsServiceUnavailable{}
}

/*
PutGroupGreetingsDefaultsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutGroupGreetingsDefaultsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults service unavailable response has a 2xx status code
func (o *PutGroupGreetingsDefaultsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults service unavailable response has a 3xx status code
func (o *PutGroupGreetingsDefaultsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults service unavailable response has a 4xx status code
func (o *PutGroupGreetingsDefaultsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put group greetings defaults service unavailable response has a 5xx status code
func (o *PutGroupGreetingsDefaultsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put group greetings defaults service unavailable response a status code equal to that given
func (o *PutGroupGreetingsDefaultsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutGroupGreetingsDefaultsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGroupGreetingsDefaultsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGroupGreetingsDefaultsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGreetingsDefaultsGatewayTimeout creates a PutGroupGreetingsDefaultsGatewayTimeout with default headers values
func NewPutGroupGreetingsDefaultsGatewayTimeout() *PutGroupGreetingsDefaultsGatewayTimeout {
	return &PutGroupGreetingsDefaultsGatewayTimeout{}
}

/*
PutGroupGreetingsDefaultsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutGroupGreetingsDefaultsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put group greetings defaults gateway timeout response has a 2xx status code
func (o *PutGroupGreetingsDefaultsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put group greetings defaults gateway timeout response has a 3xx status code
func (o *PutGroupGreetingsDefaultsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put group greetings defaults gateway timeout response has a 4xx status code
func (o *PutGroupGreetingsDefaultsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put group greetings defaults gateway timeout response has a 5xx status code
func (o *PutGroupGreetingsDefaultsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put group greetings defaults gateway timeout response a status code equal to that given
func (o *PutGroupGreetingsDefaultsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutGroupGreetingsDefaultsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGroupGreetingsDefaultsGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}/greetings/defaults][%d] putGroupGreetingsDefaultsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGroupGreetingsDefaultsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGreetingsDefaultsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
