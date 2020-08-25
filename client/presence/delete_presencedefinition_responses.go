// Code generated by go-swagger; DO NOT EDIT.

package presence

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeletePresencedefinitionReader is a Reader for the DeletePresencedefinition structure.
type DeletePresencedefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePresencedefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeletePresencedefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePresencedefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePresencedefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePresencedefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeletePresencedefinitionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeletePresencedefinitionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeletePresencedefinitionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePresencedefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeletePresencedefinitionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeletePresencedefinitionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePresencedefinitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePresencedefinitionBadRequest creates a DeletePresencedefinitionBadRequest with default headers values
func NewDeletePresencedefinitionBadRequest() *DeletePresencedefinitionBadRequest {
	return &DeletePresencedefinitionBadRequest{}
}

/*DeletePresencedefinitionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeletePresencedefinitionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePresencedefinitionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionUnauthorized creates a DeletePresencedefinitionUnauthorized with default headers values
func NewDeletePresencedefinitionUnauthorized() *DeletePresencedefinitionUnauthorized {
	return &DeletePresencedefinitionUnauthorized{}
}

/*DeletePresencedefinitionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeletePresencedefinitionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePresencedefinitionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionForbidden creates a DeletePresencedefinitionForbidden with default headers values
func NewDeletePresencedefinitionForbidden() *DeletePresencedefinitionForbidden {
	return &DeletePresencedefinitionForbidden{}
}

/*DeletePresencedefinitionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeletePresencedefinitionForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionForbidden  %+v", 403, o.Payload)
}

func (o *DeletePresencedefinitionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionNotFound creates a DeletePresencedefinitionNotFound with default headers values
func NewDeletePresencedefinitionNotFound() *DeletePresencedefinitionNotFound {
	return &DeletePresencedefinitionNotFound{}
}

/*DeletePresencedefinitionNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeletePresencedefinitionNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionNotFound  %+v", 404, o.Payload)
}

func (o *DeletePresencedefinitionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionRequestEntityTooLarge creates a DeletePresencedefinitionRequestEntityTooLarge with default headers values
func NewDeletePresencedefinitionRequestEntityTooLarge() *DeletePresencedefinitionRequestEntityTooLarge {
	return &DeletePresencedefinitionRequestEntityTooLarge{}
}

/*DeletePresencedefinitionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeletePresencedefinitionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeletePresencedefinitionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionUnsupportedMediaType creates a DeletePresencedefinitionUnsupportedMediaType with default headers values
func NewDeletePresencedefinitionUnsupportedMediaType() *DeletePresencedefinitionUnsupportedMediaType {
	return &DeletePresencedefinitionUnsupportedMediaType{}
}

/*DeletePresencedefinitionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeletePresencedefinitionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeletePresencedefinitionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionTooManyRequests creates a DeletePresencedefinitionTooManyRequests with default headers values
func NewDeletePresencedefinitionTooManyRequests() *DeletePresencedefinitionTooManyRequests {
	return &DeletePresencedefinitionTooManyRequests{}
}

/*DeletePresencedefinitionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeletePresencedefinitionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeletePresencedefinitionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionInternalServerError creates a DeletePresencedefinitionInternalServerError with default headers values
func NewDeletePresencedefinitionInternalServerError() *DeletePresencedefinitionInternalServerError {
	return &DeletePresencedefinitionInternalServerError{}
}

/*DeletePresencedefinitionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeletePresencedefinitionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePresencedefinitionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionServiceUnavailable creates a DeletePresencedefinitionServiceUnavailable with default headers values
func NewDeletePresencedefinitionServiceUnavailable() *DeletePresencedefinitionServiceUnavailable {
	return &DeletePresencedefinitionServiceUnavailable{}
}

/*DeletePresencedefinitionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeletePresencedefinitionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeletePresencedefinitionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionGatewayTimeout creates a DeletePresencedefinitionGatewayTimeout with default headers values
func NewDeletePresencedefinitionGatewayTimeout() *DeletePresencedefinitionGatewayTimeout {
	return &DeletePresencedefinitionGatewayTimeout{}
}

/*DeletePresencedefinitionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeletePresencedefinitionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeletePresencedefinitionGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinitionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeletePresencedefinitionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeletePresencedefinitionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePresencedefinitionDefault creates a DeletePresencedefinitionDefault with default headers values
func NewDeletePresencedefinitionDefault(code int) *DeletePresencedefinitionDefault {
	return &DeletePresencedefinitionDefault{
		_statusCode: code,
	}
}

/*DeletePresencedefinitionDefault handles this case with default header values.

successful operation
*/
type DeletePresencedefinitionDefault struct {
	_statusCode int
}

// Code gets the status code for the delete presencedefinition default response
func (o *DeletePresencedefinitionDefault) Code() int {
	return o._statusCode
}

func (o *DeletePresencedefinitionDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/presencedefinitions/{presenceId}][%d] deletePresencedefinition default ", o._statusCode)
}

func (o *DeletePresencedefinitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}