// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutGamificationStatusReader is a Reader for the PutGamificationStatus structure.
type PutGamificationStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGamificationStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutGamificationStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGamificationStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutGamificationStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutGamificationStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutGamificationStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutGamificationStatusRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutGamificationStatusRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutGamificationStatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutGamificationStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutGamificationStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutGamificationStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutGamificationStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutGamificationStatusOK creates a PutGamificationStatusOK with default headers values
func NewPutGamificationStatusOK() *PutGamificationStatusOK {
	return &PutGamificationStatusOK{}
}

/*
PutGamificationStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type PutGamificationStatusOK struct {
	Payload *models.GamificationStatus
}

// IsSuccess returns true when this put gamification status o k response has a 2xx status code
func (o *PutGamificationStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put gamification status o k response has a 3xx status code
func (o *PutGamificationStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status o k response has a 4xx status code
func (o *PutGamificationStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put gamification status o k response has a 5xx status code
func (o *PutGamificationStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status o k response a status code equal to that given
func (o *PutGamificationStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutGamificationStatusOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusOK  %+v", 200, o.Payload)
}

func (o *PutGamificationStatusOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusOK  %+v", 200, o.Payload)
}

func (o *PutGamificationStatusOK) GetPayload() *models.GamificationStatus {
	return o.Payload
}

func (o *PutGamificationStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GamificationStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusBadRequest creates a PutGamificationStatusBadRequest with default headers values
func NewPutGamificationStatusBadRequest() *PutGamificationStatusBadRequest {
	return &PutGamificationStatusBadRequest{}
}

/*
PutGamificationStatusBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutGamificationStatusBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status bad request response has a 2xx status code
func (o *PutGamificationStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status bad request response has a 3xx status code
func (o *PutGamificationStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status bad request response has a 4xx status code
func (o *PutGamificationStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status bad request response has a 5xx status code
func (o *PutGamificationStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status bad request response a status code equal to that given
func (o *PutGamificationStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutGamificationStatusBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusBadRequest  %+v", 400, o.Payload)
}

func (o *PutGamificationStatusBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusBadRequest  %+v", 400, o.Payload)
}

func (o *PutGamificationStatusBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusUnauthorized creates a PutGamificationStatusUnauthorized with default headers values
func NewPutGamificationStatusUnauthorized() *PutGamificationStatusUnauthorized {
	return &PutGamificationStatusUnauthorized{}
}

/*
PutGamificationStatusUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutGamificationStatusUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status unauthorized response has a 2xx status code
func (o *PutGamificationStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status unauthorized response has a 3xx status code
func (o *PutGamificationStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status unauthorized response has a 4xx status code
func (o *PutGamificationStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status unauthorized response has a 5xx status code
func (o *PutGamificationStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status unauthorized response a status code equal to that given
func (o *PutGamificationStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutGamificationStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGamificationStatusUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGamificationStatusUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusForbidden creates a PutGamificationStatusForbidden with default headers values
func NewPutGamificationStatusForbidden() *PutGamificationStatusForbidden {
	return &PutGamificationStatusForbidden{}
}

/*
PutGamificationStatusForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutGamificationStatusForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status forbidden response has a 2xx status code
func (o *PutGamificationStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status forbidden response has a 3xx status code
func (o *PutGamificationStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status forbidden response has a 4xx status code
func (o *PutGamificationStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status forbidden response has a 5xx status code
func (o *PutGamificationStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status forbidden response a status code equal to that given
func (o *PutGamificationStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutGamificationStatusForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusForbidden  %+v", 403, o.Payload)
}

func (o *PutGamificationStatusForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusForbidden  %+v", 403, o.Payload)
}

func (o *PutGamificationStatusForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusNotFound creates a PutGamificationStatusNotFound with default headers values
func NewPutGamificationStatusNotFound() *PutGamificationStatusNotFound {
	return &PutGamificationStatusNotFound{}
}

/*
PutGamificationStatusNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutGamificationStatusNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status not found response has a 2xx status code
func (o *PutGamificationStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status not found response has a 3xx status code
func (o *PutGamificationStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status not found response has a 4xx status code
func (o *PutGamificationStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status not found response has a 5xx status code
func (o *PutGamificationStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status not found response a status code equal to that given
func (o *PutGamificationStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutGamificationStatusNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusNotFound  %+v", 404, o.Payload)
}

func (o *PutGamificationStatusNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusNotFound  %+v", 404, o.Payload)
}

func (o *PutGamificationStatusNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusRequestTimeout creates a PutGamificationStatusRequestTimeout with default headers values
func NewPutGamificationStatusRequestTimeout() *PutGamificationStatusRequestTimeout {
	return &PutGamificationStatusRequestTimeout{}
}

/*
PutGamificationStatusRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutGamificationStatusRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status request timeout response has a 2xx status code
func (o *PutGamificationStatusRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status request timeout response has a 3xx status code
func (o *PutGamificationStatusRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status request timeout response has a 4xx status code
func (o *PutGamificationStatusRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status request timeout response has a 5xx status code
func (o *PutGamificationStatusRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status request timeout response a status code equal to that given
func (o *PutGamificationStatusRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutGamificationStatusRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutGamificationStatusRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutGamificationStatusRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusRequestEntityTooLarge creates a PutGamificationStatusRequestEntityTooLarge with default headers values
func NewPutGamificationStatusRequestEntityTooLarge() *PutGamificationStatusRequestEntityTooLarge {
	return &PutGamificationStatusRequestEntityTooLarge{}
}

/*
PutGamificationStatusRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutGamificationStatusRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status request entity too large response has a 2xx status code
func (o *PutGamificationStatusRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status request entity too large response has a 3xx status code
func (o *PutGamificationStatusRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status request entity too large response has a 4xx status code
func (o *PutGamificationStatusRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status request entity too large response has a 5xx status code
func (o *PutGamificationStatusRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status request entity too large response a status code equal to that given
func (o *PutGamificationStatusRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutGamificationStatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGamificationStatusRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGamificationStatusRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusUnsupportedMediaType creates a PutGamificationStatusUnsupportedMediaType with default headers values
func NewPutGamificationStatusUnsupportedMediaType() *PutGamificationStatusUnsupportedMediaType {
	return &PutGamificationStatusUnsupportedMediaType{}
}

/*
PutGamificationStatusUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutGamificationStatusUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status unsupported media type response has a 2xx status code
func (o *PutGamificationStatusUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status unsupported media type response has a 3xx status code
func (o *PutGamificationStatusUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status unsupported media type response has a 4xx status code
func (o *PutGamificationStatusUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status unsupported media type response has a 5xx status code
func (o *PutGamificationStatusUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status unsupported media type response a status code equal to that given
func (o *PutGamificationStatusUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutGamificationStatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGamificationStatusUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGamificationStatusUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusTooManyRequests creates a PutGamificationStatusTooManyRequests with default headers values
func NewPutGamificationStatusTooManyRequests() *PutGamificationStatusTooManyRequests {
	return &PutGamificationStatusTooManyRequests{}
}

/*
PutGamificationStatusTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutGamificationStatusTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status too many requests response has a 2xx status code
func (o *PutGamificationStatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status too many requests response has a 3xx status code
func (o *PutGamificationStatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status too many requests response has a 4xx status code
func (o *PutGamificationStatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put gamification status too many requests response has a 5xx status code
func (o *PutGamificationStatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put gamification status too many requests response a status code equal to that given
func (o *PutGamificationStatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutGamificationStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGamificationStatusTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGamificationStatusTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusInternalServerError creates a PutGamificationStatusInternalServerError with default headers values
func NewPutGamificationStatusInternalServerError() *PutGamificationStatusInternalServerError {
	return &PutGamificationStatusInternalServerError{}
}

/*
PutGamificationStatusInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutGamificationStatusInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status internal server error response has a 2xx status code
func (o *PutGamificationStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status internal server error response has a 3xx status code
func (o *PutGamificationStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status internal server error response has a 4xx status code
func (o *PutGamificationStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put gamification status internal server error response has a 5xx status code
func (o *PutGamificationStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put gamification status internal server error response a status code equal to that given
func (o *PutGamificationStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutGamificationStatusInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGamificationStatusInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGamificationStatusInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusServiceUnavailable creates a PutGamificationStatusServiceUnavailable with default headers values
func NewPutGamificationStatusServiceUnavailable() *PutGamificationStatusServiceUnavailable {
	return &PutGamificationStatusServiceUnavailable{}
}

/*
PutGamificationStatusServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutGamificationStatusServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status service unavailable response has a 2xx status code
func (o *PutGamificationStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status service unavailable response has a 3xx status code
func (o *PutGamificationStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status service unavailable response has a 4xx status code
func (o *PutGamificationStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put gamification status service unavailable response has a 5xx status code
func (o *PutGamificationStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put gamification status service unavailable response a status code equal to that given
func (o *PutGamificationStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutGamificationStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGamificationStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGamificationStatusServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationStatusGatewayTimeout creates a PutGamificationStatusGatewayTimeout with default headers values
func NewPutGamificationStatusGatewayTimeout() *PutGamificationStatusGatewayTimeout {
	return &PutGamificationStatusGatewayTimeout{}
}

/*
PutGamificationStatusGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutGamificationStatusGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put gamification status gateway timeout response has a 2xx status code
func (o *PutGamificationStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put gamification status gateway timeout response has a 3xx status code
func (o *PutGamificationStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put gamification status gateway timeout response has a 4xx status code
func (o *PutGamificationStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put gamification status gateway timeout response has a 5xx status code
func (o *PutGamificationStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put gamification status gateway timeout response a status code equal to that given
func (o *PutGamificationStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutGamificationStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGamificationStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/status][%d] putGamificationStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGamificationStatusGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
