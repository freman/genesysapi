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

// PatchConversationsCallbackReader is a Reader for the PatchConversationsCallback structure.
type PatchConversationsCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCallbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCallbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCallbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCallbackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsCallbackRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCallbackRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCallbackUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCallbackTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCallbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCallbackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCallbackGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCallbackOK creates a PatchConversationsCallbackOK with default headers values
func NewPatchConversationsCallbackOK() *PatchConversationsCallbackOK {
	return &PatchConversationsCallbackOK{}
}

/*PatchConversationsCallbackOK handles this case with default header values.

successful operation
*/
type PatchConversationsCallbackOK struct {
	Payload *models.Conversation
}

func (o *PatchConversationsCallbackOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsCallbackOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *PatchConversationsCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackBadRequest creates a PatchConversationsCallbackBadRequest with default headers values
func NewPatchConversationsCallbackBadRequest() *PatchConversationsCallbackBadRequest {
	return &PatchConversationsCallbackBadRequest{}
}

/*PatchConversationsCallbackBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCallbackBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallbackBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackUnauthorized creates a PatchConversationsCallbackUnauthorized with default headers values
func NewPatchConversationsCallbackUnauthorized() *PatchConversationsCallbackUnauthorized {
	return &PatchConversationsCallbackUnauthorized{}
}

/*PatchConversationsCallbackUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCallbackUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallbackUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackForbidden creates a PatchConversationsCallbackForbidden with default headers values
func NewPatchConversationsCallbackForbidden() *PatchConversationsCallbackForbidden {
	return &PatchConversationsCallbackForbidden{}
}

/*PatchConversationsCallbackForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCallbackForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallbackForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackNotFound creates a PatchConversationsCallbackNotFound with default headers values
func NewPatchConversationsCallbackNotFound() *PatchConversationsCallbackNotFound {
	return &PatchConversationsCallbackNotFound{}
}

/*PatchConversationsCallbackNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsCallbackNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallbackNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackRequestTimeout creates a PatchConversationsCallbackRequestTimeout with default headers values
func NewPatchConversationsCallbackRequestTimeout() *PatchConversationsCallbackRequestTimeout {
	return &PatchConversationsCallbackRequestTimeout{}
}

/*PatchConversationsCallbackRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsCallbackRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallbackRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackRequestEntityTooLarge creates a PatchConversationsCallbackRequestEntityTooLarge with default headers values
func NewPatchConversationsCallbackRequestEntityTooLarge() *PatchConversationsCallbackRequestEntityTooLarge {
	return &PatchConversationsCallbackRequestEntityTooLarge{}
}

/*PatchConversationsCallbackRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchConversationsCallbackRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallbackRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackUnsupportedMediaType creates a PatchConversationsCallbackUnsupportedMediaType with default headers values
func NewPatchConversationsCallbackUnsupportedMediaType() *PatchConversationsCallbackUnsupportedMediaType {
	return &PatchConversationsCallbackUnsupportedMediaType{}
}

/*PatchConversationsCallbackUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCallbackUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallbackUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackTooManyRequests creates a PatchConversationsCallbackTooManyRequests with default headers values
func NewPatchConversationsCallbackTooManyRequests() *PatchConversationsCallbackTooManyRequests {
	return &PatchConversationsCallbackTooManyRequests{}
}

/*PatchConversationsCallbackTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsCallbackTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallbackTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackInternalServerError creates a PatchConversationsCallbackInternalServerError with default headers values
func NewPatchConversationsCallbackInternalServerError() *PatchConversationsCallbackInternalServerError {
	return &PatchConversationsCallbackInternalServerError{}
}

/*PatchConversationsCallbackInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCallbackInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallbackInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackServiceUnavailable creates a PatchConversationsCallbackServiceUnavailable with default headers values
func NewPatchConversationsCallbackServiceUnavailable() *PatchConversationsCallbackServiceUnavailable {
	return &PatchConversationsCallbackServiceUnavailable{}
}

/*PatchConversationsCallbackServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCallbackServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallbackServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbackGatewayTimeout creates a PatchConversationsCallbackGatewayTimeout with default headers values
func NewPatchConversationsCallbackGatewayTimeout() *PatchConversationsCallbackGatewayTimeout {
	return &PatchConversationsCallbackGatewayTimeout{}
}

/*PatchConversationsCallbackGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsCallbackGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsCallbackGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks/{conversationId}][%d] patchConversationsCallbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallbackGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbackGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
