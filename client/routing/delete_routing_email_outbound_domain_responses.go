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

// DeleteRoutingEmailOutboundDomainReader is a Reader for the DeleteRoutingEmailOutboundDomain structure.
type DeleteRoutingEmailOutboundDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingEmailOutboundDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingEmailOutboundDomainNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingEmailOutboundDomainBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingEmailOutboundDomainUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingEmailOutboundDomainForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingEmailOutboundDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingEmailOutboundDomainRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingEmailOutboundDomainRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingEmailOutboundDomainUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingEmailOutboundDomainTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingEmailOutboundDomainInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingEmailOutboundDomainServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingEmailOutboundDomainGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingEmailOutboundDomainNoContent creates a DeleteRoutingEmailOutboundDomainNoContent with default headers values
func NewDeleteRoutingEmailOutboundDomainNoContent() *DeleteRoutingEmailOutboundDomainNoContent {
	return &DeleteRoutingEmailOutboundDomainNoContent{}
}

/*
DeleteRoutingEmailOutboundDomainNoContent describes a response with status code 204, with default header values.

Operation was successful.
*/
type DeleteRoutingEmailOutboundDomainNoContent struct {
}

// IsSuccess returns true when this delete routing email outbound domain no content response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing email outbound domain no content response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain no content response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing email outbound domain no content response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain no content response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteRoutingEmailOutboundDomainNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainNoContent ", 204)
}

func (o *DeleteRoutingEmailOutboundDomainNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainNoContent ", 204)
}

