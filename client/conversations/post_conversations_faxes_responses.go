// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostConversationsFaxesReader is a Reader for the PostConversationsFaxes structure.
type PostConversationsFaxesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsFaxesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsFaxesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsFaxesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsFaxesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsFaxesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsFaxesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsFaxesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsFaxesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsFaxesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsFaxesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsFaxesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsFaxesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsFaxesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsFaxesOK creates a PostConversationsFaxesOK with default headers values
func NewPostConversationsFaxesOK() *PostConversationsFaxesOK {
	return &PostConversationsFaxesOK{}
}

/*PostConversationsFaxesOK handles this case with default header values.

successful operation
*/
type PostConversationsFaxesOK struct {
	Payload *models.FaxSendResponse
}

func (o *PostConversationsFaxesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesOK  %+v", 200, o.Payload)
}

func (o *PostConversationsFaxesOK) GetPayload() *models.FaxSendResponse {
	return o.Payload
}

func (o *PostConversationsFaxesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FaxSendResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesBadRequest creates a PostConversationsFaxesBadRequest with default headers values
func NewPostConversationsFaxesBadRequest() *PostConversationsFaxesBadRequest {
	return &PostConversationsFaxesBadRequest{}
}

/*PostConversationsFaxesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsFaxesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsFaxesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesUnauthorized creates a PostConversationsFaxesUnauthorized with default headers values
func NewPostConversationsFaxesUnauthorized() *PostConversationsFaxesUnauthorized {
	return &PostConversationsFaxesUnauthorized{}
}

/*PostConversationsFaxesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsFaxesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsFaxesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesForbidden creates a PostConversationsFaxesForbidden with default headers values
func NewPostConversationsFaxesForbidden() *PostConversationsFaxesForbidden {
	return &PostConversationsFaxesForbidden{}
}

/*PostConversationsFaxesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsFaxesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsFaxesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesNotFound creates a PostConversationsFaxesNotFound with default headers values
func NewPostConversationsFaxesNotFound() *PostConversationsFaxesNotFound {
	return &PostConversationsFaxesNotFound{}
}

/*PostConversationsFaxesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsFaxesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsFaxesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesRequestTimeout creates a PostConversationsFaxesRequestTimeout with default headers values
func NewPostConversationsFaxesRequestTimeout() *PostConversationsFaxesRequestTimeout {
	return &PostConversationsFaxesRequestTimeout{}
}

/*PostConversationsFaxesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsFaxesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsFaxesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesRequestEntityTooLarge creates a PostConversationsFaxesRequestEntityTooLarge with default headers values
func NewPostConversationsFaxesRequestEntityTooLarge() *PostConversationsFaxesRequestEntityTooLarge {
	return &PostConversationsFaxesRequestEntityTooLarge{}
}

/*PostConversationsFaxesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsFaxesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsFaxesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesUnsupportedMediaType creates a PostConversationsFaxesUnsupportedMediaType with default headers values
func NewPostConversationsFaxesUnsupportedMediaType() *PostConversationsFaxesUnsupportedMediaType {
	return &PostConversationsFaxesUnsupportedMediaType{}
}

/*PostConversationsFaxesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsFaxesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsFaxesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesTooManyRequests creates a PostConversationsFaxesTooManyRequests with default headers values
func NewPostConversationsFaxesTooManyRequests() *PostConversationsFaxesTooManyRequests {
	return &PostConversationsFaxesTooManyRequests{}
}

/*PostConversationsFaxesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsFaxesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsFaxesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesInternalServerError creates a PostConversationsFaxesInternalServerError with default headers values
func NewPostConversationsFaxesInternalServerError() *PostConversationsFaxesInternalServerError {
	return &PostConversationsFaxesInternalServerError{}
}

/*PostConversationsFaxesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsFaxesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsFaxesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesServiceUnavailable creates a PostConversationsFaxesServiceUnavailable with default headers values
func NewPostConversationsFaxesServiceUnavailable() *PostConversationsFaxesServiceUnavailable {
	return &PostConversationsFaxesServiceUnavailable{}
}

/*PostConversationsFaxesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsFaxesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsFaxesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsFaxesGatewayTimeout creates a PostConversationsFaxesGatewayTimeout with default headers values
func NewPostConversationsFaxesGatewayTimeout() *PostConversationsFaxesGatewayTimeout {
	return &PostConversationsFaxesGatewayTimeout{}
}

/*PostConversationsFaxesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsFaxesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsFaxesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/faxes][%d] postConversationsFaxesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsFaxesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsFaxesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
