// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetPresenceSourcesReader is a Reader for the GetPresenceSources structure.
type GetPresenceSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPresenceSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPresenceSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPresenceSourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPresenceSourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPresenceSourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPresenceSourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetPresenceSourcesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetPresenceSourcesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetPresenceSourcesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPresenceSourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPresenceSourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetPresenceSourcesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetPresenceSourcesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPresenceSourcesOK creates a GetPresenceSourcesOK with default headers values
func NewGetPresenceSourcesOK() *GetPresenceSourcesOK {
	return &GetPresenceSourcesOK{}
}

/*
GetPresenceSourcesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPresenceSourcesOK struct {
	Payload *models.SourceEntityListing
}

// IsSuccess returns true when this get presence sources o k response has a 2xx status code
func (o *GetPresenceSourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get presence sources o k response has a 3xx status code
func (o *GetPresenceSourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources o k response has a 4xx status code
func (o *GetPresenceSourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presence sources o k response has a 5xx status code
func (o *GetPresenceSourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources o k response a status code equal to that given
func (o *GetPresenceSourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPresenceSourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesOK  %+v", 200, o.Payload)
}

func (o *GetPresenceSourcesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesOK  %+v", 200, o.Payload)
}

func (o *GetPresenceSourcesOK) GetPayload() *models.SourceEntityListing {
	return o.Payload
}

func (o *GetPresenceSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesBadRequest creates a GetPresenceSourcesBadRequest with default headers values
func NewGetPresenceSourcesBadRequest() *GetPresenceSourcesBadRequest {
	return &GetPresenceSourcesBadRequest{}
}

/*
GetPresenceSourcesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetPresenceSourcesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources bad request response has a 2xx status code
func (o *GetPresenceSourcesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources bad request response has a 3xx status code
func (o *GetPresenceSourcesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources bad request response has a 4xx status code
func (o *GetPresenceSourcesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources bad request response has a 5xx status code
func (o *GetPresenceSourcesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources bad request response a status code equal to that given
func (o *GetPresenceSourcesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPresenceSourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetPresenceSourcesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetPresenceSourcesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesUnauthorized creates a GetPresenceSourcesUnauthorized with default headers values
func NewGetPresenceSourcesUnauthorized() *GetPresenceSourcesUnauthorized {
	return &GetPresenceSourcesUnauthorized{}
}

/*
GetPresenceSourcesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetPresenceSourcesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources unauthorized response has a 2xx status code
func (o *GetPresenceSourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources unauthorized response has a 3xx status code
func (o *GetPresenceSourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources unauthorized response has a 4xx status code
func (o *GetPresenceSourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources unauthorized response has a 5xx status code
func (o *GetPresenceSourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources unauthorized response a status code equal to that given
func (o *GetPresenceSourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPresenceSourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPresenceSourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPresenceSourcesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesForbidden creates a GetPresenceSourcesForbidden with default headers values
func NewGetPresenceSourcesForbidden() *GetPresenceSourcesForbidden {
	return &GetPresenceSourcesForbidden{}
}

/*
GetPresenceSourcesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetPresenceSourcesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources forbidden response has a 2xx status code
func (o *GetPresenceSourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources forbidden response has a 3xx status code
func (o *GetPresenceSourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources forbidden response has a 4xx status code
func (o *GetPresenceSourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources forbidden response has a 5xx status code
func (o *GetPresenceSourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources forbidden response a status code equal to that given
func (o *GetPresenceSourcesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPresenceSourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetPresenceSourcesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetPresenceSourcesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesNotFound creates a GetPresenceSourcesNotFound with default headers values
func NewGetPresenceSourcesNotFound() *GetPresenceSourcesNotFound {
	return &GetPresenceSourcesNotFound{}
}

/*
GetPresenceSourcesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetPresenceSourcesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources not found response has a 2xx status code
func (o *GetPresenceSourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources not found response has a 3xx status code
func (o *GetPresenceSourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources not found response has a 4xx status code
func (o *GetPresenceSourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources not found response has a 5xx status code
func (o *GetPresenceSourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources not found response a status code equal to that given
func (o *GetPresenceSourcesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPresenceSourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetPresenceSourcesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetPresenceSourcesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesRequestTimeout creates a GetPresenceSourcesRequestTimeout with default headers values
func NewGetPresenceSourcesRequestTimeout() *GetPresenceSourcesRequestTimeout {
	return &GetPresenceSourcesRequestTimeout{}
}

/*
GetPresenceSourcesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetPresenceSourcesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources request timeout response has a 2xx status code
func (o *GetPresenceSourcesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources request timeout response has a 3xx status code
func (o *GetPresenceSourcesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources request timeout response has a 4xx status code
func (o *GetPresenceSourcesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources request timeout response has a 5xx status code
func (o *GetPresenceSourcesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources request timeout response a status code equal to that given
func (o *GetPresenceSourcesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetPresenceSourcesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetPresenceSourcesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetPresenceSourcesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesRequestEntityTooLarge creates a GetPresenceSourcesRequestEntityTooLarge with default headers values
func NewGetPresenceSourcesRequestEntityTooLarge() *GetPresenceSourcesRequestEntityTooLarge {
	return &GetPresenceSourcesRequestEntityTooLarge{}
}

/*
GetPresenceSourcesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetPresenceSourcesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources request entity too large response has a 2xx status code
func (o *GetPresenceSourcesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources request entity too large response has a 3xx status code
func (o *GetPresenceSourcesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources request entity too large response has a 4xx status code
func (o *GetPresenceSourcesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources request entity too large response has a 5xx status code
func (o *GetPresenceSourcesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources request entity too large response a status code equal to that given
func (o *GetPresenceSourcesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetPresenceSourcesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetPresenceSourcesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetPresenceSourcesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesUnsupportedMediaType creates a GetPresenceSourcesUnsupportedMediaType with default headers values
func NewGetPresenceSourcesUnsupportedMediaType() *GetPresenceSourcesUnsupportedMediaType {
	return &GetPresenceSourcesUnsupportedMediaType{}
}

/*
GetPresenceSourcesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetPresenceSourcesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources unsupported media type response has a 2xx status code
func (o *GetPresenceSourcesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources unsupported media type response has a 3xx status code
func (o *GetPresenceSourcesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources unsupported media type response has a 4xx status code
func (o *GetPresenceSourcesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources unsupported media type response has a 5xx status code
func (o *GetPresenceSourcesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources unsupported media type response a status code equal to that given
func (o *GetPresenceSourcesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetPresenceSourcesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPresenceSourcesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetPresenceSourcesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesTooManyRequests creates a GetPresenceSourcesTooManyRequests with default headers values
func NewGetPresenceSourcesTooManyRequests() *GetPresenceSourcesTooManyRequests {
	return &GetPresenceSourcesTooManyRequests{}
}

/*
GetPresenceSourcesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetPresenceSourcesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources too many requests response has a 2xx status code
func (o *GetPresenceSourcesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources too many requests response has a 3xx status code
func (o *GetPresenceSourcesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources too many requests response has a 4xx status code
func (o *GetPresenceSourcesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get presence sources too many requests response has a 5xx status code
func (o *GetPresenceSourcesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get presence sources too many requests response a status code equal to that given
func (o *GetPresenceSourcesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetPresenceSourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPresenceSourcesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPresenceSourcesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesInternalServerError creates a GetPresenceSourcesInternalServerError with default headers values
func NewGetPresenceSourcesInternalServerError() *GetPresenceSourcesInternalServerError {
	return &GetPresenceSourcesInternalServerError{}
}

/*
GetPresenceSourcesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetPresenceSourcesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources internal server error response has a 2xx status code
func (o *GetPresenceSourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources internal server error response has a 3xx status code
func (o *GetPresenceSourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources internal server error response has a 4xx status code
func (o *GetPresenceSourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presence sources internal server error response has a 5xx status code
func (o *GetPresenceSourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get presence sources internal server error response a status code equal to that given
func (o *GetPresenceSourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPresenceSourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPresenceSourcesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPresenceSourcesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesServiceUnavailable creates a GetPresenceSourcesServiceUnavailable with default headers values
func NewGetPresenceSourcesServiceUnavailable() *GetPresenceSourcesServiceUnavailable {
	return &GetPresenceSourcesServiceUnavailable{}
}

/*
GetPresenceSourcesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetPresenceSourcesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources service unavailable response has a 2xx status code
func (o *GetPresenceSourcesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources service unavailable response has a 3xx status code
func (o *GetPresenceSourcesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources service unavailable response has a 4xx status code
func (o *GetPresenceSourcesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presence sources service unavailable response has a 5xx status code
func (o *GetPresenceSourcesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get presence sources service unavailable response a status code equal to that given
func (o *GetPresenceSourcesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetPresenceSourcesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPresenceSourcesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetPresenceSourcesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresenceSourcesGatewayTimeout creates a GetPresenceSourcesGatewayTimeout with default headers values
func NewGetPresenceSourcesGatewayTimeout() *GetPresenceSourcesGatewayTimeout {
	return &GetPresenceSourcesGatewayTimeout{}
}

/*
GetPresenceSourcesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetPresenceSourcesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get presence sources gateway timeout response has a 2xx status code
func (o *GetPresenceSourcesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get presence sources gateway timeout response has a 3xx status code
func (o *GetPresenceSourcesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get presence sources gateway timeout response has a 4xx status code
func (o *GetPresenceSourcesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get presence sources gateway timeout response has a 5xx status code
func (o *GetPresenceSourcesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get presence sources gateway timeout response a status code equal to that given
func (o *GetPresenceSourcesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetPresenceSourcesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetPresenceSourcesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/presence/sources][%d] getPresenceSourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetPresenceSourcesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetPresenceSourcesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
