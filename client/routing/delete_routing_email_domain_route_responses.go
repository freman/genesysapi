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

// DeleteRoutingEmailDomainRouteReader is a Reader for the DeleteRoutingEmailDomainRoute structure.
type DeleteRoutingEmailDomainRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingEmailDomainRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoutingEmailDomainRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingEmailDomainRouteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingEmailDomainRouteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingEmailDomainRouteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingEmailDomainRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingEmailDomainRouteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteRoutingEmailDomainRouteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingEmailDomainRouteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingEmailDomainRouteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingEmailDomainRouteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingEmailDomainRouteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingEmailDomainRouteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingEmailDomainRouteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingEmailDomainRouteOK creates a DeleteRoutingEmailDomainRouteOK with default headers values
func NewDeleteRoutingEmailDomainRouteOK() *DeleteRoutingEmailDomainRouteOK {
	return &DeleteRoutingEmailDomainRouteOK{}
}

/*DeleteRoutingEmailDomainRouteOK handles this case with default header values.

Operation was successful.
*/
type DeleteRoutingEmailDomainRouteOK struct {
}

func (o *DeleteRoutingEmailDomainRouteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteOK ", 200)
}

func (o *DeleteRoutingEmailDomainRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingEmailDomainRouteBadRequest creates a DeleteRoutingEmailDomainRouteBadRequest with default headers values
func NewDeleteRoutingEmailDomainRouteBadRequest() *DeleteRoutingEmailDomainRouteBadRequest {
	return &DeleteRoutingEmailDomainRouteBadRequest{}
}

/*DeleteRoutingEmailDomainRouteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingEmailDomainRouteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteUnauthorized creates a DeleteRoutingEmailDomainRouteUnauthorized with default headers values
func NewDeleteRoutingEmailDomainRouteUnauthorized() *DeleteRoutingEmailDomainRouteUnauthorized {
	return &DeleteRoutingEmailDomainRouteUnauthorized{}
}

/*DeleteRoutingEmailDomainRouteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingEmailDomainRouteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteForbidden creates a DeleteRoutingEmailDomainRouteForbidden with default headers values
func NewDeleteRoutingEmailDomainRouteForbidden() *DeleteRoutingEmailDomainRouteForbidden {
	return &DeleteRoutingEmailDomainRouteForbidden{}
}

/*DeleteRoutingEmailDomainRouteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingEmailDomainRouteForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteNotFound creates a DeleteRoutingEmailDomainRouteNotFound with default headers values
func NewDeleteRoutingEmailDomainRouteNotFound() *DeleteRoutingEmailDomainRouteNotFound {
	return &DeleteRoutingEmailDomainRouteNotFound{}
}

/*DeleteRoutingEmailDomainRouteNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRoutingEmailDomainRouteNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteRequestTimeout creates a DeleteRoutingEmailDomainRouteRequestTimeout with default headers values
func NewDeleteRoutingEmailDomainRouteRequestTimeout() *DeleteRoutingEmailDomainRouteRequestTimeout {
	return &DeleteRoutingEmailDomainRouteRequestTimeout{}
}

/*DeleteRoutingEmailDomainRouteRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingEmailDomainRouteRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteConflict creates a DeleteRoutingEmailDomainRouteConflict with default headers values
func NewDeleteRoutingEmailDomainRouteConflict() *DeleteRoutingEmailDomainRouteConflict {
	return &DeleteRoutingEmailDomainRouteConflict{}
}

/*DeleteRoutingEmailDomainRouteConflict handles this case with default header values.

Conflict
*/
type DeleteRoutingEmailDomainRouteConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteConflict  %+v", 409, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteRequestEntityTooLarge creates a DeleteRoutingEmailDomainRouteRequestEntityTooLarge with default headers values
func NewDeleteRoutingEmailDomainRouteRequestEntityTooLarge() *DeleteRoutingEmailDomainRouteRequestEntityTooLarge {
	return &DeleteRoutingEmailDomainRouteRequestEntityTooLarge{}
}

/*DeleteRoutingEmailDomainRouteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRoutingEmailDomainRouteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteUnsupportedMediaType creates a DeleteRoutingEmailDomainRouteUnsupportedMediaType with default headers values
func NewDeleteRoutingEmailDomainRouteUnsupportedMediaType() *DeleteRoutingEmailDomainRouteUnsupportedMediaType {
	return &DeleteRoutingEmailDomainRouteUnsupportedMediaType{}
}

/*DeleteRoutingEmailDomainRouteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingEmailDomainRouteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteTooManyRequests creates a DeleteRoutingEmailDomainRouteTooManyRequests with default headers values
func NewDeleteRoutingEmailDomainRouteTooManyRequests() *DeleteRoutingEmailDomainRouteTooManyRequests {
	return &DeleteRoutingEmailDomainRouteTooManyRequests{}
}

/*DeleteRoutingEmailDomainRouteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingEmailDomainRouteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteInternalServerError creates a DeleteRoutingEmailDomainRouteInternalServerError with default headers values
func NewDeleteRoutingEmailDomainRouteInternalServerError() *DeleteRoutingEmailDomainRouteInternalServerError {
	return &DeleteRoutingEmailDomainRouteInternalServerError{}
}

/*DeleteRoutingEmailDomainRouteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingEmailDomainRouteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteServiceUnavailable creates a DeleteRoutingEmailDomainRouteServiceUnavailable with default headers values
func NewDeleteRoutingEmailDomainRouteServiceUnavailable() *DeleteRoutingEmailDomainRouteServiceUnavailable {
	return &DeleteRoutingEmailDomainRouteServiceUnavailable{}
}

/*DeleteRoutingEmailDomainRouteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingEmailDomainRouteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingEmailDomainRouteGatewayTimeout creates a DeleteRoutingEmailDomainRouteGatewayTimeout with default headers values
func NewDeleteRoutingEmailDomainRouteGatewayTimeout() *DeleteRoutingEmailDomainRouteGatewayTimeout {
	return &DeleteRoutingEmailDomainRouteGatewayTimeout{}
}

/*DeleteRoutingEmailDomainRouteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRoutingEmailDomainRouteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingEmailDomainRouteGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/email/domains/{domainName}/routes/{routeId}][%d] deleteRoutingEmailDomainRouteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingEmailDomainRouteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingEmailDomainRouteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
