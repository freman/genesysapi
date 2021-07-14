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

// DeleteOutboundContactlistContactsReader is a Reader for the DeleteOutboundContactlistContacts structure.
type DeleteOutboundContactlistContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundContactlistContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundContactlistContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundContactlistContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundContactlistContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundContactlistContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundContactlistContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOutboundContactlistContactsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundContactlistContactsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundContactlistContactsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundContactlistContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundContactlistContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundContactlistContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundContactlistContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundContactlistContactsOK creates a DeleteOutboundContactlistContactsOK with default headers values
func NewDeleteOutboundContactlistContactsOK() *DeleteOutboundContactlistContactsOK {
	return &DeleteOutboundContactlistContactsOK{}
}

/*DeleteOutboundContactlistContactsOK handles this case with default header values.

Contacts Deleted.
*/
type DeleteOutboundContactlistContactsOK struct {
}

func (o *DeleteOutboundContactlistContactsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsOK ", 200)
}

func (o *DeleteOutboundContactlistContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundContactlistContactsBadRequest creates a DeleteOutboundContactlistContactsBadRequest with default headers values
func NewDeleteOutboundContactlistContactsBadRequest() *DeleteOutboundContactlistContactsBadRequest {
	return &DeleteOutboundContactlistContactsBadRequest{}
}

/*DeleteOutboundContactlistContactsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundContactlistContactsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundContactlistContactsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsUnauthorized creates a DeleteOutboundContactlistContactsUnauthorized with default headers values
func NewDeleteOutboundContactlistContactsUnauthorized() *DeleteOutboundContactlistContactsUnauthorized {
	return &DeleteOutboundContactlistContactsUnauthorized{}
}

/*DeleteOutboundContactlistContactsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundContactlistContactsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundContactlistContactsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsForbidden creates a DeleteOutboundContactlistContactsForbidden with default headers values
func NewDeleteOutboundContactlistContactsForbidden() *DeleteOutboundContactlistContactsForbidden {
	return &DeleteOutboundContactlistContactsForbidden{}
}

/*DeleteOutboundContactlistContactsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundContactlistContactsForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundContactlistContactsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsNotFound creates a DeleteOutboundContactlistContactsNotFound with default headers values
func NewDeleteOutboundContactlistContactsNotFound() *DeleteOutboundContactlistContactsNotFound {
	return &DeleteOutboundContactlistContactsNotFound{}
}

/*DeleteOutboundContactlistContactsNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundContactlistContactsNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundContactlistContactsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsRequestTimeout creates a DeleteOutboundContactlistContactsRequestTimeout with default headers values
func NewDeleteOutboundContactlistContactsRequestTimeout() *DeleteOutboundContactlistContactsRequestTimeout {
	return &DeleteOutboundContactlistContactsRequestTimeout{}
}

/*DeleteOutboundContactlistContactsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundContactlistContactsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundContactlistContactsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsRequestEntityTooLarge creates a DeleteOutboundContactlistContactsRequestEntityTooLarge with default headers values
func NewDeleteOutboundContactlistContactsRequestEntityTooLarge() *DeleteOutboundContactlistContactsRequestEntityTooLarge {
	return &DeleteOutboundContactlistContactsRequestEntityTooLarge{}
}

/*DeleteOutboundContactlistContactsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundContactlistContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundContactlistContactsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsUnsupportedMediaType creates a DeleteOutboundContactlistContactsUnsupportedMediaType with default headers values
func NewDeleteOutboundContactlistContactsUnsupportedMediaType() *DeleteOutboundContactlistContactsUnsupportedMediaType {
	return &DeleteOutboundContactlistContactsUnsupportedMediaType{}
}

/*DeleteOutboundContactlistContactsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundContactlistContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundContactlistContactsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsTooManyRequests creates a DeleteOutboundContactlistContactsTooManyRequests with default headers values
func NewDeleteOutboundContactlistContactsTooManyRequests() *DeleteOutboundContactlistContactsTooManyRequests {
	return &DeleteOutboundContactlistContactsTooManyRequests{}
}

/*DeleteOutboundContactlistContactsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundContactlistContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundContactlistContactsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsInternalServerError creates a DeleteOutboundContactlistContactsInternalServerError with default headers values
func NewDeleteOutboundContactlistContactsInternalServerError() *DeleteOutboundContactlistContactsInternalServerError {
	return &DeleteOutboundContactlistContactsInternalServerError{}
}

/*DeleteOutboundContactlistContactsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundContactlistContactsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundContactlistContactsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsServiceUnavailable creates a DeleteOutboundContactlistContactsServiceUnavailable with default headers values
func NewDeleteOutboundContactlistContactsServiceUnavailable() *DeleteOutboundContactlistContactsServiceUnavailable {
	return &DeleteOutboundContactlistContactsServiceUnavailable{}
}

/*DeleteOutboundContactlistContactsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundContactlistContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundContactlistContactsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistContactsGatewayTimeout creates a DeleteOutboundContactlistContactsGatewayTimeout with default headers values
func NewDeleteOutboundContactlistContactsGatewayTimeout() *DeleteOutboundContactlistContactsGatewayTimeout {
	return &DeleteOutboundContactlistContactsGatewayTimeout{}
}

/*DeleteOutboundContactlistContactsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundContactlistContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}/contacts][%d] deleteOutboundContactlistContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundContactlistContactsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
