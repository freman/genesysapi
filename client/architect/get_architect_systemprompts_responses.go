// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetArchitectSystempromptsReader is a Reader for the GetArchitectSystemprompts structure.
type GetArchitectSystempromptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectSystempromptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectSystempromptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectSystempromptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectSystempromptsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectSystempromptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectSystempromptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectSystempromptsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectSystempromptsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectSystempromptsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectSystempromptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectSystempromptsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectSystempromptsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectSystempromptsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectSystempromptsOK creates a GetArchitectSystempromptsOK with default headers values
func NewGetArchitectSystempromptsOK() *GetArchitectSystempromptsOK {
	return &GetArchitectSystempromptsOK{}
}

/*
GetArchitectSystempromptsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArchitectSystempromptsOK struct {
	Payload *models.SystemPromptEntityListing
}

// IsSuccess returns true when this get architect systemprompts o k response has a 2xx status code
func (o *GetArchitectSystempromptsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get architect systemprompts o k response has a 3xx status code
func (o *GetArchitectSystempromptsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts o k response has a 4xx status code
func (o *GetArchitectSystempromptsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompts o k response has a 5xx status code
func (o *GetArchitectSystempromptsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts o k response a status code equal to that given
func (o *GetArchitectSystempromptsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArchitectSystempromptsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsOK  %+v", 200, o.Payload)
}

func (o *GetArchitectSystempromptsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsOK  %+v", 200, o.Payload)
}

func (o *GetArchitectSystempromptsOK) GetPayload() *models.SystemPromptEntityListing {
	return o.Payload
}

func (o *GetArchitectSystempromptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemPromptEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsBadRequest creates a GetArchitectSystempromptsBadRequest with default headers values
func NewGetArchitectSystempromptsBadRequest() *GetArchitectSystempromptsBadRequest {
	return &GetArchitectSystempromptsBadRequest{}
}

/*
GetArchitectSystempromptsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectSystempromptsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts bad request response has a 2xx status code
func (o *GetArchitectSystempromptsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts bad request response has a 3xx status code
func (o *GetArchitectSystempromptsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts bad request response has a 4xx status code
func (o *GetArchitectSystempromptsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts bad request response has a 5xx status code
func (o *GetArchitectSystempromptsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts bad request response a status code equal to that given
func (o *GetArchitectSystempromptsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArchitectSystempromptsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectSystempromptsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectSystempromptsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsUnauthorized creates a GetArchitectSystempromptsUnauthorized with default headers values
func NewGetArchitectSystempromptsUnauthorized() *GetArchitectSystempromptsUnauthorized {
	return &GetArchitectSystempromptsUnauthorized{}
}

/*
GetArchitectSystempromptsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectSystempromptsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts unauthorized response has a 2xx status code
func (o *GetArchitectSystempromptsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts unauthorized response has a 3xx status code
func (o *GetArchitectSystempromptsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts unauthorized response has a 4xx status code
func (o *GetArchitectSystempromptsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts unauthorized response has a 5xx status code
func (o *GetArchitectSystempromptsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts unauthorized response a status code equal to that given
func (o *GetArchitectSystempromptsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArchitectSystempromptsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectSystempromptsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectSystempromptsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsForbidden creates a GetArchitectSystempromptsForbidden with default headers values
func NewGetArchitectSystempromptsForbidden() *GetArchitectSystempromptsForbidden {
	return &GetArchitectSystempromptsForbidden{}
}

/*
GetArchitectSystempromptsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectSystempromptsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts forbidden response has a 2xx status code
func (o *GetArchitectSystempromptsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts forbidden response has a 3xx status code
func (o *GetArchitectSystempromptsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts forbidden response has a 4xx status code
func (o *GetArchitectSystempromptsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts forbidden response has a 5xx status code
func (o *GetArchitectSystempromptsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts forbidden response a status code equal to that given
func (o *GetArchitectSystempromptsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArchitectSystempromptsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectSystempromptsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectSystempromptsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsNotFound creates a GetArchitectSystempromptsNotFound with default headers values
func NewGetArchitectSystempromptsNotFound() *GetArchitectSystempromptsNotFound {
	return &GetArchitectSystempromptsNotFound{}
}

/*
GetArchitectSystempromptsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetArchitectSystempromptsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts not found response has a 2xx status code
func (o *GetArchitectSystempromptsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts not found response has a 3xx status code
func (o *GetArchitectSystempromptsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts not found response has a 4xx status code
func (o *GetArchitectSystempromptsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts not found response has a 5xx status code
func (o *GetArchitectSystempromptsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts not found response a status code equal to that given
func (o *GetArchitectSystempromptsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArchitectSystempromptsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectSystempromptsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectSystempromptsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsRequestTimeout creates a GetArchitectSystempromptsRequestTimeout with default headers values
func NewGetArchitectSystempromptsRequestTimeout() *GetArchitectSystempromptsRequestTimeout {
	return &GetArchitectSystempromptsRequestTimeout{}
}

/*
GetArchitectSystempromptsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectSystempromptsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts request timeout response has a 2xx status code
func (o *GetArchitectSystempromptsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts request timeout response has a 3xx status code
func (o *GetArchitectSystempromptsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts request timeout response has a 4xx status code
func (o *GetArchitectSystempromptsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts request timeout response has a 5xx status code
func (o *GetArchitectSystempromptsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts request timeout response a status code equal to that given
func (o *GetArchitectSystempromptsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetArchitectSystempromptsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectSystempromptsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectSystempromptsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsRequestEntityTooLarge creates a GetArchitectSystempromptsRequestEntityTooLarge with default headers values
func NewGetArchitectSystempromptsRequestEntityTooLarge() *GetArchitectSystempromptsRequestEntityTooLarge {
	return &GetArchitectSystempromptsRequestEntityTooLarge{}
}

/*
GetArchitectSystempromptsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetArchitectSystempromptsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts request entity too large response has a 2xx status code
func (o *GetArchitectSystempromptsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts request entity too large response has a 3xx status code
func (o *GetArchitectSystempromptsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts request entity too large response has a 4xx status code
func (o *GetArchitectSystempromptsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts request entity too large response has a 5xx status code
func (o *GetArchitectSystempromptsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts request entity too large response a status code equal to that given
func (o *GetArchitectSystempromptsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetArchitectSystempromptsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectSystempromptsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectSystempromptsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsUnsupportedMediaType creates a GetArchitectSystempromptsUnsupportedMediaType with default headers values
func NewGetArchitectSystempromptsUnsupportedMediaType() *GetArchitectSystempromptsUnsupportedMediaType {
	return &GetArchitectSystempromptsUnsupportedMediaType{}
}

/*
GetArchitectSystempromptsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectSystempromptsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts unsupported media type response has a 2xx status code
func (o *GetArchitectSystempromptsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts unsupported media type response has a 3xx status code
func (o *GetArchitectSystempromptsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts unsupported media type response has a 4xx status code
func (o *GetArchitectSystempromptsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts unsupported media type response has a 5xx status code
func (o *GetArchitectSystempromptsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts unsupported media type response a status code equal to that given
func (o *GetArchitectSystempromptsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetArchitectSystempromptsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectSystempromptsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectSystempromptsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsTooManyRequests creates a GetArchitectSystempromptsTooManyRequests with default headers values
func NewGetArchitectSystempromptsTooManyRequests() *GetArchitectSystempromptsTooManyRequests {
	return &GetArchitectSystempromptsTooManyRequests{}
}

/*
GetArchitectSystempromptsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectSystempromptsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts too many requests response has a 2xx status code
func (o *GetArchitectSystempromptsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts too many requests response has a 3xx status code
func (o *GetArchitectSystempromptsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts too many requests response has a 4xx status code
func (o *GetArchitectSystempromptsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompts too many requests response has a 5xx status code
func (o *GetArchitectSystempromptsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompts too many requests response a status code equal to that given
func (o *GetArchitectSystempromptsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetArchitectSystempromptsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectSystempromptsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectSystempromptsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsInternalServerError creates a GetArchitectSystempromptsInternalServerError with default headers values
func NewGetArchitectSystempromptsInternalServerError() *GetArchitectSystempromptsInternalServerError {
	return &GetArchitectSystempromptsInternalServerError{}
}

/*
GetArchitectSystempromptsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectSystempromptsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts internal server error response has a 2xx status code
func (o *GetArchitectSystempromptsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts internal server error response has a 3xx status code
func (o *GetArchitectSystempromptsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts internal server error response has a 4xx status code
func (o *GetArchitectSystempromptsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompts internal server error response has a 5xx status code
func (o *GetArchitectSystempromptsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect systemprompts internal server error response a status code equal to that given
func (o *GetArchitectSystempromptsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArchitectSystempromptsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectSystempromptsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectSystempromptsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsServiceUnavailable creates a GetArchitectSystempromptsServiceUnavailable with default headers values
func NewGetArchitectSystempromptsServiceUnavailable() *GetArchitectSystempromptsServiceUnavailable {
	return &GetArchitectSystempromptsServiceUnavailable{}
}

/*
GetArchitectSystempromptsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectSystempromptsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts service unavailable response has a 2xx status code
func (o *GetArchitectSystempromptsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts service unavailable response has a 3xx status code
func (o *GetArchitectSystempromptsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts service unavailable response has a 4xx status code
func (o *GetArchitectSystempromptsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompts service unavailable response has a 5xx status code
func (o *GetArchitectSystempromptsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect systemprompts service unavailable response a status code equal to that given
func (o *GetArchitectSystempromptsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetArchitectSystempromptsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectSystempromptsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectSystempromptsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptsGatewayTimeout creates a GetArchitectSystempromptsGatewayTimeout with default headers values
func NewGetArchitectSystempromptsGatewayTimeout() *GetArchitectSystempromptsGatewayTimeout {
	return &GetArchitectSystempromptsGatewayTimeout{}
}

/*
GetArchitectSystempromptsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetArchitectSystempromptsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompts gateway timeout response has a 2xx status code
func (o *GetArchitectSystempromptsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompts gateway timeout response has a 3xx status code
func (o *GetArchitectSystempromptsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompts gateway timeout response has a 4xx status code
func (o *GetArchitectSystempromptsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompts gateway timeout response has a 5xx status code
func (o *GetArchitectSystempromptsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect systemprompts gateway timeout response a status code equal to that given
func (o *GetArchitectSystempromptsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetArchitectSystempromptsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectSystempromptsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts][%d] getArchitectSystempromptsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectSystempromptsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
