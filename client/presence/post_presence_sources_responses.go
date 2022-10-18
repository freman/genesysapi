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

// PostPresenceSourcesReader is a Reader for the PostPresenceSources structure.
type PostPresenceSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPresenceSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPresenceSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPresenceSourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPresenceSourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostPresenceSourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPresenceSourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostPresenceSourcesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPresenceSourcesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostPresenceSourcesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostPresenceSourcesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostPresenceSourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPresenceSourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostPresenceSourcesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostPresenceSourcesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostPresenceSourcesOK creates a PostPresenceSourcesOK with default headers values
func NewPostPresenceSourcesOK() *PostPresenceSourcesOK {
	return &PostPresenceSourcesOK{}
}

/*PostPresenceSourcesOK handles this case with default header values.

successful operation
*/
type PostPresenceSourcesOK struct {
	Payload *models.Source
}

func (o *PostPresenceSourcesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesOK  %+v", 200, o.Payload)
}

func (o *PostPresenceSourcesOK) GetPayload() *models.Source {
	return o.Payload
}

func (o *PostPresenceSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Source)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesBadRequest creates a PostPresenceSourcesBadRequest with default headers values
func NewPostPresenceSourcesBadRequest() *PostPresenceSourcesBadRequest {
	return &PostPresenceSourcesBadRequest{}
}

/*PostPresenceSourcesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostPresenceSourcesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesBadRequest  %+v", 400, o.Payload)
}

func (o *PostPresenceSourcesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesUnauthorized creates a PostPresenceSourcesUnauthorized with default headers values
func NewPostPresenceSourcesUnauthorized() *PostPresenceSourcesUnauthorized {
	return &PostPresenceSourcesUnauthorized{}
}

/*PostPresenceSourcesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostPresenceSourcesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPresenceSourcesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesForbidden creates a PostPresenceSourcesForbidden with default headers values
func NewPostPresenceSourcesForbidden() *PostPresenceSourcesForbidden {
	return &PostPresenceSourcesForbidden{}
}

/*PostPresenceSourcesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostPresenceSourcesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesForbidden  %+v", 403, o.Payload)
}

func (o *PostPresenceSourcesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesNotFound creates a PostPresenceSourcesNotFound with default headers values
func NewPostPresenceSourcesNotFound() *PostPresenceSourcesNotFound {
	return &PostPresenceSourcesNotFound{}
}

/*PostPresenceSourcesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostPresenceSourcesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesNotFound  %+v", 404, o.Payload)
}

func (o *PostPresenceSourcesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesRequestTimeout creates a PostPresenceSourcesRequestTimeout with default headers values
func NewPostPresenceSourcesRequestTimeout() *PostPresenceSourcesRequestTimeout {
	return &PostPresenceSourcesRequestTimeout{}
}

/*PostPresenceSourcesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostPresenceSourcesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostPresenceSourcesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesConflict creates a PostPresenceSourcesConflict with default headers values
func NewPostPresenceSourcesConflict() *PostPresenceSourcesConflict {
	return &PostPresenceSourcesConflict{}
}

/*PostPresenceSourcesConflict handles this case with default header values.

Conflict
*/
type PostPresenceSourcesConflict struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesConflict  %+v", 409, o.Payload)
}

func (o *PostPresenceSourcesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesRequestEntityTooLarge creates a PostPresenceSourcesRequestEntityTooLarge with default headers values
func NewPostPresenceSourcesRequestEntityTooLarge() *PostPresenceSourcesRequestEntityTooLarge {
	return &PostPresenceSourcesRequestEntityTooLarge{}
}

/*PostPresenceSourcesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostPresenceSourcesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostPresenceSourcesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesUnsupportedMediaType creates a PostPresenceSourcesUnsupportedMediaType with default headers values
func NewPostPresenceSourcesUnsupportedMediaType() *PostPresenceSourcesUnsupportedMediaType {
	return &PostPresenceSourcesUnsupportedMediaType{}
}

/*PostPresenceSourcesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostPresenceSourcesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostPresenceSourcesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesTooManyRequests creates a PostPresenceSourcesTooManyRequests with default headers values
func NewPostPresenceSourcesTooManyRequests() *PostPresenceSourcesTooManyRequests {
	return &PostPresenceSourcesTooManyRequests{}
}

/*PostPresenceSourcesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostPresenceSourcesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostPresenceSourcesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesInternalServerError creates a PostPresenceSourcesInternalServerError with default headers values
func NewPostPresenceSourcesInternalServerError() *PostPresenceSourcesInternalServerError {
	return &PostPresenceSourcesInternalServerError{}
}

/*PostPresenceSourcesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostPresenceSourcesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostPresenceSourcesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesServiceUnavailable creates a PostPresenceSourcesServiceUnavailable with default headers values
func NewPostPresenceSourcesServiceUnavailable() *PostPresenceSourcesServiceUnavailable {
	return &PostPresenceSourcesServiceUnavailable{}
}

/*PostPresenceSourcesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostPresenceSourcesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostPresenceSourcesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPresenceSourcesGatewayTimeout creates a PostPresenceSourcesGatewayTimeout with default headers values
func NewPostPresenceSourcesGatewayTimeout() *PostPresenceSourcesGatewayTimeout {
	return &PostPresenceSourcesGatewayTimeout{}
}

/*PostPresenceSourcesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostPresenceSourcesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostPresenceSourcesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/presence/sources][%d] postPresenceSourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostPresenceSourcesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostPresenceSourcesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
