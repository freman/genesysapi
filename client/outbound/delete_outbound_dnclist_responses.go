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

// DeleteOutboundDnclistReader is a Reader for the DeleteOutboundDnclist structure.
type DeleteOutboundDnclistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundDnclistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundDnclistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundDnclistBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundDnclistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundDnclistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundDnclistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundDnclistRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundDnclistUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundDnclistTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundDnclistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundDnclistServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundDnclistGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundDnclistOK creates a DeleteOutboundDnclistOK with default headers values
func NewDeleteOutboundDnclistOK() *DeleteOutboundDnclistOK {
	return &DeleteOutboundDnclistOK{}
}

/*DeleteOutboundDnclistOK handles this case with default header values.

Operation was successful.
*/
type DeleteOutboundDnclistOK struct {
}

func (o *DeleteOutboundDnclistOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistOK ", 200)
}

func (o *DeleteOutboundDnclistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundDnclistBadRequest creates a DeleteOutboundDnclistBadRequest with default headers values
func NewDeleteOutboundDnclistBadRequest() *DeleteOutboundDnclistBadRequest {
	return &DeleteOutboundDnclistBadRequest{}
}

/*DeleteOutboundDnclistBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundDnclistBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundDnclistBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistUnauthorized creates a DeleteOutboundDnclistUnauthorized with default headers values
func NewDeleteOutboundDnclistUnauthorized() *DeleteOutboundDnclistUnauthorized {
	return &DeleteOutboundDnclistUnauthorized{}
}

/*DeleteOutboundDnclistUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundDnclistUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundDnclistUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistForbidden creates a DeleteOutboundDnclistForbidden with default headers values
func NewDeleteOutboundDnclistForbidden() *DeleteOutboundDnclistForbidden {
	return &DeleteOutboundDnclistForbidden{}
}

/*DeleteOutboundDnclistForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundDnclistForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundDnclistForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistNotFound creates a DeleteOutboundDnclistNotFound with default headers values
func NewDeleteOutboundDnclistNotFound() *DeleteOutboundDnclistNotFound {
	return &DeleteOutboundDnclistNotFound{}
}

/*DeleteOutboundDnclistNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundDnclistNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundDnclistNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistRequestEntityTooLarge creates a DeleteOutboundDnclistRequestEntityTooLarge with default headers values
func NewDeleteOutboundDnclistRequestEntityTooLarge() *DeleteOutboundDnclistRequestEntityTooLarge {
	return &DeleteOutboundDnclistRequestEntityTooLarge{}
}

/*DeleteOutboundDnclistRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundDnclistRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundDnclistRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistUnsupportedMediaType creates a DeleteOutboundDnclistUnsupportedMediaType with default headers values
func NewDeleteOutboundDnclistUnsupportedMediaType() *DeleteOutboundDnclistUnsupportedMediaType {
	return &DeleteOutboundDnclistUnsupportedMediaType{}
}

/*DeleteOutboundDnclistUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundDnclistUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundDnclistUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistTooManyRequests creates a DeleteOutboundDnclistTooManyRequests with default headers values
func NewDeleteOutboundDnclistTooManyRequests() *DeleteOutboundDnclistTooManyRequests {
	return &DeleteOutboundDnclistTooManyRequests{}
}

/*DeleteOutboundDnclistTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOutboundDnclistTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundDnclistTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistInternalServerError creates a DeleteOutboundDnclistInternalServerError with default headers values
func NewDeleteOutboundDnclistInternalServerError() *DeleteOutboundDnclistInternalServerError {
	return &DeleteOutboundDnclistInternalServerError{}
}

/*DeleteOutboundDnclistInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundDnclistInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundDnclistInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistServiceUnavailable creates a DeleteOutboundDnclistServiceUnavailable with default headers values
func NewDeleteOutboundDnclistServiceUnavailable() *DeleteOutboundDnclistServiceUnavailable {
	return &DeleteOutboundDnclistServiceUnavailable{}
}

/*DeleteOutboundDnclistServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundDnclistServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundDnclistServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundDnclistGatewayTimeout creates a DeleteOutboundDnclistGatewayTimeout with default headers values
func NewDeleteOutboundDnclistGatewayTimeout() *DeleteOutboundDnclistGatewayTimeout {
	return &DeleteOutboundDnclistGatewayTimeout{}
}

/*DeleteOutboundDnclistGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundDnclistGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundDnclistGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/dnclists/{dncListId}][%d] deleteOutboundDnclistGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundDnclistGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundDnclistGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
