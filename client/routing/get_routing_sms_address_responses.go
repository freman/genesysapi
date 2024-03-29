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

// GetRoutingSmsAddressReader is a Reader for the GetRoutingSmsAddress structure.
type GetRoutingSmsAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSmsAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSmsAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSmsAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSmsAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSmsAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSmsAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingSmsAddressRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSmsAddressRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSmsAddressUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSmsAddressTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSmsAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSmsAddressServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSmsAddressGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSmsAddressOK creates a GetRoutingSmsAddressOK with default headers values
func NewGetRoutingSmsAddressOK() *GetRoutingSmsAddressOK {
	return &GetRoutingSmsAddressOK{}
}

/*
GetRoutingSmsAddressOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingSmsAddressOK struct {
	Payload *models.SmsAddress
}

// IsSuccess returns true when this get routing sms address o k response has a 2xx status code
func (o *GetRoutingSmsAddressOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing sms address o k response has a 3xx status code
func (o *GetRoutingSmsAddressOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address o k response has a 4xx status code
func (o *GetRoutingSmsAddressOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing sms address o k response has a 5xx status code
func (o *GetRoutingSmsAddressOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address o k response a status code equal to that given
func (o *GetRoutingSmsAddressOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingSmsAddressOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSmsAddressOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSmsAddressOK) GetPayload() *models.SmsAddress {
	return o.Payload
}

func (o *GetRoutingSmsAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SmsAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressBadRequest creates a GetRoutingSmsAddressBadRequest with default headers values
func NewGetRoutingSmsAddressBadRequest() *GetRoutingSmsAddressBadRequest {
	return &GetRoutingSmsAddressBadRequest{}
}

/*
GetRoutingSmsAddressBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSmsAddressBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address bad request response has a 2xx status code
func (o *GetRoutingSmsAddressBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address bad request response has a 3xx status code
func (o *GetRoutingSmsAddressBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address bad request response has a 4xx status code
func (o *GetRoutingSmsAddressBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address bad request response has a 5xx status code
func (o *GetRoutingSmsAddressBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address bad request response a status code equal to that given
func (o *GetRoutingSmsAddressBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingSmsAddressBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSmsAddressBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSmsAddressBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressUnauthorized creates a GetRoutingSmsAddressUnauthorized with default headers values
func NewGetRoutingSmsAddressUnauthorized() *GetRoutingSmsAddressUnauthorized {
	return &GetRoutingSmsAddressUnauthorized{}
}

/*
GetRoutingSmsAddressUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSmsAddressUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address unauthorized response has a 2xx status code
func (o *GetRoutingSmsAddressUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address unauthorized response has a 3xx status code
func (o *GetRoutingSmsAddressUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address unauthorized response has a 4xx status code
func (o *GetRoutingSmsAddressUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address unauthorized response has a 5xx status code
func (o *GetRoutingSmsAddressUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address unauthorized response a status code equal to that given
func (o *GetRoutingSmsAddressUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingSmsAddressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSmsAddressUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSmsAddressUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressForbidden creates a GetRoutingSmsAddressForbidden with default headers values
func NewGetRoutingSmsAddressForbidden() *GetRoutingSmsAddressForbidden {
	return &GetRoutingSmsAddressForbidden{}
}

/*
GetRoutingSmsAddressForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSmsAddressForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address forbidden response has a 2xx status code
func (o *GetRoutingSmsAddressForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address forbidden response has a 3xx status code
func (o *GetRoutingSmsAddressForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address forbidden response has a 4xx status code
func (o *GetRoutingSmsAddressForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address forbidden response has a 5xx status code
func (o *GetRoutingSmsAddressForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address forbidden response a status code equal to that given
func (o *GetRoutingSmsAddressForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingSmsAddressForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSmsAddressForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSmsAddressForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressNotFound creates a GetRoutingSmsAddressNotFound with default headers values
func NewGetRoutingSmsAddressNotFound() *GetRoutingSmsAddressNotFound {
	return &GetRoutingSmsAddressNotFound{}
}

/*
GetRoutingSmsAddressNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingSmsAddressNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address not found response has a 2xx status code
func (o *GetRoutingSmsAddressNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address not found response has a 3xx status code
func (o *GetRoutingSmsAddressNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address not found response has a 4xx status code
func (o *GetRoutingSmsAddressNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address not found response has a 5xx status code
func (o *GetRoutingSmsAddressNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address not found response a status code equal to that given
func (o *GetRoutingSmsAddressNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingSmsAddressNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSmsAddressNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSmsAddressNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressRequestTimeout creates a GetRoutingSmsAddressRequestTimeout with default headers values
func NewGetRoutingSmsAddressRequestTimeout() *GetRoutingSmsAddressRequestTimeout {
	return &GetRoutingSmsAddressRequestTimeout{}
}

/*
GetRoutingSmsAddressRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingSmsAddressRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address request timeout response has a 2xx status code
func (o *GetRoutingSmsAddressRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address request timeout response has a 3xx status code
func (o *GetRoutingSmsAddressRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address request timeout response has a 4xx status code
func (o *GetRoutingSmsAddressRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address request timeout response has a 5xx status code
func (o *GetRoutingSmsAddressRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address request timeout response a status code equal to that given
func (o *GetRoutingSmsAddressRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingSmsAddressRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSmsAddressRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSmsAddressRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressRequestEntityTooLarge creates a GetRoutingSmsAddressRequestEntityTooLarge with default headers values
func NewGetRoutingSmsAddressRequestEntityTooLarge() *GetRoutingSmsAddressRequestEntityTooLarge {
	return &GetRoutingSmsAddressRequestEntityTooLarge{}
}

/*
GetRoutingSmsAddressRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingSmsAddressRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address request entity too large response has a 2xx status code
func (o *GetRoutingSmsAddressRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address request entity too large response has a 3xx status code
func (o *GetRoutingSmsAddressRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address request entity too large response has a 4xx status code
func (o *GetRoutingSmsAddressRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address request entity too large response has a 5xx status code
func (o *GetRoutingSmsAddressRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address request entity too large response a status code equal to that given
func (o *GetRoutingSmsAddressRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingSmsAddressRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSmsAddressRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSmsAddressRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressUnsupportedMediaType creates a GetRoutingSmsAddressUnsupportedMediaType with default headers values
func NewGetRoutingSmsAddressUnsupportedMediaType() *GetRoutingSmsAddressUnsupportedMediaType {
	return &GetRoutingSmsAddressUnsupportedMediaType{}
}

/*
GetRoutingSmsAddressUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSmsAddressUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address unsupported media type response has a 2xx status code
func (o *GetRoutingSmsAddressUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address unsupported media type response has a 3xx status code
func (o *GetRoutingSmsAddressUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address unsupported media type response has a 4xx status code
func (o *GetRoutingSmsAddressUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address unsupported media type response has a 5xx status code
func (o *GetRoutingSmsAddressUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address unsupported media type response a status code equal to that given
func (o *GetRoutingSmsAddressUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingSmsAddressUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSmsAddressUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSmsAddressUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressTooManyRequests creates a GetRoutingSmsAddressTooManyRequests with default headers values
func NewGetRoutingSmsAddressTooManyRequests() *GetRoutingSmsAddressTooManyRequests {
	return &GetRoutingSmsAddressTooManyRequests{}
}

/*
GetRoutingSmsAddressTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingSmsAddressTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address too many requests response has a 2xx status code
func (o *GetRoutingSmsAddressTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address too many requests response has a 3xx status code
func (o *GetRoutingSmsAddressTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address too many requests response has a 4xx status code
func (o *GetRoutingSmsAddressTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing sms address too many requests response has a 5xx status code
func (o *GetRoutingSmsAddressTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing sms address too many requests response a status code equal to that given
func (o *GetRoutingSmsAddressTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingSmsAddressTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSmsAddressTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSmsAddressTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressInternalServerError creates a GetRoutingSmsAddressInternalServerError with default headers values
func NewGetRoutingSmsAddressInternalServerError() *GetRoutingSmsAddressInternalServerError {
	return &GetRoutingSmsAddressInternalServerError{}
}

/*
GetRoutingSmsAddressInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSmsAddressInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address internal server error response has a 2xx status code
func (o *GetRoutingSmsAddressInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address internal server error response has a 3xx status code
func (o *GetRoutingSmsAddressInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address internal server error response has a 4xx status code
func (o *GetRoutingSmsAddressInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing sms address internal server error response has a 5xx status code
func (o *GetRoutingSmsAddressInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing sms address internal server error response a status code equal to that given
func (o *GetRoutingSmsAddressInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingSmsAddressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSmsAddressInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSmsAddressInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressServiceUnavailable creates a GetRoutingSmsAddressServiceUnavailable with default headers values
func NewGetRoutingSmsAddressServiceUnavailable() *GetRoutingSmsAddressServiceUnavailable {
	return &GetRoutingSmsAddressServiceUnavailable{}
}

/*
GetRoutingSmsAddressServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSmsAddressServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address service unavailable response has a 2xx status code
func (o *GetRoutingSmsAddressServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address service unavailable response has a 3xx status code
func (o *GetRoutingSmsAddressServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address service unavailable response has a 4xx status code
func (o *GetRoutingSmsAddressServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing sms address service unavailable response has a 5xx status code
func (o *GetRoutingSmsAddressServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing sms address service unavailable response a status code equal to that given
func (o *GetRoutingSmsAddressServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingSmsAddressServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSmsAddressServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSmsAddressServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSmsAddressGatewayTimeout creates a GetRoutingSmsAddressGatewayTimeout with default headers values
func NewGetRoutingSmsAddressGatewayTimeout() *GetRoutingSmsAddressGatewayTimeout {
	return &GetRoutingSmsAddressGatewayTimeout{}
}

/*
GetRoutingSmsAddressGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingSmsAddressGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing sms address gateway timeout response has a 2xx status code
func (o *GetRoutingSmsAddressGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing sms address gateway timeout response has a 3xx status code
func (o *GetRoutingSmsAddressGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing sms address gateway timeout response has a 4xx status code
func (o *GetRoutingSmsAddressGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing sms address gateway timeout response has a 5xx status code
func (o *GetRoutingSmsAddressGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing sms address gateway timeout response a status code equal to that given
func (o *GetRoutingSmsAddressGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingSmsAddressGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSmsAddressGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/sms/addresses/{addressId}][%d] getRoutingSmsAddressGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSmsAddressGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSmsAddressGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
