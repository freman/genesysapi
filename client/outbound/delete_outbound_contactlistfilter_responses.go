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

// DeleteOutboundContactlistfilterReader is a Reader for the DeleteOutboundContactlistfilter structure.
type DeleteOutboundContactlistfilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundContactlistfilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOutboundContactlistfilterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundContactlistfilterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundContactlistfilterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundContactlistfilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundContactlistfilterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundContactlistfilterRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundContactlistfilterUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundContactlistfilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundContactlistfilterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundContactlistfilterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundContactlistfilterGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundContactlistfilterNoContent creates a DeleteOutboundContactlistfilterNoContent with default headers values
func NewDeleteOutboundContactlistfilterNoContent() *DeleteOutboundContactlistfilterNoContent {
	return &DeleteOutboundContactlistfilterNoContent{}
}

/*DeleteOutboundContactlistfilterNoContent handles this case with default header values.

Contact list filter deleted
*/
type DeleteOutboundContactlistfilterNoContent struct {
}

func (o *DeleteOutboundContactlistfilterNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterNoContent ", 204)
}

func (o *DeleteOutboundContactlistfilterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundContactlistfilterBadRequest creates a DeleteOutboundContactlistfilterBadRequest with default headers values
func NewDeleteOutboundContactlistfilterBadRequest() *DeleteOutboundContactlistfilterBadRequest {
	return &DeleteOutboundContactlistfilterBadRequest{}
}

/*DeleteOutboundContactlistfilterBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundContactlistfilterBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundContactlistfilterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterUnauthorized creates a DeleteOutboundContactlistfilterUnauthorized with default headers values
func NewDeleteOutboundContactlistfilterUnauthorized() *DeleteOutboundContactlistfilterUnauthorized {
	return &DeleteOutboundContactlistfilterUnauthorized{}
}

/*DeleteOutboundContactlistfilterUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundContactlistfilterUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundContactlistfilterUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterForbidden creates a DeleteOutboundContactlistfilterForbidden with default headers values
func NewDeleteOutboundContactlistfilterForbidden() *DeleteOutboundContactlistfilterForbidden {
	return &DeleteOutboundContactlistfilterForbidden{}
}

/*DeleteOutboundContactlistfilterForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundContactlistfilterForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundContactlistfilterForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterNotFound creates a DeleteOutboundContactlistfilterNotFound with default headers values
func NewDeleteOutboundContactlistfilterNotFound() *DeleteOutboundContactlistfilterNotFound {
	return &DeleteOutboundContactlistfilterNotFound{}
}

/*DeleteOutboundContactlistfilterNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundContactlistfilterNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundContactlistfilterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterRequestEntityTooLarge creates a DeleteOutboundContactlistfilterRequestEntityTooLarge with default headers values
func NewDeleteOutboundContactlistfilterRequestEntityTooLarge() *DeleteOutboundContactlistfilterRequestEntityTooLarge {
	return &DeleteOutboundContactlistfilterRequestEntityTooLarge{}
}

/*DeleteOutboundContactlistfilterRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundContactlistfilterRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundContactlistfilterRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterUnsupportedMediaType creates a DeleteOutboundContactlistfilterUnsupportedMediaType with default headers values
func NewDeleteOutboundContactlistfilterUnsupportedMediaType() *DeleteOutboundContactlistfilterUnsupportedMediaType {
	return &DeleteOutboundContactlistfilterUnsupportedMediaType{}
}

/*DeleteOutboundContactlistfilterUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundContactlistfilterUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundContactlistfilterUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterTooManyRequests creates a DeleteOutboundContactlistfilterTooManyRequests with default headers values
func NewDeleteOutboundContactlistfilterTooManyRequests() *DeleteOutboundContactlistfilterTooManyRequests {
	return &DeleteOutboundContactlistfilterTooManyRequests{}
}

/*DeleteOutboundContactlistfilterTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOutboundContactlistfilterTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundContactlistfilterTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterInternalServerError creates a DeleteOutboundContactlistfilterInternalServerError with default headers values
func NewDeleteOutboundContactlistfilterInternalServerError() *DeleteOutboundContactlistfilterInternalServerError {
	return &DeleteOutboundContactlistfilterInternalServerError{}
}

/*DeleteOutboundContactlistfilterInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundContactlistfilterInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundContactlistfilterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterServiceUnavailable creates a DeleteOutboundContactlistfilterServiceUnavailable with default headers values
func NewDeleteOutboundContactlistfilterServiceUnavailable() *DeleteOutboundContactlistfilterServiceUnavailable {
	return &DeleteOutboundContactlistfilterServiceUnavailable{}
}

/*DeleteOutboundContactlistfilterServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundContactlistfilterServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundContactlistfilterServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistfilterGatewayTimeout creates a DeleteOutboundContactlistfilterGatewayTimeout with default headers values
func NewDeleteOutboundContactlistfilterGatewayTimeout() *DeleteOutboundContactlistfilterGatewayTimeout {
	return &DeleteOutboundContactlistfilterGatewayTimeout{}
}

/*DeleteOutboundContactlistfilterGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundContactlistfilterGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistfilterGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlistfilters/{contactListFilterId}][%d] deleteOutboundContactlistfilterGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundContactlistfilterGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistfilterGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
