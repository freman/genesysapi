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

// PostRoutingEmailOutboundDomainsSimulatedReader is a Reader for the PostRoutingEmailOutboundDomainsSimulated structure.
type PostRoutingEmailOutboundDomainsSimulatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingEmailOutboundDomainsSimulatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostRoutingEmailOutboundDomainsSimulatedAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoutingEmailOutboundDomainsSimulatedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingEmailOutboundDomainsSimulatedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingEmailOutboundDomainsSimulatedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingEmailOutboundDomainsSimulatedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRoutingEmailOutboundDomainsSimulatedRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingEmailOutboundDomainsSimulatedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingEmailOutboundDomainsSimulatedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingEmailOutboundDomainsSimulatedServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingEmailOutboundDomainsSimulatedGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRoutingEmailOutboundDomainsSimulatedAccepted creates a PostRoutingEmailOutboundDomainsSimulatedAccepted with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedAccepted() *PostRoutingEmailOutboundDomainsSimulatedAccepted {
	return &PostRoutingEmailOutboundDomainsSimulatedAccepted{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedAccepted describes a response with status code 202, with default header values.

Creation request was successful. Need to wait for creation to complete
*/
type PostRoutingEmailOutboundDomainsSimulatedAccepted struct {
	Payload *models.EmailOutboundDomainResult
}

// IsSuccess returns true when this post routing email outbound domains simulated accepted response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post routing email outbound domains simulated accepted response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated accepted response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email outbound domains simulated accepted response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated accepted response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedAccepted  %+v", 202, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedAccepted  %+v", 202, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) GetPayload() *models.EmailOutboundDomainResult {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailOutboundDomainResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedBadRequest creates a PostRoutingEmailOutboundDomainsSimulatedBadRequest with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedBadRequest() *PostRoutingEmailOutboundDomainsSimulatedBadRequest {
	return &PostRoutingEmailOutboundDomainsSimulatedBadRequest{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingEmailOutboundDomainsSimulatedBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated bad request response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated bad request response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated bad request response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated bad request response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated bad request response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedUnauthorized creates a PostRoutingEmailOutboundDomainsSimulatedUnauthorized with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedUnauthorized() *PostRoutingEmailOutboundDomainsSimulatedUnauthorized {
	return &PostRoutingEmailOutboundDomainsSimulatedUnauthorized{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingEmailOutboundDomainsSimulatedUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated unauthorized response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated unauthorized response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated unauthorized response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated unauthorized response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated unauthorized response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedForbidden creates a PostRoutingEmailOutboundDomainsSimulatedForbidden with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedForbidden() *PostRoutingEmailOutboundDomainsSimulatedForbidden {
	return &PostRoutingEmailOutboundDomainsSimulatedForbidden{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingEmailOutboundDomainsSimulatedForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated forbidden response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated forbidden response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated forbidden response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated forbidden response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated forbidden response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedNotFound creates a PostRoutingEmailOutboundDomainsSimulatedNotFound with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedNotFound() *PostRoutingEmailOutboundDomainsSimulatedNotFound {
	return &PostRoutingEmailOutboundDomainsSimulatedNotFound{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRoutingEmailOutboundDomainsSimulatedNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated not found response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated not found response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated not found response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated not found response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated not found response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedRequestTimeout creates a PostRoutingEmailOutboundDomainsSimulatedRequestTimeout with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedRequestTimeout() *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout {
	return &PostRoutingEmailOutboundDomainsSimulatedRequestTimeout{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRoutingEmailOutboundDomainsSimulatedRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated request timeout response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated request timeout response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated request timeout response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated request timeout response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated request timeout response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge creates a PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge() *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge {
	return &PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated request entity too large response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated request entity too large response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated request entity too large response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated request entity too large response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated request entity too large response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType creates a PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType() *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType {
	return &PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated unsupported media type response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated unsupported media type response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated unsupported media type response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated unsupported media type response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated unsupported media type response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedTooManyRequests creates a PostRoutingEmailOutboundDomainsSimulatedTooManyRequests with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedTooManyRequests() *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests {
	return &PostRoutingEmailOutboundDomainsSimulatedTooManyRequests{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRoutingEmailOutboundDomainsSimulatedTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated too many requests response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated too many requests response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated too many requests response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing email outbound domains simulated too many requests response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing email outbound domains simulated too many requests response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedInternalServerError creates a PostRoutingEmailOutboundDomainsSimulatedInternalServerError with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedInternalServerError() *PostRoutingEmailOutboundDomainsSimulatedInternalServerError {
	return &PostRoutingEmailOutboundDomainsSimulatedInternalServerError{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingEmailOutboundDomainsSimulatedInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated internal server error response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated internal server error response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated internal server error response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email outbound domains simulated internal server error response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing email outbound domains simulated internal server error response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedServiceUnavailable creates a PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedServiceUnavailable() *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable {
	return &PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated service unavailable response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated service unavailable response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated service unavailable response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email outbound domains simulated service unavailable response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing email outbound domains simulated service unavailable response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailOutboundDomainsSimulatedGatewayTimeout creates a PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout with default headers values
func NewPostRoutingEmailOutboundDomainsSimulatedGatewayTimeout() *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout {
	return &PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout{}
}

/*
PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing email outbound domains simulated gateway timeout response has a 2xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing email outbound domains simulated gateway timeout response has a 3xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing email outbound domains simulated gateway timeout response has a 4xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing email outbound domains simulated gateway timeout response has a 5xx status code
func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing email outbound domains simulated gateway timeout response a status code equal to that given
func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/outbound/domains/simulated][%d] postRoutingEmailOutboundDomainsSimulatedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailOutboundDomainsSimulatedGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}