func (o *DeleteRoutingEmailOutboundDomainNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingEmailOutboundDomainBadRequest creates a DeleteRoutingEmailOutboundDomainBadRequest with default headers values
func NewDeleteRoutingEmailOutboundDomainBadRequest() *DeleteRoutingEmailOutboundDomainBadRequest {
	return &DeleteRoutingEmailOutboundDomainBadRequest{}
}

/*
DeleteRoutingEmailOutboundDomainBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingEmailOutboundDomainBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain bad request response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain bad request response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain bad request response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain bad request response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain bad request response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRoutingEmailOutboundDomainBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainUnauthorized creates a DeleteRoutingEmailOutboundDomainUnauthorized with default headers values
func NewDeleteRoutingEmailOutboundDomainUnauthorized() *DeleteRoutingEmailOutboundDomainUnauthorized {
	return &DeleteRoutingEmailOutboundDomainUnauthorized{}
}

/*
DeleteRoutingEmailOutboundDomainUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingEmailOutboundDomainUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain unauthorized response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain unauthorized response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain unauthorized response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain unauthorized response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain unauthorized response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoutingEmailOutboundDomainUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainForbidden creates a DeleteRoutingEmailOutboundDomainForbidden with default headers values
func NewDeleteRoutingEmailOutboundDomainForbidden() *DeleteRoutingEmailOutboundDomainForbidden {
	return &DeleteRoutingEmailOutboundDomainForbidden{}
}

/*
DeleteRoutingEmailOutboundDomainForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingEmailOutboundDomainForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain forbidden response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain forbidden response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain forbidden response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain forbidden response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain forbidden response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoutingEmailOutboundDomainForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainNotFound creates a DeleteRoutingEmailOutboundDomainNotFound with default headers values
func NewDeleteRoutingEmailOutboundDomainNotFound() *DeleteRoutingEmailOutboundDomainNotFound {
	return &DeleteRoutingEmailOutboundDomainNotFound{}
}

/*
DeleteRoutingEmailOutboundDomainNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteRoutingEmailOutboundDomainNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain not found response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain not found response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain not found response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain not found response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain not found response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoutingEmailOutboundDomainNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainRequestTimeout creates a DeleteRoutingEmailOutboundDomainRequestTimeout with default headers values
func NewDeleteRoutingEmailOutboundDomainRequestTimeout() *DeleteRoutingEmailOutboundDomainRequestTimeout {
	return &DeleteRoutingEmailOutboundDomainRequestTimeout{}
}

/*
DeleteRoutingEmailOutboundDomainRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingEmailOutboundDomainRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain request timeout response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain request timeout response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain request timeout response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain request timeout response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain request timeout response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainRequestEntityTooLarge creates a DeleteRoutingEmailOutboundDomainRequestEntityTooLarge with default headers values
func NewDeleteRoutingEmailOutboundDomainRequestEntityTooLarge() *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge {
	return &DeleteRoutingEmailOutboundDomainRequestEntityTooLarge{}
}

/*
DeleteRoutingEmailOutboundDomainRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteRoutingEmailOutboundDomainRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain request entity too large response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain request entity too large response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain request entity too large response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain request entity too large response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain request entity too large response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainUnsupportedMediaType creates a DeleteRoutingEmailOutboundDomainUnsupportedMediaType with default headers values
func NewDeleteRoutingEmailOutboundDomainUnsupportedMediaType() *DeleteRoutingEmailOutboundDomainUnsupportedMediaType {
	return &DeleteRoutingEmailOutboundDomainUnsupportedMediaType{}
}

/*
DeleteRoutingEmailOutboundDomainUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingEmailOutboundDomainUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain unsupported media type response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain unsupported media type response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain unsupported media type response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain unsupported media type response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain unsupported media type response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainTooManyRequests creates a DeleteRoutingEmailOutboundDomainTooManyRequests with default headers values
func NewDeleteRoutingEmailOutboundDomainTooManyRequests() *DeleteRoutingEmailOutboundDomainTooManyRequests {
	return &DeleteRoutingEmailOutboundDomainTooManyRequests{}
}

/*
DeleteRoutingEmailOutboundDomainTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingEmailOutboundDomainTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain too many requests response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain too many requests response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain too many requests response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing email outbound domain too many requests response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing email outbound domain too many requests response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainInternalServerError creates a DeleteRoutingEmailOutboundDomainInternalServerError with default headers values
func NewDeleteRoutingEmailOutboundDomainInternalServerError() *DeleteRoutingEmailOutboundDomainInternalServerError {
	return &DeleteRoutingEmailOutboundDomainInternalServerError{}
}

/*
DeleteRoutingEmailOutboundDomainInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingEmailOutboundDomainInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain internal server error response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain internal server error response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain internal server error response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing email outbound domain internal server error response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing email outbound domain internal server error response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoutingEmailOutboundDomainInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainServiceUnavailable creates a DeleteRoutingEmailOutboundDomainServiceUnavailable with default headers values
func NewDeleteRoutingEmailOutboundDomainServiceUnavailable() *DeleteRoutingEmailOutboundDomainServiceUnavailable {
	return &DeleteRoutingEmailOutboundDomainServiceUnavailable{}
}

/*
DeleteRoutingEmailOutboundDomainServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingEmailOutboundDomainServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain service unavailable response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain service unavailable response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain service unavailable response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing email outbound domain service unavailable response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing email outbound domain service unavailable response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailOutboundDomainGatewayTimeout creates a DeleteRoutingEmailOutboundDomainGatewayTimeout with default headers values
func NewDeleteRoutingEmailOutboundDomainGatewayTimeout() *DeleteRoutingEmailOutboundDomainGatewayTimeout {
	return &DeleteRoutingEmailOutboundDomainGatewayTimeout{}
}

/*
DeleteRoutingEmailOutboundDomainGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteRoutingEmailOutboundDomainGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing email outbound domain gateway timeout response has a 2xx status code
func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing email outbound domain gateway timeout response has a 3xx status code
func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing email outbound domain gateway timeout response has a 4xx status code
func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing email outbound domain gateway timeout response has a 5xx status code
func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing email outbound domain gateway timeout response a status code equal to that given
func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/outbound/domains/{domainId}][%d] deleteRoutingEmailOutboundDomainGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailOutboundDomainGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
