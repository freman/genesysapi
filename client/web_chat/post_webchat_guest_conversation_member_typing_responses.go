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

// PostWebchatGuestConversationMemberTypingReader is a Reader for the PostWebchatGuestConversationMemberTyping structure.
type PostWebchatGuestConversationMemberTypingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWebchatGuestConversationMemberTypingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWebchatGuestConversationMemberTypingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWebchatGuestConversationMemberTypingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWebchatGuestConversationMemberTypingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWebchatGuestConversationMemberTypingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWebchatGuestConversationMemberTypingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWebchatGuestConversationMemberTypingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWebchatGuestConversationMemberTypingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWebchatGuestConversationMemberTypingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWebchatGuestConversationMemberTypingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWebchatGuestConversationMemberTypingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWebchatGuestConversationMemberTypingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWebchatGuestConversationMemberTypingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWebchatGuestConversationMemberTypingOK creates a PostWebchatGuestConversationMemberTypingOK with default headers values
func NewPostWebchatGuestConversationMemberTypingOK() *PostWebchatGuestConversationMemberTypingOK {
	return &PostWebchatGuestConversationMemberTypingOK{}
}

/*PostWebchatGuestConversationMemberTypingOK handles this case with default header values.

successful operation
*/
type PostWebchatGuestConversationMemberTypingOK struct {
	Payload *models.WebChatTyping
}

func (o *PostWebchatGuestConversationMemberTypingOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingOK  %+v", 200, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingOK) GetPayload() *models.WebChatTyping {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatTyping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingBadRequest creates a PostWebchatGuestConversationMemberTypingBadRequest with default headers values
func NewPostWebchatGuestConversationMemberTypingBadRequest() *PostWebchatGuestConversationMemberTypingBadRequest {
	return &PostWebchatGuestConversationMemberTypingBadRequest{}
}

/*PostWebchatGuestConversationMemberTypingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWebchatGuestConversationMemberTypingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingBadRequest  %+v", 400, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingUnauthorized creates a PostWebchatGuestConversationMemberTypingUnauthorized with default headers values
func NewPostWebchatGuestConversationMemberTypingUnauthorized() *PostWebchatGuestConversationMemberTypingUnauthorized {
	return &PostWebchatGuestConversationMemberTypingUnauthorized{}
}

/*PostWebchatGuestConversationMemberTypingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWebchatGuestConversationMemberTypingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingForbidden creates a PostWebchatGuestConversationMemberTypingForbidden with default headers values
func NewPostWebchatGuestConversationMemberTypingForbidden() *PostWebchatGuestConversationMemberTypingForbidden {
	return &PostWebchatGuestConversationMemberTypingForbidden{}
}

/*PostWebchatGuestConversationMemberTypingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostWebchatGuestConversationMemberTypingForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingForbidden  %+v", 403, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingNotFound creates a PostWebchatGuestConversationMemberTypingNotFound with default headers values
func NewPostWebchatGuestConversationMemberTypingNotFound() *PostWebchatGuestConversationMemberTypingNotFound {
	return &PostWebchatGuestConversationMemberTypingNotFound{}
}

/*PostWebchatGuestConversationMemberTypingNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostWebchatGuestConversationMemberTypingNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingNotFound  %+v", 404, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingRequestTimeout creates a PostWebchatGuestConversationMemberTypingRequestTimeout with default headers values
func NewPostWebchatGuestConversationMemberTypingRequestTimeout() *PostWebchatGuestConversationMemberTypingRequestTimeout {
	return &PostWebchatGuestConversationMemberTypingRequestTimeout{}
}

/*PostWebchatGuestConversationMemberTypingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWebchatGuestConversationMemberTypingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingRequestEntityTooLarge creates a PostWebchatGuestConversationMemberTypingRequestEntityTooLarge with default headers values
func NewPostWebchatGuestConversationMemberTypingRequestEntityTooLarge() *PostWebchatGuestConversationMemberTypingRequestEntityTooLarge {
	return &PostWebchatGuestConversationMemberTypingRequestEntityTooLarge{}
}

/*PostWebchatGuestConversationMemberTypingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostWebchatGuestConversationMemberTypingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingUnsupportedMediaType creates a PostWebchatGuestConversationMemberTypingUnsupportedMediaType with default headers values
func NewPostWebchatGuestConversationMemberTypingUnsupportedMediaType() *PostWebchatGuestConversationMemberTypingUnsupportedMediaType {
	return &PostWebchatGuestConversationMemberTypingUnsupportedMediaType{}
}

/*PostWebchatGuestConversationMemberTypingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWebchatGuestConversationMemberTypingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingTooManyRequests creates a PostWebchatGuestConversationMemberTypingTooManyRequests with default headers values
func NewPostWebchatGuestConversationMemberTypingTooManyRequests() *PostWebchatGuestConversationMemberTypingTooManyRequests {
	return &PostWebchatGuestConversationMemberTypingTooManyRequests{}
}

/*PostWebchatGuestConversationMemberTypingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWebchatGuestConversationMemberTypingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingInternalServerError creates a PostWebchatGuestConversationMemberTypingInternalServerError with default headers values
func NewPostWebchatGuestConversationMemberTypingInternalServerError() *PostWebchatGuestConversationMemberTypingInternalServerError {
	return &PostWebchatGuestConversationMemberTypingInternalServerError{}
}

/*PostWebchatGuestConversationMemberTypingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWebchatGuestConversationMemberTypingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingServiceUnavailable creates a PostWebchatGuestConversationMemberTypingServiceUnavailable with default headers values
func NewPostWebchatGuestConversationMemberTypingServiceUnavailable() *PostWebchatGuestConversationMemberTypingServiceUnavailable {
	return &PostWebchatGuestConversationMemberTypingServiceUnavailable{}
}

/*PostWebchatGuestConversationMemberTypingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWebchatGuestConversationMemberTypingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWebchatGuestConversationMemberTypingGatewayTimeout creates a PostWebchatGuestConversationMemberTypingGatewayTimeout with default headers values
func NewPostWebchatGuestConversationMemberTypingGatewayTimeout() *PostWebchatGuestConversationMemberTypingGatewayTimeout {
	return &PostWebchatGuestConversationMemberTypingGatewayTimeout{}
}

/*PostWebchatGuestConversationMemberTypingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostWebchatGuestConversationMemberTypingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostWebchatGuestConversationMemberTypingGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}/typing][%d] postWebchatGuestConversationMemberTypingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWebchatGuestConversationMemberTypingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWebchatGuestConversationMemberTypingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
