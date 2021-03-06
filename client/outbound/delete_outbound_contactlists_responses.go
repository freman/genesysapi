// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteOutboundContactlistsReader is a Reader for the DeleteOutboundContactlists structure.
type DeleteOutboundContactlistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundContactlistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOutboundContactlistsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundContactlistsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundContactlistsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundContactlistsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundContactlistsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteOutboundContactlistsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundContactlistsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundContactlistsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundContactlistsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundContactlistsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundContactlistsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundContactlistsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundContactlistsNoContent creates a DeleteOutboundContactlistsNoContent with default headers values
func NewDeleteOutboundContactlistsNoContent() *DeleteOutboundContactlistsNoContent {
	return &DeleteOutboundContactlistsNoContent{}
}

/*DeleteOutboundContactlistsNoContent handles this case with default header values.

Contact lists accepted for delete.
*/
type DeleteOutboundContactlistsNoContent struct {
}

func (o *DeleteOutboundContactlistsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsNoContent ", 204)
}

func (o *DeleteOutboundContactlistsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundContactlistsBadRequest creates a DeleteOutboundContactlistsBadRequest with default headers values
func NewDeleteOutboundContactlistsBadRequest() *DeleteOutboundContactlistsBadRequest {
	return &DeleteOutboundContactlistsBadRequest{}
}

/*DeleteOutboundContactlistsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundContactlistsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundContactlistsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsUnauthorized creates a DeleteOutboundContactlistsUnauthorized with default headers values
func NewDeleteOutboundContactlistsUnauthorized() *DeleteOutboundContactlistsUnauthorized {
	return &DeleteOutboundContactlistsUnauthorized{}
}

/*DeleteOutboundContactlistsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundContactlistsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundContactlistsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsForbidden creates a DeleteOutboundContactlistsForbidden with default headers values
func NewDeleteOutboundContactlistsForbidden() *DeleteOutboundContactlistsForbidden {
	return &DeleteOutboundContactlistsForbidden{}
}

/*DeleteOutboundContactlistsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundContactlistsForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundContactlistsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsNotFound creates a DeleteOutboundContactlistsNotFound with default headers values
func NewDeleteOutboundContactlistsNotFound() *DeleteOutboundContactlistsNotFound {
	return &DeleteOutboundContactlistsNotFound{}
}

/*DeleteOutboundContactlistsNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundContactlistsNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundContactlistsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsConflict creates a DeleteOutboundContactlistsConflict with default headers values
func NewDeleteOutboundContactlistsConflict() *DeleteOutboundContactlistsConflict {
	return &DeleteOutboundContactlistsConflict{}
}

/*DeleteOutboundContactlistsConflict handles this case with default header values.

Conflict
*/
type DeleteOutboundContactlistsConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsConflict  %+v", 409, o.Payload)
}

func (o *DeleteOutboundContactlistsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsRequestEntityTooLarge creates a DeleteOutboundContactlistsRequestEntityTooLarge with default headers values
func NewDeleteOutboundContactlistsRequestEntityTooLarge() *DeleteOutboundContactlistsRequestEntityTooLarge {
	return &DeleteOutboundContactlistsRequestEntityTooLarge{}
}

/*DeleteOutboundContactlistsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundContactlistsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundContactlistsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsUnsupportedMediaType creates a DeleteOutboundContactlistsUnsupportedMediaType with default headers values
func NewDeleteOutboundContactlistsUnsupportedMediaType() *DeleteOutboundContactlistsUnsupportedMediaType {
	return &DeleteOutboundContactlistsUnsupportedMediaType{}
}

/*DeleteOutboundContactlistsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundContactlistsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundContactlistsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsTooManyRequests creates a DeleteOutboundContactlistsTooManyRequests with default headers values
func NewDeleteOutboundContactlistsTooManyRequests() *DeleteOutboundContactlistsTooManyRequests {
	return &DeleteOutboundContactlistsTooManyRequests{}
}

/*DeleteOutboundContactlistsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOutboundContactlistsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundContactlistsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsInternalServerError creates a DeleteOutboundContactlistsInternalServerError with default headers values
func NewDeleteOutboundContactlistsInternalServerError() *DeleteOutboundContactlistsInternalServerError {
	return &DeleteOutboundContactlistsInternalServerError{}
}

/*DeleteOutboundContactlistsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundContactlistsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundContactlistsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsServiceUnavailable creates a DeleteOutboundContactlistsServiceUnavailable with default headers values
func NewDeleteOutboundContactlistsServiceUnavailable() *DeleteOutboundContactlistsServiceUnavailable {
	return &DeleteOutboundContactlistsServiceUnavailable{}
}

/*DeleteOutboundContactlistsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundContactlistsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundContactlistsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistsGatewayTimeout creates a DeleteOutboundContactlistsGatewayTimeout with default headers values
func NewDeleteOutboundContactlistsGatewayTimeout() *DeleteOutboundContactlistsGatewayTimeout {
	return &DeleteOutboundContactlistsGatewayTimeout{}
}

/*DeleteOutboundContactlistsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundContactlistsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistsGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists][%d] deleteOutboundContactlistsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundContactlistsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
