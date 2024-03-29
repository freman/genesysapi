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

// PostRoutingEmailDomainRoutesReader is a Reader for the PostRoutingEmailDomainRoutes structure.
type PostRoutingEmailDomainRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingEmailDomainRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRoutingEmailDomainRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoutingEmailDomainRoutesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingEmailDomainRoutesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingEmailDomainRoutesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingEmailDomainRoutesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRoutingEmailDomainRoutesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingEmailDomainRoutesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingEmailDomainRoutesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingEmailDomainRoutesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingEmailDomainRoutesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingEmailDomainRoutesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingEmailDomainRoutesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRoutingEmailDomainRoutesOK creates a PostRoutingEmailDomainRoutesOK with default headers values
func NewPostRoutingEmailDomainRoutesOK() *PostRoutingEmailDomainRoutesOK {
	return &PostRoutingEmailDomainRoutesOK{}
}

/*
PostRoutingEmailDomainRoutesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostRoutingEmailDomainRoutesOK struct {
	Payload *models.InboundRoute
}

// IsSuccess returns true when this post routing email domain routes o k response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post routing email domain routes o k response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes o k response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email domain routes o k response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes o k response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostRoutingEmailDomainRoutesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesOK  %+v", 200, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesOK  %+v", 200, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesOK) GetPayload() *models.InboundRoute {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InboundRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesBadRequest creates a PostRoutingEmailDomainRoutesBadRequest with default headers values
func NewPostRoutingEmailDomainRoutesBadRequest() *PostRoutingEmailDomainRoutesBadRequest {
	return &PostRoutingEmailDomainRoutesBadRequest{}
}

/*
PostRoutingEmailDomainRoutesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingEmailDomainRoutesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes bad request response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes bad request response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes bad request response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes bad request response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes bad request response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRoutingEmailDomainRoutesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesUnauthorized creates a PostRoutingEmailDomainRoutesUnauthorized with default headers values
func NewPostRoutingEmailDomainRoutesUnauthorized() *PostRoutingEmailDomainRoutesUnauthorized {
	return &PostRoutingEmailDomainRoutesUnauthorized{}
}

/*
PostRoutingEmailDomainRoutesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingEmailDomainRoutesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes unauthorized response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes unauthorized response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes unauthorized response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes unauthorized response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes unauthorized response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRoutingEmailDomainRoutesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesForbidden creates a PostRoutingEmailDomainRoutesForbidden with default headers values
func NewPostRoutingEmailDomainRoutesForbidden() *PostRoutingEmailDomainRoutesForbidden {
	return &PostRoutingEmailDomainRoutesForbidden{}
}

/*
PostRoutingEmailDomainRoutesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingEmailDomainRoutesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes forbidden response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes forbidden response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes forbidden response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes forbidden response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes forbidden response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRoutingEmailDomainRoutesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesNotFound creates a PostRoutingEmailDomainRoutesNotFound with default headers values
func NewPostRoutingEmailDomainRoutesNotFound() *PostRoutingEmailDomainRoutesNotFound {
	return &PostRoutingEmailDomainRoutesNotFound{}
}

/*
PostRoutingEmailDomainRoutesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRoutingEmailDomainRoutesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes not found response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes not found response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes not found response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes not found response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes not found response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRoutingEmailDomainRoutesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesRequestTimeout creates a PostRoutingEmailDomainRoutesRequestTimeout with default headers values
func NewPostRoutingEmailDomainRoutesRequestTimeout() *PostRoutingEmailDomainRoutesRequestTimeout {
	return &PostRoutingEmailDomainRoutesRequestTimeout{}
}

/*
PostRoutingEmailDomainRoutesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRoutingEmailDomainRoutesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes request timeout response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes request timeout response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes request timeout response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes request timeout response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes request timeout response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRoutingEmailDomainRoutesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesRequestEntityTooLarge creates a PostRoutingEmailDomainRoutesRequestEntityTooLarge with default headers values
func NewPostRoutingEmailDomainRoutesRequestEntityTooLarge() *PostRoutingEmailDomainRoutesRequestEntityTooLarge {
	return &PostRoutingEmailDomainRoutesRequestEntityTooLarge{}
}

/*
PostRoutingEmailDomainRoutesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostRoutingEmailDomainRoutesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes request entity too large response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes request entity too large response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes request entity too large response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes request entity too large response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes request entity too large response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesUnsupportedMediaType creates a PostRoutingEmailDomainRoutesUnsupportedMediaType with default headers values
func NewPostRoutingEmailDomainRoutesUnsupportedMediaType() *PostRoutingEmailDomainRoutesUnsupportedMediaType {
	return &PostRoutingEmailDomainRoutesUnsupportedMediaType{}
}

/*
PostRoutingEmailDomainRoutesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingEmailDomainRoutesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes unsupported media type response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes unsupported media type response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes unsupported media type response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes unsupported media type response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes unsupported media type response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesTooManyRequests creates a PostRoutingEmailDomainRoutesTooManyRequests with default headers values
func NewPostRoutingEmailDomainRoutesTooManyRequests() *PostRoutingEmailDomainRoutesTooManyRequests {
	return &PostRoutingEmailDomainRoutesTooManyRequests{}
}

/*
PostRoutingEmailDomainRoutesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRoutingEmailDomainRoutesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes too many requests response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes too many requests response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes too many requests response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email domain routes too many requests response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email domain routes too many requests response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRoutingEmailDomainRoutesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesInternalServerError creates a PostRoutingEmailDomainRoutesInternalServerError with default headers values
func NewPostRoutingEmailDomainRoutesInternalServerError() *PostRoutingEmailDomainRoutesInternalServerError {
	return &PostRoutingEmailDomainRoutesInternalServerError{}
}

/*
PostRoutingEmailDomainRoutesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingEmailDomainRoutesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes internal server error response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes internal server error response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes internal server error response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email domain routes internal server error response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing email domain routes internal server error response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRoutingEmailDomainRoutesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesServiceUnavailable creates a PostRoutingEmailDomainRoutesServiceUnavailable with default headers values
func NewPostRoutingEmailDomainRoutesServiceUnavailable() *PostRoutingEmailDomainRoutesServiceUnavailable {
	return &PostRoutingEmailDomainRoutesServiceUnavailable{}
}

/*
PostRoutingEmailDomainRoutesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingEmailDomainRoutesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes service unavailable response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes service unavailable response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes service unavailable response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email domain routes service unavailable response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing email domain routes service unavailable response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRoutingEmailDomainRoutesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainRoutesGatewayTimeout creates a PostRoutingEmailDomainRoutesGatewayTimeout with default headers values
func NewPostRoutingEmailDomainRoutesGatewayTimeout() *PostRoutingEmailDomainRoutesGatewayTimeout {
	return &PostRoutingEmailDomainRoutesGatewayTimeout{}
}

/*
PostRoutingEmailDomainRoutesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRoutingEmailDomainRoutesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email domain routes gateway timeout response has a 2xx status code
func (o *PostRoutingEmailDomainRoutesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email domain routes gateway timeout response has a 3xx status code
func (o *PostRoutingEmailDomainRoutesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email domain routes gateway timeout response has a 4xx status code
func (o *PostRoutingEmailDomainRoutesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email domain routes gateway timeout response has a 5xx status code
func (o *PostRoutingEmailDomainRoutesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing email domain routes gateway timeout response a status code equal to that given
func (o *PostRoutingEmailDomainRoutesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRoutingEmailDomainRoutesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains/{domainName}/routes][%d] postRoutingEmailDomainRoutesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingEmailDomainRoutesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainRoutesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
