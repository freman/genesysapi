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

// DeleteWebchatGuestConversationMemberReader is a Reader for the DeleteWebchatGuestConversationMember structure.
type DeleteWebchatGuestConversationMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebchatGuestConversationMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWebchatGuestConversationMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWebchatGuestConversationMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWebchatGuestConversationMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWebchatGuestConversationMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWebchatGuestConversationMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteWebchatGuestConversationMemberRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWebchatGuestConversationMemberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWebchatGuestConversationMemberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWebchatGuestConversationMemberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWebchatGuestConversationMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWebchatGuestConversationMemberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWebchatGuestConversationMemberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWebchatGuestConversationMemberNoContent creates a DeleteWebchatGuestConversationMemberNoContent with default headers values
func NewDeleteWebchatGuestConversationMemberNoContent() *DeleteWebchatGuestConversationMemberNoContent {
	return &DeleteWebchatGuestConversationMemberNoContent{}
}

/*DeleteWebchatGuestConversationMemberNoContent handles this case with default header values.

Operation was successful.
*/
type DeleteWebchatGuestConversationMemberNoContent struct {
}

func (o *DeleteWebchatGuestConversationMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberNoContent ", 204)
}

func (o *DeleteWebchatGuestConversationMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebchatGuestConversationMemberBadRequest creates a DeleteWebchatGuestConversationMemberBadRequest with default headers values
func NewDeleteWebchatGuestConversationMemberBadRequest() *DeleteWebchatGuestConversationMemberBadRequest {
	return &DeleteWebchatGuestConversationMemberBadRequest{}
}

/*DeleteWebchatGuestConversationMemberBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWebchatGuestConversationMemberBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberUnauthorized creates a DeleteWebchatGuestConversationMemberUnauthorized with default headers values
func NewDeleteWebchatGuestConversationMemberUnauthorized() *DeleteWebchatGuestConversationMemberUnauthorized {
	return &DeleteWebchatGuestConversationMemberUnauthorized{}
}

/*DeleteWebchatGuestConversationMemberUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWebchatGuestConversationMemberUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberForbidden creates a DeleteWebchatGuestConversationMemberForbidden with default headers values
func NewDeleteWebchatGuestConversationMemberForbidden() *DeleteWebchatGuestConversationMemberForbidden {
	return &DeleteWebchatGuestConversationMemberForbidden{}
}

/*DeleteWebchatGuestConversationMemberForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWebchatGuestConversationMemberForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberNotFound creates a DeleteWebchatGuestConversationMemberNotFound with default headers values
func NewDeleteWebchatGuestConversationMemberNotFound() *DeleteWebchatGuestConversationMemberNotFound {
	return &DeleteWebchatGuestConversationMemberNotFound{}
}

/*DeleteWebchatGuestConversationMemberNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWebchatGuestConversationMemberNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberRequestTimeout creates a DeleteWebchatGuestConversationMemberRequestTimeout with default headers values
func NewDeleteWebchatGuestConversationMemberRequestTimeout() *DeleteWebchatGuestConversationMemberRequestTimeout {
	return &DeleteWebchatGuestConversationMemberRequestTimeout{}
}

/*DeleteWebchatGuestConversationMemberRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteWebchatGuestConversationMemberRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberRequestEntityTooLarge creates a DeleteWebchatGuestConversationMemberRequestEntityTooLarge with default headers values
func NewDeleteWebchatGuestConversationMemberRequestEntityTooLarge() *DeleteWebchatGuestConversationMemberRequestEntityTooLarge {
	return &DeleteWebchatGuestConversationMemberRequestEntityTooLarge{}
}

/*DeleteWebchatGuestConversationMemberRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWebchatGuestConversationMemberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberUnsupportedMediaType creates a DeleteWebchatGuestConversationMemberUnsupportedMediaType with default headers values
func NewDeleteWebchatGuestConversationMemberUnsupportedMediaType() *DeleteWebchatGuestConversationMemberUnsupportedMediaType {
	return &DeleteWebchatGuestConversationMemberUnsupportedMediaType{}
}

/*DeleteWebchatGuestConversationMemberUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWebchatGuestConversationMemberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberTooManyRequests creates a DeleteWebchatGuestConversationMemberTooManyRequests with default headers values
func NewDeleteWebchatGuestConversationMemberTooManyRequests() *DeleteWebchatGuestConversationMemberTooManyRequests {
	return &DeleteWebchatGuestConversationMemberTooManyRequests{}
}

/*DeleteWebchatGuestConversationMemberTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteWebchatGuestConversationMemberTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberInternalServerError creates a DeleteWebchatGuestConversationMemberInternalServerError with default headers values
func NewDeleteWebchatGuestConversationMemberInternalServerError() *DeleteWebchatGuestConversationMemberInternalServerError {
	return &DeleteWebchatGuestConversationMemberInternalServerError{}
}

/*DeleteWebchatGuestConversationMemberInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWebchatGuestConversationMemberInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberServiceUnavailable creates a DeleteWebchatGuestConversationMemberServiceUnavailable with default headers values
func NewDeleteWebchatGuestConversationMemberServiceUnavailable() *DeleteWebchatGuestConversationMemberServiceUnavailable {
	return &DeleteWebchatGuestConversationMemberServiceUnavailable{}
}

/*DeleteWebchatGuestConversationMemberServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWebchatGuestConversationMemberServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebchatGuestConversationMemberGatewayTimeout creates a DeleteWebchatGuestConversationMemberGatewayTimeout with default headers values
func NewDeleteWebchatGuestConversationMemberGatewayTimeout() *DeleteWebchatGuestConversationMemberGatewayTimeout {
	return &DeleteWebchatGuestConversationMemberGatewayTimeout{}
}

/*DeleteWebchatGuestConversationMemberGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWebchatGuestConversationMemberGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWebchatGuestConversationMemberGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/webchat/guest/conversations/{conversationId}/members/{memberId}][%d] deleteWebchatGuestConversationMemberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWebchatGuestConversationMemberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWebchatGuestConversationMemberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
