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

// PostOutboundContactlistClearReader is a Reader for the PostOutboundContactlistClear structure.
type PostOutboundContactlistClearReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundContactlistClearReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostOutboundContactlistClearNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundContactlistClearBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundContactlistClearUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundContactlistClearForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundContactlistClearNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundContactlistClearRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundContactlistClearUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundContactlistClearTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundContactlistClearInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundContactlistClearServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundContactlistClearGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundContactlistClearNoContent creates a PostOutboundContactlistClearNoContent with default headers values
func NewPostOutboundContactlistClearNoContent() *PostOutboundContactlistClearNoContent {
	return &PostOutboundContactlistClearNoContent{}
}

/*PostOutboundContactlistClearNoContent handles this case with default header values.

Contacts will be deleted.
*/
type PostOutboundContactlistClearNoContent struct {
}

func (o *PostOutboundContactlistClearNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearNoContent ", 204)
}

func (o *PostOutboundContactlistClearNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOutboundContactlistClearBadRequest creates a PostOutboundContactlistClearBadRequest with default headers values
func NewPostOutboundContactlistClearBadRequest() *PostOutboundContactlistClearBadRequest {
	return &PostOutboundContactlistClearBadRequest{}
}

/*PostOutboundContactlistClearBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundContactlistClearBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundContactlistClearBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearUnauthorized creates a PostOutboundContactlistClearUnauthorized with default headers values
func NewPostOutboundContactlistClearUnauthorized() *PostOutboundContactlistClearUnauthorized {
	return &PostOutboundContactlistClearUnauthorized{}
}

/*PostOutboundContactlistClearUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundContactlistClearUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundContactlistClearUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearForbidden creates a PostOutboundContactlistClearForbidden with default headers values
func NewPostOutboundContactlistClearForbidden() *PostOutboundContactlistClearForbidden {
	return &PostOutboundContactlistClearForbidden{}
}

/*PostOutboundContactlistClearForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundContactlistClearForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundContactlistClearForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearNotFound creates a PostOutboundContactlistClearNotFound with default headers values
func NewPostOutboundContactlistClearNotFound() *PostOutboundContactlistClearNotFound {
	return &PostOutboundContactlistClearNotFound{}
}

/*PostOutboundContactlistClearNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundContactlistClearNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundContactlistClearNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearRequestEntityTooLarge creates a PostOutboundContactlistClearRequestEntityTooLarge with default headers values
func NewPostOutboundContactlistClearRequestEntityTooLarge() *PostOutboundContactlistClearRequestEntityTooLarge {
	return &PostOutboundContactlistClearRequestEntityTooLarge{}
}

/*PostOutboundContactlistClearRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOutboundContactlistClearRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundContactlistClearRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearUnsupportedMediaType creates a PostOutboundContactlistClearUnsupportedMediaType with default headers values
func NewPostOutboundContactlistClearUnsupportedMediaType() *PostOutboundContactlistClearUnsupportedMediaType {
	return &PostOutboundContactlistClearUnsupportedMediaType{}
}

/*PostOutboundContactlistClearUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundContactlistClearUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundContactlistClearUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearTooManyRequests creates a PostOutboundContactlistClearTooManyRequests with default headers values
func NewPostOutboundContactlistClearTooManyRequests() *PostOutboundContactlistClearTooManyRequests {
	return &PostOutboundContactlistClearTooManyRequests{}
}

/*PostOutboundContactlistClearTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostOutboundContactlistClearTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundContactlistClearTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearInternalServerError creates a PostOutboundContactlistClearInternalServerError with default headers values
func NewPostOutboundContactlistClearInternalServerError() *PostOutboundContactlistClearInternalServerError {
	return &PostOutboundContactlistClearInternalServerError{}
}

/*PostOutboundContactlistClearInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundContactlistClearInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundContactlistClearInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearServiceUnavailable creates a PostOutboundContactlistClearServiceUnavailable with default headers values
func NewPostOutboundContactlistClearServiceUnavailable() *PostOutboundContactlistClearServiceUnavailable {
	return &PostOutboundContactlistClearServiceUnavailable{}
}

/*PostOutboundContactlistClearServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundContactlistClearServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundContactlistClearServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundContactlistClearGatewayTimeout creates a PostOutboundContactlistClearGatewayTimeout with default headers values
func NewPostOutboundContactlistClearGatewayTimeout() *PostOutboundContactlistClearGatewayTimeout {
	return &PostOutboundContactlistClearGatewayTimeout{}
}

/*PostOutboundContactlistClearGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundContactlistClearGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundContactlistClearGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/contactlists/{contactListId}/clear][%d] postOutboundContactlistClearGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundContactlistClearGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundContactlistClearGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
