// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteContentmanagementShareReader is a Reader for the DeleteContentmanagementShare structure.
type DeleteContentmanagementShareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteContentmanagementShareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteContentmanagementShareBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteContentmanagementShareUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteContentmanagementShareForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteContentmanagementShareNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteContentmanagementShareRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteContentmanagementShareUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteContentmanagementShareTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteContentmanagementShareInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteContentmanagementShareServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteContentmanagementShareGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteContentmanagementShareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteContentmanagementShareBadRequest creates a DeleteContentmanagementShareBadRequest with default headers values
func NewDeleteContentmanagementShareBadRequest() *DeleteContentmanagementShareBadRequest {
	return &DeleteContentmanagementShareBadRequest{}
}

/*DeleteContentmanagementShareBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteContentmanagementShareBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteContentmanagementShareBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareUnauthorized creates a DeleteContentmanagementShareUnauthorized with default headers values
func NewDeleteContentmanagementShareUnauthorized() *DeleteContentmanagementShareUnauthorized {
	return &DeleteContentmanagementShareUnauthorized{}
}

/*DeleteContentmanagementShareUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteContentmanagementShareUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteContentmanagementShareUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareForbidden creates a DeleteContentmanagementShareForbidden with default headers values
func NewDeleteContentmanagementShareForbidden() *DeleteContentmanagementShareForbidden {
	return &DeleteContentmanagementShareForbidden{}
}

/*DeleteContentmanagementShareForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteContentmanagementShareForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareForbidden  %+v", 403, o.Payload)
}

func (o *DeleteContentmanagementShareForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareNotFound creates a DeleteContentmanagementShareNotFound with default headers values
func NewDeleteContentmanagementShareNotFound() *DeleteContentmanagementShareNotFound {
	return &DeleteContentmanagementShareNotFound{}
}

/*DeleteContentmanagementShareNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteContentmanagementShareNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareNotFound  %+v", 404, o.Payload)
}

func (o *DeleteContentmanagementShareNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareRequestEntityTooLarge creates a DeleteContentmanagementShareRequestEntityTooLarge with default headers values
func NewDeleteContentmanagementShareRequestEntityTooLarge() *DeleteContentmanagementShareRequestEntityTooLarge {
	return &DeleteContentmanagementShareRequestEntityTooLarge{}
}

/*DeleteContentmanagementShareRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteContentmanagementShareRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteContentmanagementShareRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareUnsupportedMediaType creates a DeleteContentmanagementShareUnsupportedMediaType with default headers values
func NewDeleteContentmanagementShareUnsupportedMediaType() *DeleteContentmanagementShareUnsupportedMediaType {
	return &DeleteContentmanagementShareUnsupportedMediaType{}
}

/*DeleteContentmanagementShareUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteContentmanagementShareUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteContentmanagementShareUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareTooManyRequests creates a DeleteContentmanagementShareTooManyRequests with default headers values
func NewDeleteContentmanagementShareTooManyRequests() *DeleteContentmanagementShareTooManyRequests {
	return &DeleteContentmanagementShareTooManyRequests{}
}

/*DeleteContentmanagementShareTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteContentmanagementShareTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteContentmanagementShareTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareInternalServerError creates a DeleteContentmanagementShareInternalServerError with default headers values
func NewDeleteContentmanagementShareInternalServerError() *DeleteContentmanagementShareInternalServerError {
	return &DeleteContentmanagementShareInternalServerError{}
}

/*DeleteContentmanagementShareInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteContentmanagementShareInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteContentmanagementShareInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareServiceUnavailable creates a DeleteContentmanagementShareServiceUnavailable with default headers values
func NewDeleteContentmanagementShareServiceUnavailable() *DeleteContentmanagementShareServiceUnavailable {
	return &DeleteContentmanagementShareServiceUnavailable{}
}

/*DeleteContentmanagementShareServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteContentmanagementShareServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteContentmanagementShareServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareGatewayTimeout creates a DeleteContentmanagementShareGatewayTimeout with default headers values
func NewDeleteContentmanagementShareGatewayTimeout() *DeleteContentmanagementShareGatewayTimeout {
	return &DeleteContentmanagementShareGatewayTimeout{}
}

/*DeleteContentmanagementShareGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteContentmanagementShareGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteContentmanagementShareGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShareGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteContentmanagementShareGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteContentmanagementShareGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteContentmanagementShareDefault creates a DeleteContentmanagementShareDefault with default headers values
func NewDeleteContentmanagementShareDefault(code int) *DeleteContentmanagementShareDefault {
	return &DeleteContentmanagementShareDefault{
		_statusCode: code,
	}
}

/*DeleteContentmanagementShareDefault handles this case with default header values.

successful operation
*/
type DeleteContentmanagementShareDefault struct {
	_statusCode int
}

// Code gets the status code for the delete contentmanagement share default response
func (o *DeleteContentmanagementShareDefault) Code() int {
	return o._statusCode
}

func (o *DeleteContentmanagementShareDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/contentmanagement/shares/{shareId}][%d] deleteContentmanagementShare default ", o._statusCode)
}

func (o *DeleteContentmanagementShareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
