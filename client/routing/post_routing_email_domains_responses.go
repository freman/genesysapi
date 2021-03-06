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

// PostRoutingEmailDomainsReader is a Reader for the PostRoutingEmailDomains structure.
type PostRoutingEmailDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingEmailDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRoutingEmailDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoutingEmailDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingEmailDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingEmailDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingEmailDomainsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingEmailDomainsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingEmailDomainsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingEmailDomainsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingEmailDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingEmailDomainsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingEmailDomainsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRoutingEmailDomainsOK creates a PostRoutingEmailDomainsOK with default headers values
func NewPostRoutingEmailDomainsOK() *PostRoutingEmailDomainsOK {
	return &PostRoutingEmailDomainsOK{}
}

/*PostRoutingEmailDomainsOK handles this case with default header values.

successful operation
*/
type PostRoutingEmailDomainsOK struct {
	Payload *models.InboundDomain
}

func (o *PostRoutingEmailDomainsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsOK  %+v", 200, o.Payload)
}

func (o *PostRoutingEmailDomainsOK) GetPayload() *models.InboundDomain {
	return o.Payload
}

func (o *PostRoutingEmailDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InboundDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsBadRequest creates a PostRoutingEmailDomainsBadRequest with default headers values
func NewPostRoutingEmailDomainsBadRequest() *PostRoutingEmailDomainsBadRequest {
	return &PostRoutingEmailDomainsBadRequest{}
}

/*PostRoutingEmailDomainsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingEmailDomainsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingEmailDomainsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsUnauthorized creates a PostRoutingEmailDomainsUnauthorized with default headers values
func NewPostRoutingEmailDomainsUnauthorized() *PostRoutingEmailDomainsUnauthorized {
	return &PostRoutingEmailDomainsUnauthorized{}
}

/*PostRoutingEmailDomainsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingEmailDomainsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingEmailDomainsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsForbidden creates a PostRoutingEmailDomainsForbidden with default headers values
func NewPostRoutingEmailDomainsForbidden() *PostRoutingEmailDomainsForbidden {
	return &PostRoutingEmailDomainsForbidden{}
}

/*PostRoutingEmailDomainsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingEmailDomainsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingEmailDomainsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsNotFound creates a PostRoutingEmailDomainsNotFound with default headers values
func NewPostRoutingEmailDomainsNotFound() *PostRoutingEmailDomainsNotFound {
	return &PostRoutingEmailDomainsNotFound{}
}

/*PostRoutingEmailDomainsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostRoutingEmailDomainsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingEmailDomainsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsRequestEntityTooLarge creates a PostRoutingEmailDomainsRequestEntityTooLarge with default headers values
func NewPostRoutingEmailDomainsRequestEntityTooLarge() *PostRoutingEmailDomainsRequestEntityTooLarge {
	return &PostRoutingEmailDomainsRequestEntityTooLarge{}
}

/*PostRoutingEmailDomainsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostRoutingEmailDomainsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingEmailDomainsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsUnsupportedMediaType creates a PostRoutingEmailDomainsUnsupportedMediaType with default headers values
func NewPostRoutingEmailDomainsUnsupportedMediaType() *PostRoutingEmailDomainsUnsupportedMediaType {
	return &PostRoutingEmailDomainsUnsupportedMediaType{}
}

/*PostRoutingEmailDomainsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingEmailDomainsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingEmailDomainsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsTooManyRequests creates a PostRoutingEmailDomainsTooManyRequests with default headers values
func NewPostRoutingEmailDomainsTooManyRequests() *PostRoutingEmailDomainsTooManyRequests {
	return &PostRoutingEmailDomainsTooManyRequests{}
}

/*PostRoutingEmailDomainsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostRoutingEmailDomainsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingEmailDomainsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsInternalServerError creates a PostRoutingEmailDomainsInternalServerError with default headers values
func NewPostRoutingEmailDomainsInternalServerError() *PostRoutingEmailDomainsInternalServerError {
	return &PostRoutingEmailDomainsInternalServerError{}
}

/*PostRoutingEmailDomainsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingEmailDomainsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingEmailDomainsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsServiceUnavailable creates a PostRoutingEmailDomainsServiceUnavailable with default headers values
func NewPostRoutingEmailDomainsServiceUnavailable() *PostRoutingEmailDomainsServiceUnavailable {
	return &PostRoutingEmailDomainsServiceUnavailable{}
}

/*PostRoutingEmailDomainsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingEmailDomainsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingEmailDomainsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingEmailDomainsGatewayTimeout creates a PostRoutingEmailDomainsGatewayTimeout with default headers values
func NewPostRoutingEmailDomainsGatewayTimeout() *PostRoutingEmailDomainsGatewayTimeout {
	return &PostRoutingEmailDomainsGatewayTimeout{}
}

/*PostRoutingEmailDomainsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostRoutingEmailDomainsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingEmailDomainsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/email/domains][%d] postRoutingEmailDomainsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingEmailDomainsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingEmailDomainsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
