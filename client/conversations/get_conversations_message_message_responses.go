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

// GetConversationsMessageMessageReader is a Reader for the GetConversationsMessageMessage structure.
type GetConversationsMessageMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessageMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessageMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessageMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessageMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessageMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessageMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessageMessageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessageMessageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessageMessageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessageMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessageMessageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessageMessageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessageMessageOK creates a GetConversationsMessageMessageOK with default headers values
func NewGetConversationsMessageMessageOK() *GetConversationsMessageMessageOK {
	return &GetConversationsMessageMessageOK{}
}

/*GetConversationsMessageMessageOK handles this case with default header values.

successful operation
*/
type GetConversationsMessageMessageOK struct {
	Payload *models.MessageData
}

func (o *GetConversationsMessageMessageOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessageMessageOK) GetPayload() *models.MessageData {
	return o.Payload
}

func (o *GetConversationsMessageMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageBadRequest creates a GetConversationsMessageMessageBadRequest with default headers values
func NewGetConversationsMessageMessageBadRequest() *GetConversationsMessageMessageBadRequest {
	return &GetConversationsMessageMessageBadRequest{}
}

/*GetConversationsMessageMessageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessageMessageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessageMessageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageUnauthorized creates a GetConversationsMessageMessageUnauthorized with default headers values
func NewGetConversationsMessageMessageUnauthorized() *GetConversationsMessageMessageUnauthorized {
	return &GetConversationsMessageMessageUnauthorized{}
}

/*GetConversationsMessageMessageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessageMessageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessageMessageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageForbidden creates a GetConversationsMessageMessageForbidden with default headers values
func NewGetConversationsMessageMessageForbidden() *GetConversationsMessageMessageForbidden {
	return &GetConversationsMessageMessageForbidden{}
}

/*GetConversationsMessageMessageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessageMessageForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessageMessageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageNotFound creates a GetConversationsMessageMessageNotFound with default headers values
func NewGetConversationsMessageMessageNotFound() *GetConversationsMessageMessageNotFound {
	return &GetConversationsMessageMessageNotFound{}
}

/*GetConversationsMessageMessageNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessageMessageNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessageMessageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageRequestEntityTooLarge creates a GetConversationsMessageMessageRequestEntityTooLarge with default headers values
func NewGetConversationsMessageMessageRequestEntityTooLarge() *GetConversationsMessageMessageRequestEntityTooLarge {
	return &GetConversationsMessageMessageRequestEntityTooLarge{}
}

/*GetConversationsMessageMessageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsMessageMessageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessageMessageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageUnsupportedMediaType creates a GetConversationsMessageMessageUnsupportedMediaType with default headers values
func NewGetConversationsMessageMessageUnsupportedMediaType() *GetConversationsMessageMessageUnsupportedMediaType {
	return &GetConversationsMessageMessageUnsupportedMediaType{}
}

/*GetConversationsMessageMessageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessageMessageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessageMessageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageTooManyRequests creates a GetConversationsMessageMessageTooManyRequests with default headers values
func NewGetConversationsMessageMessageTooManyRequests() *GetConversationsMessageMessageTooManyRequests {
	return &GetConversationsMessageMessageTooManyRequests{}
}

/*GetConversationsMessageMessageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsMessageMessageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessageMessageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageInternalServerError creates a GetConversationsMessageMessageInternalServerError with default headers values
func NewGetConversationsMessageMessageInternalServerError() *GetConversationsMessageMessageInternalServerError {
	return &GetConversationsMessageMessageInternalServerError{}
}

/*GetConversationsMessageMessageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessageMessageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessageMessageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageServiceUnavailable creates a GetConversationsMessageMessageServiceUnavailable with default headers values
func NewGetConversationsMessageMessageServiceUnavailable() *GetConversationsMessageMessageServiceUnavailable {
	return &GetConversationsMessageMessageServiceUnavailable{}
}

/*GetConversationsMessageMessageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessageMessageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessageMessageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessageMessageGatewayTimeout creates a GetConversationsMessageMessageGatewayTimeout with default headers values
func NewGetConversationsMessageMessageGatewayTimeout() *GetConversationsMessageMessageGatewayTimeout {
	return &GetConversationsMessageMessageGatewayTimeout{}
}

/*GetConversationsMessageMessageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessageMessageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessageMessageGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messages/{conversationId}/messages/{messageId}][%d] getConversationsMessageMessageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessageMessageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessageMessageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
