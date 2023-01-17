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

// GetRoutingEmailDomainReader is a Reader for the GetRoutingEmailDomain structure.
type GetRoutingEmailDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingEmailDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingEmailDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingEmailDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingEmailDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingEmailDomainForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingEmailDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingEmailDomainRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingEmailDomainRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingEmailDomainUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingEmailDomainTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingEmailDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingEmailDomainServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingEmailDomainGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingEmailDomainOK creates a GetRoutingEmailDomainOK with default headers values
func NewGetRoutingEmailDomainOK() *GetRoutingEmailDomainOK {
	return &GetRoutingEmailDomainOK{}
}

/*
GetRoutingEmailDomainOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingEmailDomainOK struct {
	Payload *models.InboundDomain
}

// IsSuccess returns true when this get routing email domain o k response has a 2xx status code
func (o *GetRoutingEmailDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing email domain o k response has a 3xx status code
func (o *GetRoutingEmailDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain o k response has a 4xx status code
func (o *GetRoutingEmailDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email domain o k response has a 5xx status code
func (o *GetRoutingEmailDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain o k response a status code equal to that given
func (o *GetRoutingEmailDomainOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingEmailDomainOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailDomainOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailDomainOK) GetPayload() *models.InboundDomain {
	return o.Payload
}

func (o *GetRoutingEmailDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InboundDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainBadRequest creates a GetRoutingEmailDomainBadRequest with default headers values
func NewGetRoutingEmailDomainBadRequest() *GetRoutingEmailDomainBadRequest {
	return &GetRoutingEmailDomainBadRequest{}
}

/*
GetRoutingEmailDomainBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingEmailDomainBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain bad request response has a 2xx status code
func (o *GetRoutingEmailDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain bad request response has a 3xx status code
func (o *GetRoutingEmailDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain bad request response has a 4xx status code
func (o *GetRoutingEmailDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain bad request response has a 5xx status code
func (o *GetRoutingEmailDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain bad request response a status code equal to that given
func (o *GetRoutingEmailDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingEmailDomainBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailDomainBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailDomainBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainUnauthorized creates a GetRoutingEmailDomainUnauthorized with default headers values
func NewGetRoutingEmailDomainUnauthorized() *GetRoutingEmailDomainUnauthorized {
	return &GetRoutingEmailDomainUnauthorized{}
}

/*
GetRoutingEmailDomainUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingEmailDomainUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain unauthorized response has a 2xx status code
func (o *GetRoutingEmailDomainUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain unauthorized response has a 3xx status code
func (o *GetRoutingEmailDomainUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain unauthorized response has a 4xx status code
func (o *GetRoutingEmailDomainUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain unauthorized response has a 5xx status code
func (o *GetRoutingEmailDomainUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain unauthorized response a status code equal to that given
func (o *GetRoutingEmailDomainUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingEmailDomainUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailDomainUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailDomainUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainForbidden creates a GetRoutingEmailDomainForbidden with default headers values
func NewGetRoutingEmailDomainForbidden() *GetRoutingEmailDomainForbidden {
	return &GetRoutingEmailDomainForbidden{}
}

/*
GetRoutingEmailDomainForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingEmailDomainForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain forbidden response has a 2xx status code
func (o *GetRoutingEmailDomainForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain forbidden response has a 3xx status code
func (o *GetRoutingEmailDomainForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain forbidden response has a 4xx status code
func (o *GetRoutingEmailDomainForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain forbidden response has a 5xx status code
func (o *GetRoutingEmailDomainForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain forbidden response a status code equal to that given
func (o *GetRoutingEmailDomainForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingEmailDomainForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailDomainForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailDomainForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainNotFound creates a GetRoutingEmailDomainNotFound with default headers values
func NewGetRoutingEmailDomainNotFound() *GetRoutingEmailDomainNotFound {
	return &GetRoutingEmailDomainNotFound{}
}

/*
GetRoutingEmailDomainNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingEmailDomainNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain not found response has a 2xx status code
func (o *GetRoutingEmailDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain not found response has a 3xx status code
func (o *GetRoutingEmailDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain not found response has a 4xx status code
func (o *GetRoutingEmailDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain not found response has a 5xx status code
func (o *GetRoutingEmailDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain not found response a status code equal to that given
func (o *GetRoutingEmailDomainNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingEmailDomainNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailDomainNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailDomainNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRequestTimeout creates a GetRoutingEmailDomainRequestTimeout with default headers values
func NewGetRoutingEmailDomainRequestTimeout() *GetRoutingEmailDomainRequestTimeout {
	return &GetRoutingEmailDomainRequestTimeout{}
}

/*
GetRoutingEmailDomainRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingEmailDomainRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain request timeout response has a 2xx status code
func (o *GetRoutingEmailDomainRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain request timeout response has a 3xx status code
func (o *GetRoutingEmailDomainRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain request timeout response has a 4xx status code
func (o *GetRoutingEmailDomainRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain request timeout response has a 5xx status code
func (o *GetRoutingEmailDomainRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain request timeout response a status code equal to that given
func (o *GetRoutingEmailDomainRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingEmailDomainRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailDomainRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailDomainRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainRequestEntityTooLarge creates a GetRoutingEmailDomainRequestEntityTooLarge with default headers values
func NewGetRoutingEmailDomainRequestEntityTooLarge() *GetRoutingEmailDomainRequestEntityTooLarge {
	return &GetRoutingEmailDomainRequestEntityTooLarge{}
}

/*
GetRoutingEmailDomainRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingEmailDomainRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain request entity too large response has a 2xx status code
func (o *GetRoutingEmailDomainRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain request entity too large response has a 3xx status code
func (o *GetRoutingEmailDomainRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain request entity too large response has a 4xx status code
func (o *GetRoutingEmailDomainRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain request entity too large response has a 5xx status code
func (o *GetRoutingEmailDomainRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain request entity too large response a status code equal to that given
func (o *GetRoutingEmailDomainRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingEmailDomainRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailDomainRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailDomainRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainUnsupportedMediaType creates a GetRoutingEmailDomainUnsupportedMediaType with default headers values
func NewGetRoutingEmailDomainUnsupportedMediaType() *GetRoutingEmailDomainUnsupportedMediaType {
	return &GetRoutingEmailDomainUnsupportedMediaType{}
}

/*
GetRoutingEmailDomainUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingEmailDomainUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain unsupported media type response has a 2xx status code
func (o *GetRoutingEmailDomainUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain unsupported media type response has a 3xx status code
func (o *GetRoutingEmailDomainUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain unsupported media type response has a 4xx status code
func (o *GetRoutingEmailDomainUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain unsupported media type response has a 5xx status code
func (o *GetRoutingEmailDomainUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain unsupported media type response a status code equal to that given
func (o *GetRoutingEmailDomainUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingEmailDomainUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailDomainUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailDomainUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainTooManyRequests creates a GetRoutingEmailDomainTooManyRequests with default headers values
func NewGetRoutingEmailDomainTooManyRequests() *GetRoutingEmailDomainTooManyRequests {
	return &GetRoutingEmailDomainTooManyRequests{}
}

/*
GetRoutingEmailDomainTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingEmailDomainTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain too many requests response has a 2xx status code
func (o *GetRoutingEmailDomainTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain too many requests response has a 3xx status code
func (o *GetRoutingEmailDomainTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain too many requests response has a 4xx status code
func (o *GetRoutingEmailDomainTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email domain too many requests response has a 5xx status code
func (o *GetRoutingEmailDomainTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email domain too many requests response a status code equal to that given
func (o *GetRoutingEmailDomainTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingEmailDomainTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailDomainTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailDomainTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainInternalServerError creates a GetRoutingEmailDomainInternalServerError with default headers values
func NewGetRoutingEmailDomainInternalServerError() *GetRoutingEmailDomainInternalServerError {
	return &GetRoutingEmailDomainInternalServerError{}
}

/*
GetRoutingEmailDomainInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingEmailDomainInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain internal server error response has a 2xx status code
func (o *GetRoutingEmailDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain internal server error response has a 3xx status code
func (o *GetRoutingEmailDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain internal server error response has a 4xx status code
func (o *GetRoutingEmailDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email domain internal server error response has a 5xx status code
func (o *GetRoutingEmailDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing email domain internal server error response a status code equal to that given
func (o *GetRoutingEmailDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingEmailDomainInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailDomainInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailDomainInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainServiceUnavailable creates a GetRoutingEmailDomainServiceUnavailable with default headers values
func NewGetRoutingEmailDomainServiceUnavailable() *GetRoutingEmailDomainServiceUnavailable {
	return &GetRoutingEmailDomainServiceUnavailable{}
}

/*
GetRoutingEmailDomainServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingEmailDomainServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain service unavailable response has a 2xx status code
func (o *GetRoutingEmailDomainServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain service unavailable response has a 3xx status code
func (o *GetRoutingEmailDomainServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain service unavailable response has a 4xx status code
func (o *GetRoutingEmailDomainServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email domain service unavailable response has a 5xx status code
func (o *GetRoutingEmailDomainServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing email domain service unavailable response a status code equal to that given
func (o *GetRoutingEmailDomainServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingEmailDomainServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailDomainServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailDomainServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainGatewayTimeout creates a GetRoutingEmailDomainGatewayTimeout with default headers values
func NewGetRoutingEmailDomainGatewayTimeout() *GetRoutingEmailDomainGatewayTimeout {
	return &GetRoutingEmailDomainGatewayTimeout{}
}

/*
GetRoutingEmailDomainGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingEmailDomainGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email domain gateway timeout response has a 2xx status code
func (o *GetRoutingEmailDomainGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email domain gateway timeout response has a 3xx status code
func (o *GetRoutingEmailDomainGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email domain gateway timeout response has a 4xx status code
func (o *GetRoutingEmailDomainGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email domain gateway timeout response has a 5xx status code
func (o *GetRoutingEmailDomainGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing email domain gateway timeout response a status code equal to that given
func (o *GetRoutingEmailDomainGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingEmailDomainGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailDomainGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains/{domainId}][%d] getRoutingEmailDomainGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailDomainGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
