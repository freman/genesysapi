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

// GetSystempresencesReader is a Reader for the GetSystempresences structure.
type GetSystempresencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystempresencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSystempresencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSystempresencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSystempresencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSystempresencesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSystempresencesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSystempresencesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSystempresencesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSystempresencesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSystempresencesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSystempresencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSystempresencesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSystempresencesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSystempresencesOK creates a GetSystempresencesOK with default headers values
func NewGetSystempresencesOK() *GetSystempresencesOK {
	return &GetSystempresencesOK{}
}

/*
GetSystempresencesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSystempresencesOK struct {
	Payload []*models.SystemPresence
}

// IsSuccess returns true when this get systempresences o k response has a 2xx status code
func (o *GetSystempresencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get systempresences o k response has a 3xx status code
func (o *GetSystempresencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences o k response has a 4xx status code
func (o *GetSystempresencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get systempresences o k response has a 5xx status code
func (o *GetSystempresencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences o k response a status code equal to that given
func (o *GetSystempresencesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSystempresencesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesOK  %+v", 200, o.Payload)
}

func (o *GetSystempresencesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesOK  %+v", 200, o.Payload)
}

func (o *GetSystempresencesOK) GetPayload() []*models.SystemPresence {
	return o.Payload
}

func (o *GetSystempresencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesBadRequest creates a GetSystempresencesBadRequest with default headers values
func NewGetSystempresencesBadRequest() *GetSystempresencesBadRequest {
	return &GetSystempresencesBadRequest{}
}

/*
GetSystempresencesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSystempresencesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences bad request response has a 2xx status code
func (o *GetSystempresencesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences bad request response has a 3xx status code
func (o *GetSystempresencesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences bad request response has a 4xx status code
func (o *GetSystempresencesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences bad request response has a 5xx status code
func (o *GetSystempresencesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences bad request response a status code equal to that given
func (o *GetSystempresencesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSystempresencesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesBadRequest  %+v", 400, o.Payload)
}

func (o *GetSystempresencesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesBadRequest  %+v", 400, o.Payload)
}

func (o *GetSystempresencesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesUnauthorized creates a GetSystempresencesUnauthorized with default headers values
func NewGetSystempresencesUnauthorized() *GetSystempresencesUnauthorized {
	return &GetSystempresencesUnauthorized{}
}

/*
GetSystempresencesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSystempresencesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences unauthorized response has a 2xx status code
func (o *GetSystempresencesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences unauthorized response has a 3xx status code
func (o *GetSystempresencesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences unauthorized response has a 4xx status code
func (o *GetSystempresencesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences unauthorized response has a 5xx status code
func (o *GetSystempresencesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences unauthorized response a status code equal to that given
func (o *GetSystempresencesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSystempresencesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSystempresencesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSystempresencesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesForbidden creates a GetSystempresencesForbidden with default headers values
func NewGetSystempresencesForbidden() *GetSystempresencesForbidden {
	return &GetSystempresencesForbidden{}
}

/*
GetSystempresencesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetSystempresencesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences forbidden response has a 2xx status code
func (o *GetSystempresencesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences forbidden response has a 3xx status code
func (o *GetSystempresencesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences forbidden response has a 4xx status code
func (o *GetSystempresencesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences forbidden response has a 5xx status code
func (o *GetSystempresencesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences forbidden response a status code equal to that given
func (o *GetSystempresencesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSystempresencesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesForbidden  %+v", 403, o.Payload)
}

func (o *GetSystempresencesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesForbidden  %+v", 403, o.Payload)
}

func (o *GetSystempresencesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesNotFound creates a GetSystempresencesNotFound with default headers values
func NewGetSystempresencesNotFound() *GetSystempresencesNotFound {
	return &GetSystempresencesNotFound{}
}

/*
GetSystempresencesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetSystempresencesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences not found response has a 2xx status code
func (o *GetSystempresencesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences not found response has a 3xx status code
func (o *GetSystempresencesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences not found response has a 4xx status code
func (o *GetSystempresencesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences not found response has a 5xx status code
func (o *GetSystempresencesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences not found response a status code equal to that given
func (o *GetSystempresencesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSystempresencesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesNotFound  %+v", 404, o.Payload)
}

func (o *GetSystempresencesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesNotFound  %+v", 404, o.Payload)
}

func (o *GetSystempresencesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesRequestTimeout creates a GetSystempresencesRequestTimeout with default headers values
func NewGetSystempresencesRequestTimeout() *GetSystempresencesRequestTimeout {
	return &GetSystempresencesRequestTimeout{}
}

/*
GetSystempresencesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSystempresencesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences request timeout response has a 2xx status code
func (o *GetSystempresencesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences request timeout response has a 3xx status code
func (o *GetSystempresencesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences request timeout response has a 4xx status code
func (o *GetSystempresencesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences request timeout response has a 5xx status code
func (o *GetSystempresencesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences request timeout response a status code equal to that given
func (o *GetSystempresencesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetSystempresencesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSystempresencesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSystempresencesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesRequestEntityTooLarge creates a GetSystempresencesRequestEntityTooLarge with default headers values
func NewGetSystempresencesRequestEntityTooLarge() *GetSystempresencesRequestEntityTooLarge {
	return &GetSystempresencesRequestEntityTooLarge{}
}

/*
GetSystempresencesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetSystempresencesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences request entity too large response has a 2xx status code
func (o *GetSystempresencesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences request entity too large response has a 3xx status code
func (o *GetSystempresencesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences request entity too large response has a 4xx status code
func (o *GetSystempresencesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences request entity too large response has a 5xx status code
func (o *GetSystempresencesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences request entity too large response a status code equal to that given
func (o *GetSystempresencesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetSystempresencesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSystempresencesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSystempresencesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesUnsupportedMediaType creates a GetSystempresencesUnsupportedMediaType with default headers values
func NewGetSystempresencesUnsupportedMediaType() *GetSystempresencesUnsupportedMediaType {
	return &GetSystempresencesUnsupportedMediaType{}
}

/*
GetSystempresencesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSystempresencesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences unsupported media type response has a 2xx status code
func (o *GetSystempresencesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences unsupported media type response has a 3xx status code
func (o *GetSystempresencesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences unsupported media type response has a 4xx status code
func (o *GetSystempresencesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences unsupported media type response has a 5xx status code
func (o *GetSystempresencesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences unsupported media type response a status code equal to that given
func (o *GetSystempresencesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetSystempresencesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSystempresencesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSystempresencesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesTooManyRequests creates a GetSystempresencesTooManyRequests with default headers values
func NewGetSystempresencesTooManyRequests() *GetSystempresencesTooManyRequests {
	return &GetSystempresencesTooManyRequests{}
}

/*
GetSystempresencesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSystempresencesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences too many requests response has a 2xx status code
func (o *GetSystempresencesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences too many requests response has a 3xx status code
func (o *GetSystempresencesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences too many requests response has a 4xx status code
func (o *GetSystempresencesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get systempresences too many requests response has a 5xx status code
func (o *GetSystempresencesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get systempresences too many requests response a status code equal to that given
func (o *GetSystempresencesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSystempresencesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSystempresencesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSystempresencesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesInternalServerError creates a GetSystempresencesInternalServerError with default headers values
func NewGetSystempresencesInternalServerError() *GetSystempresencesInternalServerError {
	return &GetSystempresencesInternalServerError{}
}

/*
GetSystempresencesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSystempresencesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences internal server error response has a 2xx status code
func (o *GetSystempresencesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences internal server error response has a 3xx status code
func (o *GetSystempresencesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences internal server error response has a 4xx status code
func (o *GetSystempresencesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get systempresences internal server error response has a 5xx status code
func (o *GetSystempresencesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get systempresences internal server error response a status code equal to that given
func (o *GetSystempresencesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSystempresencesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystempresencesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSystempresencesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesServiceUnavailable creates a GetSystempresencesServiceUnavailable with default headers values
func NewGetSystempresencesServiceUnavailable() *GetSystempresencesServiceUnavailable {
	return &GetSystempresencesServiceUnavailable{}
}

/*
GetSystempresencesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSystempresencesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences service unavailable response has a 2xx status code
func (o *GetSystempresencesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences service unavailable response has a 3xx status code
func (o *GetSystempresencesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences service unavailable response has a 4xx status code
func (o *GetSystempresencesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get systempresences service unavailable response has a 5xx status code
func (o *GetSystempresencesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get systempresences service unavailable response a status code equal to that given
func (o *GetSystempresencesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSystempresencesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSystempresencesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSystempresencesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSystempresencesGatewayTimeout creates a GetSystempresencesGatewayTimeout with default headers values
func NewGetSystempresencesGatewayTimeout() *GetSystempresencesGatewayTimeout {
	return &GetSystempresencesGatewayTimeout{}
}

/*
GetSystempresencesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetSystempresencesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get systempresences gateway timeout response has a 2xx status code
func (o *GetSystempresencesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get systempresences gateway timeout response has a 3xx status code
func (o *GetSystempresencesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get systempresences gateway timeout response has a 4xx status code
func (o *GetSystempresencesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get systempresences gateway timeout response has a 5xx status code
func (o *GetSystempresencesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get systempresences gateway timeout response a status code equal to that given
func (o *GetSystempresencesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetSystempresencesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSystempresencesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/systempresences][%d] getSystempresencesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSystempresencesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSystempresencesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
