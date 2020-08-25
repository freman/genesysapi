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

// PostConversationsCallReader is a Reader for the PostConversationsCall structure.
type PostConversationsCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsCallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsCallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsCallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsCallNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsCallRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsCallUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsCallTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsCallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsCallServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsCallGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsCallOK creates a PostConversationsCallOK with default headers values
func NewPostConversationsCallOK() *PostConversationsCallOK {
	return &PostConversationsCallOK{}
}

/*PostConversationsCallOK handles this case with default header values.

successful operation
*/
type PostConversationsCallOK struct {
	Payload *models.Conversation
}

func (o *PostConversationsCallOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallOK  %+v", 200, o.Payload)
}

func (o *PostConversationsCallOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *PostConversationsCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallBadRequest creates a PostConversationsCallBadRequest with default headers values
func NewPostConversationsCallBadRequest() *PostConversationsCallBadRequest {
	return &PostConversationsCallBadRequest{}
}

/*PostConversationsCallBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsCallBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsCallBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallUnauthorized creates a PostConversationsCallUnauthorized with default headers values
func NewPostConversationsCallUnauthorized() *PostConversationsCallUnauthorized {
	return &PostConversationsCallUnauthorized{}
}

/*PostConversationsCallUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsCallUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsCallUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallForbidden creates a PostConversationsCallForbidden with default headers values
func NewPostConversationsCallForbidden() *PostConversationsCallForbidden {
	return &PostConversationsCallForbidden{}
}

/*PostConversationsCallForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsCallForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsCallForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallNotFound creates a PostConversationsCallNotFound with default headers values
func NewPostConversationsCallNotFound() *PostConversationsCallNotFound {
	return &PostConversationsCallNotFound{}
}

/*PostConversationsCallNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsCallNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsCallNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallRequestEntityTooLarge creates a PostConversationsCallRequestEntityTooLarge with default headers values
func NewPostConversationsCallRequestEntityTooLarge() *PostConversationsCallRequestEntityTooLarge {
	return &PostConversationsCallRequestEntityTooLarge{}
}

/*PostConversationsCallRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsCallRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsCallRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallUnsupportedMediaType creates a PostConversationsCallUnsupportedMediaType with default headers values
func NewPostConversationsCallUnsupportedMediaType() *PostConversationsCallUnsupportedMediaType {
	return &PostConversationsCallUnsupportedMediaType{}
}

/*PostConversationsCallUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsCallUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsCallUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallTooManyRequests creates a PostConversationsCallTooManyRequests with default headers values
func NewPostConversationsCallTooManyRequests() *PostConversationsCallTooManyRequests {
	return &PostConversationsCallTooManyRequests{}
}

/*PostConversationsCallTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsCallTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsCallTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallInternalServerError creates a PostConversationsCallInternalServerError with default headers values
func NewPostConversationsCallInternalServerError() *PostConversationsCallInternalServerError {
	return &PostConversationsCallInternalServerError{}
}

/*PostConversationsCallInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsCallInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsCallInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallServiceUnavailable creates a PostConversationsCallServiceUnavailable with default headers values
func NewPostConversationsCallServiceUnavailable() *PostConversationsCallServiceUnavailable {
	return &PostConversationsCallServiceUnavailable{}
}

/*PostConversationsCallServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsCallServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsCallServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallGatewayTimeout creates a PostConversationsCallGatewayTimeout with default headers values
func NewPostConversationsCallGatewayTimeout() *PostConversationsCallGatewayTimeout {
	return &PostConversationsCallGatewayTimeout{}
}

/*PostConversationsCallGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsCallGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsCallGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}][%d] postConversationsCallGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsCallGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}