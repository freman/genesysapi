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

// PostConversationsMessagingIntegrationsTwitterReader is a Reader for the PostConversationsMessagingIntegrationsTwitter structure.
type PostConversationsMessagingIntegrationsTwitterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsMessagingIntegrationsTwitterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsMessagingIntegrationsTwitterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostConversationsMessagingIntegrationsTwitterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsMessagingIntegrationsTwitterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsMessagingIntegrationsTwitterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsMessagingIntegrationsTwitterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsMessagingIntegrationsTwitterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsMessagingIntegrationsTwitterRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsMessagingIntegrationsTwitterUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsMessagingIntegrationsTwitterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsMessagingIntegrationsTwitterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsMessagingIntegrationsTwitterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsMessagingIntegrationsTwitterGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsMessagingIntegrationsTwitterOK creates a PostConversationsMessagingIntegrationsTwitterOK with default headers values
func NewPostConversationsMessagingIntegrationsTwitterOK() *PostConversationsMessagingIntegrationsTwitterOK {
	return &PostConversationsMessagingIntegrationsTwitterOK{}
}

/*PostConversationsMessagingIntegrationsTwitterOK handles this case with default header values.

successful operation
*/
type PostConversationsMessagingIntegrationsTwitterOK struct {
	Payload *models.TwitterIntegration
}

func (o *PostConversationsMessagingIntegrationsTwitterOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterOK  %+v", 200, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterOK) GetPayload() *models.TwitterIntegration {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TwitterIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterAccepted creates a PostConversationsMessagingIntegrationsTwitterAccepted with default headers values
func NewPostConversationsMessagingIntegrationsTwitterAccepted() *PostConversationsMessagingIntegrationsTwitterAccepted {
	return &PostConversationsMessagingIntegrationsTwitterAccepted{}
}

/*PostConversationsMessagingIntegrationsTwitterAccepted handles this case with default header values.

Accepted - The integration creation is in progress.
*/
type PostConversationsMessagingIntegrationsTwitterAccepted struct {
	Payload *models.TwitterIntegration
}

func (o *PostConversationsMessagingIntegrationsTwitterAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterAccepted  %+v", 202, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterAccepted) GetPayload() *models.TwitterIntegration {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TwitterIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterBadRequest creates a PostConversationsMessagingIntegrationsTwitterBadRequest with default headers values
func NewPostConversationsMessagingIntegrationsTwitterBadRequest() *PostConversationsMessagingIntegrationsTwitterBadRequest {
	return &PostConversationsMessagingIntegrationsTwitterBadRequest{}
}

/*PostConversationsMessagingIntegrationsTwitterBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsMessagingIntegrationsTwitterBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterUnauthorized creates a PostConversationsMessagingIntegrationsTwitterUnauthorized with default headers values
func NewPostConversationsMessagingIntegrationsTwitterUnauthorized() *PostConversationsMessagingIntegrationsTwitterUnauthorized {
	return &PostConversationsMessagingIntegrationsTwitterUnauthorized{}
}

/*PostConversationsMessagingIntegrationsTwitterUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsMessagingIntegrationsTwitterUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterForbidden creates a PostConversationsMessagingIntegrationsTwitterForbidden with default headers values
func NewPostConversationsMessagingIntegrationsTwitterForbidden() *PostConversationsMessagingIntegrationsTwitterForbidden {
	return &PostConversationsMessagingIntegrationsTwitterForbidden{}
}

/*PostConversationsMessagingIntegrationsTwitterForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsMessagingIntegrationsTwitterForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterNotFound creates a PostConversationsMessagingIntegrationsTwitterNotFound with default headers values
func NewPostConversationsMessagingIntegrationsTwitterNotFound() *PostConversationsMessagingIntegrationsTwitterNotFound {
	return &PostConversationsMessagingIntegrationsTwitterNotFound{}
}

/*PostConversationsMessagingIntegrationsTwitterNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsMessagingIntegrationsTwitterNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterRequestTimeout creates a PostConversationsMessagingIntegrationsTwitterRequestTimeout with default headers values
func NewPostConversationsMessagingIntegrationsTwitterRequestTimeout() *PostConversationsMessagingIntegrationsTwitterRequestTimeout {
	return &PostConversationsMessagingIntegrationsTwitterRequestTimeout{}
}

/*PostConversationsMessagingIntegrationsTwitterRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsMessagingIntegrationsTwitterRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge creates a PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge with default headers values
func NewPostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge() *PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge {
	return &PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge{}
}

/*PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterUnsupportedMediaType creates a PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType with default headers values
func NewPostConversationsMessagingIntegrationsTwitterUnsupportedMediaType() *PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType {
	return &PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType{}
}

/*PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterTooManyRequests creates a PostConversationsMessagingIntegrationsTwitterTooManyRequests with default headers values
func NewPostConversationsMessagingIntegrationsTwitterTooManyRequests() *PostConversationsMessagingIntegrationsTwitterTooManyRequests {
	return &PostConversationsMessagingIntegrationsTwitterTooManyRequests{}
}

/*PostConversationsMessagingIntegrationsTwitterTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsMessagingIntegrationsTwitterTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterInternalServerError creates a PostConversationsMessagingIntegrationsTwitterInternalServerError with default headers values
func NewPostConversationsMessagingIntegrationsTwitterInternalServerError() *PostConversationsMessagingIntegrationsTwitterInternalServerError {
	return &PostConversationsMessagingIntegrationsTwitterInternalServerError{}
}

/*PostConversationsMessagingIntegrationsTwitterInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsMessagingIntegrationsTwitterInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterServiceUnavailable creates a PostConversationsMessagingIntegrationsTwitterServiceUnavailable with default headers values
func NewPostConversationsMessagingIntegrationsTwitterServiceUnavailable() *PostConversationsMessagingIntegrationsTwitterServiceUnavailable {
	return &PostConversationsMessagingIntegrationsTwitterServiceUnavailable{}
}

/*PostConversationsMessagingIntegrationsTwitterServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsMessagingIntegrationsTwitterServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsTwitterGatewayTimeout creates a PostConversationsMessagingIntegrationsTwitterGatewayTimeout with default headers values
func NewPostConversationsMessagingIntegrationsTwitterGatewayTimeout() *PostConversationsMessagingIntegrationsTwitterGatewayTimeout {
	return &PostConversationsMessagingIntegrationsTwitterGatewayTimeout{}
}

/*PostConversationsMessagingIntegrationsTwitterGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsMessagingIntegrationsTwitterGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsTwitterGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/twitter][%d] postConversationsMessagingIntegrationsTwitterGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsTwitterGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsTwitterGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
