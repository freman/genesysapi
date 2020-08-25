// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteLocationReader is a Reader for the DeleteLocation structure.
type DeleteLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLocationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteLocationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteLocationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteLocationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteLocationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteLocationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLocationNoContent creates a DeleteLocationNoContent with default headers values
func NewDeleteLocationNoContent() *DeleteLocationNoContent {
	return &DeleteLocationNoContent{}
}

/*DeleteLocationNoContent handles this case with default header values.

The location was deleted successfully
*/
type DeleteLocationNoContent struct {
}

func (o *DeleteLocationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationNoContent ", 204)
}

func (o *DeleteLocationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLocationBadRequest creates a DeleteLocationBadRequest with default headers values
func NewDeleteLocationBadRequest() *DeleteLocationBadRequest {
	return &DeleteLocationBadRequest{}
}

/*DeleteLocationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteLocationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLocationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationUnauthorized creates a DeleteLocationUnauthorized with default headers values
func NewDeleteLocationUnauthorized() *DeleteLocationUnauthorized {
	return &DeleteLocationUnauthorized{}
}

/*DeleteLocationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteLocationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLocationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationForbidden creates a DeleteLocationForbidden with default headers values
func NewDeleteLocationForbidden() *DeleteLocationForbidden {
	return &DeleteLocationForbidden{}
}

/*DeleteLocationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteLocationForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLocationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationNotFound creates a DeleteLocationNotFound with default headers values
func NewDeleteLocationNotFound() *DeleteLocationNotFound {
	return &DeleteLocationNotFound{}
}

/*DeleteLocationNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteLocationNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLocationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationRequestEntityTooLarge creates a DeleteLocationRequestEntityTooLarge with default headers values
func NewDeleteLocationRequestEntityTooLarge() *DeleteLocationRequestEntityTooLarge {
	return &DeleteLocationRequestEntityTooLarge{}
}

/*DeleteLocationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteLocationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteLocationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationUnsupportedMediaType creates a DeleteLocationUnsupportedMediaType with default headers values
func NewDeleteLocationUnsupportedMediaType() *DeleteLocationUnsupportedMediaType {
	return &DeleteLocationUnsupportedMediaType{}
}

/*DeleteLocationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteLocationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteLocationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationTooManyRequests creates a DeleteLocationTooManyRequests with default headers values
func NewDeleteLocationTooManyRequests() *DeleteLocationTooManyRequests {
	return &DeleteLocationTooManyRequests{}
}

/*DeleteLocationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteLocationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteLocationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationInternalServerError creates a DeleteLocationInternalServerError with default headers values
func NewDeleteLocationInternalServerError() *DeleteLocationInternalServerError {
	return &DeleteLocationInternalServerError{}
}

/*DeleteLocationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteLocationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLocationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationServiceUnavailable creates a DeleteLocationServiceUnavailable with default headers values
func NewDeleteLocationServiceUnavailable() *DeleteLocationServiceUnavailable {
	return &DeleteLocationServiceUnavailable{}
}

/*DeleteLocationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteLocationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLocationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocationGatewayTimeout creates a DeleteLocationGatewayTimeout with default headers values
func NewDeleteLocationGatewayTimeout() *DeleteLocationGatewayTimeout {
	return &DeleteLocationGatewayTimeout{}
}

/*DeleteLocationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteLocationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLocationGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/locations/{locationId}][%d] deleteLocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteLocationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLocationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}