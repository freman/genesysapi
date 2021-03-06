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

// PostConversationParticipantSecureivrsessionsReader is a Reader for the PostConversationParticipantSecureivrsessions structure.
type PostConversationParticipantSecureivrsessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationParticipantSecureivrsessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationParticipantSecureivrsessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationParticipantSecureivrsessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationParticipantSecureivrsessionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationParticipantSecureivrsessionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationParticipantSecureivrsessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationParticipantSecureivrsessionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationParticipantSecureivrsessionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationParticipantSecureivrsessionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationParticipantSecureivrsessionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationParticipantSecureivrsessionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationParticipantSecureivrsessionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationParticipantSecureivrsessionsOK creates a PostConversationParticipantSecureivrsessionsOK with default headers values
func NewPostConversationParticipantSecureivrsessionsOK() *PostConversationParticipantSecureivrsessionsOK {
	return &PostConversationParticipantSecureivrsessionsOK{}
}

/*PostConversationParticipantSecureivrsessionsOK handles this case with default header values.

successful operation
*/
type PostConversationParticipantSecureivrsessionsOK struct {
	Payload *models.SecureSession
}

func (o *PostConversationParticipantSecureivrsessionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsOK  %+v", 200, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsOK) GetPayload() *models.SecureSession {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecureSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsBadRequest creates a PostConversationParticipantSecureivrsessionsBadRequest with default headers values
func NewPostConversationParticipantSecureivrsessionsBadRequest() *PostConversationParticipantSecureivrsessionsBadRequest {
	return &PostConversationParticipantSecureivrsessionsBadRequest{}
}

/*PostConversationParticipantSecureivrsessionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationParticipantSecureivrsessionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsUnauthorized creates a PostConversationParticipantSecureivrsessionsUnauthorized with default headers values
func NewPostConversationParticipantSecureivrsessionsUnauthorized() *PostConversationParticipantSecureivrsessionsUnauthorized {
	return &PostConversationParticipantSecureivrsessionsUnauthorized{}
}

/*PostConversationParticipantSecureivrsessionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationParticipantSecureivrsessionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsForbidden creates a PostConversationParticipantSecureivrsessionsForbidden with default headers values
func NewPostConversationParticipantSecureivrsessionsForbidden() *PostConversationParticipantSecureivrsessionsForbidden {
	return &PostConversationParticipantSecureivrsessionsForbidden{}
}

/*PostConversationParticipantSecureivrsessionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationParticipantSecureivrsessionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsNotFound creates a PostConversationParticipantSecureivrsessionsNotFound with default headers values
func NewPostConversationParticipantSecureivrsessionsNotFound() *PostConversationParticipantSecureivrsessionsNotFound {
	return &PostConversationParticipantSecureivrsessionsNotFound{}
}

/*PostConversationParticipantSecureivrsessionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationParticipantSecureivrsessionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsRequestEntityTooLarge creates a PostConversationParticipantSecureivrsessionsRequestEntityTooLarge with default headers values
func NewPostConversationParticipantSecureivrsessionsRequestEntityTooLarge() *PostConversationParticipantSecureivrsessionsRequestEntityTooLarge {
	return &PostConversationParticipantSecureivrsessionsRequestEntityTooLarge{}
}

/*PostConversationParticipantSecureivrsessionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationParticipantSecureivrsessionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsUnsupportedMediaType creates a PostConversationParticipantSecureivrsessionsUnsupportedMediaType with default headers values
func NewPostConversationParticipantSecureivrsessionsUnsupportedMediaType() *PostConversationParticipantSecureivrsessionsUnsupportedMediaType {
	return &PostConversationParticipantSecureivrsessionsUnsupportedMediaType{}
}

/*PostConversationParticipantSecureivrsessionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationParticipantSecureivrsessionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsTooManyRequests creates a PostConversationParticipantSecureivrsessionsTooManyRequests with default headers values
func NewPostConversationParticipantSecureivrsessionsTooManyRequests() *PostConversationParticipantSecureivrsessionsTooManyRequests {
	return &PostConversationParticipantSecureivrsessionsTooManyRequests{}
}

/*PostConversationParticipantSecureivrsessionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationParticipantSecureivrsessionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsInternalServerError creates a PostConversationParticipantSecureivrsessionsInternalServerError with default headers values
func NewPostConversationParticipantSecureivrsessionsInternalServerError() *PostConversationParticipantSecureivrsessionsInternalServerError {
	return &PostConversationParticipantSecureivrsessionsInternalServerError{}
}

/*PostConversationParticipantSecureivrsessionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationParticipantSecureivrsessionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsServiceUnavailable creates a PostConversationParticipantSecureivrsessionsServiceUnavailable with default headers values
func NewPostConversationParticipantSecureivrsessionsServiceUnavailable() *PostConversationParticipantSecureivrsessionsServiceUnavailable {
	return &PostConversationParticipantSecureivrsessionsServiceUnavailable{}
}

/*PostConversationParticipantSecureivrsessionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationParticipantSecureivrsessionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationParticipantSecureivrsessionsGatewayTimeout creates a PostConversationParticipantSecureivrsessionsGatewayTimeout with default headers values
func NewPostConversationParticipantSecureivrsessionsGatewayTimeout() *PostConversationParticipantSecureivrsessionsGatewayTimeout {
	return &PostConversationParticipantSecureivrsessionsGatewayTimeout{}
}

/*PostConversationParticipantSecureivrsessionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationParticipantSecureivrsessionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationParticipantSecureivrsessionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions][%d] postConversationParticipantSecureivrsessionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationParticipantSecureivrsessionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationParticipantSecureivrsessionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
