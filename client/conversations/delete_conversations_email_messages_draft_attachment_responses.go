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

// DeleteConversationsEmailMessagesDraftAttachmentReader is a Reader for the DeleteConversationsEmailMessagesDraftAttachment structure.
type DeleteConversationsEmailMessagesDraftAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationsEmailMessagesDraftAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConversationsEmailMessagesDraftAttachmentOK creates a DeleteConversationsEmailMessagesDraftAttachmentOK with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentOK() *DeleteConversationsEmailMessagesDraftAttachmentOK {
	return &DeleteConversationsEmailMessagesDraftAttachmentOK{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentOK handles this case with default header values.

Operation was successful.
*/
type DeleteConversationsEmailMessagesDraftAttachmentOK struct {
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentOK ", 200)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentBadRequest creates a DeleteConversationsEmailMessagesDraftAttachmentBadRequest with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentBadRequest() *DeleteConversationsEmailMessagesDraftAttachmentBadRequest {
	return &DeleteConversationsEmailMessagesDraftAttachmentBadRequest{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationsEmailMessagesDraftAttachmentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentUnauthorized creates a DeleteConversationsEmailMessagesDraftAttachmentUnauthorized with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentUnauthorized() *DeleteConversationsEmailMessagesDraftAttachmentUnauthorized {
	return &DeleteConversationsEmailMessagesDraftAttachmentUnauthorized{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationsEmailMessagesDraftAttachmentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentForbidden creates a DeleteConversationsEmailMessagesDraftAttachmentForbidden with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentForbidden() *DeleteConversationsEmailMessagesDraftAttachmentForbidden {
	return &DeleteConversationsEmailMessagesDraftAttachmentForbidden{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationsEmailMessagesDraftAttachmentForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentNotFound creates a DeleteConversationsEmailMessagesDraftAttachmentNotFound with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentNotFound() *DeleteConversationsEmailMessagesDraftAttachmentNotFound {
	return &DeleteConversationsEmailMessagesDraftAttachmentNotFound{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteConversationsEmailMessagesDraftAttachmentNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentRequestTimeout creates a DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentRequestTimeout() *DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout {
	return &DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge creates a DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge() *DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge {
	return &DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType creates a DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType() *DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType {
	return &DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentTooManyRequests creates a DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentTooManyRequests() *DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests {
	return &DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentInternalServerError creates a DeleteConversationsEmailMessagesDraftAttachmentInternalServerError with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentInternalServerError() *DeleteConversationsEmailMessagesDraftAttachmentInternalServerError {
	return &DeleteConversationsEmailMessagesDraftAttachmentInternalServerError{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationsEmailMessagesDraftAttachmentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable creates a DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable() *DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable {
	return &DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout creates a DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout with default headers values
func NewDeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout() *DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout {
	return &DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout{}
}

/*DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/emails/{conversationId}/messages/draft/attachments/{attachmentId}][%d] deleteConversationsEmailMessagesDraftAttachmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsEmailMessagesDraftAttachmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
