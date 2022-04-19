// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostIntegrationsActionExecuteReader is a Reader for the PostIntegrationsActionExecute structure.
type PostIntegrationsActionExecuteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsActionExecuteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsActionExecuteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsActionExecuteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsActionExecuteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsActionExecuteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsActionExecuteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPostIntegrationsActionExecuteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostIntegrationsActionExecuteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsActionExecuteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsActionExecuteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsActionExecuteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsActionExecuteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsActionExecuteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsActionExecuteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsActionExecuteOK creates a PostIntegrationsActionExecuteOK with default headers values
func NewPostIntegrationsActionExecuteOK() *PostIntegrationsActionExecuteOK {
	return &PostIntegrationsActionExecuteOK{}
}

/*PostIntegrationsActionExecuteOK handles this case with default header values.

successful operation
*/
type PostIntegrationsActionExecuteOK struct {
	Payload interface{}
}

func (o *PostIntegrationsActionExecuteOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionExecuteOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteBadRequest creates a PostIntegrationsActionExecuteBadRequest with default headers values
func NewPostIntegrationsActionExecuteBadRequest() *PostIntegrationsActionExecuteBadRequest {
	return &PostIntegrationsActionExecuteBadRequest{}
}

/*PostIntegrationsActionExecuteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsActionExecuteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionExecuteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteUnauthorized creates a PostIntegrationsActionExecuteUnauthorized with default headers values
func NewPostIntegrationsActionExecuteUnauthorized() *PostIntegrationsActionExecuteUnauthorized {
	return &PostIntegrationsActionExecuteUnauthorized{}
}

/*PostIntegrationsActionExecuteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsActionExecuteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionExecuteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteForbidden creates a PostIntegrationsActionExecuteForbidden with default headers values
func NewPostIntegrationsActionExecuteForbidden() *PostIntegrationsActionExecuteForbidden {
	return &PostIntegrationsActionExecuteForbidden{}
}

/*PostIntegrationsActionExecuteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsActionExecuteForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionExecuteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteNotFound creates a PostIntegrationsActionExecuteNotFound with default headers values
func NewPostIntegrationsActionExecuteNotFound() *PostIntegrationsActionExecuteNotFound {
	return &PostIntegrationsActionExecuteNotFound{}
}

/*PostIntegrationsActionExecuteNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostIntegrationsActionExecuteNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionExecuteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteMethodNotAllowed creates a PostIntegrationsActionExecuteMethodNotAllowed with default headers values
func NewPostIntegrationsActionExecuteMethodNotAllowed() *PostIntegrationsActionExecuteMethodNotAllowed {
	return &PostIntegrationsActionExecuteMethodNotAllowed{}
}

/*PostIntegrationsActionExecuteMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type PostIntegrationsActionExecuteMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostIntegrationsActionExecuteMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteRequestTimeout creates a PostIntegrationsActionExecuteRequestTimeout with default headers values
func NewPostIntegrationsActionExecuteRequestTimeout() *PostIntegrationsActionExecuteRequestTimeout {
	return &PostIntegrationsActionExecuteRequestTimeout{}
}

/*PostIntegrationsActionExecuteRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostIntegrationsActionExecuteRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsActionExecuteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteRequestEntityTooLarge creates a PostIntegrationsActionExecuteRequestEntityTooLarge with default headers values
func NewPostIntegrationsActionExecuteRequestEntityTooLarge() *PostIntegrationsActionExecuteRequestEntityTooLarge {
	return &PostIntegrationsActionExecuteRequestEntityTooLarge{}
}

/*PostIntegrationsActionExecuteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostIntegrationsActionExecuteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionExecuteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteUnsupportedMediaType creates a PostIntegrationsActionExecuteUnsupportedMediaType with default headers values
func NewPostIntegrationsActionExecuteUnsupportedMediaType() *PostIntegrationsActionExecuteUnsupportedMediaType {
	return &PostIntegrationsActionExecuteUnsupportedMediaType{}
}

/*PostIntegrationsActionExecuteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsActionExecuteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionExecuteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteTooManyRequests creates a PostIntegrationsActionExecuteTooManyRequests with default headers values
func NewPostIntegrationsActionExecuteTooManyRequests() *PostIntegrationsActionExecuteTooManyRequests {
	return &PostIntegrationsActionExecuteTooManyRequests{}
}

/*PostIntegrationsActionExecuteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostIntegrationsActionExecuteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionExecuteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteInternalServerError creates a PostIntegrationsActionExecuteInternalServerError with default headers values
func NewPostIntegrationsActionExecuteInternalServerError() *PostIntegrationsActionExecuteInternalServerError {
	return &PostIntegrationsActionExecuteInternalServerError{}
}

/*PostIntegrationsActionExecuteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsActionExecuteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionExecuteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteServiceUnavailable creates a PostIntegrationsActionExecuteServiceUnavailable with default headers values
func NewPostIntegrationsActionExecuteServiceUnavailable() *PostIntegrationsActionExecuteServiceUnavailable {
	return &PostIntegrationsActionExecuteServiceUnavailable{}
}

/*PostIntegrationsActionExecuteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsActionExecuteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionExecuteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionExecuteGatewayTimeout creates a PostIntegrationsActionExecuteGatewayTimeout with default headers values
func NewPostIntegrationsActionExecuteGatewayTimeout() *PostIntegrationsActionExecuteGatewayTimeout {
	return &PostIntegrationsActionExecuteGatewayTimeout{}
}

/*PostIntegrationsActionExecuteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostIntegrationsActionExecuteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionExecuteGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions/{actionId}/execute][%d] postIntegrationsActionExecuteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionExecuteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionExecuteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
