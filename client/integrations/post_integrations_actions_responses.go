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

// PostIntegrationsActionsReader is a Reader for the PostIntegrationsActions structure.
type PostIntegrationsActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIntegrationsActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIntegrationsActionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostIntegrationsActionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostIntegrationsActionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostIntegrationsActionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostIntegrationsActionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostIntegrationsActionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostIntegrationsActionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostIntegrationsActionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostIntegrationsActionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostIntegrationsActionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostIntegrationsActionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostIntegrationsActionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIntegrationsActionsOK creates a PostIntegrationsActionsOK with default headers values
func NewPostIntegrationsActionsOK() *PostIntegrationsActionsOK {
	return &PostIntegrationsActionsOK{}
}

/*PostIntegrationsActionsOK handles this case with default header values.

successful operation
*/
type PostIntegrationsActionsOK struct {
	Payload *models.Action
}

func (o *PostIntegrationsActionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsOK  %+v", 200, o.Payload)
}

func (o *PostIntegrationsActionsOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PostIntegrationsActionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsBadRequest creates a PostIntegrationsActionsBadRequest with default headers values
func NewPostIntegrationsActionsBadRequest() *PostIntegrationsActionsBadRequest {
	return &PostIntegrationsActionsBadRequest{}
}

/*PostIntegrationsActionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostIntegrationsActionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostIntegrationsActionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsUnauthorized creates a PostIntegrationsActionsUnauthorized with default headers values
func NewPostIntegrationsActionsUnauthorized() *PostIntegrationsActionsUnauthorized {
	return &PostIntegrationsActionsUnauthorized{}
}

/*PostIntegrationsActionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostIntegrationsActionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostIntegrationsActionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsForbidden creates a PostIntegrationsActionsForbidden with default headers values
func NewPostIntegrationsActionsForbidden() *PostIntegrationsActionsForbidden {
	return &PostIntegrationsActionsForbidden{}
}

/*PostIntegrationsActionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostIntegrationsActionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsForbidden  %+v", 403, o.Payload)
}

func (o *PostIntegrationsActionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsNotFound creates a PostIntegrationsActionsNotFound with default headers values
func NewPostIntegrationsActionsNotFound() *PostIntegrationsActionsNotFound {
	return &PostIntegrationsActionsNotFound{}
}

/*PostIntegrationsActionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostIntegrationsActionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsNotFound  %+v", 404, o.Payload)
}

func (o *PostIntegrationsActionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsRequestTimeout creates a PostIntegrationsActionsRequestTimeout with default headers values
func NewPostIntegrationsActionsRequestTimeout() *PostIntegrationsActionsRequestTimeout {
	return &PostIntegrationsActionsRequestTimeout{}
}

/*PostIntegrationsActionsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostIntegrationsActionsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostIntegrationsActionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsRequestEntityTooLarge creates a PostIntegrationsActionsRequestEntityTooLarge with default headers values
func NewPostIntegrationsActionsRequestEntityTooLarge() *PostIntegrationsActionsRequestEntityTooLarge {
	return &PostIntegrationsActionsRequestEntityTooLarge{}
}

/*PostIntegrationsActionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostIntegrationsActionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostIntegrationsActionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsUnsupportedMediaType creates a PostIntegrationsActionsUnsupportedMediaType with default headers values
func NewPostIntegrationsActionsUnsupportedMediaType() *PostIntegrationsActionsUnsupportedMediaType {
	return &PostIntegrationsActionsUnsupportedMediaType{}
}

/*PostIntegrationsActionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostIntegrationsActionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostIntegrationsActionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsTooManyRequests creates a PostIntegrationsActionsTooManyRequests with default headers values
func NewPostIntegrationsActionsTooManyRequests() *PostIntegrationsActionsTooManyRequests {
	return &PostIntegrationsActionsTooManyRequests{}
}

/*PostIntegrationsActionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostIntegrationsActionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostIntegrationsActionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsInternalServerError creates a PostIntegrationsActionsInternalServerError with default headers values
func NewPostIntegrationsActionsInternalServerError() *PostIntegrationsActionsInternalServerError {
	return &PostIntegrationsActionsInternalServerError{}
}

/*PostIntegrationsActionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostIntegrationsActionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostIntegrationsActionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsServiceUnavailable creates a PostIntegrationsActionsServiceUnavailable with default headers values
func NewPostIntegrationsActionsServiceUnavailable() *PostIntegrationsActionsServiceUnavailable {
	return &PostIntegrationsActionsServiceUnavailable{}
}

/*PostIntegrationsActionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostIntegrationsActionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostIntegrationsActionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIntegrationsActionsGatewayTimeout creates a PostIntegrationsActionsGatewayTimeout with default headers values
func NewPostIntegrationsActionsGatewayTimeout() *PostIntegrationsActionsGatewayTimeout {
	return &PostIntegrationsActionsGatewayTimeout{}
}

/*PostIntegrationsActionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostIntegrationsActionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostIntegrationsActionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/integrations/actions][%d] postIntegrationsActionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostIntegrationsActionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostIntegrationsActionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
