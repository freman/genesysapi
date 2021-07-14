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

// PostConversationsMessagingIntegrationsLineReader is a Reader for the PostConversationsMessagingIntegrationsLine structure.
type PostConversationsMessagingIntegrationsLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsMessagingIntegrationsLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsMessagingIntegrationsLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostConversationsMessagingIntegrationsLineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsMessagingIntegrationsLineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsMessagingIntegrationsLineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsMessagingIntegrationsLineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsMessagingIntegrationsLineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsMessagingIntegrationsLineRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsMessagingIntegrationsLineRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsMessagingIntegrationsLineUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsMessagingIntegrationsLineTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsMessagingIntegrationsLineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsMessagingIntegrationsLineServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsMessagingIntegrationsLineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsMessagingIntegrationsLineOK creates a PostConversationsMessagingIntegrationsLineOK with default headers values
func NewPostConversationsMessagingIntegrationsLineOK() *PostConversationsMessagingIntegrationsLineOK {
	return &PostConversationsMessagingIntegrationsLineOK{}
}

/*PostConversationsMessagingIntegrationsLineOK handles this case with default header values.

successful operation
*/
type PostConversationsMessagingIntegrationsLineOK struct {
	Payload *models.LineIntegration
}

func (o *PostConversationsMessagingIntegrationsLineOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineOK  %+v", 200, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineOK) GetPayload() *models.LineIntegration {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineAccepted creates a PostConversationsMessagingIntegrationsLineAccepted with default headers values
func NewPostConversationsMessagingIntegrationsLineAccepted() *PostConversationsMessagingIntegrationsLineAccepted {
	return &PostConversationsMessagingIntegrationsLineAccepted{}
}

/*PostConversationsMessagingIntegrationsLineAccepted handles this case with default header values.

Accepted - If async is true, the integration creation in progress.
*/
type PostConversationsMessagingIntegrationsLineAccepted struct {
	Payload *models.LineIntegration
}

func (o *PostConversationsMessagingIntegrationsLineAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineAccepted  %+v", 202, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineAccepted) GetPayload() *models.LineIntegration {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineBadRequest creates a PostConversationsMessagingIntegrationsLineBadRequest with default headers values
func NewPostConversationsMessagingIntegrationsLineBadRequest() *PostConversationsMessagingIntegrationsLineBadRequest {
	return &PostConversationsMessagingIntegrationsLineBadRequest{}
}

/*PostConversationsMessagingIntegrationsLineBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsMessagingIntegrationsLineBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineUnauthorized creates a PostConversationsMessagingIntegrationsLineUnauthorized with default headers values
func NewPostConversationsMessagingIntegrationsLineUnauthorized() *PostConversationsMessagingIntegrationsLineUnauthorized {
	return &PostConversationsMessagingIntegrationsLineUnauthorized{}
}

/*PostConversationsMessagingIntegrationsLineUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsMessagingIntegrationsLineUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineForbidden creates a PostConversationsMessagingIntegrationsLineForbidden with default headers values
func NewPostConversationsMessagingIntegrationsLineForbidden() *PostConversationsMessagingIntegrationsLineForbidden {
	return &PostConversationsMessagingIntegrationsLineForbidden{}
}

/*PostConversationsMessagingIntegrationsLineForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsMessagingIntegrationsLineForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineNotFound creates a PostConversationsMessagingIntegrationsLineNotFound with default headers values
func NewPostConversationsMessagingIntegrationsLineNotFound() *PostConversationsMessagingIntegrationsLineNotFound {
	return &PostConversationsMessagingIntegrationsLineNotFound{}
}

/*PostConversationsMessagingIntegrationsLineNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsMessagingIntegrationsLineNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineRequestTimeout creates a PostConversationsMessagingIntegrationsLineRequestTimeout with default headers values
func NewPostConversationsMessagingIntegrationsLineRequestTimeout() *PostConversationsMessagingIntegrationsLineRequestTimeout {
	return &PostConversationsMessagingIntegrationsLineRequestTimeout{}
}

/*PostConversationsMessagingIntegrationsLineRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsMessagingIntegrationsLineRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineRequestEntityTooLarge creates a PostConversationsMessagingIntegrationsLineRequestEntityTooLarge with default headers values
func NewPostConversationsMessagingIntegrationsLineRequestEntityTooLarge() *PostConversationsMessagingIntegrationsLineRequestEntityTooLarge {
	return &PostConversationsMessagingIntegrationsLineRequestEntityTooLarge{}
}

/*PostConversationsMessagingIntegrationsLineRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsMessagingIntegrationsLineRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineUnsupportedMediaType creates a PostConversationsMessagingIntegrationsLineUnsupportedMediaType with default headers values
func NewPostConversationsMessagingIntegrationsLineUnsupportedMediaType() *PostConversationsMessagingIntegrationsLineUnsupportedMediaType {
	return &PostConversationsMessagingIntegrationsLineUnsupportedMediaType{}
}

/*PostConversationsMessagingIntegrationsLineUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsMessagingIntegrationsLineUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineTooManyRequests creates a PostConversationsMessagingIntegrationsLineTooManyRequests with default headers values
func NewPostConversationsMessagingIntegrationsLineTooManyRequests() *PostConversationsMessagingIntegrationsLineTooManyRequests {
	return &PostConversationsMessagingIntegrationsLineTooManyRequests{}
}

/*PostConversationsMessagingIntegrationsLineTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsMessagingIntegrationsLineTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineInternalServerError creates a PostConversationsMessagingIntegrationsLineInternalServerError with default headers values
func NewPostConversationsMessagingIntegrationsLineInternalServerError() *PostConversationsMessagingIntegrationsLineInternalServerError {
	return &PostConversationsMessagingIntegrationsLineInternalServerError{}
}

/*PostConversationsMessagingIntegrationsLineInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsMessagingIntegrationsLineInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineServiceUnavailable creates a PostConversationsMessagingIntegrationsLineServiceUnavailable with default headers values
func NewPostConversationsMessagingIntegrationsLineServiceUnavailable() *PostConversationsMessagingIntegrationsLineServiceUnavailable {
	return &PostConversationsMessagingIntegrationsLineServiceUnavailable{}
}

/*PostConversationsMessagingIntegrationsLineServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsMessagingIntegrationsLineServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsLineGatewayTimeout creates a PostConversationsMessagingIntegrationsLineGatewayTimeout with default headers values
func NewPostConversationsMessagingIntegrationsLineGatewayTimeout() *PostConversationsMessagingIntegrationsLineGatewayTimeout {
	return &PostConversationsMessagingIntegrationsLineGatewayTimeout{}
}

/*PostConversationsMessagingIntegrationsLineGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsMessagingIntegrationsLineGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsLineGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/line][%d] postConversationsMessagingIntegrationsLineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsLineGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsLineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
