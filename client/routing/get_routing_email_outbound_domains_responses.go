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

// GetRoutingEmailOutboundDomainsReader is a Reader for the GetRoutingEmailOutboundDomains structure.
type GetRoutingEmailOutboundDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingEmailOutboundDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingEmailOutboundDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingEmailOutboundDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingEmailOutboundDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingEmailOutboundDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingEmailOutboundDomainsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingEmailOutboundDomainsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingEmailOutboundDomainsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingEmailOutboundDomainsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingEmailOutboundDomainsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingEmailOutboundDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingEmailOutboundDomainsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingEmailOutboundDomainsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingEmailOutboundDomainsOK creates a GetRoutingEmailOutboundDomainsOK with default headers values
func NewGetRoutingEmailOutboundDomainsOK() *GetRoutingEmailOutboundDomainsOK {
	return &GetRoutingEmailOutboundDomainsOK{}
}

/*
GetRoutingEmailOutboundDomainsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingEmailOutboundDomainsOK struct {
	Payload *models.OutboundDomainEntityListing
}

// IsSuccess returns true when this get routing email outbound domains o k response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing email outbound domains o k response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains o k response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email outbound domains o k response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains o k response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingEmailOutboundDomainsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsOK) GetPayload() *models.OutboundDomainEntityListing {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutboundDomainEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsBadRequest creates a GetRoutingEmailOutboundDomainsBadRequest with default headers values
func NewGetRoutingEmailOutboundDomainsBadRequest() *GetRoutingEmailOutboundDomainsBadRequest {
	return &GetRoutingEmailOutboundDomainsBadRequest{}
}

/*
GetRoutingEmailOutboundDomainsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingEmailOutboundDomainsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains bad request response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains bad request response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains bad request response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains bad request response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains bad request response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingEmailOutboundDomainsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsUnauthorized creates a GetRoutingEmailOutboundDomainsUnauthorized with default headers values
func NewGetRoutingEmailOutboundDomainsUnauthorized() *GetRoutingEmailOutboundDomainsUnauthorized {
	return &GetRoutingEmailOutboundDomainsUnauthorized{}
}

/*
GetRoutingEmailOutboundDomainsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingEmailOutboundDomainsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains unauthorized response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains unauthorized response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains unauthorized response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains unauthorized response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains unauthorized response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingEmailOutboundDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsForbidden creates a GetRoutingEmailOutboundDomainsForbidden with default headers values
func NewGetRoutingEmailOutboundDomainsForbidden() *GetRoutingEmailOutboundDomainsForbidden {
	return &GetRoutingEmailOutboundDomainsForbidden{}
}

/*
GetRoutingEmailOutboundDomainsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingEmailOutboundDomainsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains forbidden response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains forbidden response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains forbidden response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains forbidden response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains forbidden response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingEmailOutboundDomainsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsNotFound creates a GetRoutingEmailOutboundDomainsNotFound with default headers values
func NewGetRoutingEmailOutboundDomainsNotFound() *GetRoutingEmailOutboundDomainsNotFound {
	return &GetRoutingEmailOutboundDomainsNotFound{}
}

/*
GetRoutingEmailOutboundDomainsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingEmailOutboundDomainsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains not found response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains not found response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains not found response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains not found response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains not found response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingEmailOutboundDomainsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsRequestTimeout creates a GetRoutingEmailOutboundDomainsRequestTimeout with default headers values
func NewGetRoutingEmailOutboundDomainsRequestTimeout() *GetRoutingEmailOutboundDomainsRequestTimeout {
	return &GetRoutingEmailOutboundDomainsRequestTimeout{}
}

/*
GetRoutingEmailOutboundDomainsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingEmailOutboundDomainsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains request timeout response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains request timeout response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains request timeout response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains request timeout response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains request timeout response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingEmailOutboundDomainsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsRequestEntityTooLarge creates a GetRoutingEmailOutboundDomainsRequestEntityTooLarge with default headers values
func NewGetRoutingEmailOutboundDomainsRequestEntityTooLarge() *GetRoutingEmailOutboundDomainsRequestEntityTooLarge {
	return &GetRoutingEmailOutboundDomainsRequestEntityTooLarge{}
}

/*
GetRoutingEmailOutboundDomainsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingEmailOutboundDomainsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains request entity too large response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains request entity too large response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains request entity too large response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains request entity too large response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains request entity too large response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsUnsupportedMediaType creates a GetRoutingEmailOutboundDomainsUnsupportedMediaType with default headers values
func NewGetRoutingEmailOutboundDomainsUnsupportedMediaType() *GetRoutingEmailOutboundDomainsUnsupportedMediaType {
	return &GetRoutingEmailOutboundDomainsUnsupportedMediaType{}
}

/*
GetRoutingEmailOutboundDomainsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingEmailOutboundDomainsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains unsupported media type response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains unsupported media type response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains unsupported media type response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains unsupported media type response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains unsupported media type response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsTooManyRequests creates a GetRoutingEmailOutboundDomainsTooManyRequests with default headers values
func NewGetRoutingEmailOutboundDomainsTooManyRequests() *GetRoutingEmailOutboundDomainsTooManyRequests {
	return &GetRoutingEmailOutboundDomainsTooManyRequests{}
}

/*
GetRoutingEmailOutboundDomainsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingEmailOutboundDomainsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains too many requests response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains too many requests response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains too many requests response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing email outbound domains too many requests response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing email outbound domains too many requests response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingEmailOutboundDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsInternalServerError creates a GetRoutingEmailOutboundDomainsInternalServerError with default headers values
func NewGetRoutingEmailOutboundDomainsInternalServerError() *GetRoutingEmailOutboundDomainsInternalServerError {
	return &GetRoutingEmailOutboundDomainsInternalServerError{}
}

/*
GetRoutingEmailOutboundDomainsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingEmailOutboundDomainsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains internal server error response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains internal server error response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains internal server error response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email outbound domains internal server error response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing email outbound domains internal server error response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingEmailOutboundDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsServiceUnavailable creates a GetRoutingEmailOutboundDomainsServiceUnavailable with default headers values
func NewGetRoutingEmailOutboundDomainsServiceUnavailable() *GetRoutingEmailOutboundDomainsServiceUnavailable {
	return &GetRoutingEmailOutboundDomainsServiceUnavailable{}
}

/*
GetRoutingEmailOutboundDomainsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingEmailOutboundDomainsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains service unavailable response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains service unavailable response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains service unavailable response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email outbound domains service unavailable response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing email outbound domains service unavailable response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailOutboundDomainsGatewayTimeout creates a GetRoutingEmailOutboundDomainsGatewayTimeout with default headers values
func NewGetRoutingEmailOutboundDomainsGatewayTimeout() *GetRoutingEmailOutboundDomainsGatewayTimeout {
	return &GetRoutingEmailOutboundDomainsGatewayTimeout{}
}

/*
GetRoutingEmailOutboundDomainsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingEmailOutboundDomainsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing email outbound domains gateway timeout response has a 2xx status code
func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing email outbound domains gateway timeout response has a 3xx status code
func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing email outbound domains gateway timeout response has a 4xx status code
func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing email outbound domains gateway timeout response has a 5xx status code
func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing email outbound domains gateway timeout response a status code equal to that given
func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/outbound/domains][%d] getRoutingEmailOutboundDomainsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailOutboundDomainsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
