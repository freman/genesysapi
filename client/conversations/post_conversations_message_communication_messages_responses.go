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

// PostConversationsMessageCommunicationMessagesReader is a Reader for the PostConversationsMessageCommunicationMessages structure.
type PostConversationsMessageCommunicationMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsMessageCommunicationMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsMessageCommunicationMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostConversationsMessageCommunicationMessagesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsMessageCommunicationMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsMessageCommunicationMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsMessageCommunicationMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsMessageCommunicationMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsMessageCommunicationMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsMessageCommunicationMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsMessageCommunicationMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsMessageCommunicationMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsMessageCommunicationMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsMessageCommunicationMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsMessageCommunicationMessagesOK creates a PostConversationsMessageCommunicationMessagesOK with default headers values
func NewPostConversationsMessageCommunicationMessagesOK() *PostConversationsMessageCommunicationMessagesOK {
	return &PostConversationsMessageCommunicationMessagesOK{}
}

/*PostConversationsMessageCommunicationMessagesOK handles this case with default header values.

successful operation
*/
type PostConversationsMessageCommunicationMessagesOK struct {
	Payload *models.MessageData
}

func (o *PostConversationsMessageCommunicationMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesOK  %+v", 200, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesOK) GetPayload() *models.MessageData {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesAccepted creates a PostConversationsMessageCommunicationMessagesAccepted with default headers values
func NewPostConversationsMessageCommunicationMessagesAccepted() *PostConversationsMessageCommunicationMessagesAccepted {
	return &PostConversationsMessageCommunicationMessagesAccepted{}
}

/*PostConversationsMessageCommunicationMessagesAccepted handles this case with default header values.

Accepted
*/
type PostConversationsMessageCommunicationMessagesAccepted struct {
}

func (o *PostConversationsMessageCommunicationMessagesAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesAccepted ", 202)
}

func (o *PostConversationsMessageCommunicationMessagesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConversationsMessageCommunicationMessagesBadRequest creates a PostConversationsMessageCommunicationMessagesBadRequest with default headers values
func NewPostConversationsMessageCommunicationMessagesBadRequest() *PostConversationsMessageCommunicationMessagesBadRequest {
	return &PostConversationsMessageCommunicationMessagesBadRequest{}
}

/*PostConversationsMessageCommunicationMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsMessageCommunicationMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesUnauthorized creates a PostConversationsMessageCommunicationMessagesUnauthorized with default headers values
func NewPostConversationsMessageCommunicationMessagesUnauthorized() *PostConversationsMessageCommunicationMessagesUnauthorized {
	return &PostConversationsMessageCommunicationMessagesUnauthorized{}
}

/*PostConversationsMessageCommunicationMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsMessageCommunicationMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesForbidden creates a PostConversationsMessageCommunicationMessagesForbidden with default headers values
func NewPostConversationsMessageCommunicationMessagesForbidden() *PostConversationsMessageCommunicationMessagesForbidden {
	return &PostConversationsMessageCommunicationMessagesForbidden{}
}

/*PostConversationsMessageCommunicationMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsMessageCommunicationMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesNotFound creates a PostConversationsMessageCommunicationMessagesNotFound with default headers values
func NewPostConversationsMessageCommunicationMessagesNotFound() *PostConversationsMessageCommunicationMessagesNotFound {
	return &PostConversationsMessageCommunicationMessagesNotFound{}
}

/*PostConversationsMessageCommunicationMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsMessageCommunicationMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesRequestEntityTooLarge creates a PostConversationsMessageCommunicationMessagesRequestEntityTooLarge with default headers values
func NewPostConversationsMessageCommunicationMessagesRequestEntityTooLarge() *PostConversationsMessageCommunicationMessagesRequestEntityTooLarge {
	return &PostConversationsMessageCommunicationMessagesRequestEntityTooLarge{}
}

/*PostConversationsMessageCommunicationMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsMessageCommunicationMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesUnsupportedMediaType creates a PostConversationsMessageCommunicationMessagesUnsupportedMediaType with default headers values
func NewPostConversationsMessageCommunicationMessagesUnsupportedMediaType() *PostConversationsMessageCommunicationMessagesUnsupportedMediaType {
	return &PostConversationsMessageCommunicationMessagesUnsupportedMediaType{}
}

/*PostConversationsMessageCommunicationMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsMessageCommunicationMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesTooManyRequests creates a PostConversationsMessageCommunicationMessagesTooManyRequests with default headers values
func NewPostConversationsMessageCommunicationMessagesTooManyRequests() *PostConversationsMessageCommunicationMessagesTooManyRequests {
	return &PostConversationsMessageCommunicationMessagesTooManyRequests{}
}

/*PostConversationsMessageCommunicationMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsMessageCommunicationMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesInternalServerError creates a PostConversationsMessageCommunicationMessagesInternalServerError with default headers values
func NewPostConversationsMessageCommunicationMessagesInternalServerError() *PostConversationsMessageCommunicationMessagesInternalServerError {
	return &PostConversationsMessageCommunicationMessagesInternalServerError{}
}

/*PostConversationsMessageCommunicationMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsMessageCommunicationMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesServiceUnavailable creates a PostConversationsMessageCommunicationMessagesServiceUnavailable with default headers values
func NewPostConversationsMessageCommunicationMessagesServiceUnavailable() *PostConversationsMessageCommunicationMessagesServiceUnavailable {
	return &PostConversationsMessageCommunicationMessagesServiceUnavailable{}
}

/*PostConversationsMessageCommunicationMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsMessageCommunicationMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageCommunicationMessagesGatewayTimeout creates a PostConversationsMessageCommunicationMessagesGatewayTimeout with default headers values
func NewPostConversationsMessageCommunicationMessagesGatewayTimeout() *PostConversationsMessageCommunicationMessagesGatewayTimeout {
	return &PostConversationsMessageCommunicationMessagesGatewayTimeout{}
}

/*PostConversationsMessageCommunicationMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsMessageCommunicationMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageCommunicationMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages][%d] postConversationsMessageCommunicationMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsMessageCommunicationMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageCommunicationMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
