// Code generated by go-swagger; DO NOT EDIT.

package coaching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetCoachingNotificationsReader is a Reader for the GetCoachingNotifications structure.
type GetCoachingNotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoachingNotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoachingNotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCoachingNotificationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCoachingNotificationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCoachingNotificationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCoachingNotificationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetCoachingNotificationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetCoachingNotificationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetCoachingNotificationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCoachingNotificationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCoachingNotificationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetCoachingNotificationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetCoachingNotificationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoachingNotificationsOK creates a GetCoachingNotificationsOK with default headers values
func NewGetCoachingNotificationsOK() *GetCoachingNotificationsOK {
	return &GetCoachingNotificationsOK{}
}

/*
GetCoachingNotificationsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCoachingNotificationsOK struct {
	Payload *models.CoachingNotificationList
}

// IsSuccess returns true when this get coaching notifications o k response has a 2xx status code
func (o *GetCoachingNotificationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get coaching notifications o k response has a 3xx status code
func (o *GetCoachingNotificationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications o k response has a 4xx status code
func (o *GetCoachingNotificationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching notifications o k response has a 5xx status code
func (o *GetCoachingNotificationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications o k response a status code equal to that given
func (o *GetCoachingNotificationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetCoachingNotificationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsOK  %+v", 200, o.Payload)
}

func (o *GetCoachingNotificationsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsOK  %+v", 200, o.Payload)
}

func (o *GetCoachingNotificationsOK) GetPayload() *models.CoachingNotificationList {
	return o.Payload
}

func (o *GetCoachingNotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingNotificationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsBadRequest creates a GetCoachingNotificationsBadRequest with default headers values
func NewGetCoachingNotificationsBadRequest() *GetCoachingNotificationsBadRequest {
	return &GetCoachingNotificationsBadRequest{}
}

/*
GetCoachingNotificationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetCoachingNotificationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications bad request response has a 2xx status code
func (o *GetCoachingNotificationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications bad request response has a 3xx status code
func (o *GetCoachingNotificationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications bad request response has a 4xx status code
func (o *GetCoachingNotificationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications bad request response has a 5xx status code
func (o *GetCoachingNotificationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications bad request response a status code equal to that given
func (o *GetCoachingNotificationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetCoachingNotificationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingNotificationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCoachingNotificationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsUnauthorized creates a GetCoachingNotificationsUnauthorized with default headers values
func NewGetCoachingNotificationsUnauthorized() *GetCoachingNotificationsUnauthorized {
	return &GetCoachingNotificationsUnauthorized{}
}

/*
GetCoachingNotificationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetCoachingNotificationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications unauthorized response has a 2xx status code
func (o *GetCoachingNotificationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications unauthorized response has a 3xx status code
func (o *GetCoachingNotificationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications unauthorized response has a 4xx status code
func (o *GetCoachingNotificationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications unauthorized response has a 5xx status code
func (o *GetCoachingNotificationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications unauthorized response a status code equal to that given
func (o *GetCoachingNotificationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetCoachingNotificationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingNotificationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCoachingNotificationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsForbidden creates a GetCoachingNotificationsForbidden with default headers values
func NewGetCoachingNotificationsForbidden() *GetCoachingNotificationsForbidden {
	return &GetCoachingNotificationsForbidden{}
}

/*
GetCoachingNotificationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetCoachingNotificationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications forbidden response has a 2xx status code
func (o *GetCoachingNotificationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications forbidden response has a 3xx status code
func (o *GetCoachingNotificationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications forbidden response has a 4xx status code
func (o *GetCoachingNotificationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications forbidden response has a 5xx status code
func (o *GetCoachingNotificationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications forbidden response a status code equal to that given
func (o *GetCoachingNotificationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetCoachingNotificationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingNotificationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsForbidden  %+v", 403, o.Payload)
}

func (o *GetCoachingNotificationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsNotFound creates a GetCoachingNotificationsNotFound with default headers values
func NewGetCoachingNotificationsNotFound() *GetCoachingNotificationsNotFound {
	return &GetCoachingNotificationsNotFound{}
}

/*
GetCoachingNotificationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetCoachingNotificationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications not found response has a 2xx status code
func (o *GetCoachingNotificationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications not found response has a 3xx status code
func (o *GetCoachingNotificationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications not found response has a 4xx status code
func (o *GetCoachingNotificationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications not found response has a 5xx status code
func (o *GetCoachingNotificationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications not found response a status code equal to that given
func (o *GetCoachingNotificationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetCoachingNotificationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingNotificationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsNotFound  %+v", 404, o.Payload)
}

func (o *GetCoachingNotificationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsRequestTimeout creates a GetCoachingNotificationsRequestTimeout with default headers values
func NewGetCoachingNotificationsRequestTimeout() *GetCoachingNotificationsRequestTimeout {
	return &GetCoachingNotificationsRequestTimeout{}
}

/*
GetCoachingNotificationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetCoachingNotificationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications request timeout response has a 2xx status code
func (o *GetCoachingNotificationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications request timeout response has a 3xx status code
func (o *GetCoachingNotificationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications request timeout response has a 4xx status code
func (o *GetCoachingNotificationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications request timeout response has a 5xx status code
func (o *GetCoachingNotificationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications request timeout response a status code equal to that given
func (o *GetCoachingNotificationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetCoachingNotificationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingNotificationsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetCoachingNotificationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsRequestEntityTooLarge creates a GetCoachingNotificationsRequestEntityTooLarge with default headers values
func NewGetCoachingNotificationsRequestEntityTooLarge() *GetCoachingNotificationsRequestEntityTooLarge {
	return &GetCoachingNotificationsRequestEntityTooLarge{}
}

/*
GetCoachingNotificationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetCoachingNotificationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications request entity too large response has a 2xx status code
func (o *GetCoachingNotificationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications request entity too large response has a 3xx status code
func (o *GetCoachingNotificationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications request entity too large response has a 4xx status code
func (o *GetCoachingNotificationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications request entity too large response has a 5xx status code
func (o *GetCoachingNotificationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications request entity too large response a status code equal to that given
func (o *GetCoachingNotificationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetCoachingNotificationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingNotificationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetCoachingNotificationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsUnsupportedMediaType creates a GetCoachingNotificationsUnsupportedMediaType with default headers values
func NewGetCoachingNotificationsUnsupportedMediaType() *GetCoachingNotificationsUnsupportedMediaType {
	return &GetCoachingNotificationsUnsupportedMediaType{}
}

/*
GetCoachingNotificationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetCoachingNotificationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications unsupported media type response has a 2xx status code
func (o *GetCoachingNotificationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications unsupported media type response has a 3xx status code
func (o *GetCoachingNotificationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications unsupported media type response has a 4xx status code
func (o *GetCoachingNotificationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications unsupported media type response has a 5xx status code
func (o *GetCoachingNotificationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications unsupported media type response a status code equal to that given
func (o *GetCoachingNotificationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetCoachingNotificationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingNotificationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetCoachingNotificationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsTooManyRequests creates a GetCoachingNotificationsTooManyRequests with default headers values
func NewGetCoachingNotificationsTooManyRequests() *GetCoachingNotificationsTooManyRequests {
	return &GetCoachingNotificationsTooManyRequests{}
}

/*
GetCoachingNotificationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetCoachingNotificationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications too many requests response has a 2xx status code
func (o *GetCoachingNotificationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications too many requests response has a 3xx status code
func (o *GetCoachingNotificationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications too many requests response has a 4xx status code
func (o *GetCoachingNotificationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coaching notifications too many requests response has a 5xx status code
func (o *GetCoachingNotificationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get coaching notifications too many requests response a status code equal to that given
func (o *GetCoachingNotificationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetCoachingNotificationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingNotificationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCoachingNotificationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsInternalServerError creates a GetCoachingNotificationsInternalServerError with default headers values
func NewGetCoachingNotificationsInternalServerError() *GetCoachingNotificationsInternalServerError {
	return &GetCoachingNotificationsInternalServerError{}
}

/*
GetCoachingNotificationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetCoachingNotificationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications internal server error response has a 2xx status code
func (o *GetCoachingNotificationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications internal server error response has a 3xx status code
func (o *GetCoachingNotificationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications internal server error response has a 4xx status code
func (o *GetCoachingNotificationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching notifications internal server error response has a 5xx status code
func (o *GetCoachingNotificationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching notifications internal server error response a status code equal to that given
func (o *GetCoachingNotificationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetCoachingNotificationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingNotificationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCoachingNotificationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsServiceUnavailable creates a GetCoachingNotificationsServiceUnavailable with default headers values
func NewGetCoachingNotificationsServiceUnavailable() *GetCoachingNotificationsServiceUnavailable {
	return &GetCoachingNotificationsServiceUnavailable{}
}

/*
GetCoachingNotificationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetCoachingNotificationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications service unavailable response has a 2xx status code
func (o *GetCoachingNotificationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications service unavailable response has a 3xx status code
func (o *GetCoachingNotificationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications service unavailable response has a 4xx status code
func (o *GetCoachingNotificationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching notifications service unavailable response has a 5xx status code
func (o *GetCoachingNotificationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching notifications service unavailable response a status code equal to that given
func (o *GetCoachingNotificationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetCoachingNotificationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingNotificationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCoachingNotificationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoachingNotificationsGatewayTimeout creates a GetCoachingNotificationsGatewayTimeout with default headers values
func NewGetCoachingNotificationsGatewayTimeout() *GetCoachingNotificationsGatewayTimeout {
	return &GetCoachingNotificationsGatewayTimeout{}
}

/*
GetCoachingNotificationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetCoachingNotificationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get coaching notifications gateway timeout response has a 2xx status code
func (o *GetCoachingNotificationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coaching notifications gateway timeout response has a 3xx status code
func (o *GetCoachingNotificationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coaching notifications gateway timeout response has a 4xx status code
func (o *GetCoachingNotificationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coaching notifications gateway timeout response has a 5xx status code
func (o *GetCoachingNotificationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get coaching notifications gateway timeout response a status code equal to that given
func (o *GetCoachingNotificationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetCoachingNotificationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingNotificationsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/coaching/notifications][%d] getCoachingNotificationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCoachingNotificationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCoachingNotificationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
