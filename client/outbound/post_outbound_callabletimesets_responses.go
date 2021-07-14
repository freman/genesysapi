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

// PostOutboundCallabletimesetsReader is a Reader for the PostOutboundCallabletimesets structure.
type PostOutboundCallabletimesetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundCallabletimesetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundCallabletimesetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundCallabletimesetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundCallabletimesetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundCallabletimesetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundCallabletimesetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundCallabletimesetsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundCallabletimesetsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundCallabletimesetsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundCallabletimesetsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundCallabletimesetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundCallabletimesetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundCallabletimesetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundCallabletimesetsOK creates a PostOutboundCallabletimesetsOK with default headers values
func NewPostOutboundCallabletimesetsOK() *PostOutboundCallabletimesetsOK {
	return &PostOutboundCallabletimesetsOK{}
}

/*PostOutboundCallabletimesetsOK handles this case with default header values.

successful operation
*/
type PostOutboundCallabletimesetsOK struct {
	Payload *models.CallableTimeSet
}

func (o *PostOutboundCallabletimesetsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundCallabletimesetsOK) GetPayload() *models.CallableTimeSet {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallableTimeSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsBadRequest creates a PostOutboundCallabletimesetsBadRequest with default headers values
func NewPostOutboundCallabletimesetsBadRequest() *PostOutboundCallabletimesetsBadRequest {
	return &PostOutboundCallabletimesetsBadRequest{}
}

/*PostOutboundCallabletimesetsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundCallabletimesetsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundCallabletimesetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsUnauthorized creates a PostOutboundCallabletimesetsUnauthorized with default headers values
func NewPostOutboundCallabletimesetsUnauthorized() *PostOutboundCallabletimesetsUnauthorized {
	return &PostOutboundCallabletimesetsUnauthorized{}
}

/*PostOutboundCallabletimesetsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundCallabletimesetsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundCallabletimesetsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsForbidden creates a PostOutboundCallabletimesetsForbidden with default headers values
func NewPostOutboundCallabletimesetsForbidden() *PostOutboundCallabletimesetsForbidden {
	return &PostOutboundCallabletimesetsForbidden{}
}

/*PostOutboundCallabletimesetsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundCallabletimesetsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundCallabletimesetsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsNotFound creates a PostOutboundCallabletimesetsNotFound with default headers values
func NewPostOutboundCallabletimesetsNotFound() *PostOutboundCallabletimesetsNotFound {
	return &PostOutboundCallabletimesetsNotFound{}
}

/*PostOutboundCallabletimesetsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundCallabletimesetsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundCallabletimesetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsRequestTimeout creates a PostOutboundCallabletimesetsRequestTimeout with default headers values
func NewPostOutboundCallabletimesetsRequestTimeout() *PostOutboundCallabletimesetsRequestTimeout {
	return &PostOutboundCallabletimesetsRequestTimeout{}
}

/*PostOutboundCallabletimesetsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundCallabletimesetsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundCallabletimesetsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsRequestEntityTooLarge creates a PostOutboundCallabletimesetsRequestEntityTooLarge with default headers values
func NewPostOutboundCallabletimesetsRequestEntityTooLarge() *PostOutboundCallabletimesetsRequestEntityTooLarge {
	return &PostOutboundCallabletimesetsRequestEntityTooLarge{}
}

/*PostOutboundCallabletimesetsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOutboundCallabletimesetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundCallabletimesetsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsUnsupportedMediaType creates a PostOutboundCallabletimesetsUnsupportedMediaType with default headers values
func NewPostOutboundCallabletimesetsUnsupportedMediaType() *PostOutboundCallabletimesetsUnsupportedMediaType {
	return &PostOutboundCallabletimesetsUnsupportedMediaType{}
}

/*PostOutboundCallabletimesetsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundCallabletimesetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundCallabletimesetsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsTooManyRequests creates a PostOutboundCallabletimesetsTooManyRequests with default headers values
func NewPostOutboundCallabletimesetsTooManyRequests() *PostOutboundCallabletimesetsTooManyRequests {
	return &PostOutboundCallabletimesetsTooManyRequests{}
}

/*PostOutboundCallabletimesetsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundCallabletimesetsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundCallabletimesetsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsInternalServerError creates a PostOutboundCallabletimesetsInternalServerError with default headers values
func NewPostOutboundCallabletimesetsInternalServerError() *PostOutboundCallabletimesetsInternalServerError {
	return &PostOutboundCallabletimesetsInternalServerError{}
}

/*PostOutboundCallabletimesetsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundCallabletimesetsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundCallabletimesetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsServiceUnavailable creates a PostOutboundCallabletimesetsServiceUnavailable with default headers values
func NewPostOutboundCallabletimesetsServiceUnavailable() *PostOutboundCallabletimesetsServiceUnavailable {
	return &PostOutboundCallabletimesetsServiceUnavailable{}
}

/*PostOutboundCallabletimesetsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundCallabletimesetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundCallabletimesetsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundCallabletimesetsGatewayTimeout creates a PostOutboundCallabletimesetsGatewayTimeout with default headers values
func NewPostOutboundCallabletimesetsGatewayTimeout() *PostOutboundCallabletimesetsGatewayTimeout {
	return &PostOutboundCallabletimesetsGatewayTimeout{}
}

/*PostOutboundCallabletimesetsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundCallabletimesetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundCallabletimesetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/callabletimesets][%d] postOutboundCallabletimesetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundCallabletimesetsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundCallabletimesetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
