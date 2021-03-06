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

// DeleteOutboundContactlistReader is a Reader for the DeleteOutboundContactlist structure.
type DeleteOutboundContactlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundContactlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundContactlistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundContactlistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundContactlistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundContactlistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundContactlistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundContactlistRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundContactlistUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundContactlistTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundContactlistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundContactlistServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundContactlistGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundContactlistOK creates a DeleteOutboundContactlistOK with default headers values
func NewDeleteOutboundContactlistOK() *DeleteOutboundContactlistOK {
	return &DeleteOutboundContactlistOK{}
}

/*DeleteOutboundContactlistOK handles this case with default header values.

Operation was successful.
*/
type DeleteOutboundContactlistOK struct {
}

func (o *DeleteOutboundContactlistOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistOK ", 200)
}

func (o *DeleteOutboundContactlistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundContactlistBadRequest creates a DeleteOutboundContactlistBadRequest with default headers values
func NewDeleteOutboundContactlistBadRequest() *DeleteOutboundContactlistBadRequest {
	return &DeleteOutboundContactlistBadRequest{}
}

/*DeleteOutboundContactlistBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundContactlistBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundContactlistBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistUnauthorized creates a DeleteOutboundContactlistUnauthorized with default headers values
func NewDeleteOutboundContactlistUnauthorized() *DeleteOutboundContactlistUnauthorized {
	return &DeleteOutboundContactlistUnauthorized{}
}

/*DeleteOutboundContactlistUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundContactlistUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundContactlistUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistForbidden creates a DeleteOutboundContactlistForbidden with default headers values
func NewDeleteOutboundContactlistForbidden() *DeleteOutboundContactlistForbidden {
	return &DeleteOutboundContactlistForbidden{}
}

/*DeleteOutboundContactlistForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundContactlistForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundContactlistForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistNotFound creates a DeleteOutboundContactlistNotFound with default headers values
func NewDeleteOutboundContactlistNotFound() *DeleteOutboundContactlistNotFound {
	return &DeleteOutboundContactlistNotFound{}
}

/*DeleteOutboundContactlistNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundContactlistNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundContactlistNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistRequestEntityTooLarge creates a DeleteOutboundContactlistRequestEntityTooLarge with default headers values
func NewDeleteOutboundContactlistRequestEntityTooLarge() *DeleteOutboundContactlistRequestEntityTooLarge {
	return &DeleteOutboundContactlistRequestEntityTooLarge{}
}

/*DeleteOutboundContactlistRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundContactlistRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundContactlistRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistUnsupportedMediaType creates a DeleteOutboundContactlistUnsupportedMediaType with default headers values
func NewDeleteOutboundContactlistUnsupportedMediaType() *DeleteOutboundContactlistUnsupportedMediaType {
	return &DeleteOutboundContactlistUnsupportedMediaType{}
}

/*DeleteOutboundContactlistUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundContactlistUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundContactlistUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistTooManyRequests creates a DeleteOutboundContactlistTooManyRequests with default headers values
func NewDeleteOutboundContactlistTooManyRequests() *DeleteOutboundContactlistTooManyRequests {
	return &DeleteOutboundContactlistTooManyRequests{}
}

/*DeleteOutboundContactlistTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOutboundContactlistTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundContactlistTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistInternalServerError creates a DeleteOutboundContactlistInternalServerError with default headers values
func NewDeleteOutboundContactlistInternalServerError() *DeleteOutboundContactlistInternalServerError {
	return &DeleteOutboundContactlistInternalServerError{}
}

/*DeleteOutboundContactlistInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundContactlistInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundContactlistInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistServiceUnavailable creates a DeleteOutboundContactlistServiceUnavailable with default headers values
func NewDeleteOutboundContactlistServiceUnavailable() *DeleteOutboundContactlistServiceUnavailable {
	return &DeleteOutboundContactlistServiceUnavailable{}
}

/*DeleteOutboundContactlistServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundContactlistServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundContactlistServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundContactlistGatewayTimeout creates a DeleteOutboundContactlistGatewayTimeout with default headers values
func NewDeleteOutboundContactlistGatewayTimeout() *DeleteOutboundContactlistGatewayTimeout {
	return &DeleteOutboundContactlistGatewayTimeout{}
}

/*DeleteOutboundContactlistGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundContactlistGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundContactlistGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/contactlists/{contactListId}][%d] deleteOutboundContactlistGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundContactlistGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundContactlistGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
