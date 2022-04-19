// Code generated by go-swagger; DO NOT EDIT.

package web_chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostWebchatGuestConversationMemberMessagesReader is a Reader for the PostWebchatGuestConversationMemberMessages structure.
type PostWebchatGuestConversationMemberMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebchatGuestConversationMemberMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWebchatGuestConversationMemberMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWebchatGuestConversationMemberMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWebchatGuestConversationMemberMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWebchatGuestConversationMemberMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWebchatGuestConversationMemberMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWebchatGuestConversationMemberMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWebchatGuestConversationMemberMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWebchatGuestConversationMemberMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWebchatGuestConversationMemberMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWebchatGuestConversationMemberMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWebchatGuestConversationMemberMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWebchatGuestConversationMemberMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWebchatGuestConversationMemberMessagesOK creates a PostWebchatGuestConversationMemberMessagesOK with default headers values
func NewPostWebchatGuestConversationMemberMessagesOK() *PostWebchatGuestConversationMemberMessagesOK {
	return &PostWebchatGuestConversationMemberMessagesOK{}
}

/*PostWebchatGuestConversationMemberMessagesOK handles this case with default header values.

successful operation
*/
type PostWebchatGuestConversationMemberMessagesOK struct {
	Payload *models.WebChatMessage
}

func (o *PostWebchatGuestConversationMemberMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesOK  %+v", 200, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesOK) GetPayload() *models.WebChatMessage {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesBadRequest creates a PostWebchatGuestConversationMemberMessagesBadRequest with default headers values
func NewPostWebchatGuestConversationMemberMessagesBadRequest() *PostWebchatGuestConversationMemberMessagesBadRequest {
	return &PostWebchatGuestConversationMemberMessagesBadRequest{}
}

/*PostWebchatGuestConversationMemberMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWebchatGuestConversationMemberMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesUnauthorized creates a PostWebchatGuestConversationMemberMessagesUnauthorized with default headers values
func NewPostWebchatGuestConversationMemberMessagesUnauthorized() *PostWebchatGuestConversationMemberMessagesUnauthorized {
	return &PostWebchatGuestConversationMemberMessagesUnauthorized{}
}

/*PostWebchatGuestConversationMemberMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWebchatGuestConversationMemberMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesForbidden creates a PostWebchatGuestConversationMemberMessagesForbidden with default headers values
func NewPostWebchatGuestConversationMemberMessagesForbidden() *PostWebchatGuestConversationMemberMessagesForbidden {
	return &PostWebchatGuestConversationMemberMessagesForbidden{}
}

/*PostWebchatGuestConversationMemberMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWebchatGuestConversationMemberMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesNotFound creates a PostWebchatGuestConversationMemberMessagesNotFound with default headers values
func NewPostWebchatGuestConversationMemberMessagesNotFound() *PostWebchatGuestConversationMemberMessagesNotFound {
	return &PostWebchatGuestConversationMemberMessagesNotFound{}
}

/*PostWebchatGuestConversationMemberMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWebchatGuestConversationMemberMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesRequestTimeout creates a PostWebchatGuestConversationMemberMessagesRequestTimeout with default headers values
func NewPostWebchatGuestConversationMemberMessagesRequestTimeout() *PostWebchatGuestConversationMemberMessagesRequestTimeout {
	return &PostWebchatGuestConversationMemberMessagesRequestTimeout{}
}

/*PostWebchatGuestConversationMemberMessagesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWebchatGuestConversationMemberMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesRequestEntityTooLarge creates a PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge with default headers values
func NewPostWebchatGuestConversationMemberMessagesRequestEntityTooLarge() *PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge {
	return &PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge{}
}

/*PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesUnsupportedMediaType creates a PostWebchatGuestConversationMemberMessagesUnsupportedMediaType with default headers values
func NewPostWebchatGuestConversationMemberMessagesUnsupportedMediaType() *PostWebchatGuestConversationMemberMessagesUnsupportedMediaType {
	return &PostWebchatGuestConversationMemberMessagesUnsupportedMediaType{}
}

/*PostWebchatGuestConversationMemberMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWebchatGuestConversationMemberMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesTooManyRequests creates a PostWebchatGuestConversationMemberMessagesTooManyRequests with default headers values
func NewPostWebchatGuestConversationMemberMessagesTooManyRequests() *PostWebchatGuestConversationMemberMessagesTooManyRequests {
	return &PostWebchatGuestConversationMemberMessagesTooManyRequests{}
}

/*PostWebchatGuestConversationMemberMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWebchatGuestConversationMemberMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesInternalServerError creates a PostWebchatGuestConversationMemberMessagesInternalServerError with default headers values
func NewPostWebchatGuestConversationMemberMessagesInternalServerError() *PostWebchatGuestConversationMemberMessagesInternalServerError {
	return &PostWebchatGuestConversationMemberMessagesInternalServerError{}
}

/*PostWebchatGuestConversationMemberMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWebchatGuestConversationMemberMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesServiceUnavailable creates a PostWebchatGuestConversationMemberMessagesServiceUnavailable with default headers values
func NewPostWebchatGuestConversationMemberMessagesServiceUnavailable() *PostWebchatGuestConversationMemberMessagesServiceUnavailable {
	return &PostWebchatGuestConversationMemberMessagesServiceUnavailable{}
}

/*PostWebchatGuestConversationMemberMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWebchatGuestConversationMemberMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberMessagesGatewayTimeout creates a PostWebchatGuestConversationMemberMessagesGatewayTimeout with default headers values
func NewPostWebchatGuestConversationMemberMessagesGatewayTimeout() *PostWebchatGuestConversationMemberMessagesGatewayTimeout {
	return &PostWebchatGuestConversationMemberMessagesGatewayTimeout{}
}

/*PostWebchatGuestConversationMemberMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWebchatGuestConversationMemberMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/messages][%d] postWebchatGuestConversationMemberMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWebchatGuestConversationMemberMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
