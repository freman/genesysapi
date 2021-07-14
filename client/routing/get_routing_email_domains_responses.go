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

// GetRoutingEmailDomainsReader is a Reader for the GetRoutingEmailDomains structure.
type GetRoutingEmailDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingEmailDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingEmailDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingEmailDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingEmailDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingEmailDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingEmailDomainsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingEmailDomainsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingEmailDomainsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingEmailDomainsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingEmailDomainsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingEmailDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingEmailDomainsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingEmailDomainsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingEmailDomainsOK creates a GetRoutingEmailDomainsOK with default headers values
func NewGetRoutingEmailDomainsOK() *GetRoutingEmailDomainsOK {
	return &GetRoutingEmailDomainsOK{}
}

/*GetRoutingEmailDomainsOK handles this case with default header values.

successful operation
*/
type GetRoutingEmailDomainsOK struct {
	Payload *models.InboundDomainEntityListing
}

func (o *GetRoutingEmailDomainsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailDomainsOK) GetPayload() *models.InboundDomainEntityListing {
	return o.Payload
}

func (o *GetRoutingEmailDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InboundDomainEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsBadRequest creates a GetRoutingEmailDomainsBadRequest with default headers values
func NewGetRoutingEmailDomainsBadRequest() *GetRoutingEmailDomainsBadRequest {
	return &GetRoutingEmailDomainsBadRequest{}
}

/*GetRoutingEmailDomainsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingEmailDomainsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailDomainsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsUnauthorized creates a GetRoutingEmailDomainsUnauthorized with default headers values
func NewGetRoutingEmailDomainsUnauthorized() *GetRoutingEmailDomainsUnauthorized {
	return &GetRoutingEmailDomainsUnauthorized{}
}

/*GetRoutingEmailDomainsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingEmailDomainsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailDomainsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsForbidden creates a GetRoutingEmailDomainsForbidden with default headers values
func NewGetRoutingEmailDomainsForbidden() *GetRoutingEmailDomainsForbidden {
	return &GetRoutingEmailDomainsForbidden{}
}

/*GetRoutingEmailDomainsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingEmailDomainsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailDomainsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsNotFound creates a GetRoutingEmailDomainsNotFound with default headers values
func NewGetRoutingEmailDomainsNotFound() *GetRoutingEmailDomainsNotFound {
	return &GetRoutingEmailDomainsNotFound{}
}

/*GetRoutingEmailDomainsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingEmailDomainsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailDomainsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsRequestTimeout creates a GetRoutingEmailDomainsRequestTimeout with default headers values
func NewGetRoutingEmailDomainsRequestTimeout() *GetRoutingEmailDomainsRequestTimeout {
	return &GetRoutingEmailDomainsRequestTimeout{}
}

/*GetRoutingEmailDomainsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingEmailDomainsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailDomainsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsRequestEntityTooLarge creates a GetRoutingEmailDomainsRequestEntityTooLarge with default headers values
func NewGetRoutingEmailDomainsRequestEntityTooLarge() *GetRoutingEmailDomainsRequestEntityTooLarge {
	return &GetRoutingEmailDomainsRequestEntityTooLarge{}
}

/*GetRoutingEmailDomainsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingEmailDomainsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailDomainsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsUnsupportedMediaType creates a GetRoutingEmailDomainsUnsupportedMediaType with default headers values
func NewGetRoutingEmailDomainsUnsupportedMediaType() *GetRoutingEmailDomainsUnsupportedMediaType {
	return &GetRoutingEmailDomainsUnsupportedMediaType{}
}

/*GetRoutingEmailDomainsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingEmailDomainsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailDomainsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsTooManyRequests creates a GetRoutingEmailDomainsTooManyRequests with default headers values
func NewGetRoutingEmailDomainsTooManyRequests() *GetRoutingEmailDomainsTooManyRequests {
	return &GetRoutingEmailDomainsTooManyRequests{}
}

/*GetRoutingEmailDomainsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingEmailDomainsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailDomainsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsInternalServerError creates a GetRoutingEmailDomainsInternalServerError with default headers values
func NewGetRoutingEmailDomainsInternalServerError() *GetRoutingEmailDomainsInternalServerError {
	return &GetRoutingEmailDomainsInternalServerError{}
}

/*GetRoutingEmailDomainsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingEmailDomainsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailDomainsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsServiceUnavailable creates a GetRoutingEmailDomainsServiceUnavailable with default headers values
func NewGetRoutingEmailDomainsServiceUnavailable() *GetRoutingEmailDomainsServiceUnavailable {
	return &GetRoutingEmailDomainsServiceUnavailable{}
}

/*GetRoutingEmailDomainsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingEmailDomainsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailDomainsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailDomainsGatewayTimeout creates a GetRoutingEmailDomainsGatewayTimeout with default headers values
func NewGetRoutingEmailDomainsGatewayTimeout() *GetRoutingEmailDomainsGatewayTimeout {
	return &GetRoutingEmailDomainsGatewayTimeout{}
}

/*GetRoutingEmailDomainsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingEmailDomainsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailDomainsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/domains][%d] getRoutingEmailDomainsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailDomainsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailDomainsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
