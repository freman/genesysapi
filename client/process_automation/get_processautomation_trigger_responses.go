// Code generated by go-swagger; DO NOT EDIT.

package process_automation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetProcessautomationTriggerReader is a Reader for the GetProcessautomationTrigger structure.
type GetProcessautomationTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessautomationTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessautomationTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessautomationTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessautomationTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessautomationTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessautomationTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetProcessautomationTriggerRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetProcessautomationTriggerRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetProcessautomationTriggerUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetProcessautomationTriggerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProcessautomationTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetProcessautomationTriggerServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetProcessautomationTriggerGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProcessautomationTriggerOK creates a GetProcessautomationTriggerOK with default headers values
func NewGetProcessautomationTriggerOK() *GetProcessautomationTriggerOK {
	return &GetProcessautomationTriggerOK{}
}

/*
GetProcessautomationTriggerOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessautomationTriggerOK struct {
	Payload *models.Trigger
}

// IsSuccess returns true when this get processautomation trigger o k response has a 2xx status code
func (o *GetProcessautomationTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get processautomation trigger o k response has a 3xx status code
func (o *GetProcessautomationTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger o k response has a 4xx status code
func (o *GetProcessautomationTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation trigger o k response has a 5xx status code
func (o *GetProcessautomationTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger o k response a status code equal to that given
func (o *GetProcessautomationTriggerOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProcessautomationTriggerOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerOK  %+v", 200, o.Payload)
}

func (o *GetProcessautomationTriggerOK) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerOK  %+v", 200, o.Payload)
}

func (o *GetProcessautomationTriggerOK) GetPayload() *models.Trigger {
	return o.Payload
}

func (o *GetProcessautomationTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Trigger)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerBadRequest creates a GetProcessautomationTriggerBadRequest with default headers values
func NewGetProcessautomationTriggerBadRequest() *GetProcessautomationTriggerBadRequest {
	return &GetProcessautomationTriggerBadRequest{}
}

/*
GetProcessautomationTriggerBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetProcessautomationTriggerBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger bad request response has a 2xx status code
func (o *GetProcessautomationTriggerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger bad request response has a 3xx status code
func (o *GetProcessautomationTriggerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger bad request response has a 4xx status code
func (o *GetProcessautomationTriggerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger bad request response has a 5xx status code
func (o *GetProcessautomationTriggerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger bad request response a status code equal to that given
func (o *GetProcessautomationTriggerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetProcessautomationTriggerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *GetProcessautomationTriggerBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *GetProcessautomationTriggerBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerUnauthorized creates a GetProcessautomationTriggerUnauthorized with default headers values
func NewGetProcessautomationTriggerUnauthorized() *GetProcessautomationTriggerUnauthorized {
	return &GetProcessautomationTriggerUnauthorized{}
}

/*
GetProcessautomationTriggerUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetProcessautomationTriggerUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger unauthorized response has a 2xx status code
func (o *GetProcessautomationTriggerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger unauthorized response has a 3xx status code
func (o *GetProcessautomationTriggerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger unauthorized response has a 4xx status code
func (o *GetProcessautomationTriggerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger unauthorized response has a 5xx status code
func (o *GetProcessautomationTriggerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger unauthorized response a status code equal to that given
func (o *GetProcessautomationTriggerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetProcessautomationTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProcessautomationTriggerUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProcessautomationTriggerUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerForbidden creates a GetProcessautomationTriggerForbidden with default headers values
func NewGetProcessautomationTriggerForbidden() *GetProcessautomationTriggerForbidden {
	return &GetProcessautomationTriggerForbidden{}
}

/*
GetProcessautomationTriggerForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetProcessautomationTriggerForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger forbidden response has a 2xx status code
func (o *GetProcessautomationTriggerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger forbidden response has a 3xx status code
func (o *GetProcessautomationTriggerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger forbidden response has a 4xx status code
func (o *GetProcessautomationTriggerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger forbidden response has a 5xx status code
func (o *GetProcessautomationTriggerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger forbidden response a status code equal to that given
func (o *GetProcessautomationTriggerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetProcessautomationTriggerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerForbidden  %+v", 403, o.Payload)
}

func (o *GetProcessautomationTriggerForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerForbidden  %+v", 403, o.Payload)
}

func (o *GetProcessautomationTriggerForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerNotFound creates a GetProcessautomationTriggerNotFound with default headers values
func NewGetProcessautomationTriggerNotFound() *GetProcessautomationTriggerNotFound {
	return &GetProcessautomationTriggerNotFound{}
}

/*
GetProcessautomationTriggerNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetProcessautomationTriggerNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger not found response has a 2xx status code
func (o *GetProcessautomationTriggerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger not found response has a 3xx status code
func (o *GetProcessautomationTriggerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger not found response has a 4xx status code
func (o *GetProcessautomationTriggerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger not found response has a 5xx status code
func (o *GetProcessautomationTriggerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger not found response a status code equal to that given
func (o *GetProcessautomationTriggerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetProcessautomationTriggerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerNotFound  %+v", 404, o.Payload)
}

func (o *GetProcessautomationTriggerNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerNotFound  %+v", 404, o.Payload)
}

func (o *GetProcessautomationTriggerNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerRequestTimeout creates a GetProcessautomationTriggerRequestTimeout with default headers values
func NewGetProcessautomationTriggerRequestTimeout() *GetProcessautomationTriggerRequestTimeout {
	return &GetProcessautomationTriggerRequestTimeout{}
}

/*
GetProcessautomationTriggerRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetProcessautomationTriggerRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger request timeout response has a 2xx status code
func (o *GetProcessautomationTriggerRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger request timeout response has a 3xx status code
func (o *GetProcessautomationTriggerRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger request timeout response has a 4xx status code
func (o *GetProcessautomationTriggerRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger request timeout response has a 5xx status code
func (o *GetProcessautomationTriggerRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger request timeout response a status code equal to that given
func (o *GetProcessautomationTriggerRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetProcessautomationTriggerRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetProcessautomationTriggerRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetProcessautomationTriggerRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerRequestEntityTooLarge creates a GetProcessautomationTriggerRequestEntityTooLarge with default headers values
func NewGetProcessautomationTriggerRequestEntityTooLarge() *GetProcessautomationTriggerRequestEntityTooLarge {
	return &GetProcessautomationTriggerRequestEntityTooLarge{}
}

/*
GetProcessautomationTriggerRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetProcessautomationTriggerRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger request entity too large response has a 2xx status code
func (o *GetProcessautomationTriggerRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger request entity too large response has a 3xx status code
func (o *GetProcessautomationTriggerRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger request entity too large response has a 4xx status code
func (o *GetProcessautomationTriggerRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger request entity too large response has a 5xx status code
func (o *GetProcessautomationTriggerRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger request entity too large response a status code equal to that given
func (o *GetProcessautomationTriggerRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetProcessautomationTriggerRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetProcessautomationTriggerRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetProcessautomationTriggerRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerUnsupportedMediaType creates a GetProcessautomationTriggerUnsupportedMediaType with default headers values
func NewGetProcessautomationTriggerUnsupportedMediaType() *GetProcessautomationTriggerUnsupportedMediaType {
	return &GetProcessautomationTriggerUnsupportedMediaType{}
}

/*
GetProcessautomationTriggerUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetProcessautomationTriggerUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger unsupported media type response has a 2xx status code
func (o *GetProcessautomationTriggerUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger unsupported media type response has a 3xx status code
func (o *GetProcessautomationTriggerUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger unsupported media type response has a 4xx status code
func (o *GetProcessautomationTriggerUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger unsupported media type response has a 5xx status code
func (o *GetProcessautomationTriggerUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger unsupported media type response a status code equal to that given
func (o *GetProcessautomationTriggerUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetProcessautomationTriggerUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetProcessautomationTriggerUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetProcessautomationTriggerUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerTooManyRequests creates a GetProcessautomationTriggerTooManyRequests with default headers values
func NewGetProcessautomationTriggerTooManyRequests() *GetProcessautomationTriggerTooManyRequests {
	return &GetProcessautomationTriggerTooManyRequests{}
}

/*
GetProcessautomationTriggerTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetProcessautomationTriggerTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger too many requests response has a 2xx status code
func (o *GetProcessautomationTriggerTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger too many requests response has a 3xx status code
func (o *GetProcessautomationTriggerTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger too many requests response has a 4xx status code
func (o *GetProcessautomationTriggerTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation trigger too many requests response has a 5xx status code
func (o *GetProcessautomationTriggerTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation trigger too many requests response a status code equal to that given
func (o *GetProcessautomationTriggerTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetProcessautomationTriggerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetProcessautomationTriggerTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetProcessautomationTriggerTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerInternalServerError creates a GetProcessautomationTriggerInternalServerError with default headers values
func NewGetProcessautomationTriggerInternalServerError() *GetProcessautomationTriggerInternalServerError {
	return &GetProcessautomationTriggerInternalServerError{}
}

/*
GetProcessautomationTriggerInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetProcessautomationTriggerInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger internal server error response has a 2xx status code
func (o *GetProcessautomationTriggerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger internal server error response has a 3xx status code
func (o *GetProcessautomationTriggerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger internal server error response has a 4xx status code
func (o *GetProcessautomationTriggerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation trigger internal server error response has a 5xx status code
func (o *GetProcessautomationTriggerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get processautomation trigger internal server error response a status code equal to that given
func (o *GetProcessautomationTriggerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetProcessautomationTriggerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProcessautomationTriggerInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProcessautomationTriggerInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerServiceUnavailable creates a GetProcessautomationTriggerServiceUnavailable with default headers values
func NewGetProcessautomationTriggerServiceUnavailable() *GetProcessautomationTriggerServiceUnavailable {
	return &GetProcessautomationTriggerServiceUnavailable{}
}

/*
GetProcessautomationTriggerServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetProcessautomationTriggerServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger service unavailable response has a 2xx status code
func (o *GetProcessautomationTriggerServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger service unavailable response has a 3xx status code
func (o *GetProcessautomationTriggerServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger service unavailable response has a 4xx status code
func (o *GetProcessautomationTriggerServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation trigger service unavailable response has a 5xx status code
func (o *GetProcessautomationTriggerServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get processautomation trigger service unavailable response a status code equal to that given
func (o *GetProcessautomationTriggerServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetProcessautomationTriggerServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetProcessautomationTriggerServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetProcessautomationTriggerServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggerGatewayTimeout creates a GetProcessautomationTriggerGatewayTimeout with default headers values
func NewGetProcessautomationTriggerGatewayTimeout() *GetProcessautomationTriggerGatewayTimeout {
	return &GetProcessautomationTriggerGatewayTimeout{}
}

/*
GetProcessautomationTriggerGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetProcessautomationTriggerGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation trigger gateway timeout response has a 2xx status code
func (o *GetProcessautomationTriggerGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation trigger gateway timeout response has a 3xx status code
func (o *GetProcessautomationTriggerGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation trigger gateway timeout response has a 4xx status code
func (o *GetProcessautomationTriggerGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation trigger gateway timeout response has a 5xx status code
func (o *GetProcessautomationTriggerGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get processautomation trigger gateway timeout response a status code equal to that given
func (o *GetProcessautomationTriggerGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetProcessautomationTriggerGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetProcessautomationTriggerGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/{triggerId}][%d] getProcessautomationTriggerGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetProcessautomationTriggerGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggerGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
