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

// PostOutboundContactlistsReader is a Reader for the PostOutboundContactlists structure.
type PostOutboundContactlistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundContactlistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundContactlistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundContactlistsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundContactlistsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundContactlistsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundContactlistsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundContactlistsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundContactlistsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundContactlistsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundContactlistsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundContactlistsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundContactlistsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundContactlistsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundContactlistsOK creates a PostOutboundContactlistsOK with default headers values
func NewPostOutboundContactlistsOK() *PostOutboundContactlistsOK {
	return &PostOutboundContactlistsOK{}
}

/*PostOutboundContactlistsOK handles this case with default header values.

successful operation
*/
type PostOutboundContactlistsOK struct {
	Payload *models.ContactList
}

func (o *PostOutboundContactlistsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundContactlistsOK) GetPayload() *models.ContactList {
	return o.Payload
}

func (o *PostOutboundContactlistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsBadRequest creates a PostOutboundContactlistsBadRequest with default headers values
func NewPostOutboundContactlistsBadRequest() *PostOutboundContactlistsBadRequest {
	return &PostOutboundContactlistsBadRequest{}
}

/*PostOutboundContactlistsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundContactlistsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundContactlistsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsUnauthorized creates a PostOutboundContactlistsUnauthorized with default headers values
func NewPostOutboundContactlistsUnauthorized() *PostOutboundContactlistsUnauthorized {
	return &PostOutboundContactlistsUnauthorized{}
}

/*PostOutboundContactlistsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundContactlistsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundContactlistsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsForbidden creates a PostOutboundContactlistsForbidden with default headers values
func NewPostOutboundContactlistsForbidden() *PostOutboundContactlistsForbidden {
	return &PostOutboundContactlistsForbidden{}
}

/*PostOutboundContactlistsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundContactlistsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundContactlistsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsNotFound creates a PostOutboundContactlistsNotFound with default headers values
func NewPostOutboundContactlistsNotFound() *PostOutboundContactlistsNotFound {
	return &PostOutboundContactlistsNotFound{}
}

/*PostOutboundContactlistsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundContactlistsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundContactlistsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsRequestTimeout creates a PostOutboundContactlistsRequestTimeout with default headers values
func NewPostOutboundContactlistsRequestTimeout() *PostOutboundContactlistsRequestTimeout {
	return &PostOutboundContactlistsRequestTimeout{}
}

/*PostOutboundContactlistsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundContactlistsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundContactlistsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsRequestEntityTooLarge creates a PostOutboundContactlistsRequestEntityTooLarge with default headers values
func NewPostOutboundContactlistsRequestEntityTooLarge() *PostOutboundContactlistsRequestEntityTooLarge {
	return &PostOutboundContactlistsRequestEntityTooLarge{}
}

/*PostOutboundContactlistsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOutboundContactlistsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundContactlistsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsUnsupportedMediaType creates a PostOutboundContactlistsUnsupportedMediaType with default headers values
func NewPostOutboundContactlistsUnsupportedMediaType() *PostOutboundContactlistsUnsupportedMediaType {
	return &PostOutboundContactlistsUnsupportedMediaType{}
}

/*PostOutboundContactlistsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundContactlistsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundContactlistsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsTooManyRequests creates a PostOutboundContactlistsTooManyRequests with default headers values
func NewPostOutboundContactlistsTooManyRequests() *PostOutboundContactlistsTooManyRequests {
	return &PostOutboundContactlistsTooManyRequests{}
}

/*PostOutboundContactlistsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundContactlistsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundContactlistsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsInternalServerError creates a PostOutboundContactlistsInternalServerError with default headers values
func NewPostOutboundContactlistsInternalServerError() *PostOutboundContactlistsInternalServerError {
	return &PostOutboundContactlistsInternalServerError{}
}

/*PostOutboundContactlistsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundContactlistsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundContactlistsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsServiceUnavailable creates a PostOutboundContactlistsServiceUnavailable with default headers values
func NewPostOutboundContactlistsServiceUnavailable() *PostOutboundContactlistsServiceUnavailable {
	return &PostOutboundContactlistsServiceUnavailable{}
}

/*PostOutboundContactlistsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundContactlistsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundContactlistsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistsGatewayTimeout creates a PostOutboundContactlistsGatewayTimeout with default headers values
func NewPostOutboundContactlistsGatewayTimeout() *PostOutboundContactlistsGatewayTimeout {
	return &PostOutboundContactlistsGatewayTimeout{}
}

/*PostOutboundContactlistsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundContactlistsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists][%d] postOutboundContactlistsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundContactlistsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
