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

// GetGamificationStatusReader is a Reader for the GetGamificationStatus structure.
type GetGamificationStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationStatusRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationStatusRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationStatusUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationStatusGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationStatusOK creates a GetGamificationStatusOK with default headers values
func NewGetGamificationStatusOK() *GetGamificationStatusOK {
	return &GetGamificationStatusOK{}
}

/*
GetGamificationStatusOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGamificationStatusOK struct {
	Payload *models.GamificationStatus
}

// IsSuccess returns true when this get gamification status o k response has a 2xx status code
func (o *GetGamificationStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gamification status o k response has a 3xx status code
func (o *GetGamificationStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status o k response has a 4xx status code
func (o *GetGamificationStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification status o k response has a 5xx status code
func (o *GetGamificationStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status o k response a status code equal to that given
func (o *GetGamificationStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGamificationStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusOK  %+v", 200, o.Payload)
}

func (o *GetGamificationStatusOK) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusOK  %+v", 200, o.Payload)
}

func (o *GetGamificationStatusOK) GetPayload() *models.GamificationStatus {
	return o.Payload
}

func (o *GetGamificationStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GamificationStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusBadRequest creates a GetGamificationStatusBadRequest with default headers values
func NewGetGamificationStatusBadRequest() *GetGamificationStatusBadRequest {
	return &GetGamificationStatusBadRequest{}
}

/*
GetGamificationStatusBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationStatusBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status bad request response has a 2xx status code
func (o *GetGamificationStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status bad request response has a 3xx status code
func (o *GetGamificationStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status bad request response has a 4xx status code
func (o *GetGamificationStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status bad request response has a 5xx status code
func (o *GetGamificationStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status bad request response a status code equal to that given
func (o *GetGamificationStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetGamificationStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationStatusBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationStatusBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusUnauthorized creates a GetGamificationStatusUnauthorized with default headers values
func NewGetGamificationStatusUnauthorized() *GetGamificationStatusUnauthorized {
	return &GetGamificationStatusUnauthorized{}
}

/*
GetGamificationStatusUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationStatusUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status unauthorized response has a 2xx status code
func (o *GetGamificationStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status unauthorized response has a 3xx status code
func (o *GetGamificationStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status unauthorized response has a 4xx status code
func (o *GetGamificationStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status unauthorized response has a 5xx status code
func (o *GetGamificationStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status unauthorized response a status code equal to that given
func (o *GetGamificationStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGamificationStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationStatusUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusForbidden creates a GetGamificationStatusForbidden with default headers values
func NewGetGamificationStatusForbidden() *GetGamificationStatusForbidden {
	return &GetGamificationStatusForbidden{}
}

/*
GetGamificationStatusForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationStatusForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status forbidden response has a 2xx status code
func (o *GetGamificationStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status forbidden response has a 3xx status code
func (o *GetGamificationStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status forbidden response has a 4xx status code
func (o *GetGamificationStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status forbidden response has a 5xx status code
func (o *GetGamificationStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status forbidden response a status code equal to that given
func (o *GetGamificationStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGamificationStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationStatusForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationStatusForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusNotFound creates a GetGamificationStatusNotFound with default headers values
func NewGetGamificationStatusNotFound() *GetGamificationStatusNotFound {
	return &GetGamificationStatusNotFound{}
}

/*
GetGamificationStatusNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetGamificationStatusNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status not found response has a 2xx status code
func (o *GetGamificationStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status not found response has a 3xx status code
func (o *GetGamificationStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status not found response has a 4xx status code
func (o *GetGamificationStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status not found response has a 5xx status code
func (o *GetGamificationStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status not found response a status code equal to that given
func (o *GetGamificationStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGamificationStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationStatusNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationStatusNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusRequestTimeout creates a GetGamificationStatusRequestTimeout with default headers values
func NewGetGamificationStatusRequestTimeout() *GetGamificationStatusRequestTimeout {
	return &GetGamificationStatusRequestTimeout{}
}

/*
GetGamificationStatusRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationStatusRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status request timeout response has a 2xx status code
func (o *GetGamificationStatusRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status request timeout response has a 3xx status code
func (o *GetGamificationStatusRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status request timeout response has a 4xx status code
func (o *GetGamificationStatusRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status request timeout response has a 5xx status code
func (o *GetGamificationStatusRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status request timeout response a status code equal to that given
func (o *GetGamificationStatusRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetGamificationStatusRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationStatusRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationStatusRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusRequestEntityTooLarge creates a GetGamificationStatusRequestEntityTooLarge with default headers values
func NewGetGamificationStatusRequestEntityTooLarge() *GetGamificationStatusRequestEntityTooLarge {
	return &GetGamificationStatusRequestEntityTooLarge{}
}

/*
GetGamificationStatusRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGamificationStatusRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status request entity too large response has a 2xx status code
func (o *GetGamificationStatusRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status request entity too large response has a 3xx status code
func (o *GetGamificationStatusRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status request entity too large response has a 4xx status code
func (o *GetGamificationStatusRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status request entity too large response has a 5xx status code
func (o *GetGamificationStatusRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status request entity too large response a status code equal to that given
func (o *GetGamificationStatusRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetGamificationStatusRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationStatusRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationStatusRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusUnsupportedMediaType creates a GetGamificationStatusUnsupportedMediaType with default headers values
func NewGetGamificationStatusUnsupportedMediaType() *GetGamificationStatusUnsupportedMediaType {
	return &GetGamificationStatusUnsupportedMediaType{}
}

/*
GetGamificationStatusUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationStatusUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status unsupported media type response has a 2xx status code
func (o *GetGamificationStatusUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status unsupported media type response has a 3xx status code
func (o *GetGamificationStatusUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status unsupported media type response has a 4xx status code
func (o *GetGamificationStatusUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status unsupported media type response has a 5xx status code
func (o *GetGamificationStatusUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status unsupported media type response a status code equal to that given
func (o *GetGamificationStatusUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetGamificationStatusUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationStatusUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationStatusUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusTooManyRequests creates a GetGamificationStatusTooManyRequests with default headers values
func NewGetGamificationStatusTooManyRequests() *GetGamificationStatusTooManyRequests {
	return &GetGamificationStatusTooManyRequests{}
}

/*
GetGamificationStatusTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationStatusTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status too many requests response has a 2xx status code
func (o *GetGamificationStatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status too many requests response has a 3xx status code
func (o *GetGamificationStatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status too many requests response has a 4xx status code
func (o *GetGamificationStatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification status too many requests response has a 5xx status code
func (o *GetGamificationStatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification status too many requests response a status code equal to that given
func (o *GetGamificationStatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetGamificationStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationStatusTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationStatusTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusInternalServerError creates a GetGamificationStatusInternalServerError with default headers values
func NewGetGamificationStatusInternalServerError() *GetGamificationStatusInternalServerError {
	return &GetGamificationStatusInternalServerError{}
}

/*
GetGamificationStatusInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationStatusInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status internal server error response has a 2xx status code
func (o *GetGamificationStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status internal server error response has a 3xx status code
func (o *GetGamificationStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status internal server error response has a 4xx status code
func (o *GetGamificationStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification status internal server error response has a 5xx status code
func (o *GetGamificationStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification status internal server error response a status code equal to that given
func (o *GetGamificationStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGamificationStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationStatusInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusServiceUnavailable creates a GetGamificationStatusServiceUnavailable with default headers values
func NewGetGamificationStatusServiceUnavailable() *GetGamificationStatusServiceUnavailable {
	return &GetGamificationStatusServiceUnavailable{}
}

/*
GetGamificationStatusServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationStatusServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status service unavailable response has a 2xx status code
func (o *GetGamificationStatusServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status service unavailable response has a 3xx status code
func (o *GetGamificationStatusServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status service unavailable response has a 4xx status code
func (o *GetGamificationStatusServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification status service unavailable response has a 5xx status code
func (o *GetGamificationStatusServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification status service unavailable response a status code equal to that given
func (o *GetGamificationStatusServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetGamificationStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationStatusServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationStatusServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationStatusGatewayTimeout creates a GetGamificationStatusGatewayTimeout with default headers values
func NewGetGamificationStatusGatewayTimeout() *GetGamificationStatusGatewayTimeout {
	return &GetGamificationStatusGatewayTimeout{}
}

/*
GetGamificationStatusGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetGamificationStatusGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification status gateway timeout response has a 2xx status code
func (o *GetGamificationStatusGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification status gateway timeout response has a 3xx status code
func (o *GetGamificationStatusGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification status gateway timeout response has a 4xx status code
func (o *GetGamificationStatusGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification status gateway timeout response has a 5xx status code
func (o *GetGamificationStatusGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification status gateway timeout response a status code equal to that given
func (o *GetGamificationStatusGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetGamificationStatusGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationStatusGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/status][%d] getGamificationStatusGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationStatusGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationStatusGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
