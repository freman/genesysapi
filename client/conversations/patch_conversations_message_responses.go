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

// PatchConversationsMessageReader is a Reader for the PatchConversationsMessage structure.
type PatchConversationsMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsMessageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsMessageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsMessageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsMessageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsMessageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsMessageOK creates a PatchConversationsMessageOK with default headers values
func NewPatchConversationsMessageOK() *PatchConversationsMessageOK {
	return &PatchConversationsMessageOK{}
}

/*PatchConversationsMessageOK handles this case with default header values.

successful operation
*/
type PatchConversationsMessageOK struct {
	Payload *models.Conversation
}

func (o *PatchConversationsMessageOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsMessageOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *PatchConversationsMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageBadRequest creates a PatchConversationsMessageBadRequest with default headers values
func NewPatchConversationsMessageBadRequest() *PatchConversationsMessageBadRequest {
	return &PatchConversationsMessageBadRequest{}
}

/*PatchConversationsMessageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsMessageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageUnauthorized creates a PatchConversationsMessageUnauthorized with default headers values
func NewPatchConversationsMessageUnauthorized() *PatchConversationsMessageUnauthorized {
	return &PatchConversationsMessageUnauthorized{}
}

/*PatchConversationsMessageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsMessageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageForbidden creates a PatchConversationsMessageForbidden with default headers values
func NewPatchConversationsMessageForbidden() *PatchConversationsMessageForbidden {
	return &PatchConversationsMessageForbidden{}
}

/*PatchConversationsMessageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsMessageForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageNotFound creates a PatchConversationsMessageNotFound with default headers values
func NewPatchConversationsMessageNotFound() *PatchConversationsMessageNotFound {
	return &PatchConversationsMessageNotFound{}
}

/*PatchConversationsMessageNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsMessageNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageRequestEntityTooLarge creates a PatchConversationsMessageRequestEntityTooLarge with default headers values
func NewPatchConversationsMessageRequestEntityTooLarge() *PatchConversationsMessageRequestEntityTooLarge {
	return &PatchConversationsMessageRequestEntityTooLarge{}
}

/*PatchConversationsMessageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchConversationsMessageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageUnsupportedMediaType creates a PatchConversationsMessageUnsupportedMediaType with default headers values
func NewPatchConversationsMessageUnsupportedMediaType() *PatchConversationsMessageUnsupportedMediaType {
	return &PatchConversationsMessageUnsupportedMediaType{}
}

/*PatchConversationsMessageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsMessageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageTooManyRequests creates a PatchConversationsMessageTooManyRequests with default headers values
func NewPatchConversationsMessageTooManyRequests() *PatchConversationsMessageTooManyRequests {
	return &PatchConversationsMessageTooManyRequests{}
}

/*PatchConversationsMessageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchConversationsMessageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageInternalServerError creates a PatchConversationsMessageInternalServerError with default headers values
func NewPatchConversationsMessageInternalServerError() *PatchConversationsMessageInternalServerError {
	return &PatchConversationsMessageInternalServerError{}
}

/*PatchConversationsMessageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsMessageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageServiceUnavailable creates a PatchConversationsMessageServiceUnavailable with default headers values
func NewPatchConversationsMessageServiceUnavailable() *PatchConversationsMessageServiceUnavailable {
	return &PatchConversationsMessageServiceUnavailable{}
}

/*PatchConversationsMessageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsMessageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessageGatewayTimeout creates a PatchConversationsMessageGatewayTimeout with default headers values
func NewPatchConversationsMessageGatewayTimeout() *PatchConversationsMessageGatewayTimeout {
	return &PatchConversationsMessageGatewayTimeout{}
}

/*PatchConversationsMessageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsMessageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessageGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messages/{conversationId}][%d] patchConversationsMessageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
