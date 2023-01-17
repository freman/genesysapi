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

// GetRoutingMessageRecipientReader is a Reader for the GetRoutingMessageRecipient structure.
type GetRoutingMessageRecipientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingMessageRecipientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingMessageRecipientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingMessageRecipientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingMessageRecipientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingMessageRecipientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingMessageRecipientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingMessageRecipientRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingMessageRecipientRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingMessageRecipientUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingMessageRecipientTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingMessageRecipientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingMessageRecipientServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingMessageRecipientGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingMessageRecipientOK creates a GetRoutingMessageRecipientOK with default headers values
func NewGetRoutingMessageRecipientOK() *GetRoutingMessageRecipientOK {
	return &GetRoutingMessageRecipientOK{}
}

/*
GetRoutingMessageRecipientOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingMessageRecipientOK struct {
	Payload *models.Recipient
}

// IsSuccess returns true when this get routing message recipient o k response has a 2xx status code
func (o *GetRoutingMessageRecipientOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing message recipient o k response has a 3xx status code
func (o *GetRoutingMessageRecipientOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient o k response has a 4xx status code
func (o *GetRoutingMessageRecipientOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipient o k response has a 5xx status code
func (o *GetRoutingMessageRecipientOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient o k response a status code equal to that given
func (o *GetRoutingMessageRecipientOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingMessageRecipientOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientOK  %+v", 200, o.Payload)
}

func (o *GetRoutingMessageRecipientOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientOK  %+v", 200, o.Payload)
}

func (o *GetRoutingMessageRecipientOK) GetPayload() *models.Recipient {
	return o.Payload
}

func (o *GetRoutingMessageRecipientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recipient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientBadRequest creates a GetRoutingMessageRecipientBadRequest with default headers values
func NewGetRoutingMessageRecipientBadRequest() *GetRoutingMessageRecipientBadRequest {
	return &GetRoutingMessageRecipientBadRequest{}
}

/*
GetRoutingMessageRecipientBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingMessageRecipientBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient bad request response has a 2xx status code
func (o *GetRoutingMessageRecipientBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient bad request response has a 3xx status code
func (o *GetRoutingMessageRecipientBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient bad request response has a 4xx status code
func (o *GetRoutingMessageRecipientBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient bad request response has a 5xx status code
func (o *GetRoutingMessageRecipientBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient bad request response a status code equal to that given
func (o *GetRoutingMessageRecipientBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingMessageRecipientBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingMessageRecipientBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingMessageRecipientBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientUnauthorized creates a GetRoutingMessageRecipientUnauthorized with default headers values
func NewGetRoutingMessageRecipientUnauthorized() *GetRoutingMessageRecipientUnauthorized {
	return &GetRoutingMessageRecipientUnauthorized{}
}

/*
GetRoutingMessageRecipientUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingMessageRecipientUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient unauthorized response has a 2xx status code
func (o *GetRoutingMessageRecipientUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient unauthorized response has a 3xx status code
func (o *GetRoutingMessageRecipientUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient unauthorized response has a 4xx status code
func (o *GetRoutingMessageRecipientUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient unauthorized response has a 5xx status code
func (o *GetRoutingMessageRecipientUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient unauthorized response a status code equal to that given
func (o *GetRoutingMessageRecipientUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingMessageRecipientUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingMessageRecipientUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingMessageRecipientUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientForbidden creates a GetRoutingMessageRecipientForbidden with default headers values
func NewGetRoutingMessageRecipientForbidden() *GetRoutingMessageRecipientForbidden {
	return &GetRoutingMessageRecipientForbidden{}
}

/*
GetRoutingMessageRecipientForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingMessageRecipientForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient forbidden response has a 2xx status code
func (o *GetRoutingMessageRecipientForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient forbidden response has a 3xx status code
func (o *GetRoutingMessageRecipientForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient forbidden response has a 4xx status code
func (o *GetRoutingMessageRecipientForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient forbidden response has a 5xx status code
func (o *GetRoutingMessageRecipientForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient forbidden response a status code equal to that given
func (o *GetRoutingMessageRecipientForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingMessageRecipientForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingMessageRecipientForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingMessageRecipientForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientNotFound creates a GetRoutingMessageRecipientNotFound with default headers values
func NewGetRoutingMessageRecipientNotFound() *GetRoutingMessageRecipientNotFound {
	return &GetRoutingMessageRecipientNotFound{}
}

/*
GetRoutingMessageRecipientNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingMessageRecipientNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient not found response has a 2xx status code
func (o *GetRoutingMessageRecipientNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient not found response has a 3xx status code
func (o *GetRoutingMessageRecipientNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient not found response has a 4xx status code
func (o *GetRoutingMessageRecipientNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient not found response has a 5xx status code
func (o *GetRoutingMessageRecipientNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient not found response a status code equal to that given
func (o *GetRoutingMessageRecipientNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingMessageRecipientNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingMessageRecipientNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingMessageRecipientNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientRequestTimeout creates a GetRoutingMessageRecipientRequestTimeout with default headers values
func NewGetRoutingMessageRecipientRequestTimeout() *GetRoutingMessageRecipientRequestTimeout {
	return &GetRoutingMessageRecipientRequestTimeout{}
}

/*
GetRoutingMessageRecipientRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingMessageRecipientRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient request timeout response has a 2xx status code
func (o *GetRoutingMessageRecipientRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient request timeout response has a 3xx status code
func (o *GetRoutingMessageRecipientRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient request timeout response has a 4xx status code
func (o *GetRoutingMessageRecipientRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient request timeout response has a 5xx status code
func (o *GetRoutingMessageRecipientRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient request timeout response a status code equal to that given
func (o *GetRoutingMessageRecipientRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingMessageRecipientRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingMessageRecipientRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingMessageRecipientRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientRequestEntityTooLarge creates a GetRoutingMessageRecipientRequestEntityTooLarge with default headers values
func NewGetRoutingMessageRecipientRequestEntityTooLarge() *GetRoutingMessageRecipientRequestEntityTooLarge {
	return &GetRoutingMessageRecipientRequestEntityTooLarge{}
}

/*
GetRoutingMessageRecipientRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingMessageRecipientRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient request entity too large response has a 2xx status code
func (o *GetRoutingMessageRecipientRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient request entity too large response has a 3xx status code
func (o *GetRoutingMessageRecipientRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient request entity too large response has a 4xx status code
func (o *GetRoutingMessageRecipientRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient request entity too large response has a 5xx status code
func (o *GetRoutingMessageRecipientRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient request entity too large response a status code equal to that given
func (o *GetRoutingMessageRecipientRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingMessageRecipientRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingMessageRecipientRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingMessageRecipientRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientUnsupportedMediaType creates a GetRoutingMessageRecipientUnsupportedMediaType with default headers values
func NewGetRoutingMessageRecipientUnsupportedMediaType() *GetRoutingMessageRecipientUnsupportedMediaType {
	return &GetRoutingMessageRecipientUnsupportedMediaType{}
}

/*
GetRoutingMessageRecipientUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingMessageRecipientUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient unsupported media type response has a 2xx status code
func (o *GetRoutingMessageRecipientUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient unsupported media type response has a 3xx status code
func (o *GetRoutingMessageRecipientUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient unsupported media type response has a 4xx status code
func (o *GetRoutingMessageRecipientUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient unsupported media type response has a 5xx status code
func (o *GetRoutingMessageRecipientUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient unsupported media type response a status code equal to that given
func (o *GetRoutingMessageRecipientUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingMessageRecipientUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingMessageRecipientUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingMessageRecipientUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientTooManyRequests creates a GetRoutingMessageRecipientTooManyRequests with default headers values
func NewGetRoutingMessageRecipientTooManyRequests() *GetRoutingMessageRecipientTooManyRequests {
	return &GetRoutingMessageRecipientTooManyRequests{}
}

/*
GetRoutingMessageRecipientTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingMessageRecipientTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient too many requests response has a 2xx status code
func (o *GetRoutingMessageRecipientTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient too many requests response has a 3xx status code
func (o *GetRoutingMessageRecipientTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient too many requests response has a 4xx status code
func (o *GetRoutingMessageRecipientTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing message recipient too many requests response has a 5xx status code
func (o *GetRoutingMessageRecipientTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing message recipient too many requests response a status code equal to that given
func (o *GetRoutingMessageRecipientTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingMessageRecipientTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingMessageRecipientTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingMessageRecipientTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientInternalServerError creates a GetRoutingMessageRecipientInternalServerError with default headers values
func NewGetRoutingMessageRecipientInternalServerError() *GetRoutingMessageRecipientInternalServerError {
	return &GetRoutingMessageRecipientInternalServerError{}
}

/*
GetRoutingMessageRecipientInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingMessageRecipientInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient internal server error response has a 2xx status code
func (o *GetRoutingMessageRecipientInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient internal server error response has a 3xx status code
func (o *GetRoutingMessageRecipientInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient internal server error response has a 4xx status code
func (o *GetRoutingMessageRecipientInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipient internal server error response has a 5xx status code
func (o *GetRoutingMessageRecipientInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing message recipient internal server error response a status code equal to that given
func (o *GetRoutingMessageRecipientInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingMessageRecipientInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingMessageRecipientInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingMessageRecipientInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientServiceUnavailable creates a GetRoutingMessageRecipientServiceUnavailable with default headers values
func NewGetRoutingMessageRecipientServiceUnavailable() *GetRoutingMessageRecipientServiceUnavailable {
	return &GetRoutingMessageRecipientServiceUnavailable{}
}

/*
GetRoutingMessageRecipientServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingMessageRecipientServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient service unavailable response has a 2xx status code
func (o *GetRoutingMessageRecipientServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient service unavailable response has a 3xx status code
func (o *GetRoutingMessageRecipientServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient service unavailable response has a 4xx status code
func (o *GetRoutingMessageRecipientServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipient service unavailable response has a 5xx status code
func (o *GetRoutingMessageRecipientServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing message recipient service unavailable response a status code equal to that given
func (o *GetRoutingMessageRecipientServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingMessageRecipientServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingMessageRecipientServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingMessageRecipientServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingMessageRecipientGatewayTimeout creates a GetRoutingMessageRecipientGatewayTimeout with default headers values
func NewGetRoutingMessageRecipientGatewayTimeout() *GetRoutingMessageRecipientGatewayTimeout {
	return &GetRoutingMessageRecipientGatewayTimeout{}
}

/*
GetRoutingMessageRecipientGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingMessageRecipientGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing message recipient gateway timeout response has a 2xx status code
func (o *GetRoutingMessageRecipientGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing message recipient gateway timeout response has a 3xx status code
func (o *GetRoutingMessageRecipientGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing message recipient gateway timeout response has a 4xx status code
func (o *GetRoutingMessageRecipientGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing message recipient gateway timeout response has a 5xx status code
func (o *GetRoutingMessageRecipientGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing message recipient gateway timeout response a status code equal to that given
func (o *GetRoutingMessageRecipientGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingMessageRecipientGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingMessageRecipientGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/message/recipients/{recipientId}][%d] getRoutingMessageRecipientGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingMessageRecipientGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingMessageRecipientGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
